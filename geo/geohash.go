/**
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     geo
 * @date        2018-04-19 14:32
 */
package geo

import (
	"bytes"
)

const (
	BASE32                = "0123456789bcdefghjkmnpqrstuvwxyz"
	MAX_LATITUDE  float64 = 90
	MIN_LATITUDE  float64 = -90
	MAX_LONGITUDE float64 = 180
	MIN_LONGITUDE float64 = -180
)

var (
	bits      = []int{16, 8, 4, 2, 1}
	base32    = []byte(BASE32)
	base32Pos = make(map[byte]int)
)

//初始化操作
func init() {
	for i, c := range base32 {
		base32Pos[c] = i
	}
}

//一个正方形区域，用左上角、右下角的经纬度表示
type SquareDistrict struct {
	MinLat float64 //最小纬度
	MinLng float64 //最小经度
	MaxLat float64 //最大纬度
	MaxLng float64 //最大经度
}

//经度方向的跨度
func (sd *SquareDistrict) LngSpan() float64 {
	return sd.MaxLng - sd.MinLng
}

//纬度方向的跨度
func (sd *SquareDistrict) LatSpan() float64 {
	return sd.MaxLat - sd.MinLat
}

//获取中心点坐标
func (sd *SquareDistrict) MidLatLng() (float64, float64) {
	return (sd.MaxLat + sd.MinLat) / 2.0, (sd.MaxLng + sd.MinLng) / 2.0
}

//正方形的边长
func (sd *SquareDistrict) BorderLength() float64 {
	return EarthDistance(sd.MaxLat, sd.MaxLng, sd.MaxLat, sd.MinLng)
}

/**
 根据指定精度返回GeoHash编码及该点所在区域，如下各精度对应的距离范围
 Lat:31.546095	Lng:120.254508
	precision:4		geoHash:wtte      	dist:33318m x 33318m
	precision:5		geoHash:wtte2     	dist:4168m x 4168m
	precision:6		geoHash:wtte2q    	dist:1042m x 1042m
	precision:7		geoHash:wtte2qy   	dist:130m x 130m
	precision:8		geoHash:wtte2qy9  	dist:32m x 32m
	precision:9		geoHash:wtte2qy9m 	dist:4m x 4m
	precision:10	geoHash:wtte2qy9m4	dist:1m x 1m
*/
func GeoHashEncode(lat, lng float64, precision int) (string, *SquareDistrict) {
	if lat < MIN_LATITUDE || lat > MAX_LATITUDE {
		return "", &SquareDistrict{}
	}
	if lng < MIN_LONGITUDE || lng > MAX_LONGITUDE {
		return "", &SquareDistrict{}
	}
	var buf bytes.Buffer
	var minLat, maxLat = MIN_LATITUDE, MAX_LATITUDE
	var minLng, maxLng = MIN_LONGITUDE, MAX_LONGITUDE
	var mid float64 = 0

	bit, ch, length, isEven := 0, 0, 0, true
	for length < precision {
		if isEven { //偶数位和奇数位处理不一样
			if mid = (minLng + maxLng) / 2; mid < lng {
				ch |= bits[bit]
				minLng = mid
			} else {
				maxLng = mid
			}
		} else {
			if mid = (minLat + maxLat) / 2; mid < lat {
				ch |= bits[bit]
				minLat = mid
			} else {
				maxLat = mid
			}
		}
		isEven = !isEven
		if bit < 4 {
			bit++
		} else {
			buf.WriteByte(base32[ch])
			length, bit, ch = length+1, 0, 0
		}
	}

	b := &SquareDistrict{MinLat: minLat, MaxLat: maxLat, MinLng: minLng, MaxLng: maxLng}
	return buf.String(), b
}

//geoHash转换为坐标点，返回的是一个格子
//当对一对经纬度转为geoHash，再次转为经纬度时只能保证是同一个格子，并不能精确的还原经纬度
func GeoHashDecode(geohash string) *SquareDistrict {
	ret := &SquareDistrict{}
	isEven := true
	lats := [2]float64{MIN_LATITUDE, MAX_LATITUDE}
	lngs := [2]float64{MIN_LONGITUDE, MAX_LONGITUDE}
	for _, c := range geohash {
		pos, ok := base32Pos[byte(c)]
		if !ok {
			return ret
		}
		for i := 0; i < 5; i++ {
			ti := pos & bits[i]
			if ti > 0 {
				ti = 0
			} else {
				ti = 1
			}
			if isEven {
				lngs[ti] = (lngs[0] + lngs[1]) / 2.0
			} else {
				lats[ti] = (lats[0] + lats[1]) / 2.0
			}
			isEven = !isEven
		}
	}
	ret.MinLat = lats[0]
	ret.MaxLat = lats[1]
	ret.MinLng = lngs[0]
	ret.MaxLng = lngs[1]
	return ret
}

//计算给定的经纬度点在指定精度下周围8个区域的geoHash编码，包括自身，一共9个点
func GetNeighborsGeoHash(lat, lng float64, precision int) []string {
	if lat < MIN_LATITUDE || lat > MAX_LATITUDE {
		return []string{}
	}
	if lng < MIN_LONGITUDE || lng > MAX_LONGITUDE {
		return []string{}
	}
	geoHashList := make([]string, 9)

	//自身的区域
	cur, b := GeoHashEncode(lat, lng, precision)
	geoHashList[0] = cur

	//上下左右四个格子
	up, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2+b.LatSpan(), (b.MinLng+b.MaxLng)/2, precision)
	down, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2-b.LatSpan(), (b.MinLng+b.MaxLng)/2, precision)
	left, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2, (b.MinLng+b.MaxLng)/2-b.LngSpan(), precision)
	right, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2, (b.MinLng+b.MaxLng)/2+b.LngSpan(), precision)

	//四个角的格子
	leftUp, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2+b.LatSpan(), (b.MinLng+b.MaxLng)/2-b.LngSpan(), precision)
	leftDown, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2-b.LatSpan(), (b.MinLng+b.MaxLng)/2-b.LngSpan(), precision)
	rightUp, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2+b.LatSpan(), (b.MinLng+b.MaxLng)/2+b.LngSpan(), precision)
	rightDown, _ := GeoHashEncode((b.MinLat+b.MaxLat)/2-b.LatSpan(), (b.MinLng+b.MaxLng)/2+b.LngSpan(), precision)

	//八个格子赋值
	geoHashList[1] = up
	geoHashList[2] = down
	geoHashList[3] = left
	geoHashList[4] = right
	geoHashList[5] = leftUp
	geoHashList[6] = leftDown
	geoHashList[7] = rightUp
	geoHashList[8] = rightDown
	return geoHashList
}
