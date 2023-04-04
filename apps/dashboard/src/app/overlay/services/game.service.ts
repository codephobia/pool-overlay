import { HttpClient } from '@angular/common/http';
import { Inject, Injectable } from '@angular/core';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';
import { GameType, IGame, Direction } from '@pool-overlay/models';

@Injectable()
export class GameService {
    private apiURL: string;
    private apiVersion: string;
    private endpoint = 'game';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
        this.apiVersion = config.environment.apiVersion;
    }

    public getGame(table: number) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}`;
        return this.http.get<IGame>(url);
    }

    public unsetPlayer(table: number, playerNum: number) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}/players`;
        return this.http.delete<IGame>(url, {
            body: { playerNum },
        });
    }

    public setPlayer(table: number, playerNum: number, playerID: number) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}/players`;
        return this.http.patch<IGame>(url, { playerNum, playerID });
    }

    public updateRaceTo(table: number, direction: Direction) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}/race-to`;
        return this.http.patch<{ raceTo: number, useFargoHotHandicap: boolean }>(url, { direction });
    }

    public resetScore(table: number) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}/score`;
        return this.http.delete<{ scoreOne: number, scoreTwo: number }>(url);
    }

    public updateScore(table: number, playerNum: number, direction: Direction) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}/score`;
        return this.http.patch<{ scoreOne: number, scoreTwo: number }>(url, { playerNum, direction });
    }

    public setGameType(table: number, type: GameType) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}/type`;
        return this.http.patch<{ type: number }>(url, { type });
    }

    public setFargoHotHandicap(table: number, useFargoHotHandicap: boolean) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}/fargo-hot-handicap`;
        return this.http.patch<{ useFargoHotHandicap: boolean }>(url, { useFargoHotHandicap });
    }

    public save(table: number) {
        const url = `${this.apiURL}/${this.apiVersion}/table/${table}/${this.endpoint}`;
        return this.http.post<void>(url, {});
    }
}
