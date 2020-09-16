export class Logger {
    private name:String

    static create(instance): Logger {
        return new Logger(instance.constructor.name)
    }

    constructor(name: string) {
        this.name = name
    }

    info(msg: any): void {
        console.log(this.format(msg))
    }

    debug(msg: any): void {
        console.debug(this.format(msg))
    }

    private format(msg: any) {
        return `[${this.name}] ${msg}`
    }
}