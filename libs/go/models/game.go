package models

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	// ErrInvalidPlayerNumber - Invalid player number.
	ErrInvalidPlayerNumber = errors.New("invalid player number")
	// ErrInvalidPlayerID - Invalid player ID.
	ErrInvalidPlayerID = errors.New("invalid player id")
	// ErrInvalidTeamNumber - Invalid team number.
	ErrInvalidTeamNumber = errors.New("invalid team number")
)

// Game is the current state of the game being played.
type Game struct {
	db *gorm.DB

	ID        uint       `json:"id,omitempty" gorm:"primarykey"`
	Type      GameType   `json:"type"`
	VsMode    GameVsMode `json:"vs_mode"`
	RaceTo    int        `json:"race_to"`
	ScoreOne  int        `json:"score_one"`
	ScoreTwo  int        `json:"score_two"`
	Completed bool       `json:"completed"`

	PlayerOneID *uint `json:"player_one_id,omitempty"`
	PlayerTwoID *uint `json:"player_two_id,omitempty"`
	TeamOneID   *uint `json:"team_one_id,omitempty"`
	TeamTwoID   *uint `json:"team_two_id,omitempty"`

	PlayerOne *Player `json:"player_one,omitempty" gorm:"foreignKey:player_one_id"`
	PlayerTwo *Player `json:"player_two,omitempty" gorm:"foreignKey:player_two_id"`
	TeamOne   *Team   `json:"team_one,omitempty" gorm:"foreignKey:team_one_id"`
	TeamTwo   *Team   `json:"team_two,omitempty" gorm:"foreignKey:team_two_id"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	mutex sync.Mutex `json:"-"`
}

// NewGame returns a new Game with default settings.
func NewGame(db *gorm.DB) *Game {
	return &Game{
		db: db,

		Type:   EightBall,
		RaceTo: 5,
		VsMode: OneVsOne,
	}
}

func (g *Game) LoadLatest() *Game {
	var latest Game
	latest.db = g.db

	result := g.db.
		Where("completed = ?", false).
		Order("id desc").
		Limit(1).
		Preload("PlayerOne", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "flag_id", "fargo_id", "fargo_rating").
				Preload("Flag", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "country", "image_path")
				})
		}).
		Preload("PlayerTwo", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "flag_id", "fargo_id", "fargo_rating").
				Preload("Flag", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "country", "image_path")
				})
		}).
		Preload("TeamOne", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name").
				Preload("Players", func(db *gorm.DB) *gorm.DB {
					return db.
						Select("id", "team_id", "player_id", "captain").
						Preload("Player", func(db *gorm.DB) *gorm.DB {
							return db.Select("id", "name", "flag_id").
								Preload("Flag", func(db *gorm.DB) *gorm.DB {
									return db.Select("id", "country", "image_path")
								})
						})
				})
		}).
		Preload("TeamTwo", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name").
				Preload("Players", func(db *gorm.DB) *gorm.DB {
					return db.
						Select("id", "team_id", "player_id", "captain").
						Preload("Player", func(db *gorm.DB) *gorm.DB {
							return db.Select("id", "name", "flag_id").
								Preload("Flag", func(db *gorm.DB) *gorm.DB {
									return db.Select("id", "country", "image_path")
								})
						})
				})
		}).
		Find(&latest)

	if result.Error == nil && result.RowsAffected == 1 {
		return &latest
	}

	return g
}

// Reset will reset a game to a new one, while keeping some information about
// the existing one.
func (g *Game) Reset() error {
	// lock
	g.mutex.Lock()

	// unset id so that on save it creates a new one in the database
	g.ID = 0
	// mark as incomplete
	g.Completed = false

	// unlock
	g.mutex.Unlock()

	// reset the score
	if err := g.ResetScore(); err != nil {
		return err
	}

	// save the new game so that it will reload
	return g.Save(false)
}

// SetType sets the Type of the current game.
func (g *Game) SetType(t GameType) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.Type = t

	return g.Save(false)
}

// SetVsMode changes the current GameVsMode of the game.
func (g *Game) SetVsMode(mode GameVsMode) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.VsMode = mode

	return g.Save(false)
}

// IncrementRaceTo increases the RaceTo of the game by one.
func (g *Game) IncrementRaceTo() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.RaceTo++

	return g.Save(false)
}

// DecrementRaceTo decreases the RaceTo of the game by one. Minimum value will
// always be 1.
func (g *Game) DecrementRaceTo() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	newRaceTo := g.RaceTo - 1

	if newRaceTo > 0 {
		g.RaceTo = newRaceTo
	}

	return g.Save(false)
}

// IncrementScore increases the score for the specified player by one.
func (g *Game) IncrementScore(playerNum int) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		newScore := g.ScoreOne + 1
		// if newScore <= g.RaceTo {
		g.ScoreOne = newScore
		// }
	case 2:
		newScore := g.ScoreTwo + 1
		// if newScore <= g.RaceTo {
		g.ScoreTwo = newScore
		// }
	default:
		return ErrInvalidPlayerNumber
	}

	return g.Save(false)
}

// DecrementScore decreases the score for the specified player by one.
func (g *Game) DecrementScore(playerNum int) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		newScore := g.ScoreOne - 1
		if newScore >= 0 {
			g.ScoreOne = newScore
		}
	case 2:
		newScore := g.ScoreTwo - 1
		if newScore >= 0 {
			g.ScoreTwo = newScore
		}
	default:
		return ErrInvalidPlayerNumber
	}

	return g.Save(false)
}

// ResetScore resets the current game score.
func (g *Game) ResetScore() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.ScoreOne = 0
	g.ScoreTwo = 0

	return g.Save(false)
}

// SetPlayer sets the player number to the specified player.
func (g *Game) SetPlayer(playerNum int, player *Player) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		g.PlayerOneID = &player.ID
		g.PlayerOne = player
	case 2:
		g.PlayerTwoID = &player.ID
		g.PlayerTwo = player
	default:
		return ErrInvalidPlayerNumber
	}

	// Unset teams when adding a player.
	g.TeamOneID = nil
	g.TeamOne = nil
	g.TeamTwoID = nil
	g.TeamTwo = nil

	return g.Save(false)
}

// UnsetPlayer unsets a player for the specified player num.
func (g *Game) UnsetPlayer(playerNum int) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		g.PlayerOneID = nil
		g.PlayerOne = nil
	case 2:
		g.PlayerTwoID = nil
		g.PlayerTwo = nil
	default:
		return ErrInvalidPlayerNumber
	}

	return g.Save(false)
}

// SetPlayerFlag sets the flag to the specified player.
func (g *Game) SetPlayerFlag(playerNum int, flag *Flag) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		// Handle if player has currently been unset.
		if g.PlayerOne == nil {
			g.PlayerOne = &Player{}
		}

		g.PlayerOne.FlagID = flag.ID
		g.PlayerOne.Flag = flag
	case 2:
		// Handle if player has currently been unset.
		if g.PlayerTwo == nil {
			g.PlayerTwo = &Player{}
		}

		g.PlayerTwo.FlagID = flag.ID
		g.PlayerTwo.Flag = flag
	default:
		return ErrInvalidPlayerNumber
	}

	return nil
}

// SetPlayerName sets the name to the specified player.
func (g *Game) SetPlayerName(playerNum int, name string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		// Handle if player has currently been unset.
		if g.PlayerOne == nil {
			g.PlayerOne = &Player{}
		}

		g.PlayerOne.Name = name
	case 2:
		// Handle if player has currently been unset.
		if g.PlayerTwo == nil {
			g.PlayerTwo = &Player{}
		}

		g.PlayerTwo.Name = name
	default:
		return ErrInvalidPlayerNumber
	}

	return nil
}

// SetTeam sets the team number to the specified team.
func (g *Game) SetTeam(teamNum int, team *Team) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch teamNum {
	case 1:
		g.TeamOneID = &team.ID
		g.TeamOne = team
	case 2:
		g.TeamTwoID = &team.ID
		g.TeamTwo = team
	default:
		return ErrInvalidTeamNumber
	}

	// Unset players when adding a team.
	g.PlayerOneID = nil
	g.PlayerOne = nil
	g.PlayerTwoID = nil
	g.PlayerTwo = nil

	return g.Save(false)
}

// Save saves the game to the database.
func (g *Game) Save(completed bool) error {
	if g.ID != 0 {
		return g.db.Model(g).Updates(map[string]interface{}{
			"type":      g.Type,
			"vs_mode":   g.VsMode,
			"race_to":   g.RaceTo,
			"score_one": g.ScoreOne,
			"score_two": g.ScoreTwo,
			"completed": completed,

			"player_one_id": g.PlayerOneID,
			"player_two_id": g.PlayerTwoID,
			"team_one_id":   g.TeamOneID,
			"team_two_id":   g.TeamTwoID,
		}).Error
	}

	return g.db.Create(g).Error
}
