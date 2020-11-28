import 'phaser';

import {ZombieMove} from "./ZombieMove";
import {MainScene} from "./MainScene";
import {Zombie} from "./Zombie";
import Sprite = Phaser.GameObjects.Sprite;

export class Gui {
    private readonly TILE_WIDTH:number = 40

    private game:Phaser.Game;
    private scene:Phaser.Scene;

    constructor(readonly xTiles:number, readonly yTiles:number) {
    }

    private zombies:{
        [id: string]: Zombie
    } = {};

    private zombieSprites:{
        [id: string]: Sprite
    } = {};

    run():void {
        console.log("Running GUI");
        this.scene = new MainScene()

        this.game = new Phaser.Game({
            type: Phaser.AUTO,
            backgroundColor: '#333333',
            width: this.xTiles * this.TILE_WIDTH,
            height: this.yTiles * this.TILE_WIDTH,
            scene: this.scene,
            parent: 'gameContent'
        });

        // @ts-ignore
        window.gw_scene = this.scene
    }

    // zombieMoved(zombieMove: ZombieMove): void {
    //
    // }

    zombieMoved(zombieMoved: ZombieMove): void {
        // console.log("Drawing", zombieMoved)

        if (this.zombies.hasOwnProperty(zombieMoved.id)) {
            const zombie = this.zombies[zombieMoved.id];
            // console.log("Existing zombie:");
            // console.log(zombie);
            this.zombieSprites[zombie.id].setPosition(zombieMoved.x * this.TILE_WIDTH, zombieMoved.y * this.TILE_WIDTH);
        } else {
            let x = zombieMoved.x * this.TILE_WIDTH
            let y = zombieMoved.y * this.TILE_WIDTH

            let sprite = this.scene.add.sprite(x, y, "skeleton");
            sprite.setOrigin(0, 0)
            sprite.setScale(0.2 , 0.2); // 40x40

            const newZombie = new Zombie(zombieMoved.id, zombieMoved.x, zombieMoved.y)
            this.zombieSprites[newZombie.id] = sprite

            // console.log("New zombie:");
            // console.log(newZombie);

            this.zombies[newZombie.id] = newZombie;
        }
    }
}