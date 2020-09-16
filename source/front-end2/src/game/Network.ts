import {Logger} from './Logger'

export class Network {
    private logger:Logger = Logger.create(this)

    private host: string;

    constructor(host: string) {
        this.host = host
    }

    connect(): void {
        this.logger.debug(`Connecting to ${this.host}`)
    }
}
