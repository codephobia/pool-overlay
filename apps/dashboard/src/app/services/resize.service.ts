import { ElementRef, Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class ResizeService {
    public resizeObservableElRef<T extends HTMLElement>(elRef: ElementRef<T>): Observable<ResizeObserverEntry[]> {
        return this.resizeObservableEl<T>(elRef.nativeElement);
    }

    public resizeObservableEl<T extends HTMLElement>(el: T): Observable<ResizeObserverEntry[]> {
        return new Observable(subscriber => {
            const ro = new ResizeObserver(entries => {
                subscriber.next(entries);
            });

            // Observe one or multiple elements
            ro.observe(el);

            return function unsubscribe() {
                ro.unobserve(el);
            }
        });
    }
}
