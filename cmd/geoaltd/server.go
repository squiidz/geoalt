package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	h3 "github.com/squiidz/h3-go"

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

func (s Server) GeoFeed(feed pb.GeoAlt_GeoFeedServer) error {
	for {
		in, err := feed.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for {
			time.Sleep(time.Second * 5)
			out, err := s.GetAlerts(feed.Context(), in)
			if err != nil {
				return err
			}
			err = feed.Send(out)
			if err != nil {
				return err
			}
		}
	}
}

func (s Server) GetAlerts(context context.Context, req *pb.GetAlertsReq) (*pb.GetAlertsResp, error) {
	c, err := checkToken(context)
	if err != nil {
		return nil, err
	}
	cellID := h3.FromGeo(h3.GeoCoord{Latitude: req.Lat, Longitude: req.Lng}, s.CellLvl)

	var alerts []*pb.Alert
	user, err := s.db.UserStore.GetUser(c.ID)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	alertIDs := s.db.AlertStore.GetUserAlertIDs(uint64(cellID), user.ID)
	for _, aid := range alertIDs {
		alert, err := s.db.AlertStore.GetAlert(uint64(cellID), user.ID, uint32(aid))
		if err != nil {
			log.Println(err)
			continue
		}
		alerts = append(alerts, s.AlertToProto(alert))
	}

	log.Printf("Get %d Alerts for user %d at lat %f lng %f", len(alerts), user.ID, req.Lat, req.Lng)
	return &pb.GetAlertsResp{
		Alerts: alerts,
	}, nil
}

func (s Server) GetActiveAlerts(ctx context.Context, req *pb.GetAlertsReq) (*pb.GetAlertsResp, error) {
	var active []*pb.Alert
	alertResp, err := s.GetAlerts(ctx, req)
	if err != nil {
		return nil, err
	}
	uCellID := h3.FromGeo(h3.GeoCoord{Latitude: req.Lat, Longitude: req.Lng}, 15)
	for _, a := range alertResp.Alerts {
		if h3.H3Index(a.Cell.RealCell) == h3.ToParent(uCellID, int(a.Cell.Resolution)) {
			active = append(active, a)
		}
	}
	return &pb.GetAlertsResp{Alerts: active}, nil
}

func (s Server) AddAlert(context context.Context, req *pb.AddAlertReq) (*pb.AddAlertResp, error) {
	c, err := checkToken(context)
	if err != nil {
		return nil, err
	}

	alert := s.AlertFromProto(c.ID, req)
	err = s.db.AlertStore.Insert(alert)
	if err != nil {
		return &pb.AddAlertResp{Ok: false}, err
	}
	log.Printf("Creating Alert for user %d at lat %f lng %f", c.ID, req.Lat, req.Lng)
	return &pb.AddAlertResp{Ok: true}, nil
}
