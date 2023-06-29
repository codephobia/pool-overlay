# Challonge Integration

Uses the Challonge API to load tournaments, update matches, and automatically update the overlays.

## Steps

1. Select Tournament
    - Get list of incomplete tournaments via Challonge.
2. Setup Tournament
    - Validates all player names are in the correct format. "FirstName LastName - FargoRating (FargoID)"
    - Validates all players in the tournament exist in the database.
    - Checks if tournament is multi-stage (v2 api), and double elimination.
    - Set Game type for each stage, and race for each stage / side.
    - Set number of max tables to use.
    - Toggle handicapped tournament.
3. Start Tournament
    - Loads the selected tournament.
    - Turns on tournament mode flag to lockout changes to overlay (nice to have).
    - Changes all tables to the settings for tournament. (Game Type, Race (A/B Side), Unset Players)
    - Loads first matches in to tables.
      - While tables are empty.
        - Get next match
        - Set empty table to game data. (players, resets score)
        - Mark match as in progress on Challonge (v2 api).
4. Scoring
    - When updating scores on the tablets, update the score on Challonge as well.
5. Match is completed
    - Player confirms match is completed on tablet.
    - Save game to database.
    - Mark match as completed with winner and score.
    - Check if another match still needs to be played.
    - Check if we should be staying at the same amount of tables.
      - If yes, get next match and load to same table, otherwise reduce the number of tables.
6. Skipping a match
    - Admin can skip a match on a table.
    - Check if another match still needs to be played.
    - Get next match and load to same table.
7. Allow chop in finals.
    - Admin has a way to submit score of 0-0 in case players chop. This allows the tournament to complete, but doesn't get entered as wins or losses into Fargo.
8. Complete tournament
    - When final game is saved, it should attempt to complete the tournament.
    - Unloads the tournament.
