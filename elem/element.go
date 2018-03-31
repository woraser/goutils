/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     elem
 * @date        2018-01-25 19:19
 */
package elem

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func MakeItemElem(d interface{}) ItemElem {
	return ItemElem{Data: d, RefVal: reflect.ValueOf(d)}
}

//基本的元素类型
type ItemElem struct {
	Data   interface{}   //数据元素
	RefVal reflect.Value //通过反射获取的value
}

//转换为bool类型
//如果是bool型则直接返回
func (ie ItemElem) ToBool() (bool, error) {
	if ie.RefVal.Kind() == reflect.Bool {
		return ie.RefVal.Bool(), nil
	}
	return false, fmt.Errorf("invalid type, data=%v", ie.Data)
}

//转换为int类型
func (ie ItemElem) ToInt() (int, error) {
	tmp, err := ie.ToInt64()
	if err != nil {
		return 0, err
	}
	return int(tmp), nil
}

//转换为int8类型
func (ie ItemElem) ToInt8() (int8, error) {
	tmp, err := ie.ToInt64()
	if err != nil {
		return 0, err
	}
	if tmp > math.MaxInt8 || tmp < math.MinInt8 {
		return 0, fmt.Errorf("toInt8 failed, %v overflow [math.MinInt8,math.MaxInt8]", ie.Data)
	}
	return int8(tmp), nil
}

//转换为int16类型
func (ie ItemElem) ToInt16() (int16, error) {
	tmp, err := ie.ToInt64()
	if err != nil {
		return 0, err
	}
	if tmp > math.MaxInt16 || tmp < math.MinInt16 {
		return 0, fmt.Errorf("toInt16 failed, %v overflow [math.MinInt16,math.MaxInt16]", ie.Data)
	}
	return int16(tmp), nil
}

//转换为int32类型
func (ie ItemElem) ToInt32() (int32, error) {
	tmp, err := ie.ToInt64()
	if err != nil {
		return 0, err
	}
	if tmp > math.MaxInt32 || tmp < math.MinInt32 {
		return 0, fmt.Errorf("toInt32 failed, %v overflow [math.MinInt32,math.MaxInt32]", ie.Data)
	}
	return int32(tmp), nil
}

//转换为int64类型
func (ie ItemElem) ToInt64() (int64, error) {
	switch ie.RefVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tmp := ie.RefVal.Int()
		return tmp, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tmp := ie.RefVal.Uint()
		if tmp > math.MaxInt64 {
			return 0, fmt.Errorf("invalid input,overflow MaxInt64,data=%v", ie.Data)
		}
		return int64(tmp), nil
	case reflect.Float32, reflect.Float64:
		tmpF := ie.RefVal.Float()
		tmp := int64(tmpF)
		return tmp, nil
	case reflect.String:
		tmp, err := strconv.ParseInt(ie.RefVal.String(), 10, 64)
		if err != nil {
			return 0, err
		}
		return tmp, nil
	default:
		return 0, fmt.Errorf("convert %v to int64 failed", ie.Data)
	}
}

//转换为uint类型
func (ie ItemElem) ToUint() (uint, error) {
	tmp, err := ie.ToUint64()
	if err != nil {
		return 0, err
	}
	return uint(tmp), nil
}

//转换为uint8类型
func (ie ItemElem) ToUint8() (uint8, error) {
	tmp, err := ie.ToUint64()
	if err != nil {
		return 0, err
	}
	if tmp > math.MaxUint8 {
		return 0, fmt.Errorf("toUint8 failed, %v overflow", ie.Data)
	}
	return uint8(tmp), nil
}

//转换为uint16类型
func (ie ItemElem) ToUint16() (uint16, error) {
	tmp, err := ie.ToUint64()
	if err != nil {
		return 0, err
	}
	if tmp > math.MaxUint16 {
		return 0, fmt.Errorf("toUint16 failed, %v overflow", ie.Data)
	}
	return uint16(tmp), nil
}

//转换为uint32类型
func (ie ItemElem) ToUint32() (uint32, error) {
	tmp, err := ie.ToUint64()
	if err != nil {
		return 0, err
	}
	if tmp > math.MaxUint32 {
		return 0, fmt.Errorf("toUint32 failed, %v overflow", ie.Data)
	}
	return uint32(tmp), nil
}

