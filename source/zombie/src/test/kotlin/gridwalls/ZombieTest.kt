package gridwalls

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class ZombieTest {
    @Test
    fun `zombie's coordinates should update correctly when it movies`() {
        assertEquals(6, move(Zombie(5, 3), Direction.RIGHT).x)
        assertEquals(4, move(Zombie(5, 3), Direction.LEFT).x)

        assertEquals(2, move(Zombie(5, 3), Direction.UP).y)
        assertEquals(4, move(Zombie(5, 3), Direction.DOWN).y)
    }
}