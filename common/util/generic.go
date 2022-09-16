package util

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

/**
 * @Author: zze
 * @Date: 2022/6/3 10:01
 * @Desc: 通用函数
 */

// UnmarshalYamlToMap
// desc: 反序列化 yaml 为 map[interface{}]interface{}
func UnmarshalYamlToMap(yamlStr string) (map[interface{}]interface{}, error) {
	var m map[interface{}]interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &m); err != nil {
		return nil, err
	}
	return m, nil
}

// MergeMap
// desc: 合并源 map 到目标 map
func MergeMap(dest, src map[interface{}]interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{}, len(dest))
	for k, v := range dest {
		out[k] = v
	}
	for k, v := range src {
		value := v
		if av, ok := out[k]; ok {
			if v, ok := v.(map[interface{}]interface{}); ok {
				if av, ok := av.(map[interface{}]interface{}); ok {
					out[k] = MergeMap(av, v)
				} else {
					out[k] = v
				}
			} else {
				out[k] = value
			}
		} else {
			out[k] = v
		}
	}
	return out
}

// MarshalObjectToYamlString
// desc: 序列化对象为 yaml 字符串
func MarshalObjectToYamlString(obj interface{}) (string, error) {
	bs, err := yaml.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("marshal obj [%#v] faild, err: %v", obj, err)
	}
	return strings.TrimSuffix(string(bs), "\n"), nil
}

// FindInIntSlice
// @desc 判断一个 int 值是否存在一个 int 切片中
// @param slice 切片
// @param val 值
// @return int 索引
// @return bool 是否存在
func FindInIntSlice(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Uint32SliceToIntSlice
// @desc 将一个 uint32 类型的切片转为 int 类型的切片
func Uint32SliceToIntSlice(s []uint32) []int {
	var ret []int
	for _, uintV := range s {
		ret = append(ret, int(uintV))
	}
	return ret
}

// IntSliceToUint32Slice
// @desc 将一个 int 类型的切片转为 uint32 类型的切片
func IntSliceToUint32Slice(s []int) []uint32 {
	var ret []uint32
	for _, intV := range s {
		ret = append(ret, uint32(intV))
	}
	return ret
}

// IntSliceContainsAll
// @desc 判断一个 int 切片中的元素是否全部存在于另一个 int 切片中
func IntSliceContainsAll(slice, contains []int) (*int, bool) {
	for _, v := range slice {
		if _, exists := FindInIntSlice(contains, v); !exists {
			return &v, false
		}
	}
	return nil, true
}

func MergeYaml(yamls ...string) (string, error) {
	finalMap := make(map[interface{}]interface{}, 0)

	for _, yamlStr := range yamls {
		tmpMap, err := UnmarshalYamlToMap(yamlStr)
		if err != nil {
			return "", err
		}

		finalMap = MergeMap(finalMap, tmpMap)
	}

	return MarshalObjectToYamlString(finalMap)
}

// DestInSlice 判断目标是否存在于指定切片中
func DestInSlice[T comparable](dest T, array []T) bool {
	for _, v := range array {
		if v == dest {
			return true
		}
	}
	return false
}
