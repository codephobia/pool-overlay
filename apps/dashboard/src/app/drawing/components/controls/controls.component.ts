import { ChangeDetectionStrategy, Component } from '@angular/core';
import { faCircle, faEraser, faPen, faUndo } from '@fortawesome/free-solid-svg-icons';
import { faCircle as faCircleRing } from '@fortawesome/free-regular-svg-icons';

import { DrawingService } from '../../services/drawing.service';

@Component({
    selector: 'pool-overlay-controls',
    templateUrl: './controls.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ControlsComponent {
    public faPen = faPen;
    public faEraser = faEraser;
    public faUndo = faUndo;
    public faCircle = faCircle;
    public faCircleRing = faCircleRing;
    public colors = this.drawingService.colors;
    public activeColor$ = this.drawingService.color$;
    public lineWidths = this.drawingService.lineWidths;
    public activeLineWidth$ = this.drawingService.lineWidth$;

    constructor(
        private drawingService: DrawingService,
    ) { }

    public setColor(color: string): void {
        this.drawingService.setColor(color);
    }

    public setLineWidth(width: number): void {
        this.drawingService.setLineWidth(width);
    }

    public undo(): void {
        this.drawingService.undo();
    }

    public clear(): void {
        this.drawingService.clear();
    }
}
