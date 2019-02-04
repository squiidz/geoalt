package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/geo/s2"

	geo "local/geoalt"
	pb "local/geoalt/geoaltsvc"
)

const cellLevel = 19

type Server struct {
	db *geo.DB
}

func New() *Server {
	db, err := geo.NewDB("badger")
	if err != nil {
		log.Fatal(err)
	}
	return &Server{db: db}
}

func (s Server) GetAlert(context context.Context, req *pb.GetAlertReq) (*pb.GetAlertResp, error) {
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(req.Lat, req.Lng))
	cellID := cell.ID().Parent(cellLevel)

	var alerts []*pb.Alert
	alertIDs := s.db.GetUserAlertIDs(uint64(cellID), req.UserId)
	for _, aid := range alertIDs {
		alert, err := s.db.GetAlert(uint64(cellID), uint32(aid))
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, &pb.Alert{
			Lat:       alert.Lat,
			Lng:       alert.Lng,
			Message:   alert.Message,
			Timestamp: alert.Timestamp,
		})
	}
	return &pb.GetAlertResp{
		Alerts: alerts,
	}, nil
}

func (s Server) CreateAlert(context context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	cell := s2.CellFromLatLng(s2.LatLngFromDegrees(req.Lat, req.Lng))
	cellID := cell.ID().Parent(cellLevel)

	alert := &geo.Alert{
		CellID:    cellID,
		Lat:       req.Lat,
		Lng:       req.Lng,
		UserID:    req.UserId,
		Message:   req.Message,
		Timestamp: time.Now().Format(time.RFC3339Nano),
	}

	err := s.db.InsertAlert(alert)
	if err != nil {
		return &pb.CreateResp{
			Status: &pb.CreateResp_Error{Error: err.Error()},
		}, err
	}

	return &pb.CreateResp{
		Status: &pb.CreateResp_Ok{Ok: true},
	}, nil
}
