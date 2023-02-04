import { IPlayer } from './player.model';
import { GameType } from './game-type.enum';
import { VsMode } from './vs-mode.enum';
import { FargoHotHandicap } from './fargo-hot-handicap';

export interface IGame {
    table: number;
    type: GameType;
    vs_mode: VsMode;
    race_to: number;
    score_one: number;
    score_two: number;
    player_one?: IPlayer;
    player_two?: IPlayer;
    use_fargo_hot_handicap: boolean;
    fargo_hot_handicap?: FargoHotHandicap;

    created_at?: string;
    updated_at?: string;
    deleted_at?: string;
}
