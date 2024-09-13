package tools

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func JsontoStruct(jsonStr string, structure interface{}) {
	json.Unmarshal([]byte(jsonStr), &structure)
}

func StructToJson(jsonStr string, structure interface{}) string {
	jsonBytes, err := json.Marshal(structure)
	if err != nil {
		fmt.Println("序列化失败")
		return ""
	}
	return string(jsonBytes)
}

/*
@func 将map转为Struct
@param= mmap  需要转换的map[string]interface
structure   转换后的结构体指针
@returnerror       错误信息
暂不支持递归转换
*/
func MapToStruct(mmap map[string]interface{}, structure interface{}) (err error) {
	defer func() {
		if errs := recover(); errs != nil {
			err = errors.New("调用出错")
		}
	}()
	ptp := reflect.TypeOf(structure)
	pv := reflect.ValueOf(structure)
	switch ptp.Kind() {
	case reflect.Ptr:
		if ptp.Elem().Kind() == reflect.Struct {
			fmt.Println("sss")
			break
		} else {
			return errors.New("需要*struct类型，却传入*" + ptp.Elem().Kind().String() + "类型")
		}
	default:
		return errors.New("需要*struct类型，却传入" + ptp.Kind().String() + "类型")
	}
	tp := ptp.Elem()
	v := pv.Elem()
	num := tp.NumField()
	for i := 0; i < num; i++ {
		name := tp.Field(i).Name
		tag := tp.Field(i).Tag.Get("map")
		if len(tag) != 0 {
			name = tag
		}
		value, ok := mmap[name]
		if !ok {
			continue
		}
		//能够设置值，且类型相同
		if v.Field(i).CanSet() {
			if v.Field(i).Type() == reflect.TypeOf(value) {
				v.Field(i).Set(reflect.ValueOf(value))
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return nil
}

/*
@func:                  struct转map
@param:

	structure:          需要转换的结构体
	res:                转换后的map[string]interface{}
	recur:              是否递归转换（是否将内部嵌套的struct也转换为map）

@return error       错误信息
*/
func StructToMap(structure interface{}, res map[string]interface{}, recur bool) error {
	//检查是否为结构体
	tp := reflect.TypeOf(structure)
	v := reflect.ValueOf(structure)
	switch tp.Kind() {
	case reflect.Struct:
	default:
		return errors.New("所需参数为struct类型，却传入" + tp.Kind().String() + "类型")
	}
	//res := make(map[string]interface{})
	num := tp.NumField()
	for i := 0; i < num; i++ {
		key := tp.Field(i).Name
		tag := tp.Field(i).Tag.Get("map")
		if tag != "" && len(tag) != 0 {
			key = tag
		}
		//private类型的不能获取到值，所以直接排除private的字段
		if v.Field(i).CanInterface() {
			res[key] = v.Field(i).Interface()
			//----------------是否继续解析内层strcut-------------
			if recur {
				ztp := reflect.TypeOf(v.Field(i).Interface())
				//内层依然是一个struct，那么递归调用这个函数
				if ztp.Kind() == reflect.Struct {
					resMap := make(map[string]interface{})
					err := StructToMap(v.Field(i).Interface(), resMap, recur)
					if err != nil {
						res[key] = ""
						return err
					} else {
						res[key] = resMap
					}
				}
			}
			//-------------------------------
		} else {
			continue
		}
	}
	return nil
}

/*
*
@func:          判断某个值是否在某个切片中
@param:

	val:        要查找的值
	slice:      寻找的切片

@return

	int         查找到的下标，如果没有找到返回-1
	error       错误信息
*/
func InSlice(val interface{}, slice interface{}) (int, error) {
	//先将interface类型的slice转换为[]interface{}类型的slice1
	valof := reflect.ValueOf(slice)
	if valof.Kind() != reflect.Slice {
		return -1, errors.New("类型错误")
	}
	slice1 := make([]interface{}, 0)
	for i := 0; i < valof.Len(); i++ {
		slice1 = append(slice1, valof.Index(i).Interface())
	}
	for i, v := range slice1 {
		if v == val {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func Urlencode(url string) string {
	res := base64.URLEncoding.EncodeToString([]byte(url))
	res = strings.Replace(res, "+", "%2B", -1)
	res = strings.Replace(res, " ", "%20", -1)
	res = strings.Replace(res, "/", "%2F", -1)
	res = strings.Replace(res, "?", "%3F", -1)
	res = strings.Replace(res, "%", "%25", -1)
	res = strings.Replace(res, "#", "%23", -1)
	res = strings.Replace(res, "&", "%26", -1)
	res = strings.Replace(res, "=", "%3D", -1)
	return res
}
