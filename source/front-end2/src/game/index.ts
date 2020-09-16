import {Game} from "./Game";
import {Network} from "./Network";

document.addEventListener('DOMContentLoaded', function() {
    console.log("Hello from DOMContentLoaded. Running game!!!");

    let network = new Network("localhost")
    new Game(network).run()
}, false);
