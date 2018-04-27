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
	bitsLen   = len(bits)
	base32    = []byte(BASE32)
	base32Pos = make(map[byte]int)
)

//初始化操作
func init() {
	for i, c := range base32 {
		base32Pos[c] = i
	}
}

//geoHash算法详见介绍：https://www.cnblogs.com/LBSer/p/3310455.html
func GeoHashEncode(lat, lng float64, precision int) (string, *GeoRectangle) {
	if lat < MIN_LATITUDE || lat > MAX_LATITUDE {
		return "", nil
	}
	if lng < MIN_LONGITUDE || lng > MAX_LONGITUDE {
		return "", nil
	}
	var buf bytes.Buffer
	var minLat, maxLat = MIN_LATITUDE, MAX_LATITUDE
	var minLng, maxLng = MIN_LONGITUDE, MAX_LONGITUDE
	var mid float64 = 0

	bit, ch, isEven := 0, 0, true
	for i := 0; i < precision; {
		if isEven { //偶数位放经度
			if mid = (minLng + maxLng) / 2; mid < lng {
				ch |= bits[bit]
				minLng = mid
			} else {
				maxLng = mid
			}
		} else { //奇数位处理纬度
			if mid = (minLat + maxLat) / 2; mid < lat {
				ch |= bits[bit]
				minLat = mid
			} else {
				maxLat = mid
			}
		}
		isEven = !isEven
		if bit < bitsLen-1 {
			bit++
		} else {
			buf.WriteByte(base32[ch])
			bit, ch = 0, 0
			i++
		}
	}

	b := &GeoRectangle{MinLat: minLat, MaxLat: maxLat, MinLng: minLng, MaxLng: maxLng}
	return buf.String(), b
}

//geoHash转换为坐标点，返回的是一个格子
//当对一对经纬度转为geoHash，再次转为经纬度时只能保证是同一个格子
//得到的只是格子四个角的经纬度坐标，并不能精确的还原之前的经纬度
func GeoHashDecode(geohash string) *GeoRectangle {
	ret := &GeoRectangle{}
	isEven := true
	lats := [2]float64{MIN_LATITUDE, MAX_LATITUDE}
	lngs := [2]float64{MIN_LONGITUDE, MAX_LONGITUDE}
	for _, ch := range geohash {
		pos, ok := base32Pos[byte(ch)]
		if !ok {
			return nil
		}
		for i := 0; i < bitsLen; i++ {
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
func GetNeighborsGeoCodes(lat, lng float64, precision int) []string {
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
