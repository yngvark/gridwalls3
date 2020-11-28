import {Game} from "./Game";
import {Network} from "./network/Network";
import {Logger} from "./Logger";
import {WebsocketHandler} from "./network/WebsocketHandler";
import {Gui} from "./Gui";
import {ZombieMoveListener} from "./ZombieMoveListener";

const log = Logger.create("index")

document.addEventListener('DOMContentLoaded', () => {
    console.log("index.ts loaded");

    const network = new Network(new WebsocketHandler("ws://localhost:8080/zombie"))
    const gui = new Gui(20, 11)
    const game = new Game(network, gui)
    const zombieMoveListener = new ZombieMoveListener(gui)

    network.addMessageListener(zombieMoveListener)

    document.getElementById("connectBtn").onclick = async () => {
        await game.run()
    }

    document.getElementById("disconnectBtn").onclick = () => {
        game.stop()
    }

    document.getElementById("sendBtn").onclick = () => {
        const msg = (document.getElementById("msg") as HTMLInputElement).value
        network.send(msg)
    }
}, false);
