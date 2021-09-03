import { ElementRef, Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

export interface Click {
    color: string;
    width: number;
    x: number;
    y: number;
    drag: boolean;
}

@Injectable()
export class DrawingService {
    private _canvas: ElementRef<HTMLCanvasElement> | null = null;
    private _context: CanvasRenderingContext2D | null = null;
    private _paint = false;
    private _clicks: Click[] = [];
    private _initialized = false;
    private _colors: string[] = [
        '#FFFFFF',
        '#7160A9',
        '#FECD0F',
        '#4BB76D',
        '#E67EB0',
        '#A38672',
    ];
    private _color = new BehaviorSubject<string>(this._colors[0]);
    private _lineWidths: number[] = [
        1,
        3,
        5,
    ];
    private _lineWidth = new BehaviorSubject<number>(this._lineWidths[0]);

    public init(canvas: ElementRef<HTMLCanvasElement> | null): void {
        this._canvas = canvas;
        this._context = canvas!.nativeElement.getContext("2d");
        this._initialized = true;
    }

    public get initialized(): boolean {
        return this._initialized;
    }

    public get color$(): Observable<string> {
        return this._color.asObservable();
    }

    public get colors(): string[] {
        return this._colors;
    }

    public setColor(color: string): void {
        this._color.next(color);
    }

    public get lineWidth$(): Observable<number> {
        return this._lineWidth.asObservable();
    }

    public get lineWidths(): number[] {
        return this._lineWidths;
    }

    public setLineWidth(width: number): void {
        this._lineWidth.next(width);
    }

    public undo(): void {
        this._removeLastLine();
        this._redraw();
    }

    public clear(): void {
        this._clicks = [];
        this._redraw();
    }

    public mousedown(event: MouseEvent): void {
        if (!this.initialized) {
            return;
        }
        this._paint = true;

        const x = event.pageX - this._canvas!.nativeElement.getBoundingClientRect().left;
        const y = event.pageY - this._canvas!.nativeElement.getBoundingClientRect().top;
        const drag = false;
        const click: Click = {
            color: this._color.value,
            width: this._lineWidth.value,
            x,
            y,
            drag
        };

        this._clicks.push(click);
    }

    public mouseup(): void {
        if (!this.initialized) {
            return;
        }

        this._paint = false;
    }

    public mousemove(event: MouseEvent): void {
        if (!this.initialized || !this._paint) {
            return;
        }

        const x = event.pageX - this._canvas!.nativeElement.getBoundingClientRect().left;
        const y = event.pageY - this._canvas!.nativeElement.getBoundingClientRect().top;
        const drag = true;
        const click: Click = {
            color: this._color.value,
            width: this._lineWidth.value,
            x,
            y,
            drag
        };

        this._clicks.push(click);

        this._redraw();
    }

    private _removeLastLine(): void {
        for (let i = this._clicks.length - 1; i >= 0; i--) {
            const click = this._clicks.pop();

            if (!click?.drag) {
                break;
            }
        }
    }

    private _redraw(): void {
        if (!this.initialized) {
            return;
        }

        this._context!.clearRect(0, 0, this._context!.canvas.width, this._context!.canvas.height);
        this._context!.lineJoin = 'round';

        for (var i = 0; i < this._clicks.length; i++) {
            this._context!.lineWidth = this._clicks[i].width;
            this._context!.strokeStyle = this._clicks[i].color;
            this._context!.beginPath();

            if (this._clicks[i].drag && i > 0) {
                this._context!.moveTo(this._clicks[i - 1].x, this._clicks[i - 1].y);
            } else {
                this._context!.moveTo(this._clicks[i].x - 1, this._clicks[i].y);
            }

            this._context!.lineTo(this._clicks[i].x, this._clicks[i].y);
            this._context!.closePath();
            this._context!.stroke();
        }
    }
}
