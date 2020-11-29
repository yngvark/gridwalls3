import cors from 'cors'

function corsPluginOriginValidator(origin:any, callback:any) {
    // Origin often undefined during machine-to-machine communications (which is what is happening nuxt server side)
    if (!origin) return callback(null, true)

    try {
        callback(null, true)
    } catch (error) {
        callback(error)
    }
}

const corsOptions = {
    credentials: true,
    origin: corsPluginOriginValidator
}

export default {
    middleware: cors(corsOptions)
}
