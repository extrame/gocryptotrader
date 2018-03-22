package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	ig "github.com/extrame/gocryptotrader/grpc"
)

type Server struct {
}

func (s *Server) GetTickers(context.Context, *empty.Empty) (*ig.EnabledExchangeCurrencies, error) {
	currencies := GetAllActiveTickers()
	var eec = new(ig.EnabledExchangeCurrencies)
	eec.ExchangeCurrencies = make([]*ig.ExchangeCurrencies, len(currencies))
	for k, v := range currencies {
		var evs = make([]*ig.Value, len(v.ExchangeValues))
		for k1, v1 := range v.ExchangeValues {
			val := ig.Value{
				&ig.CurrencyPair{
					v1.Pair.Delimiter,
					v1.Pair.FirstCurrency.String(),
					v1.Pair.SecondCurrency.String(),
				},
				v1.CurrencyPair,
				v1.Last,
				v1.High,
				v1.Low,
				v1.Bid,
				v1.Ask,
				v1.Volume,
				v1.PriceATH,
			}
			evs[k1] = &val
		}
		var ec = ig.ExchangeCurrencies{v.ExchangeName, evs}
		eec.ExchangeCurrencies[k] = &ec
	}
	return eec, nil
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
