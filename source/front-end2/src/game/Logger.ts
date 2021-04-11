export class Logger {
    private readonly name:string

    static create(source:unknown): Logger {
        if (source instanceof Object) {
            return new Logger(source.constructor.name)
        }

        if (source instanceof String) {
            return new Logger(source.toString())
        }

        return new Logger("")
    }

    constructor(name: string) {
        this.name = name
    }

    error(msg: unknown): void {
        console.error(this.getPrefix(), msg)
    }

    warn(msg: unknown): void {
        console.warn(this.getPrefix(), msg)
    }

    info(msg: unknown, object?: unknown): void {
        if (object == undefined) {
            console.log(this.getPrefix(), msg)
        } else {
            console.log(this.getPrefix(), msg, object)
        }
    }

    debug(msg: unknown, object?: unknown): void {
        console.debug(this.getPrefix(), msg, object)
    }

    private getPrefix(): string {
        return `[${this.name}]`
    }
}