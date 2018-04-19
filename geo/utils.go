/**
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     geo
 * @date        2018-04-19 14:31
 */
package geo

import (
	"math"
)

//计算地球上的曲线距离，返回值为米
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378137 //地球半径
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * float64(radius)
}
