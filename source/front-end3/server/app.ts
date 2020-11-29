import express from 'express'
import cors from './cors'

export function getApp(backendUrl: string): express.Application {
    const app : express.Application = express()

    app.use('/', express.static('dist'))
    app.use(cors.middleware)

    app.get('/config', (req, res) => {
        let config = {
            backendUrl
        }
        res.send(config)
    })

    return app
}
