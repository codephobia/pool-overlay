import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { environment } from '../../environments/environment';
import { IGame } from '../models/game.model';

@Injectable({ providedIn: 'root' })
export class APIService {
    private _apiURL = environment.apiURL;
    private _apiVersion = environment.apiVersion;

    constructor(private _http: HttpClient) { }

    getGame(): Observable<IGame> {
        const url = `${this._apiURL}/${this._apiVersion}/game`;
        return this._http.get<IGame>(url);
    }
}
