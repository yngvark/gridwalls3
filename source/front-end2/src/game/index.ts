import {Game} from "./Game";
import {Network} from "./Network";
import {Logger} from "./Logger";

const log = Logger.create("index")

document.addEventListener('DOMContentLoaded', () => {
    console.log("Hello from DOMContentLoaded. Running game!!!");

    // const network = new Network("ws://localhost:8080/myws/echo")
    const network = new Network("ws://localhost:8080/echo")
    const game = new Game(network)

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

