import {Logger} from '../Logger'

export class WebsocketHandler {
    private log:Logger = Logger.create(this)
    private websocket:WebSocket
    private readonly host: string;

    private onOpen: (event: Event) => void
    private onError: (event: Event) => void
    private onClose: (event: Event) => void
    private onMessage: (event: Event) => void
    private onMessageCaller: any;

    constructor(host: string) {
        this.host = host
    }

    setOnMessage(caller:any, fn: (event: MessageEvent) => void) {
        this.onMessageCaller = caller
        this.onMessage = fn
    }

    async connect(): Promise<void> {
        this.log.debug(`Connecting to ${this.host}`)

        // noinspection UnnecessaryLocalVariableJS
        const promise:Promise<void> = new Promise((resolve, reject) => {
            const websocket = new WebSocket(this.host)

            websocket.onopen = (event:Event) => {
                this.log.info("Connected!")
                this.websocket = websocket
                if (this.onOpen) this.onOpen(event)
                resolve()
            }

            websocket.onerror = (event:Event) => {
                this.log.error(event)
                if (this.onError) this.onError(event)
                reject(event)
            }

            websocket.onclose = (event:CloseEvent) => {
                this.log.info("Disconnected!")
                if (this.onClose) this.onClose(event)
                this.websocket = null
            }

            websocket.onmessage = (event) => {
                if (this.onMessage) this.onMessage.call(this.onMessageCaller, event)
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
