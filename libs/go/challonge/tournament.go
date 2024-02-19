package challonge

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"time"
)

type TournamentsResp []struct {
	Tournament Tournament `json:"tournament"`
}

type TournamentResp struct {
	Tournament Tournament `json:"tournament"`
}

type MatchesResp []struct {
	Match Match `json:"match"`
}

type ParticipantsResp []struct {
	Participant Participant `json:"participant"`
}

// Tournament is a tournament bracket on the Challonge account.
type Tournament struct {
	ID                int        `json:"id"`
	Name              string     `json:"name"`
	URL               string     `json:"url"`
	TournamentType    string     `json:"tournament_type"`
	ParticipantsCount int        `json:"participants_count"`
	Rounds            int        `json:"rounds"`
	Matches           []*Match   `json:"matches"`
	CreatedAt         time.Time  `json:"created_at"`
	CompletedAt       *time.Time `json:"completed_at"`

	Participants []*Participant `json:"participants"`
}

// Validate checks to make sure the tournament is in a valid format.
func (t *Tournament) Validate() error {
	err := t.verifyPlayerNames()
	if err != nil {
		log.Printf("Error: %s", err)
	}

	return nil
}

// GetNextMatch returns the first incomplete match in the tournament bracket.
func (t *Tournament) GetNextMatch() *Match {
	sortedMatches := t.getMatchesByPlayOrder()

	for i := range sortedMatches {
		match := sortedMatches[i]
		if match.State == "open" && match.UnderwayAt == nil && match.Player1ID != nil && match.Player2ID != nil {
			log.Printf("found next match: %d", match.ID)

			return match
		}
	}

	return nil
}

// HasMoreMatches returns if there are more matches that need to be completed in the tournament.
func (t *Tournament) HasMoreMatches() bool {
	for i := range t.Matches {
		if t.Matches[i].State != "complete" {
			return true
		}
	}
	return false
}

// CountParallelMatches returns a count of the maximum number of matches that could be played
// in parallel.
func (t *Tournament) CountParallelMatches() int {
	sortedMatches := t.getMatchesByPlayOrder()
	activePlayers := make(map[int]bool)
	parallelCount := 0

	for i := range sortedMatches {
		match := sortedMatches[i]
		if match.State != "complete" {
			// Check if players are already in a match
			if _, exists := activePlayers[*match.Player1ID]; !exists {
				if _, exists := activePlayers[*match.Player2ID]; !exists {
					// Mark players as active
					activePlayers[*match.Player1ID] = true
					activePlayers[*match.Player2ID] = true
					parallelCount++
				}
			}
		}
	}
	return parallelCount
}

// CompleteIfPossible attempts to complete the tournament.
func (t *Tournament) CompleteIfPossible(username, apiKey string) error {
	for i := range t.Matches {
		match := t.Matches[i]
		if match.State != "complete" {
			return fmt.Errorf("cannot complete tournament, there are still incomplete matches")
		}
	}

	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/finalize.json", t.ID)
	req, err := http.NewRequest("POST", url, nil)
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

	return nil
}

// getMatches loads all the matches for the tournament.
func (t *Tournament) getMatches(username, apiKey string) error {
	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/matches.json", t.ID)
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

	var matchesResp MatchesResp
	if err := json.NewDecoder(resp.Body).Decode(&matchesResp); err != nil {
		return err
	}

	var matches []*Match
	for i := range matchesResp {
		match := matchesResp[i]
		matches = append(matches, &match.Match)
	}

	t.Matches = matches

	return nil
}

// updateMatches updates all the matches for the tournament.
func (t *Tournament) updateMatches(username, apiKey string) error {
	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/matches.json", t.ID)
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

	var matchesResp MatchesResp
	if err := json.NewDecoder(resp.Body).Decode(&matchesResp); err != nil {
		return err
	}

	var matches []*Match
	for i := range matchesResp {
		match := matchesResp[i]
		matches = append(matches, &match.Match)
	}

	for _, match := range t.Matches {
		for _, updatedMatch := range matches {
			if match.ID == updatedMatch.ID {
				match.State = updatedMatch.State
				match.Player1ID = updatedMatch.Player1ID
				match.Player2ID = updatedMatch.Player2ID
				match.ScoresCsv = updatedMatch.ScoresCsv
				break
			}
		}
	}

	return nil
}

