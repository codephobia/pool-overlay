package game

import (
	"sync"

	"github.com/codephobia/pool-overlay/apps/api/pkg/models"
)

// Game is the current state of the game being played.
type Game struct {
	Type      Type           `json:"type"`
	VsMode    VsMode         `json:"vs_mode"`
	RaceTo    int            `json:"race_to"`
	Score     Score          `json:"score"`
	PlayerOne *models.Player `json:"player_one,omitempty"`
	PlayerTwo *models.Player `json:"player_two,omitempty"`
	TeamOne   *models.Team   `json:"team_one,omitempty"`
	TeamTwo   *models.Team   `json:"team_two,omitempty"`

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
func (g *Game) SetType(t Type) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.Type = t
}

// SetVsMode changes the current VsMode of the game.
func (g *Game) SetVsMode(mode VsMode) {
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
		newScore := g.Score.One + 1
		if newScore <= g.RaceTo {
			g.Score.One = newScore
		}
	case 2:
		newScore := g.Score.Two + 1
		if newScore <= g.RaceTo {
			g.Score.Two = newScore
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
		newScore := g.Score.One - 1
		if newScore >= 0 {
			g.Score.One = newScore
		}
	case 2:
		newScore := g.Score.Two - 1
		if newScore >= 0 {
			g.Score.Two = newScore
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

	g.Score.One = 0
	g.Score.Two = 0
}

// SetPlayer sets the player number to the specified player.
func (g *Game) SetPlayer(playerNum int, player *models.Player) error {
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
func (g *Game) SetTeam(teamNum int, team *models.Team) error {
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
