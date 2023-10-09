package chi

import (
	"log"
	"net/http"

	"github.com/fbriansyah/micro-broker-service/internal/port"
)

type ChiAdapter struct {
	brokerService port.BrokerServicePort
}

type ChiAdapterConfig struct {
	ServerAddress string
}

func NewChiAdapter(brokerService port.BrokerServicePort) *ChiAdapter {

	return &ChiAdapter{
		brokerService: brokerService,
	}
}

func (adapter *ChiAdapter) Run(config ChiAdapterConfig) {
	srv := &http.Server{
		Addr:    config.ServerAddress,
		Handler: adapter.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
