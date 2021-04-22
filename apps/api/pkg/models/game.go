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
	ID       uint       `json:"id,omitempty" gorm:"primarykey"`
	Type     GameType   `json:"type"`
	VsMode   GameVsMode `json:"vs_mode"`
	RaceTo   int        `json:"race_to"`
	ScoreOne int        `json:"score_one"`
	ScoreTwo int        `json:"score_two"`

	PlayerOneID uint `json:"player_one_id,omitempty"`
	PlayerTwoID uint `json:"player_two_id,omitempty"`
	TeamOneID   uint `json:"team_one_id,omitempty"`
	TeamTwoID   uint `json:"team_two_id,omitempty"`

	PlayerOne *Player `json:"player_one,omitempty" gorm:"foreignKey:id"`
	PlayerTwo *Player `json:"player_two,omitempty" gorm:"foreignKey:id"`
	TeamOne   *Team   `json:"team_one,omitempty" gorm:"foreignKey:id"`
	TeamTwo   *Team   `json:"team_two,omitempty" gorm:"foreignKey:id"`

	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	mutex sync.Mutex `json:"-"`
}

// NewGame returns a new Game with default settings.
func NewGame() *Game {
	return &Game{
		Type:   EightBall,
		RaceTo: 5,
		VsMode: OneVsOne,
	}
}

// SetType sets the Type of the current game.
func (g *Game) SetType(t GameType) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.Type = t
}

// SetVsMode changes the current GameVsMode of the game.
func (g *Game) SetVsMode(mode GameVsMode) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.VsMode = mode

	// TODO: MIGHT WANT TO UNSET TEAMS / PLAYERS BASED ON MODE CHANGE
}

// IncrementRaceTo increases the RaceTo of the game by one.
func (g *Game) IncrementRaceTo() {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.RaceTo++
}

// DecrementRaceTo decreases the RaceTo of the game by one. Minimum value will
// always be 1.
func (g *Game) DecrementRaceTo() {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	newRaceTo := g.RaceTo - 1

	if newRaceTo > 0 {
		g.RaceTo = newRaceTo
	}
}

// IncrementScore increases the score for the specified player by one.
func (g *Game) IncrementScore(playerNum int) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		newScore := g.ScoreOne + 1
		if newScore <= g.RaceTo {
			g.ScoreOne = newScore
		}
	case 2:
		newScore := g.ScoreTwo + 1
		if newScore <= g.RaceTo {
			g.ScoreTwo = newScore
		}
	default:
		return ErrInvalidPlayerNumber
	}

	return nil
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

	return nil
}

// ResetScore resets the current game score.
func (g *Game) ResetScore() {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.ScoreOne = 0
	g.ScoreTwo = 0
}

// SetPlayer sets the player number to the specified player.
func (g *Game) SetPlayer(playerNum int, player *Player) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch playerNum {
	case 1:
		g.PlayerOne = player
	case 2:
		g.PlayerTwo = player
	default:
		return ErrInvalidPlayerNumber
	}

	// Unset teams when adding a player.
	g.TeamOne = nil
	g.TeamTwo = nil

	return nil
}

// SetTeam sets the team number to the specified team.
func (g *Game) SetTeam(teamNum int, team *Team) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch teamNum {
	case 1:
		g.TeamOne = team
	case 2:
		g.TeamTwo = team
	default:
		return ErrInvalidTeamNumber
	}

	// Unset players when adding a team.
	g.PlayerOne = nil
	g.PlayerTwo = nil

	return nil
}
