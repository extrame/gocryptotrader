package main

import (
	"context"
	"flag"
	"log"

	"github.com/extrame/gocryptotrader/config"
	ig "github.com/extrame/gocryptotrader/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "", "config file path, if not specified, will use config.json as default")
	flag.Parse()

	if configFile == "" {
		configFile = config.ConfigFile
	}

	cfg := config.GetConfig()
	err := cfg.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config file(%s): %s", configFile, err)
	}

	log.Printf("Connecting to websocket host: %s", cfg.Grpc.ListenAddress)

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(cfg.Grpc.ListenAddress, grpc.WithInsecure())

	if err != nil {
		log.Println("Unable to connect to grpc server", err)
		return
	}
	log.Println("Connected to grpc!")

	client := ig.NewGoCryptoTraderServiceClient(conn)

	log.Println("Getting tickers..")
	curres, err := client.GetTickers(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Got tickers!", curres)
}
