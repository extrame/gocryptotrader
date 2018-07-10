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

type utServer struct {
	server   ig.GoCryptoTraderService_UpdateTickerServer
	cancelCh chan bool
}

type Server struct {
	utServers []utServer
}

func (s *Server) GetTickers(context.Context, *empty.Empty) (*ig.AllEnabledExchangeCurrencies, error) {
	all := GetAllActiveTickers()
	return &all, nil
}

func (s *Server) GetTicker(c context.Context, req *ig.SpecificTicker) (*ticker.Price, error) {
	return GetSpecificTicker(req.Currency, req.ExchangeName, req.AssetType)
}

func (s *Server) UpdateTicker(empty *empty.Empty, server ig.GoCryptoTraderService_UpdateTickerServer) error {
	var c = make(chan bool)
	s.utServers = append(s.utServers, utServer{
		server:   server,
		cancelCh: c,
	})
	select {
	case <-c:
		return nil
	}
	return nil
}

//StartGrpcServer start the grpc server
func StartGrpcServer(addr string) (server *Server, err error) {
	var lis net.Listener
	if lis, err = net.Listen("tcp", addr); err == nil {
		s := grpc.NewServer()
		server = new(Server)
		ig.RegisterGoCryptoTraderServiceServer(s, server)
		go func() {
			err := s.Serve(lis)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	return
}
