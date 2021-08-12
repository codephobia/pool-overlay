import { Inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { EnvironmentConfig, ENV_CONFIG } from '../../models/environment-config.model';
import { IFlag } from '@pool-overlay/models';

@Injectable()
export class FlagsService {
    private apiURL: string;
    private endpoint = 'flags';

    constructor(
        @Inject(ENV_CONFIG) config: EnvironmentConfig,
        private http: HttpClient,
    ) {
        this.apiURL = config.environment.apiURL;
    }

    public find(): Observable<IFlag[]> {
        const url = `${this.apiURL}/${this.endpoint}`;
        return this.http.get<IFlag[]>(url);
    }
}
