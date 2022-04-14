const elementsSpritesheetPath = "src/features/create-map/elements.json"

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
    .load(setup)

function setup() {

    const elementsSpritesheet = PIXI.Loader.shared.resources[elementsSpritesheetPath]

    let rows = Math.ceil(screenHeight / 48)
    let cols = Math.ceil(screenWidth / 48)

    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            const grass = new PIXI.Sprite(elementsSpritesheet.textures["tile000.png"])
            grass.y = row * 48
            grass.x = col * 48
            app.stage.addChild(grass);
        }
    }

    // grass.visible = false
}
