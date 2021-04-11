import io.ktor.http.cio.websocket.*
import kotlinx.coroutines.channels.SendChannel
import java.net.http.WebSocket
import java.util.*

class ZombieMoveGenerator {
    suspend fun sendTo(outgoing: SendChannel<Frame>) {
        println("sendTo")

        while (true) {
            val pos = rndPosition()

            val msg = "Zombie moved to ${pos}"
            println("Sending: ${msg}")
            outgoing.send(Frame.Text(msg))
            Thread.sleep(3000);
        }
    }

    private fun rndPosition(): Int {
        return Random().nextInt()
    }

}