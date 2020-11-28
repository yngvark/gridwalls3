import {Gui} from "./Gui";
import {ZombieMove} from "./ZombieMove";
import {MessageListener} from "./network/MessageListener";

export class ZombieMoveListener implements MessageListener {
    private gui: Gui;

    constructor(gui: Gui) {
        this.gui = gui
    }

    messageReceived(msg: string):void {
        const zombieMoved:ZombieMove = JSON.parse(msg);
        this.gui.zombieMoved(zombieMoved);
    }
}