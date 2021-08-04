import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './components/app/app.component';
import { SideNavComponent } from './components/side-nav/side-nav.component';
import { ENV_CONFIG } from './models/environment-config.model';
import { environment } from '../environments/environment';

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
        AppRoutingModule,
    ],
    providers: [{
        provide: ENV_CONFIG,
        useValue: { environment },
    }],
    bootstrap: [AppComponent],
})
export class AppModule { }