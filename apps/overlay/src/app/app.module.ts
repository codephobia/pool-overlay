import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { ScoreboardModule } from '@pool-overlay/scoreboard';
import { AppComponent } from './components/app/app.component';
import { environment } from '../environments/environment';

const COMPONENTS = [
    AppComponent,
];

@NgModule({
    declarations: COMPONENTS,
    imports: [
        BrowserModule,
        HttpClientModule,
        BrowserAnimationsModule,
        ScoreboardModule.withConfig({ environment }),
    ],
    providers: [],
    bootstrap: [AppComponent],
})
export class AppModule { }
