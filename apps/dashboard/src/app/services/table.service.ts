import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../models/environment-config.model';
import { ICount } from '../models/count.model';

@Injectable({ providedIn: 'root' })
export class TableService {
    private apiURL: string;
    private apiVersion: string;
    private endpoint = 'table';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
        this.apiVersion = config.environment.apiVersion;
    }

    public count(): Observable<ICount> {
        let url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/count`;
        return this.http.get<ICount>(url);
    }
}
