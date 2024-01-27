package challonge

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MatchResp struct {
	Match Match `json:"match"`
}

// Match is a set to be played within a Tournament.
type Match struct {
	ID                 int        `json:"id"`
	TournamentID       int        `json:"tournament_id"`
	State              string     `json:"state"`
	Player1ID          *int       `json:"player1_id"`
	Player2ID          *int       `json:"player2_id"`
	ScoresCsv          string     `json:"scores_csv"`
	Round              int        `json:"round"`
	Identifier         string     `json:"identifier"`
	SuggestedPlayOrder int        `json:"suggested_play_order"`
	UnderwayAt         *time.Time `json:"underway_at"`
}

// SetInProgress sets the match to in progress on Challonge.
func (m *Match) SetInProgress(username string, apiKey string) error {
	matchURL := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/matches/%d.json", m.TournamentID, m.ID)
	req, err := http.NewRequest("PUT", matchURL, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, apiKey)

	underwayAt := time.Now()
	score := "0-0"

	params := url.Values{}
	params.Add("match[scores_csv]", score)
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	m.ScoresCsv = score
	m.UnderwayAt = &underwayAt

	return nil
}

// UpdateScore updates the score for a match.
func (m *Match) UpdateScore(username string, apiKey string, scoresCsv string) error {
	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/matches/%d.json", m.TournamentID, m.ID)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, apiKey)

	params := fmt.Sprintf("match[scores_csv]=%s", scoresCsv)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(strings.NewReader(params))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	m.ScoresCsv = scoresCsv

	return nil
}

// ReportWinner reports the winner of the match to challonge.
func (m *Match) ReportWinner(username string, apiKey string, playerID int, scoresCsv string) error {
	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/matches/%d.json", m.TournamentID, m.ID)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, apiKey)

	params := fmt.Sprintf("match[winner_id]=%d&match[scores_csv]=%s", playerID, scoresCsv)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(strings.NewReader(params))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// refresh match
	return m.Refresh(username, apiKey)
}

// IsOnASide returns if the round is on the A side or not.
func (m *Match) IsOnASide() bool {
	return m.Round > 0
}

// Refreshes a match with new data from Challonge.
// Should only be used when a match was completed.
func (m *Match) Refresh(username string, apiKey string) error {
	log.Printf("refreshing match: %d", m.ID)

	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/matches/%d.json", m.TournamentID, m.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var matchResp MatchResp
	if err := json.NewDecoder(resp.Body).Decode(&matchResp); err != nil {
		return err
	}

	match := matchResp.Match
	m.State = match.State
	m.Player1ID = match.Player1ID
	m.Player2ID = match.Player2ID
	m.ScoresCsv = match.ScoresCsv
	m.UnderwayAt = match.UnderwayAt

	return nil
}
