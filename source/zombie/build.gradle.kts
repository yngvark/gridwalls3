import com.github.jengelman.gradle.plugins.shadow.tasks.ShadowJar

plugins {
    kotlin("jvm") version "1.3.72"

    application

    // Adds "shadowJar" task for generating JAR file
    id("com.github.johnrengelman.shadow") version "5.2.0"

    // Plugin for generating a docker image
    id("com.google.cloud.tools.jib") version "1.3.0"
}

group = "gridwalls"
version = "1.0-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_11

repositories {
    mavenCentral()
    jcenter()
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))

    // Testing
    testImplementation("org.junit.jupiter:junit-jupiter:5.6.1")
}

tasks {
    compileKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
    compileTestKotlin {
        kotlinOptions.jvmTarget = "1.8"
    }
}

tasks.test {
    useJUnitPlatform()
}

//tasks {
//    named<ShadowJar>("shadowJar") {
////        archiveBaseName.set("shadow")
//        //mergeServiceFiles()
//        manifest {
//            attributes(mapOf("Main-Class" to "gridwalls.MainKt"))
//        }
//    }
//}
//
//tasks {
//    build {
//        dependsOn(shadowJar)
//    }
//}

application {
    mainClassName = "gridwalls.MainKt"
}

// Configuration for creating a docker image
jib {
    container {
        mainClass = "gridwalls.MainKt"
    }
    from {
        image = "openjdk:14-jdk"
    }
    to {
        image = "yngvark/gridwalls-zombie"
    }
}

println("Using java version: " + JavaVersion.current())
