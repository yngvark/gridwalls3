import { config } from 'dotenv'
config()

import {getApp} from "./app"
import express from "express"
import http from "http"
import https from "https"
import fs from "fs"

const PORT = getEnv("PORT")
const BACKEND_URL = getEnv("BACKEND_URL")

function getEnv(name:string) {
    let e = process.env[name]
    if (!e) throw new Error(`Missing environment variable: ${name}`)

    return e
}

let useHttps = false
if (process.env.CERTIFICATE_FILE || process.env.KEY_FILE) {
    if (!process.env.CERTIFICATE_FILE || !process.env.KEY_FILE) {
        throw new Error('SSL requires both a certificate and a key')
    }

    useHttps = true
}

let app:express.Application = getApp(BACKEND_URL)
let protocol:string
let server : http.Server | https.Server

if (!useHttps) {
    protocol = 'http'
    server = http.createServer(app)
} else {
    const certificate = fs.readFileSync(process.env.CERTIFICATE_FILE!, 'utf8')
    const privateKey = fs.readFileSync(process.env.KEY_FILE!, 'utf8')

    let credentials =  {
        key: privateKey,
        cert: certificate
    }

    protocol = 'https'
    server = https.createServer(credentials, app)
}

server.listen(PORT, () => {
    console.log(`Example app listening at ${protocol}://localhost:${PORT}`)
})
