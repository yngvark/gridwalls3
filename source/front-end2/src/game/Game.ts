import {Network} from "./Network";

export class Game {
    private network: Network;
    constructor(network: Network) {
        this.network = network
    }

    run(): void {
        console.log("Hello from Game 33")
        this.network.connect()
    }
}