// verifyPlayerNames makes sure all the player names in Challonge are in the correct format.
func (t *Tournament) verifyPlayerNames() error {
	playerNames := t.playerNames()

	// This regex assumes names do not contain '-' or '(' and the Fargo ID is numeric
	// re := regexp.MustCompile(`^[^-]+ - [^(-]+ \(\d+\)$`)
	re := regexp.MustCompile(`^(.+) - (\d+) \((\d+)\)$`)

	var invalidNames []string
	for _, name := range playerNames {
		if !re.MatchString(name) {
			invalidNames = append(invalidNames, name)
		}
	}

	if len(invalidNames) > 0 {
		return fmt.Errorf("the following player names are not in the expected format: %v", invalidNames)
	}

	return nil
}

// PlayerNames returns all the player names of participants in the tournament.
func (t *Tournament) playerNames() []string {
	var playerNames []string
	for _, participant := range t.Participants {
		if participant.Name != "" {
			playerNames = append(playerNames, participant.Name)
		}
	}
	return playerNames
}

func (t *Tournament) getParticipants(username, apiKey string) error {
	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d/participants.json", t.ID)
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

	var participantsResp ParticipantsResp
	err = json.NewDecoder(resp.Body).Decode(&participantsResp)
	if err != nil {
		return err
	}

	var participants []*Participant
	for _, pr := range participantsResp {
		participant := pr.Participant
		participants = append(participants, &participant)
	}

	t.Participants = participants

	return nil
}

// Returns if the match is the last possible match in the tournament
func (t *Tournament) matchIsDoubleDip(match *Match) bool {
	if t.TournamentType == "double elimination" && match.SuggestedPlayOrder == t.getTotalMatchCount() {
		return true
	}

	return false
}

// Returns the total number of matches possible, including the double dip
func (t *Tournament) getTotalMatchCount() int {
	// TODO: handle single elimination here
	return ((t.ParticipantsCount - 1) * 2) + 1
}

func (t *Tournament) getMatchesByPlayOrder() []*Match {
	sortedMatches := make([]*Match, len(t.Matches))
	copy(sortedMatches, t.Matches)

	sort.Slice(sortedMatches, func(i, j int) bool {
		return sortedMatches[i].SuggestedPlayOrder < sortedMatches[j].SuggestedPlayOrder
	})

	return sortedMatches
}

// getLatestTournaments returns the latest tournaments for the specified Challonge account.
func getLatestTournaments(username, apiKey string) ([]*Tournament, error) {
	url := "https://api.challonge.com/v1/tournaments.json?state=all&created_after=2022-01-01"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tournamentsResp TournamentsResp
	err = json.NewDecoder(resp.Body).Decode(&tournamentsResp)
	if err != nil {
		return nil, err
	}

	var tournaments []*Tournament

	// Filter off completed tournaments
	for _, tournament := range tournamentsResp {
		currentTournament := tournament.Tournament

		if currentTournament.CompletedAt == nil {
			tournaments = append(tournaments, &currentTournament)
		}
	}

	// Sort tournaments by created date (most recent first)
	sort.Slice(tournaments, func(i, j int) bool {
		return tournaments[i].CreatedAt.After(tournaments[j].CreatedAt)
	})

	return tournaments, nil
}

// getTournamentByID returns a tournament by tournament id.
func getTournamentByID(tournamentID int, username, apiKey string) (*Tournament, error) {
	url := fmt.Sprintf("https://api.challonge.com/v1/tournaments/%d.json", tournamentID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tournamentResp TournamentResp
	err = json.NewDecoder(resp.Body).Decode(&tournamentResp)
	if err != nil {
		return nil, err
	}

	tournament := tournamentResp.Tournament

	// Retrieve the matches for the tournament
	err = tournament.getMatches(username, apiKey)
	if err != nil {
		return nil, err
	}

	err = tournament.getParticipants(username, apiKey)
	if err != nil {
		return nil, err
	}

	return &tournament, nil
}
