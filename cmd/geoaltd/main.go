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
	cellLvl  = flag.Int("cellLvl", 19, "S2 Cell level for indexing")
	dbpath   = flag.String("dbpath", ".", "database directory path")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	srvr := New(*dbpath, *cellLvl)
	if *fake {
		PopulateDB(srvr.db)
	}
	gs := grpc.NewServer()
	geoaltsvc.RegisterGeoAltServer(gs, srvr)
	gs.Serve(lis)
}
