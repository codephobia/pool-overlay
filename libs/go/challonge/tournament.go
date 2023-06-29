package challonge

import (
	"encoding/json"
	"errors"
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
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	URL            string     `json:"url"`
	TournamentType string     `json:"tournament_type"`
	Rounds         int        `json:"rounds"`
	Matches        []*Match   `json:"matches"`
	CreatedAt      time.Time  `json:"created_at"`
	CompletedAt    *time.Time `json:"completed_at"`

	Participants []*Participant `json:"participants"`
}

var (
	// ErrNoIncompleteTournaments - No incomplete tournaments.
	ErrNoIncompleteTournaments = errors.New("no incomplete tournaments")
)

// Validate checks to make sure the tournament is in a valid format.
func (t *Tournament) Validate() error {
	err := t.verifyPlayerNames()
	if err != nil {
		return err
	}

	return nil
}

// GetNextMatch returns the first incomplete match in the tournament bracket.
func (t *Tournament) GetNextMatch() *Match {
	sortedMatches := t.getMatchesByPlayOrder()

	log.Printf("sorted matches: %+v", sortedMatches)

	for i := range sortedMatches {
		match := sortedMatches[i]

		log.Printf("match id: %d", match.ID)

		if match.State == "open" && match.UnderwayAt == nil {
			log.Printf("found next match: %d", match.ID)

			return match
		}
	}

	log.Print("didn't find a match")

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

// CompleteIfPossible attempts to complete the tournament.
func (t *Tournament) CompleteIfPossible(apiKey string) error {
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
	req.SetBasicAuth(apiKey, "")

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
	err = json.NewDecoder(resp.Body).Decode(&matchesResp)
	if err != nil {
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

// verifyPlayerNames makes sure all the player names in Challonge are in the correct format.
func (t *Tournament) verifyPlayerNames() error {
	playerNames := t.playerNames()

	// This regex assumes names do not contain '-' or '(' and the Fargo ID is numeric
	re := regexp.MustCompile(`^[^-]+ - [^(-]+ \(\d+\)$`)

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

func (t *Tournament) getMatchesByPlayOrder() []*Match {
	sortedMatches := make([]*Match, len(t.Matches))
	copy(sortedMatches, t.Matches)

	sort.Slice(sortedMatches, func(i, j int) bool {
		return sortedMatches[i].SuggestedPlayOrder < sortedMatches[j].SuggestedPlayOrder
	})

	return sortedMatches
}

// getMatchesByRound returns all the matches in a tournament, sorted in order of appearance on the bracket.
// func (t *Tournament) getMatchesByRound() []*Match {
// 	matchesByRound := make(map[int][]*Match)

// 	// Group matches by round
// 	for i := range t.Matches {
// 		match := t.Matches[i]
// 		matchesByRound[match.Round] = append(matchesByRound[match.Round], match)
// 	}

// 	// Sort matches within each round
// 	for _, matches := range matchesByRound {
// 		sort.Slice(matches, func(i, j int) bool {
// 			return matches[i].Identifier < matches[j].Identifier
// 		})
// 	}

// 	// Flatten matches into a single slice, sorted by round and position
// 	var sortedMatches []*Match
// 	for i := 1; i <= t.Rounds; i++ {
// 		sortedMatches = append(sortedMatches, matchesByRound[i]...)
// 	}

// 	return sortedMatches
// }

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
		if tournament.Tournament.CompletedAt == nil {
			tournaments = append(tournaments, &tournament.Tournament)
		}
	}

	// Return error if no incomplete tournaments
	if len(tournaments) == 0 {
		return nil, ErrNoIncompleteTournaments
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
