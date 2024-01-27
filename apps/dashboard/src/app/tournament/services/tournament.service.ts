import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';
import { Tournament, TournamentSettings } from '@pool-overlay/models';

@Injectable()
export class TournamentsService {
    private apiURL: string;
    private apiVersion: string;
    private endpoint = 'tournament';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
        this.apiVersion = config.environment.apiVersion;
    }

    public getCurrent(): Observable<Tournament> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}`;
        return this.http.get<Tournament>(url);
    }

    public getList(): Observable<Tournament[]> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/list`;
        return this.http.get<Tournament[]>(url);
    }

    public getById(tournamentId: number): Observable<Tournament> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/${tournamentId}`;
        return this.http.get<Tournament>(url);
    }

    public load(id: number, settings: TournamentSettings): Observable<void> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/load`;
        return this.http.post<void>(url, { id, settings });
    }

    public unload(): Observable<void> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/unload`;
        return this.http.post<void>(url, {});
    }
}
