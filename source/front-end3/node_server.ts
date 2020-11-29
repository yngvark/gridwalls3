import express from 'express';

function runServer() {
    const app : express.Application = express()
    const port = 3001

    app.use('/', express.static('dist'))

    app.get('/config', (req, res) => {
        let a:string = "hello"
        res.send(a)
    })

    app.listen(port, () => {
        console.log(`Example app listening at http://localhost:${port}`)
    })
}

runServer()
