import { config } from 'dotenv'
config()

import {getApp} from "./app"
import express from "express"
import http from "http"
import https from "https"

const PORT = getEnv("PORT")
const BACKEND_URL = getEnv("BACKEND_URL")

function getEnv(name:string) {
    let e = process.env[name]
    if (!e) throw new Error(`Missing environment variable: ${name}`)

    return e
}

let app:express.Application = getApp(BACKEND_URL)
let server : http.Server | https.Server = http.createServer(app)

server.listen(PORT, () => {
    console.log(`Listening at http://localhost:${PORT}`)
})