//转换为uint64类型
func (ie ItemElem) ToUint64() (uint64, error) {
	switch ie.RefVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tmp := ie.RefVal.Int()
		if tmp < 0 {
			return 0, fmt.Errorf("invalid failed,data=%v", ie.Data)
		}
		return uint64(tmp), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tmp := ie.RefVal.Uint()
		return tmp, nil
	case reflect.Float32, reflect.Float64:
		tmpF := ie.RefVal.Float()
		if tmpF < 0 {
			return 0, fmt.Errorf("invalid failed,data=%v", ie.Data)
		}
		tmp := uint64(tmpF)
		return tmp, nil
	case reflect.String:
		tmp, err := strconv.ParseUint(ie.RefVal.String(), 10, 64)
		if err != nil {
			return 0, err
		}
		return tmp, nil
	default:
		return 0, fmt.Errorf("convert %v to uint64 failed", ie.Data)
	}
}

//转换为string类型
func (ie ItemElem) ToString() string {
	switch ie.RefVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tmp := strconv.FormatInt(ie.RefVal.Int(), 10)
		return tmp
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tmp := strconv.FormatUint(ie.RefVal.Uint(), 10)
		return tmp
	case reflect.Float32, reflect.Float64:
		tmp := strconv.FormatFloat(ie.RefVal.Float(), 'f', -1, 64)
		return tmp
	default:
		return ie.RefVal.String()
	}
}

//转换为float类型
func (ie ItemElem) ToFloat32() (float32, error) {
	f64, err := ie.ToFloat64()
	if err != nil {
		return 0, err
	}
	if f64 > math.MaxFloat32 {
		return 0, fmt.Errorf("toFloat32 failed, %v overflow", ie.Data)
	}
	return float32(f64), nil
}

//转换为float64类型
func (ie ItemElem) ToFloat64() (float64, error) {
	switch ie.RefVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		tmp := ie.RefVal.Int()
		return float64(tmp), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		tmp := ie.RefVal.Uint()
		return float64(tmp), nil
	case reflect.Float32, reflect.Float64:
		tmp := ie.RefVal.Float()
		return tmp, nil
	case reflect.String:
		tmp, err := strconv.ParseFloat(ie.RefVal.String(), 64)
		if err != nil {
			return 0, err
		}
		return tmp, nil
	default:
		return 0, fmt.Errorf("convert %v to float64 failed", ie.Data)
	}
}

//转换为slice类型
//原始数据若为array/slice，则直接返回
//原始数据为map时只返回[]value
//原始数据若为数字、字符串等简单类型则将其放到slice中返回，即强制转为slice
//否则，报错
func (ie ItemElem) ToSlice() ([]ItemElem, error) {
	switch ie.RefVal.Kind() {
	case reflect.Slice, reflect.Array: //in为slice类型
		vlen := ie.RefVal.Len()
		ret := make([]ItemElem, vlen)
		for i := 0; i < vlen; i++ {
			ret[i] = MakeItemElem(ie.RefVal.Index(i).Interface())
		}
		return ret, nil
	case reflect.Map: //in为map类型
		var ret []ItemElem
		ks := ie.RefVal.MapKeys()
		for _, k := range ks {
			kiface := ie.RefVal.MapIndex(k).Interface()
			ret = append(ret, MakeItemElem(kiface))
		}
		return ret, nil
	case reflect.String: //字符串类型
		tmp := []byte(ie.RefVal.String())
		var ret []ItemElem
		for _, t := range tmp {
			ret = append(ret, MakeItemElem(t))
		}
		return ret, nil
	default: //其他的类型一律强制转为slice
		return []ItemElem{ie}, nil
	}
}

//转换为map类型
//如果原始数据是map则直接返回
//如果是json字符串则尝试去解析
//否则，报错
func (ie ItemElem) ToMap() (map[ItemElem]ItemElem, error) {
	ret := make(map[ItemElem]ItemElem)
	if ie.RefVal.Kind() == reflect.Map {
		ks := ie.RefVal.MapKeys()
		for _, k := range ks {
			kiface := MakeItemElem(k.Interface())
			viface := ie.RefVal.MapIndex(k).Interface()
			ret[kiface] = MakeItemElem(viface)
		}
		return ret, nil
	}
	if ie.RefVal.Kind() == reflect.String {
		str := ie.RefVal.String()
		var vmap interface{}
		err := json.Unmarshal([]byte(str), &vmap)
		if err != nil {
			return ret, err
		}
		inRefVal := reflect.ValueOf(vmap)
		if inRefVal.Kind() == reflect.Map {
			ks := inRefVal.MapKeys()
			for _, k := range ks {
				kiface := MakeItemElem(k.Interface())
				viface := inRefVal.MapIndex(k).Interface()
				ret[kiface] = MakeItemElem(viface)
			}
			return ret, nil
		}
	}
	return ret, fmt.Errorf("cannot convert %v to map", ie.Data)
}

