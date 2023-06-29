import { GameType } from './game-type.enum';

export interface TournamentSettings {
    game_type: GameType;
    show_overlay: boolean;
    show_flags: boolean;
    show_fargo: boolean;
    show_score: boolean;
    is_handicapped: boolean;
    a_side_race_to: number;
    b_side_race_to?: number;
}
