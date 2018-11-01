package util

import (
	"reflect"
	"sort"
)

// GetAffPrice 获取每个转化的价格
func GetAffPrice(date, affName, clickType string) (price float32) {
	price = 6.8
	if affName == "" || affName == "test_affName" {
		price = 0.0
	}
	return
}

// GetOperatorPrice 获取运营商分成价格
func GetOperatorPrice(operator string) (price float32) {
	priceMap := map[string]float32{"20402": 2.199 * 1.17, "20404": 2.484 * 1.17, "20408": 2.365 * 1.17,
		"20416": 2.192 * 1.17}
	return priceMap[operator]
}

// Duplicate 列表去重
func Duplicate(a []string) (ret []string) {
	sort.Strings(a)
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).String())
	}
	return ret
}
