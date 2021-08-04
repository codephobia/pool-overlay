import { IPlayer } from './player.model';
import { GameType } from './game-type.enum';
import { VsMode } from './vs-mode.enum';

export interface IGame {
    type: GameType;
    vs_mode: VsMode;
    race_to: number;
    score_one: number;
    score_two: number;
    player_one: IPlayer;
    player_two: IPlayer;
}
