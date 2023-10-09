package main

import (
	grpcclient "github.com/fbriansyah/micro-broker-service/internal/adapter/client/grpc"
	"github.com/fbriansyah/micro-broker-service/internal/adapter/server/chi"
	"github.com/fbriansyah/micro-broker-service/internal/application"
	"github.com/fbriansyah/micro-broker-service/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %v", err.Error())
	}

	var opts []grpc.DialOption

	authConn, err := grpc.Dial(config.AuthServerAddress, opts...)
	if err != nil {
		log.Fatal().Msgf("cannot load authConn: %v", err.Error())
	}
	authClient := grpcclient.NewAuthClientAdapter(authConn)
	defer authConn.Close()

	sessionConn, err := grpc.Dial(config.SessionServerAddress, opts...)
	if err != nil {
		log.Fatal().Msgf("cannot load sessionConn: %v", err.Error())
	}
	sessionClient := grpcclient.NewSessionClientAdapter(sessionConn)
	defer sessionConn.Close()

	paymentConn, err := grpc.Dial(config.SessionServerAddress, opts...)
	if err != nil {
		log.Fatal().Msgf("cannot load paymentConn: %v", err.Error())
	}
	paymentClient := grpcclient.NewPaymentClientAdapter(paymentConn)
	defer paymentConn.Close()

	service := application.NewBrokerSerice(application.BrokerClientConfig{
		AuthClient:    authClient,
		SessionClient: sessionClient,
		PaymentClient: paymentClient,
	})

	log.Info().Msgf("Server Run at %v", config.HTTPServerAddress)

	httpAdapter := chi.NewChiAdapter(service)
	httpAdapter.Run(chi.ChiAdapterConfig{
		ServerAddress: config.HTTPServerAddress,
	})
}
