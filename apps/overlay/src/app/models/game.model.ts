import { IPlayer } from './player.model';
import { GameType } from './game-type.enum';
import { VsMode } from './vs-mode.enum';
import { IScore } from './score.model';

export interface IGame {
  type: GameType;
  vs_mode: VsMode;
  race_to: number;
  score: IScore;
  player_one: IPlayer;
  player_two: IPlayer;
}
