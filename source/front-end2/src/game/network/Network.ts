import {WebsocketHandler} from "./WebsocketHandler";
import {MessageListener} from "./MessageListener";

export class Network {
    private websocketHandler:WebsocketHandler
    private messageListeners:MessageListener[]

    constructor(websocketHandler: WebsocketHandler) {
        this.websocketHandler = websocketHandler
        this.websocketHandler.setOnMessage(this, this.onMessage)

        this.messageListeners = []
    }

    async connect(): Promise<void> {
        return this.websocketHandler.connect()
    }

    send(msg:string): void {
        this.websocketHandler.send(msg)
    }

    disconnect():void {
        this.websocketHandler.disconnect()
    }

    addMessageListener(m: MessageListener):void {
        this.messageListeners.push(m)
    }

    private onMessage(e:MessageEvent): void {
        for (const l of this.messageListeners) {
            l.messageReceived(e.data)
        }
    }
}
