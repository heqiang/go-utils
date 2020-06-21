/*
 *    Copyright 2020 Chen Quan
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package group

type Grouper interface {
	Size() int
	Get(int) interface{}
	IsEmpty() bool
}

func By(grouper Grouper, functionBy func(obj interface{}) interface{}) (result map[interface{}][]interface{}) {
	result = make(map[interface{}][]interface{})
	if grouper.IsEmpty() {
		return
	}
	for i := 0; i < grouper.Size(); i++ {
		obj := grouper.Get(i)
		// 分组的key
		key := functionBy(obj)
		if _, ok := result[key]; !ok {
			result[key] = []interface{}{obj}
		} else {
			result[key] = append(result[key], obj)
		}
	}
	return
}

func BySlice(grouper []interface{}, functionBy func(obj interface{}) interface{}) (result map[interface{}][]interface{}) {
	result = make(map[interface{}][]interface{})
	if len(grouper) == 0 {
		return
	}
	for _, obj := range grouper {
		// 分组的key
		key := functionBy(obj)
		if _, ok := result[key]; !ok {
			result[key] = []interface{}{obj}
		} else {
			result[key] = append(result[key], obj)
		}
	}
	return
}

func BySliceCount(grouper []interface{}, functionBy func(obj interface{}) interface{}) (result map[interface{}]int) {
	result = map[interface{}]int{}
	if len(grouper) == 0 {
		return
	}
	for _, obj := range grouper {
		// 分组的key
		key := functionBy(obj)
		if _, ok := result[key]; !ok {
			result[key] = 1
		} else {
			result[key] += 1
		}
	}
	return
}

func ByCount(grouper Grouper, functionBy func(obj interface{}) interface{}) (result map[interface{}]int) {
	result = map[interface{}]int{}
	if grouper.IsEmpty() {
		return
	}
	for i := 0; i < grouper.Size(); i++ {
		obj := grouper.Get(i)
		// 分组的key
		key := functionBy(obj)
		if _, ok := result[key]; !ok {
			result[key] = 1
		} else {
			result[key] += 1
		}
	}
	return
}
