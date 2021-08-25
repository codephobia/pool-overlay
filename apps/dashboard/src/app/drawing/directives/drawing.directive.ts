import { AfterViewInit, ChangeDetectorRef, Directive, ElementRef, HostBinding } from '@angular/core';

export interface Click {
    color: string;
    x: number;
    y: number;
    drag: boolean;
}

@Directive({
    selector: '[poolOverlayDrawing]',
})
export class DrawingDirective implements AfterViewInit {
    @HostBinding('attr.width')
    public width = 300;

    @HostBinding('attr.height')
    public height = 168.75;

    private paint = false;
    private canvas: ElementRef<HTMLCanvasElement> | null = null;
    private context: CanvasRenderingContext2D | null = null;
    private color = 'FF0000';
    private clicks: Click[] = [];

    constructor(
        elRef: ElementRef,
        private cd: ChangeDetectorRef,
    ) {
        console.log('init');
        this.canvas = elRef;
        this.context = elRef.nativeElement.getContext("2d");

        this.canvas.nativeElement.onmousedown = this.mousedown.bind(this);
        this.canvas.nativeElement.onmouseup = this.mouseup.bind(this);
        this.canvas.nativeElement.onmousemove = this.mousemove.bind(this);

    }

    public ngAfterViewInit(): void {
        this.width = this.canvas!.nativeElement.parentElement?.getBoundingClientRect().width ?? this.width;
        this.height = this.canvas!.nativeElement.parentElement?.getBoundingClientRect().height ?? this.height;

        this.cd.detectChanges();
    }

    public setColor(color: string): void {
        this.color = color;
    }

    public clear(): void {
        this.clicks = [];
        this.redraw();
    }

    private mousedown(event: MouseEvent): void {
        if (!this.canvas) {
            return;
        }

        this.paint = true;

        const x = event.pageX - this.canvas.nativeElement.getBoundingClientRect().left;
        const y = event.pageY - this.canvas.nativeElement.getBoundingClientRect().top;
        const drag = false;
        const click: Click = {
            color: this.color,
            x,
            y,
            drag
        };

        this.clicks.push(click);
    }

    private mouseup(): void {
        this.paint = false;
    }

    public mousemove(event: MouseEvent): void {
        if (!this.paint || !this.canvas) {
            return;
        }

        const x = event.pageX - this.canvas.nativeElement.getBoundingClientRect().left;
        const y = event.pageY - this.canvas.nativeElement.getBoundingClientRect().top;
        const drag = true;
        const click: Click = {
            color: this.color,
            x,
            y,
            drag
        };

        this.clicks.push(click);

        this.redraw();
    }

    private redraw(): void {
        if (!this.context) {
            return;
        }

        this.context.clearRect(0, 0, this.context.canvas.width, this.context.canvas.height);

        this.context.lineJoin = 'round';
        this.context.lineWidth = 5;

        for (var i = 0; i < this.clicks.length; i++) {
            this.context.strokeStyle = '#' + this.clicks[i].color;
            this.context.beginPath();

            if (this.clicks[i].drag && i > 0) {
                this.context.moveTo(this.clicks[i - 1].x, this.clicks[i - 1].y);
            } else {
                this.context.moveTo(this.clicks[i].x - 1, this.clicks[i].y);
            }

            this.context.lineTo(this.clicks[i].x, this.clicks[i].y);
            this.context.closePath();
            this.context.stroke();
        }
    }
}
