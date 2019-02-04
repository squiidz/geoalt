package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/squiidz/geoalt/geoaltsvc"

	"google.golang.org/grpc"
)

var (
	grpcPort = flag.Int("grpc", 9000, "GRPC port")
	fake     = flag.Bool("fake", false, "load fake data")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	srvr := New()
	if *fake {
		PopulateDB(srvr.db)
	}
	gs := grpc.NewServer()
	geoaltsvc.RegisterGeoAltServer(gs, srvr)
	gs.Serve(lis)
}
