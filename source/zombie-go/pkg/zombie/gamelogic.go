package zombie

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/network"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/worldmap"
)

type GameLogic struct {
	msgSender            network.MessageSender
	stopGamelogicChannel chan bool
	generator            *Generator
}

func NewGameLogic(s network.MessageSender, stopGamelogicChannel chan bool) *GameLogic {
	m := worldmap.New(20, 10)                                        //nolint:gomnd
	zombie := NewZombie("1", 10, 5, m, rand.New(rand.NewSource(45))) //nolint:gosec,gomnd

	return &GameLogic{
		msgSender:            s,
		stopGamelogicChannel: stopGamelogicChannel,
		generator:            NewGenerator(zombie),
	}
}

func (l *GameLogic) Run() {
	fmt.Println("Starting to generate...")

	ticker := time.NewTicker(time.Second * 1) //nolint:gomnd
	defer ticker.Stop()

	for {
		select {
		case <-l.stopGamelogicChannel:
			log.Println("Zombie generator stopped.")
			return
		case <-ticker.C:
			zombieMove, err := l.generator.Next()
			if err != nil {
				fmt.Println("could not generate next message: %w", err)
				return
			}

			zombieMoveJSON, err := json.Marshal(zombieMove)
			if err != nil {
				fmt.Println("could not marshal zombie move: %w", err)
				return
			}

			err = l.msgSender.SendMsg(string(zombieMoveJSON))
			if err != nil {
				l.stopGamelogicChannel <- true
				return
			}
		}
	}
}

func (l *GameLogic) NotifyListeneres(msg string) {
	fmt.Printf("Gamelogic received msg: %s\n", msg)

	if msg == "start" {
		go l.Run()
	}
}
