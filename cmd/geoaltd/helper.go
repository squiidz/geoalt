package main

import (
	"crypto/sha256"
	"errors"
	"log"
	"time"

	pb "github.com/squiidz/geoalt/geoaltsvc"
	h3 "github.com/uber/h3-go"

	jwt "github.com/dgrijalva/jwt-go"
	geo "github.com/squiidz/geoalt"
)

const Secret = "somenotsosecretsecret"

type Claim struct {
	ID    uint32 `json:"id"`
	Email string `json:"email"`
}

func (c *Claim) Valid() error {
	if c.Email != "" {
		return nil
	}
	return errors.New("Invalid claim")
}

func hashPassword(p string) string {
	h := sha256.New()
	h.Write([]byte(p))
	return string(h.Sum(nil))
}

func genToken(u *geo.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claim{
		ID:    u.ID,
		Email: u.Email,
	})
	signToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return signToken, nil
}

func tokenIsValid(t string) (*Claim, bool) {
	var claim Claim
	_, err := jwt.ParseWithClaims(t, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		log.Println(err)
		return &claim, false
	}
	return &claim, true
}

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
	}
}

func (s *Server) AlertFromProto(userID uint32, req *pb.CreateAlertReq) *geo.Alert {
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
