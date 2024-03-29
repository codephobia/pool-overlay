package models

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/codephobia/pool-overlay/libs/go/utils"
	"gorm.io/gorm"
)

var (
	// ErrInvalidPlayerNumber - Invalid player number.
	ErrInvalidPlayerNumber = errors.New("invalid player number")
	// ErrInvalidPlayerID - Invalid player ID.
	ErrInvalidPlayerID = errors.New("invalid player id")
	// ErrInvalidTeamNumber - Invalid team number.
	ErrInvalidTeamNumber = errors.New("invalid team number")
	// ErrInvalidRaceTo - Invalid race to number.
	ErrInvalidRaceTo = errors.New("invalid race to number")
	// The minimum number of games required to use handicapping.
	minHandicapRaceTo = 2
	// The maximum number of games allowed to use handicapping.
	maxHandicapRaceTo = 11
)

// Game is the current state of the game being played.
type Game struct {
	db *gorm.DB

	ID        uint       `json:"id,omitempty" gorm:"primarykey"`
	Table     int        `json:"table" gorm:"column:table_num"`
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

	UseFargoHotHandicap bool              `json:"use_fargo_hot_handicap"`
	FargoHotHandicapID  *uint             `json:"fargo_hot_handicap_id,omitempty"`
	FargoHotHandicap    *FargoHotHandicap `json:"fargo_hot_handicap,omitempty" gorm:"foreignKey:fargo_hot_handicap_id"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	mutex sync.Mutex `json:"-"`
}

// NewGame returns a new Game with default settings.
func NewGame(db *gorm.DB, table int) *Game {
	return &Game{
		db: db,

		Table:  table,
		Type:   EightBall,
		RaceTo: 5,
		VsMode: OneVsOne,

		UseFargoHotHandicap: false,
	}
}

func (g *Game) LoadLatest(table int) *Game {
	var latest Game
	latest.db = g.db

	result := g.db.
		Where("completed = ? AND table_num = ?", false, table).
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
		Preload("FargoHotHandicap", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "race_to", "difference_start", "difference_end", "wins_higher", "wins_lower")
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

// SetRaceTo sets the race to a specified amount.
func (g *Game) SetRaceTo(raceTo int) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.RaceTo = raceTo

	// TODO: Reach out to Fargo and see if we can account for all races.
	// Make sure we aren't using handicapping if outside of the race threshold.
	if g.RaceTo > maxHandicapRaceTo {
		if err := g.SetUseFargoHotHandicap(false); err != nil {
			return err
		}
	}

	// Update fargo hot handicap if we are using it.
	if g.UseFargoHotHandicap {
		if err := g.updateFargoHotHandicap(); err != nil {
			return err
		}
	}

	return g.Save(false)
}

// IncrementRaceTo increases the RaceTo of the game by one.
func (g *Game) IncrementRaceTo() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.RaceTo++

	// TODO: Reach out to Fargo and see if we can account for all races.
	// Make sure we aren't using handicapping if outside of the race threshold.
	if g.RaceTo > maxHandicapRaceTo {
		if err := g.SetUseFargoHotHandicap(false); err != nil {
			return err
		}
	}

	// Update fargo hot handicap if we are using it.
	if g.UseFargoHotHandicap {
		if err := g.updateFargoHotHandicap(); err != nil {
			return err
		}
	}

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

	// Make sure we aren't using handicapping if outside of the race threshold.
	if g.RaceTo < minHandicapRaceTo {
		if err := g.SetUseFargoHotHandicap(false); err != nil {
			return err
		}
	}

	// Update fargo hot handicap if we are using it.
	if g.UseFargoHotHandicap {
		if err := g.updateFargoHotHandicap(); err != nil {
			return err
		}
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

	// Update fargo hot handicap if we are using it.
	if g.UseFargoHotHandicap {
		if err := g.updateFargoHotHandicap(); err != nil {
			return err
		}
	}

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

	// Update fargo hot handicap if we are using it.
	if g.UseFargoHotHandicap {
		if err := g.updateFargoHotHandicap(); err != nil {
			return err
		}
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

// SetUseFargoHotHandicap updates if we are using fargo hot handicap for the
// current game.
func (g *Game) SetUseFargoHotHandicap(use bool) error {
	// Make sure race to is within allowed threshold.
	if use && (g.RaceTo < minHandicapRaceTo || g.RaceTo > maxHandicapRaceTo) {
		return ErrInvalidRaceTo
	}

	g.UseFargoHotHandicap = use

	if use {
		// Update handicap if both players are set.
		if g.PlayerOne != nil && g.PlayerTwo != nil {
			if err := g.updateFargoHotHandicap(); err != nil {
				return err
			}
		}
	} else {
		g.FargoHotHandicapID = nil
		g.FargoHotHandicap = nil
	}

	return g.Save(false)
}

// SOMETHING IN HERE IS CAUSING AN ERROR.

// Updates the handicap to be used for the current game.
// We don't need to call game.Save here because anything that calls this will
// end up saving the game anyway.
func (g *Game) updateFargoHotHandicap() error {
	// If both players are not set, unset the fargo hot handicap.
	if g.PlayerOne == nil || g.PlayerTwo == nil {
		g.FargoHotHandicapID = nil
		g.FargoHotHandicap = nil

		return nil
	}

	// get the handicaps for the current race to.
	var handicaps FargoHotHandicaps
	if err := handicaps.LoadByRaceTo(g.db, g.RaceTo); err != nil {
		return err
	}

	log.Printf("handicaps length: %d", len(handicaps))

	// Figure out difference in handicap between both players.
	diff := utils.Abs(int(g.PlayerOne.FargoRating - g.PlayerTwo.FargoRating))

	log.Printf("handicap diff: %d", diff)

	// Loop through handicaps to find the appropriate handicap based on fargo
	// difference.
	var currentHandicap FargoHotHandicap
	for _, handicap := range handicaps {
		currentHandicap = handicap

		log.Printf("current handicap id: %d", currentHandicap.ID)

		// Check if this is the last handicap or the difference end is less than
		// or equal to the rating difference.
		if handicap.DifferenceEnd == nil || int(*handicap.DifferenceEnd) >= diff {
			break
		}
	}

	log.Printf("handicap id: %+v", currentHandicap.ID)

	// Set the fargo handicap to the game.
	g.FargoHotHandicapID = &currentHandicap.ID
	g.FargoHotHandicap = &currentHandicap

	return nil
}

// WinningPlayerNum returns the winning player number (1 or 2) or if no winner,
// returns 0.
func (g *Game) WinningPlayerNum() int {
	// If not using handicap just take the highest score.
	if !g.UseFargoHotHandicap && g.ScoreOne != g.ScoreTwo && (g.ScoreOne == g.RaceTo || g.ScoreTwo == g.RaceTo) {
		if g.ScoreOne > g.ScoreTwo {
			return 1
		} else {
			return 2
		}
	} else if g.UseFargoHotHandicap {
		playerOneFargo := g.PlayerOne.FargoRating
		playerTwoFargo := g.PlayerTwo.FargoRating
		winsHigher := g.FargoHotHandicap.WinsHigher
		winsLower := g.FargoHotHandicap.WinsLower

		var playerOneRaceTo uint
		if playerOneFargo > playerTwoFargo {
			playerOneRaceTo = winsHigher
		} else {
			playerOneRaceTo = winsLower
		}

		var playerTwoRaceTo uint
		if playerTwoFargo > playerOneFargo {
			playerTwoRaceTo = winsHigher
		} else {
			playerTwoRaceTo = winsLower
		}

		// TODO: SHOULD MAYBE TEST THAT BOTH DONT EQUAL

		if g.ScoreOne == int(playerOneRaceTo) {
			return 1
		} else if g.ScoreTwo == int(playerTwoRaceTo) {
			return 2
		}
	}

	return 0
}

// Save saves the game to the database.
func (g *Game) Save(completed bool) error {
	log.Printf("%+v", g)

	if g.ID != 0 {
		return g.db.Model(g).Updates(map[string]interface{}{
			"table_num": g.Table,
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

			"use_fargo_hot_handicap": g.UseFargoHotHandicap,
			"fargo_hot_handicap_id":  g.FargoHotHandicapID,
		}).Error
	}

	return g.db.Create(g).Error
}
