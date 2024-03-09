import { NgModule, isDevMode } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { StoreRouterConnectingModule, routerReducer } from '@ngrx/router-store';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './components/app/app.component';
import { SideNavComponent } from './components/side-nav/side-nav.component';
import { ENV_CONFIG } from './models/environment-config.model';
import { environment } from '../environments/environment';
import * as fromTables from './core/tables';

const COMPONENTS = [
    AppComponent,
    SideNavComponent,
];

@NgModule({
    declarations: [
        ...COMPONENTS
    ],
    imports: [
        BrowserModule,
        BrowserAnimationsModule,
        HttpClientModule,
        FontAwesomeModule,
        AppRoutingModule,
        StoreModule.forRoot({
            [fromTables.stateKey]: fromTables.reducer,
            router: routerReducer,
        }),
        EffectsModule.forRoot([
            fromTables.TablesEffects,
        ]),
        StoreRouterConnectingModule.forRoot(),
        StoreDevtoolsModule.instrument({
            maxAge: 25,
            logOnly: !isDevMode(),
            autoPause: true,
            trace: false,
        }),
    ],
    providers: [{
        provide: ENV_CONFIG,
        useValue: { environment },
    }],
    bootstrap: [AppComponent],
})
export class AppModule { }
