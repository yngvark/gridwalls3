package gridwalls

import ZombieMoveGenerator
import io.ktor.application.*
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.features.json.*
import io.ktor.client.features.logging.*
import io.ktor.features.*
import io.ktor.gson.*
import io.ktor.http.*
import io.ktor.http.cio.websocket.*
import io.ktor.http.cio.websocket.Frame
import io.ktor.response.*
import io.ktor.routing.*
import io.ktor.websocket.*
import kotlinx.coroutines.async
import kotlinx.coroutines.channels.ClosedReceiveChannelException
import java.time.Duration

fun main(args: Array<String>): Unit = io.ktor.server.netty.EngineMain.main(args)

@Suppress("unused") // Referenced in application.conf
@kotlin.jvm.JvmOverloads
fun Application.module(testing: Boolean = false) {
    val zombieMoveGenerator = ZombieMoveGenerator()

    install(ContentNegotiation) {
        gson {
        }
    }

    val client = HttpClient(CIO) {
        install(JsonFeature) {
            serializer = GsonSerializer()
        }
        install(Logging) {
            level = LogLevel.HEADERS
        }
    }

//    runBlocking {
        // Sample for making a HTTP Client request
        /*
        val message = client.post<JsonSampleClass> {
            url("http://127.0.0.1:8080/path/to/endpoint")
            contentType(ContentType.Application.Json)
            body = JsonSampleClass(hello = "world")
        }
        */
//    }
//
    install(io.ktor.websocket.WebSockets) {
        pingPeriod = Duration.ofSeconds(15)
        timeout = Duration.ofSeconds(15)
        maxFrameSize = Long.MAX_VALUE
        masking = false
    }

    routing {
        get("/") {
            call.respondText("HELLO WORLD!", contentType = ContentType.Text.Plain)
        }

        get("/json/gson") {
            call.respond(mapOf("hello" to "world"))
        }

//        webSocket("/myws/echo") {
//            System.out.println("Websocket: Incoming")
//            send(Frame.Text("Hi from server"))
//
//            for (frame in incoming) {
//                when (frame) {
//                    is Frame.Text -> {
//                        val text = frame.readText()
//                        System.out.println("Websocket msg received: $text")
//
//                        outgoing.send(Frame.Text("YOU SAID: $text"))
//                        if (text.equals("bye", ignoreCase = true)) {
//                            System.out.println("Good bye websocket")
//                            close(CloseReason(CloseReason.Codes.NORMAL, "Client said BYE"))
//                        }
//                    }
//                }
//            }
//        }
        webSocket("/myws/echo") {
            println("onConnect")
            async { zombieMoveGenerator.sendTo(outgoing) }

            try {
                println("Reading all...")
                for (frame in incoming) {
                    println("ReadText...")
                    val text = (frame as Frame.Text).readText()
                    println("onMessage: ${text}")

                    outgoing.send(Frame.Text(text))
                }
            } catch (e: ClosedReceiveChannelException) {
                println("onClose ${closeReason.await()}")
            } catch (e: Throwable) {
                println("onError ${closeReason.await()}")
                e.printStackTrace()
            }

            println("HEY HERE")
        }
    }
}

data class JsonSampleClass(val hello: String)

