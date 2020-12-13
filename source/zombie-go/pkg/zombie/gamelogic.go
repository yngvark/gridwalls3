package zombie

import (
	"encoding/json"
	"go.uber.org/zap"
	"math/rand"
	"time"

	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/network"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/worldmap"
)

type GameLogic struct {
	log                  *zap.SugaredLogger
	msgSender            network.MessageSender
	stopGamelogicChannel chan bool
	generator            *Generator
}

func NewGameLogic(logger *zap.SugaredLogger, messageSender network.MessageSender, stopGamelogicChannel chan bool) *GameLogic {
	m := worldmap.New(20, 10)                                        //nolint:gomnd
	zombie := NewZombie("1", 10, 5, m, rand.New(rand.NewSource(45))) //nolint:gosec,gomnd

	return &GameLogic{
		log:                  logger,
		msgSender:            messageSender,
		stopGamelogicChannel: stopGamelogicChannel,
		generator:            NewGenerator(zombie),
	}
}

func (l *GameLogic) Run() {
	l.log.Info("Starting to generate...")

	ticker := time.NewTicker(time.Second * 1) //nolint:gomnd
	defer ticker.Stop()

	for {
		select {
		case <-l.stopGamelogicChannel:
			l.log.Info("Zombie generator stopped.")
			return
		case <-ticker.C:
			zombieMove, err := l.generator.Next()
			if err != nil {
				l.log.Info("could not generate next message: %w", err)
				return
			}

			zombieMoveJSON, err := json.Marshal(zombieMove)
			if err != nil {
				l.log.Info("could not marshal zombie move: %w", err)
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
	l.log.Info("Gamelogic received msg: %s\n", msg)

	if msg == "start" {
		go l.Run()
	}
}
