package utils

import (
	"math"
)

// 地球半径，单位米
const R = 6367000

// 返回的距离单位为米
func CalGeoDistance(latA, lonA, latB, lonB float64) float64 {
	c := math.Sin(latA)*math.Sin(latB)*math.Cos(lonA-lonB) + math.Cos(latA)*math.Cos(latB)
	return R * math.Acos(c) * math.Pi / 180
}
