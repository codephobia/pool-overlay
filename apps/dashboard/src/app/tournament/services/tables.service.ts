import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';

@Injectable()
export class TablesService {
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

    public swap(tableOne: number, tableTwo: number): Observable<void> {
        const url = `${this.apiURL}/${this.apiVersion}/${this.endpoint}/${tableOne}/swap/${tableTwo}`;
        return this.http.get<void>(url);
    }
}
