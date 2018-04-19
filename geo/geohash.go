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
计算方法：https://www.cnblogs.com/LBSer/p/3310455.html
地球纬度区间是[-90,90]， 北京北海公园的纬度是39.928167，可以通过下面算法对纬度39.928167进行逼近编码:
1）区间[-90,90]进行二分为[-90,0),[0,90]，称为左右区间，可以确定39.928167属于右区间[0,90]，给标记为1；
2）接着将区间[0,90]进行二分为 [0,45),[45,90]，可以确定39.928167属于左区间 [0,45)，给标记为0；
3）递归上述过程39.928167总是属于某个区间[a,b]。随着每次迭代区间[a,b]总在缩小，并越来越逼近39.928167；
4）如果给定的纬度x（39.928167）属于左区间，则记录0，如果属于右区间则记录1，这样随着算法的进行会产生一个序列1011100，序列的长度跟给定的区间划分次数有关。
通过上述计算，纬度产生的编码为10111 00011，经度产生的编码为11010 01011。偶数位放经度，奇数位放纬度，把2串编码组合生成新串：11100 11101 00100 01111。
最后使用用0-9、b-z（去掉a, i, l, o）这32个字母进行base32编码，首先将11100 11101 00100 01111转成十进制，对应着28、29、4、15，十进制对应的编码就是wx4g。
*/
func GeoHashEncode(lat, lng float64, precision int) (string, *SquareDistrict) {
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
		} else { //奇数位放纬度
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
	for _, ch := range geohash {
		pos, ok := base32Pos[byte(ch)]
		if !ok {
			return ret
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
	up, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2+b.LatSpan(),
		(b.MinLng+b.MaxLng)/2,
		precision,
	)
	down, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2-b.LatSpan(),
		(b.MinLng+b.MaxLng)/2,
		precision,
	)
	left, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2,
		(b.MinLng+b.MaxLng)/2-b.LngSpan(),
		precision,
	)
	right, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2,
		(b.MinLng+b.MaxLng)/2+b.LngSpan(),
		precision,
	)

	//四个角的格子
	leftUp, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2+b.LatSpan(),
		(b.MinLng+b.MaxLng)/2-b.LngSpan(),
		precision,
	)
	leftDown, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2-b.LatSpan(),
		(b.MinLng+b.MaxLng)/2-b.LngSpan(),
		precision,
	)
	rightUp, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2+b.LatSpan(),
		(b.MinLng+b.MaxLng)/2+b.LngSpan(),
		precision,
	)
	rightDown, _ := GeoHashEncode(
		(b.MinLat+b.MaxLat)/2-b.LatSpan(),
		(b.MinLng+b.MaxLng)/2+b.LngSpan(),
		precision,
	)

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
