import { Injectable, OnDestroy } from '@angular/core';
import { Subscription, fromEvent, Observable, BehaviorSubject } from 'rxjs';
import { distinctUntilChanged, map } from 'rxjs/operators';

export interface WindowDimensions {
    width: number;
    height: number;
}

export interface WindowScroll {
    offsetY: number;
    offsetX: number;
}

@Injectable({
    providedIn: 'root'
})
export class WindowService implements OnDestroy {
    public width$: Observable<number>;
    public height$: Observable<number>;
    public scrollX$: Observable<number>;
    public scrollY$: Observable<number>;
    private _resizeSub: Subscription;
    private _scrollSub: Subscription;

    constructor() {
        const windowResize$ = new BehaviorSubject<WindowDimensions>(this._getWindowDimensions);
        const windowScroll$ = new BehaviorSubject<WindowScroll>(this._getWindowScroll);

        this.width$ = windowResize$.pipe(
            map(window => window.width),
            distinctUntilChanged(),
        );

        this.height$ = windowResize$.pipe(
            map(window => window.height),
            distinctUntilChanged(),
        );

        this.scrollX$ = windowScroll$.pipe(
            map(window => window.offsetX),
            distinctUntilChanged(),
        );

        this.scrollY$ = windowScroll$.pipe(
            map(window => window.offsetY),
            distinctUntilChanged(),
        );

        this._resizeSub = fromEvent(window, 'resize').pipe(
            distinctUntilChanged(),
            map(() => this._getWindowDimensions),
        ).subscribe(windowResize$);

        this._scrollSub = fromEvent(window, 'scroll').pipe(
            distinctUntilChanged(),
            map(() => this._getWindowScroll),
        ).subscribe(windowScroll$);
    }

    public ngOnDestroy(): void {
        if (this._resizeSub) {
            this._resizeSub.unsubscribe();
        }
        if (this._scrollSub) {
            this._scrollSub.unsubscribe();
        }
    }

    private get _getWindowScroll(): WindowScroll {
        return {
            offsetX: window.pageXOffset,
            offsetY: window.pageYOffset,
        };
    }

    private get _getWindowDimensions(): WindowDimensions {
        return {
            width: window.innerWidth,
            height: window.innerHeight,
        };
    }
}
