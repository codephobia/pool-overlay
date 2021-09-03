import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { SharedModule } from '../shared/shared.module';
import { DrawingRoutingModule } from './drawing-routing.module';
import { HomePageComponent } from './containers/home/home-page.component';
import { ScreenComponent } from './components/screen/screen.component';
import { CanvasComponent } from './components/canvas/canvas.component';
import { DrawingDirective } from './directives/drawing.directive';
import { ControlsComponent } from './components/controls/controls.component';
import { DrawingService } from './services/drawing.service';

const COMPONENTS = [
    HomePageComponent,
    ScreenComponent,
    CanvasComponent,
    ControlsComponent,
];

const DIRECTIVES = [
    DrawingDirective,
];

const SERVICES = [
    DrawingService,
];

@NgModule({
    imports: [
        CommonModule,
        HttpClientModule,
        FontAwesomeModule,
        SharedModule,
        DrawingRoutingModule,
    ],
    exports: [],
    declarations: [
        ...COMPONENTS,
        ...DIRECTIVES,
    ],
    providers: [
        ...SERVICES,
    ],
})
export class DrawingModule { }
