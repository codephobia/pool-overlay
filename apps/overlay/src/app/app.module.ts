import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { ScoreboardComponent } from './components/scoreboard/scoreboard.component';
import { CharacterRotaterComponent } from './components/character-rotater/character-rotater.component';
import { FlagComponent } from './components/flag/flag.component';

const COMPONENTS = [
    AppComponent,
    ScoreboardComponent,
    CharacterRotaterComponent,
    FlagComponent,
];

@NgModule({
    declarations: COMPONENTS,
    imports: [BrowserModule, HttpClientModule, BrowserAnimationsModule],
    providers: [],
    bootstrap: [AppComponent],
})
export class AppModule { }
