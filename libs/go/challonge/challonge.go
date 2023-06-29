package challonge

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/codephobia/pool-overlay/libs/go/events"
	"github.com/codephobia/pool-overlay/libs/go/models"
	overlayPkg "github.com/codephobia/pool-overlay/libs/go/overlay"
	"github.com/codephobia/pool-overlay/libs/go/state"
	"gorm.io/gorm"
)

// Challonge stores the current challonge instance.
type Challonge struct {
	config  *Config
	db      *gorm.DB
	overlay *overlayPkg.Overlay
	tables  map[int]*state.State

	Tournament *Tournament `json:"tournament"`
	Settings   *Settings   `json:"settings"`

	// PlayersMap is a map of challonge player ids to players in the db.
	PlayersMap map[int]*models.Player `json:"-"`

	// CurrentMatches is a map of table number to match.
	CurrentMatches map[int]*Match `json:"-"`
}

type Settings struct {
	MaxTables     int             `json:"max_tables"`
	GameType      models.GameType `json:"game_type"`
	ShowOverlay   bool            `json:"show_overlay"`
	ShowFlags     bool            `json:"show_flags"`
	ShowFargo     bool            `json:"show_fargo"`
	ShowScore     bool            `json:"show_score"`
	IsHandicapped bool            `json:"is_handicapped"`
	ASideRaceTo   int             `json:"a_side_race_to"`
	BSideRaceTo   int             `json:"b_side_race_to"`
}

// NewChallonge returns a new Challonge.
func NewChallonge(apiKey string, username string, db *gorm.DB, overlay *overlayPkg.Overlay, tables map[int]*state.State) *Challonge {
	return &Challonge{
		config:  NewConfig(apiKey, username),
		db:      db,
		overlay: overlay,
		tables:  tables,

		PlayersMap:     make(map[int]*models.Player),
		CurrentMatches: make(map[int]*Match),
	}
}

// InTournamentMode returns if a tournament is currently loaded.
func (c *Challonge) InTournamentMode() bool {
	return c.Tournament != nil
}

// GetTournamentList returns a list of incomplete tournaments on Challonge account.
func (c *Challonge) GetTournamentList() ([]*Tournament, error) {
	return getLatestTournaments(c.config.Username, c.config.APIKey)
}

// GetTournamentByID returns a tournament by id.
func (c *Challonge) GetTournamentByID(id int) (*Tournament, error) {
	return getTournamentByID(id, c.config.Username, c.config.APIKey)
}

// loadTournament loads a tournament by id.
func (c *Challonge) LoadTournament(id int, settings *Settings) error {
	// get tournament
	tournament, err := getTournamentByID(id, c.config.Username, c.config.APIKey)
	if err != nil {
		return err
	}

	// validate tournament
	err = tournament.Validate()
	if err != nil {
		return err
	}

	// load tournament
	c.Tournament = tournament
	c.Settings = settings

	// initialize tournament
	if err := c.initializeTournament(); err != nil {
		c.UnloadTournament()
		return err
	}

	// map players and load from database
	if err := c.mapPlayers(); err != nil {
		c.UnloadTournament()
		return err
	}

	// load initial matches
	if err := c.fillTables(); err != nil {
		c.UnloadTournament()
		return err
	}

	return nil
}

// UnloadTournament unsets the current tournament.
func (c *Challonge) UnloadTournament() {
	c.Tournament = nil
	c.Settings = nil
	c.PlayersMap = make(map[int]*models.Player)
	c.CurrentMatches = make(map[int]*Match)
}

// Initializes all the settings for the tables and overlays.
func (c *Challonge) initializeTournament() error {
	log.Printf("*** SETTINGS 3 ***: %+v", c.Settings)

	for i := 1; i <= c.Settings.MaxTables; i++ {
		// **********
		// ** GAME **
		// **********
		// Reset the table.
		c.tables[i].Game.ResetScore()
		c.tables[i].Game.UnsetPlayer(1)
		c.tables[i].Game.UnsetPlayer(2)

		// Set g)ame settings.
		c.tables[i].Game.SetType(c.Settings.GameType)
		c.tables[i].Game.SetVsMode(models.OneVsOne)
		c.tables[i].Game.SetUseFargoHotHandicap(c.Settings.IsHandicapped)
		c.tables[i].Game.SetRaceTo(c.Settings.ASideRaceTo)

		// Generate game message to broadcast to overlay.
		gameMessage, err := overlayPkg.NewEvent(
			events.GameEventType,
			events.NewGameEventPayload(c.tables[i].Game),
		).ToBytes()
		if err != nil {
			return err
		}

		// Broadcast new game state to overlay.
		c.overlay.Broadcast <- gameMessage

		// *************
		// ** OVERLAY **
		// *************
		c.tables[i].Overlay.SetHidden(!c.Settings.ShowOverlay)
		c.tables[i].Overlay.SetFlags(c.Settings.ShowFlags)
		c.tables[i].Overlay.SetFargo(c.Settings.ShowFargo)
		c.tables[i].Overlay.SetScore(c.Settings.ShowScore)

		// Generate overlay state message to broadcast to overlay.
		overlayMessage, err := overlayPkg.NewEvent(
			events.OverlayStateEventType,
			c.tables[i].Overlay,
		).ToBytes()
		if err != nil {
			return err
		}

		// Broadcast new overlay state to overlay.
		c.overlay.Broadcast <- overlayMessage
	}

	return nil
}

// Maps the players by their challonge id to a databased player.
func (c *Challonge) mapPlayers() error {
	for _, participant := range c.Tournament.Participants {
		fargoObservableID, err := getFargoObservableIDFromParticpantName(participant.Name)
		if err != nil {
			return err
		}

		var player models.Player
		if err := player.LoadByFargoObservableID(c.db, fargoObservableID); err != nil {
			return err
		}

		c.PlayersMap[participant.ID] = &player
	}

	return nil
}

// Gets a Fargo observable ID from their name.
func getFargoObservableIDFromParticpantName(name string) (int, error) {
	pattern := `\(.*?(\d+)\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(name)
	if len(matches) > 1 {
		secondNumberStr := matches[1]
		secondNumber, err := strconv.Atoi(secondNumberStr)
		if err != nil {
			return 0, err
		}
		return secondNumber, nil
	}

	return 0, fmt.Errorf("no observable fargo id found in player: %s", name)
}

// Fills all tables with an initial match.
func (c *Challonge) fillTables() error {
	for i := 1; i <= c.Settings.MaxTables; i++ {
		log.Printf("filling table: %d", i)

		match := c.Tournament.GetNextMatch()

		if match == nil {
			return fmt.Errorf("no next match found")
		}

		log.Printf("using match: %d", match.ID)

		// add that match to the current matches for that table
		c.CurrentMatches[i] = match

		log.Printf("player 1 challonge id: %d", *match.Player1ID)
		log.Printf("player 2 challonge id: %d", *match.Player2ID)

		log.Printf("player 1 db id: %d", c.PlayersMap[*match.Player1ID].ID)
		log.Printf("player 2 db id: %d", c.PlayersMap[*match.Player2ID].ID)

		// load players to the overlay for that table
		c.tables[i].Game.SetPlayer(1, c.PlayersMap[*match.Player1ID])
		c.tables[i].Game.SetPlayer(2, c.PlayersMap[*match.Player2ID])

		// mark match as in progress on challonge if possible
		if err := match.SetInProgress(c.config.Username, c.config.APIKey); err != nil {
			return err
		}
	}

	return nil
}
