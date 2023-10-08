package application

import "github.com/fbriansyah/micro-broker-service/internal/port"

type BrokerService struct {
	autClient     port.AuthAdapterPort
	sessionClient port.SessionPort
}
