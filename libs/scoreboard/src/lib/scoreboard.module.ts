import { ModuleWithProviders, NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { EnvironmentConfig, ENV_CONFIG } from './models/environment-config.model';
import { APIService } from './services/api.service';
import { ScoreboardComponent } from './components/scoreboard/scoreboard.component';
import { CharacterRotaterComponent } from './components/character-rotater/character-rotater.component';
import { FlagComponent } from './components/flag/flag.component';

const COMPONENTS = [
    ScoreboardComponent,
    CharacterRotaterComponent,
    FlagComponent,
];

const SERVICES = [
    APIService,
];

export interface ScoreboardModuleConfig {
    environment: EnvironmentConfig;
}

@NgModule({
    imports: [CommonModule],
    declarations: [
        ...COMPONENTS,
    ],
    providers: [
        ...SERVICES,
    ],
    exports: [
        ...COMPONENTS,
    ],
})
export class ScoreboardModule {
    static withConfig(config: ScoreboardModuleConfig): ModuleWithProviders<ScoreboardModule> {
        return {
            ngModule: ScoreboardModule,
            providers: [
                {
                    provide: ENV_CONFIG,
                    useValue: config.environment,
                }
            ],
        };
    }
}
