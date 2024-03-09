import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';
import { IPlayer } from '@pool-overlay/models';
import { ICount } from '../../models/count.model';

export interface PlayerFindOptions {
    page: number;
    search?: string;
}

export interface PlayerCountOptions {
    search?: string;
}

@Injectable()
export class PlayersService {
    private apiURL: string;
    private apiVersion: string;
    private endpoint = 'players';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
        this.apiVersion = config.environment.apiVersion;
    }

    public find({ page, search }: PlayerFindOptions = { page: 1 }): Observable<IPlayer[]> {
        let url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}?page=${page}`;

        if (search) {
            url = url + `&search=${search}`;
        }

        return this.http.get<IPlayer[]>(url);
    }

    public findByID(id: number): Observable<IPlayer> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/${id}`;
        return this.http.get<IPlayer>(url);
    }

    public count({ search }: PlayerCountOptions = {}): Observable<ICount> {
        let url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/count`;

        if (search) {
            url = url + `?search=${search}`;
        }

        return this.http.get<ICount>(url);
    }

    public create(player: Omit<IPlayer, 'id'>) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}`;
        return this.http.post<IPlayer>(url, player);
    }

    public update({ id, ...player }: IPlayer) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/${id}`;
        return this.http.patch<IPlayer>(url, player);
    }

    public delete(playerId: number) {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/${playerId}`;
        return this.http.delete<void>(url);
    }
}
