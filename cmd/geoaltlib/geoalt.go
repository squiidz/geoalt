package geoaltlib

import (
	"context"
	"errors"
	"log"

	pb "github.com/squiidz/geoalt/geoaltsvc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GeoClt struct {
	pb.GeoAltClient
	Token  string
	Alerts []*pb.Alert
}

func NewClient(geoAddr string) *GeoClt {
	conn, err := grpc.Dial(geoAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clt := pb.NewGeoAltClient(conn)
	return &GeoClt{GeoAltClient: clt}
}

func (g *GeoClt) Authenticate(email, password string) error {
	resp, err := g.Login(context.Background(), &pb.LoginReq{Email: email, Password: password})
	if err != nil {
		return err
	}
	g.Token = resp.Token
	return nil
}

func (g *GeoClt) GetAlerts(lat, lng float64) error {
	if g.Token == "" {
		return errors.New("You need to be authenticated")
	}

	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", g.Token)
	resp, err := g.GetAlert(ctx, &pb.GetAlertReq{
		Lat: lat,
		Lng: lng,
	})
	if err != nil {
		return err
	}
	g.Alerts = resp.Alerts
	return nil
}

