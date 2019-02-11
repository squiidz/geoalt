package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	h3 "github.com/uber/h3-go"

	"google.golang.org/grpc/metadata"

	"github.com/squiidz/geoalt/geoaltsvc"
	pb "github.com/squiidz/geoalt/geoaltsvc"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	grpcAddr = flag.String("grpcAddr", "localhost:9000", "grpc address and port")
	userID   = flag.Uint("userID", 1, "user id to use")
	lat      = flag.Float64("lat", 0, "latitude")
	lng      = flag.Float64("lng", 0, "longitude")
)

type Coord struct {
	Lat float64
	Lng float64
}

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
			token, _ := cmd.Flags().GetString("token")
			ctx := context.Background()
			ctx = metadata.AppendToOutgoingContext(ctx, "token", token)
			resp, err := gclt.GetAlert(ctx, &pb.GetAlertReq{
				UserId: uid,
				Lat:    lat,
				Lng:    lng,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[!] %d alerts found in Sector [!]\n", len(resp.Alerts))
			for _, a := range resp.Alerts {
				if InAlert(&Coord{lat, lng}, a) {
					fmt.Println("----------------------------")
					fmt.Printf("Message: %s\nLat: %f\nLng: %f\nMinCell: %d\nRes: %d\nTs: %s\n", a.Message, a.Center.Lat, a.Center.Lng, a.Size.Cell, a.Size.Resolution, a.Timestamp)
					fmt.Println("Borders:", a.Borders)
				}
			}
		},
	}
	fetch.Flags().Uint32("userID", 1, "-userID #id")
	fetch.Flags().Float64("lat", 0, "-lat 2.33")
	fetch.Flags().Float64("lng", 0, "-lng 3.44")
	fetch.Flags().String("token", "", "-token tokenString")

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
			token, _ := cmd.Flags().GetString("token")
			eph, _ := cmd.Flags().GetBool("eph")
			size, _ := cmd.Flags().GetUint32("size")
			ctx := context.Background()
			ctx = metadata.AppendToOutgoingContext(ctx, "token", token)
			resp, err := gclt.CreateAlert(ctx, &pb.CreateAlertReq{
				UserId:     uid,
				Lat:        lat,
				Lng:        lng,
				Message:    msg,
				Ephemeral:  eph,
				Resolution: size,
			})
			if err != nil {
				log.Fatal(err)
			}
			if resp.Ok {
				fmt.Println("Message created successfully")
			}
		},
	}
	create.Flags().Uint32("userID", 1, "-userID #id")
	create.Flags().Float64("lat", 0, "-lat 2.33")
	create.Flags().Float64("lng", 0, "-lng 3.44")
	create.Flags().String("msg", "", "-msg message content")
	create.Flags().String("token", "", "-token tokenString")
	create.Flags().Bool("eph", false, "-eph true/false")
	create.Flags().Uint32("size", 7, "-size [7..15]")

	var rootCmd = &cobra.Command{Use: "geoclt"}
	rootCmd.AddCommand(fetch, create)
	rootCmd.Execute()
}

func InAlert(coord *Coord, alt *geoaltsvc.Alert) bool {
	a := h3.ToParent(h3.H3Index(alt.Size.Cell), int(alt.Size.Resolution))
	u := h3.FromGeo(h3.GeoCoord{Latitude: coord.Lat, Longitude: coord.Lng}, int(alt.Size.Resolution))
	return a == u
}
