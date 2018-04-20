/**
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     geo
 * @date        2018-04-19 14:31
 */
package geo

import (
	"math"
)

const (
	//地球半径
	EARTH_RADIUS = 6378137
)

//计算两个经纬度间的中间位置
func MidPoint(lat1, lng1, lat2, lng2 float64) (float64, float64) {
	lat1Arc := lat1 * math.Pi / 180.0
	lat2Arc := lat2 * math.Pi / 180.0
	lng1Arc := lng1 * math.Pi / 180.0
	diffLng := (lng2 - lng1) * math.Pi / 180.0

	bx := math.Cos(lat2Arc) * math.Cos(diffLng)
	by := math.Cos(lat2Arc) * math.Sin(diffLng)

	lat3Rad := math.Atan2(
		math.Sin(lat1Arc)+math.Sin(lat2Arc),
		math.Sqrt(math.Pow(math.Cos(lat1Arc)+bx, 2)+math.Pow(by, 2)),
	)
	lng3Rad := lng1Arc + math.Atan2(by, math.Cos(lat1Arc)+bx)

	lat3 := lat3Rad * 180.0 / math.Pi
	lng3 := lng3Rad * 180.0 / math.Pi

	return lat3, lng3
}

//在指定距离、角度上，返回另一个经纬度坐标
//lat、lng：源经纬度
//dist：距离，单位米
//angle：角度，如"45"
func PointAtDistAndAngle(lat, lng, dist, angle float64) (float64, float64) {
	dr := dist / EARTH_RADIUS

	angle = angle * (math.Pi / 180.0)

	lat1 := lat * (math.Pi / 180.0)
	lng1 := lng * (math.Pi / 180.0)

	lat2_part1 := math.Sin(lat1) * math.Cos(dr)
	lat2_part2 := math.Cos(lat1) * math.Sin(dr) * math.Cos(angle)

	lat2 := math.Asin(lat2_part1 + lat2_part2)

	lng2_part1 := math.Sin(angle) * math.Sin(dr) * math.Cos(lat1)
	lng2_part2 := math.Cos(dr) - (math.Sin(lat1) * math.Sin(lat2))

	lng2 := lng1 + math.Atan2(lng2_part1, lng2_part2)
	lng2 = math.Mod(lng2+3*math.Pi, 2*math.Pi) - math.Pi

	lat2 = lat2 * (180.0 / math.Pi)
	lng2 = lng2 * (180.0 / math.Pi)
	return lat2, lng2
}

//计算地球上的曲线距离，返回值为米
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * float64(EARTH_RADIUS)
}
