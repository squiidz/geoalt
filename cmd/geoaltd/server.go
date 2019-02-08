package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	h3 "github.com/uber/h3-go"
	"google.golang.org/grpc/metadata"

	geo "github.com/squiidz/geoalt"
	pb "github.com/squiidz/geoalt/geoaltsvc"
)

type Server struct {
	CellLvl int
	db      *geo.DB
}

func New(dbpath string, cellLvl int) *Server {
	udb := fmt.Sprintf("%s/userdb", dbpath)
	adb := fmt.Sprintf("%s/alertdb", dbpath)
	db, err := geo.NewDB(udb, adb)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{db: db, CellLvl: cellLvl}
}

func (s Server) Register(context context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	_, err := s.db.UserStore.GetUserByEmail(req.Email)
	if err == nil {
		return nil, errors.New("User already exist")
	}
	user := &geo.User{
		Email:     req.Email,
		Password:  hashPassword(req.Password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
	}
	err = s.db.UserStore.Insert(user)
	if err != nil {
		return nil, err
	}
	token, err := genToken(user)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		Token: token,
	}, nil
}

func (s Server) Login(context context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	u, err := s.db.UserStore.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if u.Password != hashPassword(req.Password) {
		return nil, errors.New("Invalid Credentials")
	}
	token, err := genToken(u)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		Token: token,
	}, nil
}

func (s Server) GetAlert(context context.Context, req *pb.GetAlertReq) (*pb.GetAlertResp, error) {
	md, ok := metadata.FromIncomingContext(context)
	if !ok || len(md.Get("token")) <= 0 {
		return nil, errors.New("No Metadata")
	}
	c, ok := tokenIsValid(md.Get("token")[0])
	if !ok {
		return nil, errors.New("Invalid Token please login")
	}
	cellID := h3.FromGeo(h3.GeoCoord{Latitude: req.Lat, Longitude: req.Lng}, s.CellLvl)
	// cell := s2.CellFromLatLng(s2.LatLngFromDegrees(req.Lat, req.Lng))
	// cellID := cell.ID().Parent(s.CellLvl)

	var alerts []*pb.Alert
	user, err := s.db.UserStore.GetUser(c.ID)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	alertIDs := s.db.AlertStore.GetUserAlertIDs(uint64(cellID), user.ID)
	for _, aid := range alertIDs {
		alert, err := s.db.AlertStore.GetAlert(uint64(cellID), user.ID, uint32(aid))
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, &pb.Alert{
			Center: &pb.Coord{
				Lat: alert.Lat,
				Lng: alert.Lng,
			},
			Borders:   geoAltBorders(alert),
			Message:   alert.Message,
			Timestamp: alert.Timestamp,
		})
	}
	log.Printf("Get %d Alerts for user %d at lat %f lng %f", len(alerts), user.ID, req.Lat, req.Lng)
	return &pb.GetAlertResp{
		Alerts: alerts,
	}, nil
}

func (s Server) CreateAlert(context context.Context, req *pb.CreateAlertReq) (*pb.CreateAlertResp, error) {
	md, ok := metadata.FromIncomingContext(context)
	if !ok || len(md.Get("token")) <= 0 {
		return nil, errors.New("No Metadata")
	}
	c, ok := tokenIsValid(md.Get("token")[0])
	if !ok {
		return nil, errors.New("Invalid Token please login")
	}
	cellID := h3.FromGeo(h3.GeoCoord{Latitude: req.Lat, Longitude: req.Lng}, s.CellLvl)
	// cell := s2.CellFromLatLng(s2.LatLngFromDegrees(req.Lat, req.Lng))
	// cellID := cell.ID().Parent(s.CellLvl)

	u, err := s.db.UserStore.GetUser(c.ID)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	alert := &geo.Alert{
		CellID:    cellID,
		Lat:       req.Lat,
		Lng:       req.Lng,
		UserID:    u.ID,
		Message:   req.Message,
		Timestamp: time.Now().Format(time.RFC3339Nano),
		Ephemeral: req.Ephemeral,
	}

	err = s.db.AlertStore.Insert(alert)
	if err != nil {
		return &pb.CreateAlertResp{Ok: false}, err
	}
	log.Printf("Creating Alert for user %d at lat %f lng %f", req.UserId, req.Lat, req.Lng)
	return &pb.CreateAlertResp{Ok: true}, nil
}
