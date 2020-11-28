import {Gui} from "./Gui";
import {ZombieMove} from "./ZombieMove";
import {MessageListener} from "./network/MessageListener";

export class ZombieMoveListener implements MessageListener {
    private gui: Gui;

    constructor(gui: Gui) {
        this.gui = gui
    }

    messageReceived(msg: string):void {
        // const zombieMoved:ZombieMove = JSON.parse(msg);
        const x:number = Math.floor(Math.random() * 20)
        const y:number = Math.floor(Math.random() * 10)

        const zombieMoved:ZombieMove = new ZombieMove(x, y)
        this.gui.zombieMoved(zombieMoved);
    }
}