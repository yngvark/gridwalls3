package gridwalls

import gridwalls.Direction.*

data class Zombie(
    val x:Int,
    val y:Int
)

fun move(zombie:Zombie, direction: Direction):Zombie {
    return when (direction) {
        LEFT -> Zombie(zombie.x - 1, zombie.y)
        RIGHT -> Zombie(zombie.x + 1, zombie.y)
        UP -> Zombie(zombie.x, zombie.y - 1)
        DOWN -> Zombie(zombie.x, zombie.y + 1)
    }
}