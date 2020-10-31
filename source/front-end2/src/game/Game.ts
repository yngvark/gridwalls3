import {Network} from "./Network";

export class Game {
    private network: Network;

    constructor(network: Network) {
        this.network = network
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
