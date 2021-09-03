import { Component } from '@angular/core';

import { DrawingService } from '../../services/drawing.service';
import { SocketService } from '../../services/socket.service';

@Component({
    selector: 'pool-overlay-canvas',
    templateUrl: './canvas.component.html',
    styles: [`
        .pb-16-9 {
            padding-bottom: 56.25%;
        }
    `],
})
export class CanvasComponent {
    constructor(
        private _drawingService: DrawingService,
        private _socketService: SocketService,
    ) {
        this._socketService.bind('CLICK', this._drawingService.addClickEvent.bind(this._drawingService));
        this._socketService.bind('UNDO', this._drawingService.undo.bind(this._drawingService));
        this._socketService.bind('CLEAR', this._drawingService.clear.bind(this._drawingService));
        this._socketService.connect();
    }
}
