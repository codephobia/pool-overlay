import { Injectable } from '@angular/core';

type EventHandlers = { [key: string]: (payload: any) => void };

@Injectable()
export class SocketService {
    private _connTimer: ReturnType<typeof setInterval> | null = null;
    private _conn: WebSocket | null = null;
    private _eventHandlers: EventHandlers = {};

    public connect(): void {
        try {
            this._conn = new WebSocket('ws://192.168.0.26:1268/latest/overlay');
            this._conn.onopen = this._onOpen.bind(this);
            this._conn.onmessage = this._onMessage.bind(this);
            this._conn.onclose = this._onClose.bind(this);
            this._conn.onerror = this._onError.bind(this);
        } catch (err) {
            console.log(err);
        }
    }

    public bind(eventType: string, callback: (payload: any) => void): void {
        if (!this._eventHandlers.hasOwnProperty(eventType)) {
            this._eventHandlers[eventType] = callback;
        }
    }

    private _onOpen(): void {
        if (this._connTimer) {
            clearInterval(this._connTimer);
            this._connTimer = null;
        }
    }

    private _onClose(): void {
        if (!this._connTimer) {
            this._connTimer = setInterval(this.connect.bind(this), 5000);
        }
    }

    private _onMessage(message: any): void {
        const { type, payload } = JSON.parse(message.data);

        if (this._eventHandlers.hasOwnProperty(type)) {
            this._eventHandlers[type](payload);
        } else {
            console.log('unknown message type: ', type);
        }
    }

    private _onError(ev: Event): void {
        console.log(ev);
    }
}
