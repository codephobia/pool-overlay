import { IFlag } from './flag.model';

export interface IPlayer {
    id: number;
    name: string;
    flag_id: number;
    flag?: IFlag;
    fargo_observable_id: number;
    fargo_id: number;
    fargo_rating: number;
}
