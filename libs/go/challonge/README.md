# Challonge Integration

Uses the Challonge API to load tournaments, update matches, and automatically update the overlays.

## Steps

1. Select Tournament
    - Get list of incomplete tournaments via Challonge.
2. Setup Tournament
    x Validates all player names are in the correct format. "FirstName LastName - FargoRating (FargoID)"
    x Validates all players in the tournament exist in the database.
    - Checks if tournament is multi-stage (v2 api), and double elimination.
    x Set Game type for each stage, and race for each stage / side.
    x Set number of max tables to use.
    x Toggle handicapped tournament.
3. Start Tournament
    x Loads the selected tournament.
    - Turns on tournament mode flag to lockout changes to overlay (nice to have).
    x Changes all tables to the settings for tournament. (Game Type, Race (A/B Side), Unset Players)
    x Loads first matches in to tables.
      x While tables are empty.
        x Get next match
        x Set empty table to game data. (players, resets score)
        - Mark match as in progress on Challonge (v2 api).
        x Set score to 0-0 on Challonge.
4. Scoring
    x When updating scores on the tablets, update the score on Challonge as well.
5. Match is completed
    x Player confirms match is completed on tablet.
    x Save game to database.
    x Mark match as completed with winner and score on Challonge.
    x Check if we should be staying at the same amount of tables. (temp solution)
    x Check if another match still needs to be played.
      x If yes, get next match and load to same table, otherwise reduce the number of tables.
6. Skipping a match
    - Admin can skip a match on a table.
    - Check if another match still needs to be played.
    - Get next match and load to same table.
7. Allow chop in finals.
    - Admin has a way to submit score of 0-0 in case players chop. This allows the tournament to complete, but doesn't get entered as wins or losses into Fargo.
8. Complete tournament
    x When final game is saved, it should attempt to complete the tournament.
    - Unloads the tournament.

Other nice to haves:
    - Admin can swap matches between two tables.
