import {WebsocketHandler} from "./WebsocketHandler";
import {MessageListener} from "./MessageListener";

export class Network {
    private websocketHandler:WebsocketHandler

    constructor(websocketHandler: WebsocketHandler) {
        this.websocketHandler = websocketHandler
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

    addMessageListener(messageListener: MessageListener):void {
        console.log("addMessageListener")
    }
}
