import { AfterViewInit, ChangeDetectorRef, Directive, ElementRef, HostBinding, OnDestroy } from '@angular/core';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';

import { WindowService } from '../../services/window.service';
import { ResizeService } from '../../services/resize.service';
import { DrawingService } from '../services/drawing.service';

@Directive({
    selector: '[poolOverlayDrawing]',
})
export class DrawingDirective implements AfterViewInit, OnDestroy {
    @HostBinding('attr.width')
    public width = 300;

    @HostBinding('attr.height')
    public height = 168.75;

    private canvas: ElementRef<HTMLCanvasElement> | null = null;
    private _destroy$ = new Subject<void>();

    constructor(
        elRef: ElementRef,
        private cd: ChangeDetectorRef,
        private windowService: WindowService,
        private resizeService: ResizeService,
        private drawingService: DrawingService,
    ) {
        this.canvas = elRef;

        this.canvas.nativeElement.onmousedown = this.drawingService.mousedown.bind(this.drawingService);
        this.canvas.nativeElement.onmouseup = this.drawingService.mouseup.bind(this.drawingService);
        this.canvas.nativeElement.onmousemove = this.drawingService.mousemove.bind(this.drawingService);
    }

    public ngAfterViewInit(): void {
        this._setCanvasSize();
        this.drawingService.init(this.canvas);

        // watch window for resize
        this.windowService.width$.pipe(
            takeUntil(this._destroy$),
        ).subscribe(() => {
            this._setCanvasSize();
        });

        // watch for parent resize
        const parentEl = this.canvas!.nativeElement.parentElement;
        if (parentEl) {
            this.resizeService.resizeObservableEl(parentEl).pipe(
                takeUntil(this._destroy$),
            ).subscribe(() => {
                this._setCanvasSize();
            });
        }

        this.cd.detectChanges();
    }

    public ngOnDestroy(): void {
        this._destroy$.next();
    }

    private _setCanvasSize(): void {
        this.width = this.canvas!.nativeElement.parentElement?.getBoundingClientRect().width ?? this.width;
        this.height = this.canvas!.nativeElement.parentElement?.getBoundingClientRect().height ?? this.height;
    }
}
