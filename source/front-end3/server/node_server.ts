import { config } from 'dotenv'
config()

import {getServer} from "./app"
import express from "express"
import https from "https"
import fs from "fs"
import path from "path"

let app:express.Application = getServer()

if (!process.env.CERTIFICATE_FILE || !process.env.KEY_FILE) {
    throw new Error('SSL requires both a certificate and a key')
}

const certificate = fs.readFileSync(process.env.CERTIFICATE_FILE, 'utf8')
const privateKey = fs.readFileSync(process.env.KEY_FILE!, 'utf8')

let credentials =  {
    key: privateKey,
    cert: certificate
}

let server = https.createServer(credentials, app)

const port = 3001

server.listen(port, () => {
    console.log(`Example app listening at https://localhost:${port}`)
})
