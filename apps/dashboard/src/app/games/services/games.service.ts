import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';
import { IGame } from '@pool-overlay/models';
import { ICount } from '../../models/count.model';

@Injectable()
export class GamesService {
    private apiURL: string;
    private apiVersion: string;
    private endpoint = 'games';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
        this.apiVersion = config.environment.apiVersion;
    }

    public find(page = 1): Observable<Partial<IGame>[]> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}?page=${page}`;
        return this.http.get<Partial<IGame>[]>(url);
    }

    public count(): Observable<ICount> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/count`;
        return this.http.get<ICount>(url);
    }
}
