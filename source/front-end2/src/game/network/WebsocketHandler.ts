import {Logger} from '../Logger'

export class WebsocketHandler {
    private log:Logger = Logger.create(this)
    private websocket:WebSocket

    private readonly host: string;

    constructor(host: string) {
        this.host = host
    }

    async connect(): Promise<void> {
        this.log.debug(`Connecting to ${this.host}`)

        // noinspection UnnecessaryLocalVariableJS
        const promise:Promise<void> = new Promise((resolve, reject) => {
            const websocket = new WebSocket(this.host)

            websocket.onopen = () => {
                this.log.info("Connected!")
                this.websocket = websocket
                resolve()
            }

            websocket.onerror = error => {
                this.log.error(error)
                reject(error)
            }

            websocket.onclose = () => {
                this.log.info("Disconnected!")
                this.websocket = null
            }

            websocket.onmessage = event => {
                this.log.info("Received message!", event.data)
            }
        })

        return promise
    }


    send(msg:string): void {
        this.log.debug("Sending", msg)
        this.websocket.send(msg)
    }

    disconnect():void {
        this.log.info("Disconnecting")
        this.websocket.close()
    }
}
