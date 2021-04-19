import { Component, OnInit } from '@angular/core';

import { ScoreboardStore } from './scoreboard.store';

@Component({
  selector: 'pool-overlay-scoreboard',
  templateUrl: './scoreboard.component.html',
  styleUrls: ['./scoreboard.component.scss'],
  providers: [ScoreboardStore],
})
export class ScoreboardComponent implements OnInit {
  public readonly vm$ = this._scoreboardStore.vm$;

  constructor(private _scoreboardStore: ScoreboardStore) {
    this._scoreboardStore.setState({
      game: null,
      hidden: false,
    });
  }

  public ngOnInit(): void {
    this._scoreboardStore.getGame();
  }
}
