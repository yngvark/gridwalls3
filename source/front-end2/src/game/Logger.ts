export class Logger {
    private readonly name:string

    static create(instance:unknown): Logger {
        if (instance instanceof Object) {
            return new Logger(instance.constructor.name)
        }

        return new Logger("")
    }

    constructor(name: string) {
        this.name = name
    }

    info(msg: unknown): void {
        console.log(this.getPrefix(), msg)
    }

    debug(msg: unknown): void {
        console.debug(this.getPrefix(), msg)
    }

    private getPrefix(): string {
        return `[${this.name}]`
    }
}