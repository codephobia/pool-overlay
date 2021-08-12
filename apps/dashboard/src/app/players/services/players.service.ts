import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';
import { IPlayer } from '@pool-overlay/models';
import { ICount } from '../../models/count.model';

@Injectable()
export class PlayersService {
    private apiURL: string;
    private endpoint = 'players';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
    }

    public find(page = 1): Observable<IPlayer[]> {
        const url = `${this.apiURL}/${this.endpoint}?page=${page}`;
        return this.http.get<IPlayer[]>(url);
    }

    public findByID(id: number): Observable<IPlayer> {
        const url = `${this.apiURL}/${this.endpoint}/${id}`;
        return this.http.get<IPlayer>(url);
    }

    public count(): Observable<ICount> {
        const url = `${this.apiURL}/${this.endpoint}/count`;
        return this.http.get<ICount>(url);
    }

    public create(player: Omit<IPlayer, 'id'>) {
        const url = `${this.apiURL}/${this.endpoint}`;
        return this.http.post<IPlayer>(url, player);
    }

    public update(player: IPlayer) {
        const url = `${this.apiURL}/${this.endpoint}`;
        return this.http.put<IPlayer>(url, player);
    }
}
