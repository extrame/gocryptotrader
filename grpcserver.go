package main

import (
	"context"
	"log"
	"net"

	"github.com/extrame/gocryptotrader/exchanges/ticker"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	ig "github.com/extrame/gocryptotrader/grpc"
)

type Server struct {
}

func (s *Server) GetTickers(context.Context, *empty.Empty) (*ig.AllEnabledExchangeCurrencies, error) {
	all := GetAllActiveTickers()
	return &all, nil
}

func (s *Server) GetTicker(c context.Context, req *ig.SpecificTicker) (*ticker.Price, error) {
	return GetSpecificTicker(req.Currency, req.ExchangeName, req.AssetType)
}

//StartGrpcServer start the grpc server
func StartGrpcServer(addr string) error {
	var err error
	var lis net.Listener
	if lis, err = net.Listen("tcp", addr); err == nil {
		s := grpc.NewServer()
		ig.RegisterGoCryptoTraderServiceServer(s, &Server{})
		err = s.Serve(lis)
	}
	if err != nil {
		log.Fatal(err)
	}
	return err
}
