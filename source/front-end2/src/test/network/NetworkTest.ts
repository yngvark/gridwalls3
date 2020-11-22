// import { Options } from '../src/Classes/Options/Options'; // this will be your custom import
import { expect } from 'chai';
import {Network} from "../../game/network/Network";
import {WebsocketHandler} from "../../game/network/WebsocketHandler";
import {MessageListener} from "../../game/network/MessageListener";

describe('Options tests', () => { // the tests container
    it('checking default options', () => { // the single test
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