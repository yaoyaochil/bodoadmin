package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// StringToUInt 将字符串转换为uint
func StringToUInt(str string) uint {
	i, _ := strconv.Atoi(str)
	return uint(i)
}

// IsToday 判断是否今日
func IsToday(timestamp int64) bool {
	inputTime := time.Unix(timestamp, 0)
	now := time.Now()

	inputYear, inputMonth, inputDay := inputTime.Date()
	nowYear, nowMonth, nowDay := now.Date()

	return inputYear == nowYear && inputMonth == nowMonth && inputDay == nowDay
}
