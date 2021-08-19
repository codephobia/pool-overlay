import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../models/environment-config.model';
import { IGame } from '@pool-overlay/models';

@Injectable()
export class APIService {
    private _apiURL: string;
    private _apiVersion: string;

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private _http: HttpClient,
    ) {
        this._apiURL = config.apiURL;
        this._apiVersion = config.apiVersion;
    }

    getGame(): Observable<IGame> {
        const url = `${this._apiURL}/${this._apiVersion}/game`;
        return this._http.get<IGame>(url);
    }
}
