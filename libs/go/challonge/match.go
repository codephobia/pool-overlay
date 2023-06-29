package challonge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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
	score := "0,0"

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
func (m *Match) UpdateScore(apiKey string, scoresCsv string) error {
	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/matches/%d.json", m.TournamentID, m.ID)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(apiKey, "")

	params := fmt.Sprintf("match[scores_csv]=%s", scoresCsv)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = ioutil.NopCloser(strings.NewReader(params))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
