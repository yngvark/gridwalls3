import { expect } from 'chai';

import {Network} from "../../src/network/Network";
import {WebsocketHandler} from "../../src/network/WebsocketHandler";
import {MessageListener} from "../../src/network/MessageListener";

describe('Network test', () => {
    it('should notify listeners when message is received', () => {
        const websocketHandler = new TestWebsocketHandler()
        const network = new Network(websocketHandler as unknown as WebsocketHandler)
        const listener = new TestMessageListener()
        network.addMessageListener(listener)

        // When
        websocketHandler.onMessage.call(network, new TestMessageEvent("hello") as MessageEvent)

        // Then
        expect(listener.msgReceived).to.equal("hello")
    });
});

class TestWebsocketHandler {
    onMessage: (event: Event) => void
    private onMessageCaller: any;

    setOnMessage(caller:any, fn: (event: MessageEvent) => void) {
        this.onMessageCaller = caller
        this.onMessage = fn
    }
}

class TestMessageListener implements MessageListener {
    msgReceived:string

    messageReceived(msg: string) {
        this.msgReceived = msg
    }
}

class TestMessageEvent {
    data:string

    constructor(msg:string) {
        this.data = msg
    }
}