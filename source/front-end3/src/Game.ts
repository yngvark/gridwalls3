import {Network} from "./network/Network";
import {Gui} from "./Gui";
import {ZombieMoveListener} from "./ZombieMoveListener";

export class Game {
    private network: Network;
    private gui: Gui;

    constructor(network: Network, gui:Gui) {
        this.network = network
        this.gui = gui
    }

    public async run(): Promise<void> {
        console.log("Running Game")
        await this.network.connect()
        this.network.send("hello")
        this.gui.run()
    }

    stop() {
        this.network.disconnect()
    }
}
