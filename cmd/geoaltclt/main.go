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
	username = flag.String("username", "jo@mail.com", "user email to use")
	password = flag.String("password", "Grenade4me", "user password")
	lat      = flag.Float64("lat", 0, "latitude")
	lng      = flag.Float64("lng", 0, "longitude")
)

type Client struct {
	token string
	pb.GeoAltClient
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clt := Client{GeoAltClient: pb.NewGeoAltClient(conn)}

	var rootCmd = &cobra.Command{Use: "geoclt"}
	rootCmd.AddCommand(clt.registerCommand(), clt.fetchCommand(), clt.createCommand())
	rootCmd.Execute()
}

func InAlert(lat, lng float64, alt *geoaltsvc.Alert) bool {
	a := h3.H3Index(alt.Cell.RealCell)
	u := h3.FromGeo(h3.GeoCoord{Latitude: lat, Longitude: lng}, int(alt.Cell.Resolution))
	fmt.Printf("\n-------------A: %d == U: %d---------------\n", a, u)
	return a == u
}

func (clt *Client) login(username, password string) {
	lresp, err := clt.Login(context.Background(), &pb.LoginReq{
		Email:    username,
		Password: password,
	})
	if err != nil {
		panic(err)
	}
	clt.token = lresp.Token
}

func (clt *Client) registerCommand() *cobra.Command {
	register := &cobra.Command{
		Use:   "register new user",
		Short: "register new user",
		Long:  "register new user",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			fname, _ := cmd.Flags().GetString("fname")
			lname, _ := cmd.Flags().GetString("lname")
			address, _ := cmd.Flags().GetString("address")
			ctx := context.Background()
			resp, err := clt.Register(ctx, &pb.RegisterReq{
				Email:     email,
				Password:  password,
				FirstName: fname,
				LastName:  lname,
				Address:   address,
			})
			if err != nil {
				panic(err)
			}
			fmt.Printf("[*] Register success [*]\nKey: %s\n", resp.Token)
		},
	}
	register.Flags().String("email", "jo@mail.com", "username email")
	register.Flags().String("password", "password123", "user password")
	register.Flags().String("fname", "jo", "user first name")
	register.Flags().String("lname", "doe", "user last name")
	register.Flags().String("address", "canada", "user address")

	return register
}

func (clt *Client) createCommand() *cobra.Command {
	create := &cobra.Command{
		Use:   "create a new message",
		Short: "create a new message",
		Long:  "create a new message",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			lat, _ := cmd.Flags().GetFloat64("lat")
			lng, _ := cmd.Flags().GetFloat64("lng")
			msg, _ := cmd.Flags().GetString("msg")
			eph, _ := cmd.Flags().GetBool("eph")
			size, _ := cmd.Flags().GetUint32("size")
			delay, _ := cmd.Flags().GetInt64("delay")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			clt.login(username, password)
			ctx := context.Background()
			ctx = metadata.AppendToOutgoingContext(ctx, "token", clt.token)
			resp, err := clt.AddAlert(ctx, &pb.AddAlertReq{
				Lat:        lat,
				Lng:        lng,
				Message:    msg,
				Ephemeral:  eph,
				Resolution: size,
				Delay:      delay,
			})
			if err != nil {
				log.Fatal(err)
			}
			if resp.Ok {
				fmt.Println("Message created successfully")
			}
		},
	}
	create.Flags().Float64("lat", 0, "-lat 2.33")
	create.Flags().Float64("lng", 0, "-lng 3.44")
	create.Flags().String("msg", "", "-msg message content")
	create.Flags().Bool("eph", false, "-eph true/false")
	create.Flags().Uint32("size", 7, "-size [7..15]")
	create.Flags().Int64("delay", 30, "delay in second")
	create.Flags().String("username", "jo@mail.com", "user email")
	create.Flags().String("password", "password123", "user password")

	return create
}

func (clt *Client) fetchCommand() *cobra.Command {
	fetch := &cobra.Command{
		Use:   "fetch the user messages",
		Short: "fetch the user messages",
		Long:  "fetch the user messages",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			lat, _ := cmd.Flags().GetFloat64("lat")
			lng, _ := cmd.Flags().GetFloat64("lng")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			clt.login(username, password)
			ctx := context.Background()
			ctx = metadata.AppendToOutgoingContext(ctx, "token", clt.token)
			resp, err := clt.GetAlert(ctx, &pb.GetAlertReq{
				Lat: lat,
				Lng: lng,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[!] %d alerts found in Sector [!]\n", len(resp.Alerts))
			for _, a := range resp.Alerts {
				if InAlert(lat, lng, a) {
					fmt.Printf("Message: %s\nLat: %f\nLng: %f\nBaseCell: %d\nIndexCell: %d\nRealCell: %d\nRes: %d\nTs: %d\n",
						a.Message, a.Center.Lat, a.Center.Lng, a.Cell.BaseCell, a.Cell.IndexCell, a.Cell.RealCell, a.Cell.Resolution, a.Timestamp)
					fmt.Println("Borders:", a.Borders)
				}
			}
		},
	}
	fetch.Flags().Float64("lat", 0, "-lat 2.33")
	fetch.Flags().Float64("lng", 0, "-lng 3.44")
	fetch.Flags().String("username", "jo@mail.com", "user email")
	fetch.Flags().String("password", "password123", "user password")

	return fetch
}
