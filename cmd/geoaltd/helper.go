package main

import (
	"time"

	pb "github.com/squiidz/geoalt/geoaltsvc"
	h3 "github.com/uber/h3-go"

	geo "github.com/squiidz/geoalt"
)

func geoAltBorders(a *geo.Alert) []*pb.Coord {
	var coords []*pb.Coord
	borders := a.Borders()
	for _, b := range borders {
		coords = append(coords, &pb.Coord{
			Lat: b.Lat,
			// need to substract 360 from longitude due to a H3-go bug
			Lng: b.Lng - 360,
		})
	}
	return coords
}

func (s *Server) AlertToProto(alert *geo.Alert) *pb.Alert {
	return &pb.Alert{
		Center: &pb.Coord{
			Lat: alert.Coord.Lat,
			Lng: alert.Coord.Lng,
		},
		Borders: geoAltBorders(alert),
		Cell: &pb.Cell{
			BaseCell:   uint64(alert.Cell.Base),
			IndexCell:  uint64(h3.ToParent(h3.H3Index(alert.Cell.Base), s.CellLvl)),
			RealCell:   uint64(alert.Cell.Real),
			Resolution: alert.Cell.Resolution,
		},
		Message:   alert.Message,
		Timestamp: alert.Timestamp,
		Delay:     alert.Delay,
		ReadAt:    alert.ReadAt,
	}
}

func (s *Server) AlertFromProto(userID uint32, req *pb.AddAlertReq) *geo.Alert {
	cellID := h3.FromGeo(h3.GeoCoord{Latitude: req.Lat, Longitude: req.Lng}, s.CellLvl)
	return &geo.Alert{
		Cell: geo.Cell{
			Base:       h3.FromGeo(h3.GeoCoord{Latitude: req.Lat, Longitude: req.Lng}, 15),
			Index:      cellID,
			Real:       h3.FromGeo(h3.GeoCoord{Latitude: req.Lat, Longitude: req.Lng}, int(req.Resolution)),
			Resolution: req.Resolution,
		},
		Coord: geo.Coord{
			Lat: req.Lat,
			Lng: req.Lng,
		},
		UserID:    userID,
		Message:   req.Message,
		Timestamp: time.Now().Unix(),
		Delay:     req.Delay,
		Ephemeral: req.Ephemeral,
	}
}
