/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package slice

// Remove 移除首次匹配到的 item 元素
func Remove[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

// IsItemInSlice 判断item是否在slice中
func IsItemInSlice[T comparable](l []T, item T) bool {
	for _, i := range l {
		if i == item {
			return true
		}
	}
	return false
}

// Unique 去重
func Unique[T comparable](list []T) []T {
	uniqueMap := make(map[T]struct{})
	uniqueList := make([]T, 0)
	for _, item := range list {
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = struct{}{}
			uniqueList = append(uniqueList, item)
		}
	}

	return uniqueList
}

// Split list to array of lists with specified length.
func Split[T any](list []T, length int) [][]T {
	listLen := len(list)

	lists := make([][]T, 0)
	if length <= 0 || listLen == 0 {
		return lists
	}

	for i := 0; i < listLen; i += length {
		if (i + length) >= listLen {
			lists = append(lists, list[i:listLen])
		} else {
			lists = append(lists, list[i:i+length])
		}
	}
	return lists
}
