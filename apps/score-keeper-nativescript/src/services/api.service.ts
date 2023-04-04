import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { Direction, IGame } from '@pool-overlay/models';

@Injectable({
    providedIn: 'root',
})
export class APIService {
    private _apiURL: string = 'http://192.168.0.26:1268';
    private _apiVersion: string = 'latest';

    constructor(
        private _http: HttpClient,
    ) { }

    public getGame(table: number): Observable<IGame> {
        const url = `${this._apiURL}/${this._apiVersion}/table/${table}/game`;
        return this._http.get<IGame>(url);
    }

    public updateScore(table: number, playerNum: number, direction: Direction) {
        const url = `${this._apiURL}/${this._apiVersion}/table/${table}/game/score`;
        return this._http.patch<{ scoreOne: number, scoreTwo: number }>(url, { playerNum, direction });
    }
}
