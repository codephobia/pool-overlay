import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './components/app/app.component';
import { CanvasComponent } from './components/canvas/canvas.component';
import { DrawingService } from './services/drawing.service';
import { SocketService } from './services/socket.service';

const COMPONENTS = [
    AppComponent,
    CanvasComponent,
];

const SERVICES = [
    SocketService,
    DrawingService,
];

@NgModule({
    declarations: [
        ...COMPONENTS,
    ],
    imports: [BrowserModule],
    providers: [
        ...SERVICES,
    ],
    bootstrap: [AppComponent],
})
export class AppModule { }
