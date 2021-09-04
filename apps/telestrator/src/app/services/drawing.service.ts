import { ElementRef, Injectable } from '@angular/core';

export interface Click {
    color: string;
    width: number;
    x: number;
    y: number;
    scale: number;
    drag: boolean;
}

@Injectable()
export class DrawingService {
    private _context: CanvasRenderingContext2D | null = null;
    private _clicks: Click[] = [];
    private _initialized = false;

    public init(canvas: ElementRef<HTMLCanvasElement> | null): void {
        this._context = canvas!.nativeElement.getContext("2d");
        this._initialized = true;
    }

    public get initialized(): boolean {
        return this._initialized;
    }

    public undo(): void {
        if (!this._clicks.length) {
            return;
        }

        this._removeLastLine();
        this._redraw();
    }

    public clear(): void {
        if (!this._clicks.length) {
            return;
        }

        this._clicks = [];
        this._redraw();
    }

    public addClickEvent(payload: Click): void {
        this._clicks = this._clicks.concat(payload);
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
                const x = this._clicks[i - 1].x * this._clicks[i - 1].scale;
                const y = this._clicks[i - 1].y * this._clicks[i - 1].scale;

                this._context!.moveTo(x, y);
            } else {
                const x = this._clicks[i].x * this._clicks[i].scale - 1;
                const y = this._clicks[i].y * this._clicks[i].scale;

                this._context!.moveTo(x, y);
            }

            const x = this._clicks[i].x * this._clicks[i].scale;
            const y = this._clicks[i].y * this._clicks[i].scale;
            this._context!.lineTo(x, y);

            this._context!.closePath();
            this._context!.stroke();
        }
    }
}
