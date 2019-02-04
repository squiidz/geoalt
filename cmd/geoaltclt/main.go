package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "local/geoalt/geoaltsvc"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	grpcAddr = flag.String("grpcAddr", "localhost:9000", "grpc address and port")
	userID   = flag.Uint("userID", 1, "user id to use")
	lat      = flag.Float64("lat", 0, "latitude")
	lng      = flag.Float64("lng", 0, "longitude")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	gclt := pb.NewGeoAltClient(conn)

	fetch := &cobra.Command{
		Use:   "fetch the user messages",
		Short: "fetch the user messages",
		Long:  "fetch the user messages",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			uid, _ := cmd.Flags().GetUint32("userID")
			lat, _ := cmd.Flags().GetFloat64("lat")
			lng, _ := cmd.Flags().GetFloat64("lng")
			resp, err := gclt.GetAlert(context.Background(), &pb.GetAlertReq{
				UserId: uid,
				Lat:    lat,
				Lng:    lng,
			})
			if err != nil {
				log.Fatal(err)
			}

			for _, a := range resp.Alerts {
				fmt.Println(a.String())
			}
		},
	}
	fetch.Flags().Uint32("userID", 1, "-userID #id")
	fetch.Flags().Float64("lat", 0, "-lat 2.33")
	fetch.Flags().Float64("lng", 0, "-lng 3.44")

	create := &cobra.Command{
		Use:   "create a new message",
		Short: "create a new message",
		Long:  "create a new message",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			uid, _ := cmd.Flags().GetUint32("userID")
			lat, _ := cmd.Flags().GetFloat64("lat")
			lng, _ := cmd.Flags().GetFloat64("lng")
			msg, _ := cmd.Flags().GetString("msg")
			resp, err := gclt.CreateAlert(context.Background(), &pb.CreateReq{
				UserId:  uid,
				Lat:     lat,
				Lng:     lng,
				Message: msg,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp.GetStatus())
		},
	}
	create.Flags().Uint32("userID", 1, "-userID #id")
	create.Flags().Float64("lat", 0, "-lat 2.33")
	create.Flags().Float64("lng", 0, "-lng 3.44")
	create.Flags().String("msg", "", "-msg mesaage content")

	var rootCmd = &cobra.Command{Use: "geoclt"}
	rootCmd.AddCommand(fetch, create)
	rootCmd.Execute()
}
