import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subject } from 'rxjs';

import { IPlayer } from '@pool-overlay/models';
import { PlayerFormStore } from '../../services/player-form.store';
import { map, takeUntil } from 'rxjs/operators';

@Component({
    selector: 'pool-overlay-player-edit-page',
    templateUrl: './player-edit-page.component.html',
    providers: [PlayerFormStore]
})
export class PlayerEditPageComponent implements OnInit, OnDestroy {
    public readonly vm$ = this.playerFormStore.vm$;
    private _destroy$ = new Subject<void>();

    constructor(
        private readonly route: ActivatedRoute,
        private readonly playerFormStore: PlayerFormStore,
    ) { }

    public ngOnInit(): void {
        this.route.params.pipe(
            map(params => params.playerId),
            takeUntil(this._destroy$),
        ).subscribe(playerId => {
            this.playerFormStore.getPlayerAndFlags(playerId);
        });
    }

    public onSubmit(event: { player: IPlayer | Omit<IPlayer, 'id'> }): void {
        this.playerFormStore.updatePlayer(event.player as IPlayer);
    }

    public ngOnDestroy(): void {
        this._destroy$.next();
    }
}
