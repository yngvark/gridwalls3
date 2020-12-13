package mainhelp

import (
	"flag"
	"fmt"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/network"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/zombie"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type mainHelp struct {
	log *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *mainHelp {
	return &mainHelp{
		log: logger,
	}
}

func (m *mainHelp) SetupGame(allowedCorsOrigins map[string]bool) {
	broadcaster := network.NewBroadcaster()
	stopGamelogicChannel := make(chan bool)

	httpHandler := NewHTTPHandler(m.log, allowedCorsOrigins, broadcaster, stopGamelogicChannel)
	http.Handle("/zombie", httpHandler)

	var messageSender network.MessageSender = httpHandler
	gameLogic := zombie.NewGameLogic(m.log, messageSender, stopGamelogicChannel)

	broadcaster.AddListener(gameLogic)
}

func (m *mainHelp) HttpListen(port string, lg *zap.SugaredLogger) {
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		out := []byte("OK")
		_, err := writer.Write(out)

		if err != nil {
			m.log.Errorf("error when responding on /health: %s\n", err)
		}
	})

	serverAddr := flag.String("addr", fmt.Sprintf(":%s", port), "http service address")
	lg.Infof("Running on %s\n", *serverAddr)

	log.Fatal(http.ListenAndServe(*serverAddr, nil))
}
