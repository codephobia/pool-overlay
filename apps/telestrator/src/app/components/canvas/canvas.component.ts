import { AfterViewInit, Component, ElementRef, ViewChild } from '@angular/core';

import { DrawingService } from '../../services/drawing.service';
import { SocketService } from '../../services/socket.service';

@Component({
    selector: 'pool-overlay-canvas',
    templateUrl: './canvas.component.html',
})
export class CanvasComponent implements AfterViewInit {
    @ViewChild('canvas', { static: false, read: ElementRef })
    public canvas!: ElementRef<HTMLCanvasElement>;

    constructor(
        private _drawingService: DrawingService,
        private _socketService: SocketService,
    ) {
        this._socketService.bind('CLICK', this._drawingService.addClickEvent.bind(this._drawingService));
        this._socketService.bind('UNDO', this._drawingService.undo.bind(this._drawingService));
        this._socketService.bind('CLEAR', this._drawingService.clear.bind(this._drawingService));
    }

    public ngAfterViewInit() {
        this._drawingService.init(this.canvas);
        this._socketService.connect();
    }
}
