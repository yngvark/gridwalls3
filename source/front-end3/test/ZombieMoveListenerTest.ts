import { expect } from 'chai';

import {ZombieMoveListener} from "../src/ZombieMoveListener";
import {Gui} from "../src/Gui";
import {ZombieMove} from "../src/ZombieMove";

describe('ZombieMoveListener test', () => { // the tests container
    it('should parse JSON correctly', () => {
        // Given
        let zombieMove:ZombieMove
        const gui = {
            zombieMoved: zm => {
                zombieMove = zm
            }
        } as Gui
        const l = new ZombieMoveListener(gui)

        // When
        l.messageReceived('{"id": "1", "x":9,"y":5}')

        // Then
        expect(new ZombieMove("1", 9, 5)).to.deep.eq(zombieMove)
    });
});
