package zombie

import (
	"fmt"
	"log"
	"time"
	"zombie-go/pkg/network"
	"zombie-go/pkg/worldmap"
)

type GameLogic struct {
	msgSender            network.MessageSender
	stopGamelogicChannel chan bool
	generator            *generator
}

func NewGameLogic(s network.MessageSender, stopGamelogicChannel chan bool) *GameLogic {
	m := worldmap.New(20, 10)     //nolint:gomnd
	zombie := newZombie(10, 5, m) //nolint:gomnd

	return &GameLogic{
		msgSender:            s,
		stopGamelogicChannel: stopGamelogicChannel,
		generator:            newGenerator(zombie),
	}
}

func (l *GameLogic) Run() {
	fmt.Println("Starting to generate...")

	ticker := time.NewTicker(time.Second * 3) //nolint:gomnd
	defer ticker.Stop()

	for {
		select {
		case <-l.stopGamelogicChannel:
			log.Println("Zombie generator stopped.")
			return
		case <-ticker.C:
			log.Println("Sending message...")

			msg, err := l.generator.next()
			if err != nil {
				fmt.Println("could not generate next message: %w", err)
				return
			}

			err = l.msgSender.SendMsg(msg)
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
