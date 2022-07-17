import { HttpClient } from '@angular/common/http';
import { Inject, Injectable } from '@angular/core';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';

@Injectable()
export class OverlayStateService {
    private apiURL: string;
    private apiVersion: string;
    private endpoint = 'overlay';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
        this.apiVersion = config.environment.apiVersion;
    }

    public toggle() {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/toggle`;
        return this.http.get<{ hidden: boolean }>(url);
    }

    public toggleFlags() {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/toggle/flags`;
        return this.http.get<{ showFlags: boolean }>(url);
    }

    public toggleFargo() {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/toggle/fargo`;
        return this.http.get<{ showFargo: boolean }>(url);
    }

    public toggleScore() {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/toggle/score`;
        return this.http.get<{ showScore: boolean }>(url);
    }
}
