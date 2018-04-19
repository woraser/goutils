/**
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     geo
 * @date        2018-04-19 14:34
 */
package geo

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
)

type TestGeoHashPoint struct {
	Lat float64
	Lng float64
}

func TestGeoHashEncode(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------------start %s------------\n", f.Name())

	//一坨的地理位置信息
	points := []TestGeoHashPoint{
		{Lat: 31.473712, Lng: 120.270125},
		{Lat: 31.552752, Lng: 120.2372409},
		{Lat: 31.575529, Lng: 120.3026},
		{Lat: 31.676580, Lng: 120.357945},
		{Lat: 31.675095, Lng: 120.309268963407},
		{Lat: 31.580200, Lng: 120.35321},
		{Lat: 31.547996, Lng: 120.365778},
		{Lat: 31.549302, Lng: 120.367247421931},
		{Lat: 31.534714, Lng: 120.295785},
		{Lat: 31.541530, Lng: 120.3526},
		{Lat: 31.546304, Lng: 120.383614},
		{Lat: 31.479810, Lng: 120.449},
		{Lat: 31.542248, Lng: 120.378810983558},
	}
	for _, pInfo := range points {
		fmt.Fprintf(os.Stdout, "Lat:%f\tLng:%f\n", pInfo.Lat, pInfo.Lng)
		for i := 4; i <= 10; i++ {
			geo, square := GeoHashEncode(pInfo.Lat, pInfo.Lng, i)
			dist := int(square.BorderLength())
			//附近9个格子
			gs := GetNeighborsGeoCodes(pInfo.Lat, pInfo.Lng, i)
			dt := fmt.Sprintf("%vm x %vm", dist, dist)
			fmt.Fprintf(os.Stdout, "\tprecision:%d\tgeoHash:%-10s\tdist:%-14s\tNeighbors:%s\n", i, geo, dt, strings.Join(gs, ","))
		}
	}
	fmt.Printf("------------end %s------------\n", f.Name())
}

func TestGeoHashDecode(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------------start %s------------\n", f.Name())

	//一坨的地理位置信息
	points := []TestGeoHashPoint{
		{Lat: 31.675095455, Lng: 120.309268963407},
		{Lat: 31.54224843, Lng: 120.378810983558},
	}
	for _, pInfo := range points {
		fmt.Fprintf(os.Stdout, "Lat:%f\tLng:%f\n", pInfo.Lat, pInfo.Lng)
		for i := 4; i <= 10; i++ {
			geo, deGeo := GeoHashEncode(pInfo.Lat, pInfo.Lng, i)
			midLat, midLng := deGeo.MidLatLng()
			fmt.Fprintf(os.Stdout, "\tprecision:%d\tgeoHash:%-10s\t\n", i, geo)
			fmt.Fprintf(os.Stdout, "\t\tmaxLat:%v\tminLat:%v\n", deGeo.MaxLat, deGeo.MinLat)
			fmt.Fprintf(os.Stdout, "\t\tminLng:%v\tmaxLng:%v\n", deGeo.MinLng, deGeo.MaxLng)
			fmt.Fprintf(os.Stdout, "\t\tmidLat:%v\tmidLng:%v\n", midLat, midLng)
			fmt.Println("\tdeGeo")
			deGeo = GeoHashDecode(geo)
			fmt.Fprintf(os.Stdout, "\t\tmaxLat:%v\tminLat:%v\n", deGeo.MaxLat, deGeo.MinLat)
			fmt.Fprintf(os.Stdout, "\t\tminLng:%v\tmaxLng:%v\n", deGeo.MinLng, deGeo.MaxLng)
			fmt.Fprintf(os.Stdout, "\t\tmidLat:%v\tmidLng:%v\n", midLat, midLng)

		}
	}
	fmt.Printf("------------end %s------------\n", f.Name())
}
