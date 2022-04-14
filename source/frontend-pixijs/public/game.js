const elementsSpritesheetPath = "src/features/create-map/elements.png"

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
    .load(setup)

function setup() {
    const grassTexture = TextureCache[elementsSpritesheetPath];
    grassTexture.frame = new Rectangle(0, 0, 48, 48)


    const rows = Math.ceil(screenHeight / 48)
    const cols = Math.ceil(screenWidth / 48)

    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            // const grass = new PIXI.Sprite(elementsSpritesheet.textures["tile000.png"])
            const grass = new Sprite(grassTexture)
            grass.y = row * 48
            grass.x = col * 48
            app.stage.addChild(grass);
        }
    }

    // grass.visible = false
}
