package com.gridwalls.zombie

import kotlin.js.ExperimentalJsExport
import kotlin.js.JsExport

@JsExport
@ExperimentalJsExport
class Zombie : ZombieInt {
    override fun hello(): String {
         return "HELLO FROM KOTLIN!"
    }
}