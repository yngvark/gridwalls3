import {Network} from "./network/Network";
import {Gui} from "./Gui";
import {ZombieMoveListener} from "./ZombieMoveListener";

export class Game {
    private network: Network;

    constructor(network: Network) {
        this.network = network
    }

    init(gui:Gui) {
        this.network.addMessageListener(new ZombieMoveListener(gui))
    }

    public async run(): Promise<void> {
        console.log("Hello from Game 33")
        await this.network.connect()
        this.network.send("hello")
    }

    stop() {
        this.network.disconnect()
    }
}
