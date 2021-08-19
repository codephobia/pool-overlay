import { InjectionToken } from '@angular/core';

export interface EnvironmentConfig {
    production: boolean;
    apiURL: string;
    apiVersion: string;
}

export const ENV_CONFIG = new InjectionToken<EnvironmentConfig>('EnvironmentConfig');
