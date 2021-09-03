import { AfterViewInit, ChangeDetectorRef, Directive, ElementRef, HostBinding } from '@angular/core';

import { DrawingService } from '../services/drawing.service';

@Directive({
    selector: '[poolOverlayDrawing]',
})
export class DrawingDirective implements AfterViewInit {
    @HostBinding('attr.width')
    public width = 300;

    @HostBinding('attr.height')
    public height = 168.75;

    private canvas: ElementRef<HTMLCanvasElement> | null = null;

    constructor(
        elRef: ElementRef,
        private cd: ChangeDetectorRef,
        private drawingService: DrawingService,
    ) {
        this.canvas = elRef;

        this.canvas.nativeElement.onmousedown = this.drawingService.mousedown.bind(this.drawingService);
        this.canvas.nativeElement.onmouseup = this.drawingService.mouseup.bind(this.drawingService);
        this.canvas.nativeElement.onmousemove = this.drawingService.mousemove.bind(this.drawingService);
    }

    public ngAfterViewInit(): void {
        this.width = this.canvas!.nativeElement.parentElement?.getBoundingClientRect().width ?? this.width;
        this.height = this.canvas!.nativeElement.parentElement?.getBoundingClientRect().height ?? this.height;
        this.drawingService.init(this.canvas);

        this.cd.detectChanges();
    }
}
