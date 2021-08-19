import { HttpClient } from '@angular/common/http';
import { Inject, Injectable } from '@angular/core';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';
import { GameType, IGame } from '@pool-overlay/models';
import { Direction } from '../models/direction.model';

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

    public getGame() {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}`;
        return this.http.get<IGame>(url);
    }

    public unsetPlayer(playerNum: number) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/players`;
        return this.http.delete<IGame>(url, {
            body: { playerNum },
        });
    }

    public setPlayer(playerNum: number, playerID: number) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/players`;
        return this.http.patch<IGame>(url, { playerNum, playerID });
    }

    public updateRaceTo(direction: Direction) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/race-to`;
        return this.http.patch<{ raceTo: number }>(url, { direction });
    }

    public resetScore() {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/score`;
        return this.http.delete<{ scoreOne: number, scoreTwo: number }>(url);
    }

    public updateScore(playerNum: number, direction: Direction) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/score`;
        return this.http.patch<{ scoreOne: number, scoreTwo: number }>(url, { playerNum, direction });
    }

    public setGameType(type: GameType) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/type`;
        return this.http.patch<{ type: number }>(url, { type });
    }

    public toggleOverlay() {
        const url = `${this.apiURL}/${this.apiVersion}/overlay/toggle`;
        return this.http.get<{ hidden: boolean }>(url);
    }
}
