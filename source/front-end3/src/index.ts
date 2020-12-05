import {Game} from "./Game";
import {Network} from "./network/Network";
import {Logger} from "./Logger";
import {WebsocketHandler} from "./network/WebsocketHandler";
import {Gui} from "./Gui";
import {ZombieMoveListener} from "./ZombieMoveListener";

const log = Logger.create("index")

document.addEventListener('DOMContentLoaded', async () => {
    console.log("index.ts loaded 33");

    let backendUrl = await getBackendUrl()
    console.log("backendUrl: " + backendUrl)

    const network = new Network(new WebsocketHandler(backendUrl))
    const gui = new Gui(20, 11)
    const game = new Game(network, gui)
    const zombieMoveListener = new ZombieMoveListener(gui)

    network.addMessageListener(zombieMoveListener)

    document.getElementById("connectBtn")!.onclick = async () => {
        await game.run()
    }

    document.getElementById("disconnectBtn")!.onclick = () => {
        game.stop()
    }

    document.getElementById("sendBtn")!.onclick = () => {
        const msg = (document.getElementById("msg") as HTMLInputElement).value
        network.send(msg)
    }
}, false);

async function getBackendUrl():Promise<string> {
    const promise = new Promise<string>((resolve, reject) => {
        let httpRequest = new XMLHttpRequest();
        httpRequest.onreadystatechange = function(res){
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let r = JSON.parse(httpRequest.responseText)
                    if (!r.backendUrl)
                        reject("backendUrl not found in server configuration: " + httpRequest.responseText)
                    resolve(r.backendUrl)
                } else {
                    reject(httpRequest.response)
                }
            }
        };
        // httpRequest.open('GET', 'http://localhost:3000/config', true);
        httpRequest.open('GET', '/config', true);
        httpRequest.send();
    })

    return promise
}