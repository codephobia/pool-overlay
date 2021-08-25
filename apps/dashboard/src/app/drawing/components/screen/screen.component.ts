import { AfterViewInit, Component, ElementRef, ViewChild } from '@angular/core';

@Component({
    selector: 'pool-overlay-screen',
    templateUrl: './screen.component.html',
})
export class ScreenComponent implements AfterViewInit {
    @ViewChild('video', { static: false, read: ElementRef })
    public video!: ElementRef<HTMLVideoElement>;

    public async ngAfterViewInit(): Promise<void> {
        try {
            await this.startCapture();
            await this.video.nativeElement.play();
        } catch (err) {
            console.error("Error: " + err);
        }
    }

    public async startCapture(): Promise<void> {
        this.video.nativeElement.srcObject = await (<any>navigator.mediaDevices).getDisplayMedia();
    }
}