//提取原始数据的长度，只有string/slice/map/array/chan
func (ie ItemElem) Len() (int, error) {
	switch ie.RefVal.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return ie.RefVal.Len(), nil
	default:
		return 0, fmt.Errorf("invalid type for len %v", ie.Data)
	}
}

//判断原始数据的类型是否为int
func (ie ItemElem) IsInt() bool {
	return ie.Kind() == reflect.Int
}

//判断原始数据的类型是否为int8
func (ie ItemElem) IsInt8() bool {
	return ie.Kind() == reflect.Int8
}

//判断原始数据的类型是否为int16
func (ie ItemElem) IsInt16() bool {
	return ie.Kind() == reflect.Int16
}

//判断原始数据的类型是否为int32
func (ie ItemElem) IsInt32() bool {
	return ie.Kind() == reflect.Int32
}

//判断原始数据的类型是否为int64
func (ie ItemElem) IsInt64() bool {
	return ie.Kind() == reflect.Int64
}

//判断原始数据的类型是否为uint
func (ie ItemElem) IsUint() bool {
	return ie.Kind() == reflect.Uint
}

//判断原始数据的类型是否为uint8
func (ie ItemElem) IsUint8() bool {
	return ie.Kind() == reflect.Uint8
}

//判断原始数据的类型是否为uint16
func (ie ItemElem) IsUint16() bool {
	return ie.Kind() == reflect.Uint16
}

//判断原始数据的类型是否为uint32
func (ie ItemElem) IsUint32() bool {
	return ie.Kind() == reflect.Uint32
}

//判断原始数据的类型是否为uint64
func (ie ItemElem) IsUint64() bool {
	return ie.Kind() == reflect.Uint64
}

//判断原始数据的类型是否为float32
func (ie ItemElem) IsFloat32() bool {
	return ie.Kind() == reflect.Float32
}

//判断原始数据的类型是否为float64
func (ie ItemElem) IsFloat64() bool {
	return ie.Kind() == reflect.Float64
}

//判断原始数据的类型是否为string
func (ie ItemElem) IsString() bool {
	return ie.Kind() == reflect.String
}

//判断原始数据的类型是否为slice
func (ie ItemElem) IsSlice() bool {
	return ie.Kind() == reflect.Slice
}

//判断原始数据的类型是否为map
func (ie ItemElem) IsMap() bool {
	return ie.Kind() == reflect.Map
}

//判断原始数据的类型是否为array
func (ie ItemElem) IsArray() bool {
	return ie.Kind() == reflect.Array
}

//判断原始数据的类型是否为chan
func (ie ItemElem) IsChan() bool {
	return ie.Kind() == reflect.Chan
}

//判断原始数据的类型是否为bool
func (ie ItemElem) IsBool() bool {
	return ie.Kind() == reflect.Bool
}

//是否为字符切片
func (ie ItemElem) IsByteSlice() bool {
	return reflect.TypeOf(ie.Data).String() == "[]uint8"
}

//是否为简单类型：int/uint/string/bool/float....
func (ie ItemElem) IsSimpleType() bool {
	return ie.IsInt() || ie.IsInt8() || ie.IsInt16() || ie.IsInt32() || ie.IsInt64() ||
		ie.IsUint() || ie.IsUint8() || ie.IsUint16() || ie.IsUint32() || ie.IsUint64() ||
		ie.IsString() || ie.IsFloat32() || ie.IsFloat64() || ie.IsBool()
}

//是否为复合类型：slice/array/map/chan
func (ie ItemElem) IsComplexType() bool {
	return ie.IsSlice() || ie.IsMap() || ie.IsArray() || ie.IsChan()
}

//原始数据的类型
func (ie ItemElem) Kind() reflect.Kind {
	return ie.RefVal.Kind()
}

//获取原始数据
func (ie ItemElem) RawData() interface{} {
	return ie.Data
}
