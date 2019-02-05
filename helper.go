package geoalt

import (
	"encoding/binary"
	"math"
	"strconv"
)

func contains(arr []int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

func float64ToBytes(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func float64fromBytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func uint32ToBytes(u uint32) []byte {
	str := strconv.Itoa(int(u))
	return []byte(str)
}

func uint32FromBytes(bytes []byte) uint32 {
	x, _ := strconv.Atoi(string(bytes))
	return uint32(x)
}
