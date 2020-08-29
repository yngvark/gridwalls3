import com.google.gson.Gson
import com.google.gson.GsonBuilder
import com.google.gson.JsonParser
import org.jetbrains.kotlin.js.parser.sourcemaps.JsonObject
import org.jetbrains.kotlin.js.parser.sourcemaps.JsonString
import org.jetbrains.kotlin.js.parser.sourcemaps.parseJson

//import com.google.gson.JsonObject

plugins {
    kotlin("multiplatform") version "1.4.0"
}
group = "com.gridwalls"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
//    jcenter()
//    maven { url = uri("https://dl.bintray.com/kotlin/kotlin-eap") }
}

kotlin {
    jvm {
        compilations.all {
            kotlinOptions.jvmTarget = "1.8"
        }
    }

    js("browser") {
        browser {
            testTask {
                useKarma {
                    useChromeHeadless()
                }
            }
            binaries.executable()
        }
    }

    sourceSets {
        val commonMain by getting
        val commonTest by getting {
            dependencies {
                implementation(kotlin("test-common"))
                implementation(kotlin("test-annotations-common"))
            }
        }
        val jvmMain by getting
        val jvmTest by getting {
            dependencies {
                implementation(kotlin("test-junit5"))
            }
        }
        val browserMain by getting
        val browserTest by getting {
            dependencies {
                implementation(kotlin("test-js"))
            }
        }

        js().compilations["main"].defaultSourceSet {
            dependencies {
                implementation(kotlin("stdlib-js"))
            }
        }
    }
}

tasks.register("npmBuild") {
    dependsOn(tasks.named("build"))

    doLast {
        makePackageJsonPublishable()
    }
}

fun makePackageJsonPublishable() {
    val packageName = "${project.name}-browser"

    // Modify Json
    val packageJson = file("$buildDir/js/packages/${packageName}/package.json")
    val outJson = file("$buildDir/js/packages/${packageName}/package2.json")

    val pj = parseJson(packageJson) as JsonObject
    pj.properties["publishConfig"] = JsonObject("registry" to JsonString("https://npm.pkg.github.com/"))
    pj.properties["repository"] = JsonString("git@github.com:yngvark/gridwalls3.git")
    pj.properties["name"] = JsonString("@yngvark/${packageName}")
    // TODO: Set version from build.gradle.kts version

    // Convert to pretty printed string
    val gson:Gson = GsonBuilder().setPrettyPrinting().create()
    val parser = JsonParser()
    val jsonObject = parser.parse(pj.toString()).asJsonObject

    val jsonPrettyPrinted:String = gson.toJson(jsonObject)

    // Write to file
    packageJson.writeText(jsonPrettyPrinted)
}