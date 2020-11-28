import {ZombieMove} from "./ZombieMove";
import {MainScene} from "./MainScene";

export class Gui {
    private game:Phaser.Game;
    private scene:MainScene;

    run():void {
        console.log("Starting game!");

        this.scene = new MainScene();

        this.game = new Phaser.Game({
            width: 800,
            height: 350,
            type: Phaser.AUTO,
            scene: this.scene
        });
    }

    // private zombies:{
    //     [key: string]: Zombie
    // } = {};

    zombieMoved(zombieMoved: ZombieMove): void {
        console.log("Drawing", zombieMoved)

        // if (this.zombies.hasOwnProperty(zombieMoved.id)) {
        //     const zombie = this.zombies[zombieMoved.id];
        //     // console.log("Existing zombie:");
        //     // console.log(zombie);
        //     zombie.sprite.setPosition(zombieMoved.coordinate.x * 15, zombieMoved.coordinate.y * 15);
        // } else {
        //     const sprite = this.scene.add.sprite(zombieMoved.coordinate.x, zombieMoved.coordinate.y, "skeleton");
        //     sprite.setScale(0.2 , 0.2);
        //
        //     const zombie:Zombie = {
        //         id: zombieMoved.id,
        //         sprite: sprite
        //     }
        //
        //     // console.log("New zombie:");
        //     // console.log(zombie);
        //
        //     this.zombies[zombie.id] = zombie;
        // }
    }
}