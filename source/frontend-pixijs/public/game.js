const elementsSpritesheetPath = "src/features/create-map/elements.png"
const playerImagepath = "src/features/player/human.png"

const Application = PIXI.Application,
    Container = PIXI.Container,
    loader = PIXI.Loader.shared,
    resources = PIXI.Loader.shared.resources,
    TextureCache = PIXI.utils.TextureCache,
    Sprite = PIXI.Sprite,
    Rectangle = PIXI.Rectangle;

let type = "WebGL";
if (!PIXI.utils.isWebGLSupported()) {
    console.log("canvas ugh")
    type = "canvas";
}

PIXI.utils.sayHello(type);

const screenWidth = 48 * 30
const screenHeight = 48 * 15

const app = new PIXI.Application({width: screenWidth, height: screenHeight});

document.body.appendChild(app.view);

PIXI.Loader.shared
    .add(elementsSpritesheetPath)
    .add(playerImagepath)
    .load(setup)

let playerSprite = null

function setup() {
    const grassTexture = TextureCache[elementsSpritesheetPath];
    grassTexture.frame = new Rectangle(0, 0, 48, 48)

    const rows = Math.ceil(screenHeight / 48)
    const cols = Math.ceil(screenWidth / 48)

    console.log("rows, cols", rows, cols)

    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            const mapSprite = new Sprite(grassTexture)
            mapSprite.y = row * 48
            mapSprite.x = col * 48

            app.stage.addChild(mapSprite);
        }
    }
    // grass.visible = false

    // Player
    const playerTexture = TextureCache[playerImagepath];
    playerSprite = new Sprite(playerTexture)
    playerSprite.x = 48 * 3
    playerSprite.y = 48 * 3

    enableMovementWithKeyboardKeys(playerSprite)
    app.stage.addChild(playerSprite)


    // Gameloop

    app.ticker.add((delta) => gameLoop(delta));
}

function gameLoop(delta) {
    playerSprite.x += playerSprite.vx
    playerSprite.y += playerSprite.vy
}

function keyboard(value) {
    const key = {};
    key.value = value;
    key.isDown = false;
    key.isUp = true;
    key.press = undefined;
    key.release = undefined;
    //The `downHandler`
    key.downHandler = (event) => {
        if (event.key === key.value) {
            if (key.isUp && key.press) {
                key.press();
            }
            key.isDown = true;
            key.isUp = false;
            event.preventDefault();
        }
    };

    //The `upHandler`
    key.upHandler = (event) => {
        if (event.key === key.value) {
            if (key.isDown && key.release) {
                key.release();
            }
            key.isDown = false;
            key.isUp = true;
            event.preventDefault();
        }
    };

    //Attach event listeners
    const downListener = key.downHandler.bind(key);
    const upListener = key.upHandler.bind(key);

    window.addEventListener("keydown", downListener, false);
    window.addEventListener("keyup", upListener, false);

    // Detach event listeners
    key.unsubscribe = () => {
        window.removeEventListener("keydown", downListener);
        window.removeEventListener("keyup", upListener);
    };

    return key;
}

function enableMovementWithKeyboardKeys(sprite) {
    sprite.vx = 0;
    sprite.vy = 0;

    console.log("enableMovementWithKeyboardKeys", sprite)
    //Capture the keyboard arrow keys
    const left = keyboard("ArrowLeft"),
        up = keyboard("ArrowUp"),
        right = keyboard("ArrowRight"),
        down = keyboard("ArrowDown");

    //Left arrow key `press` method
    left.press = () => {
        //Change the cat's velocity when the key is pressed
        sprite.vx = -5;
        sprite.vy = 0;
    };

    //Left arrow key `release` method
    left.release = () => {
        //If the left arrow has been released, and the right arrow isn't down,
        //and the cat isn't moving vertically:
        //Stop the cat
        if (!right.isDown && sprite.vy === 0) {
            sprite.vx = 0;
        }
    };

    //Up
    up.press = () => {
        sprite.vy = -5;
        sprite.vx = 0;
    };
    up.release = () => {
        if (!down.isDown && sprite.vx === 0) {
            sprite.vy = 0;
        }
    };

    //Right
    right.press = () => {
        sprite.vx = 5;
        sprite.vy = 0;
    };
    right.release = () => {
        if (!left.isDown && sprite.vy === 0) {
            sprite.vx = 0;
        }
    };

    //Down
    down.press = () => {
        sprite.vy = 5;
        sprite.vx = 0;
    };
    down.release = () => {
        if (!up.isDown && sprite.vx === 0) {
            sprite.vy = 0;
        }
    };
}