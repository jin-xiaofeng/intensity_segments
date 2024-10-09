package main

import (
	"fmt"
	"sort"
)

// IntensitySegments 用于管理不同区间的强度值。
type IntensitySegments struct {
	segments map[int]int
}

// NewIntensitySegments 初始化并返回一个新的 IntensitySegments 实例。
func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		segments: make(map[int]int),
	}
}

// Add 方法在指定范围内增加强度值。
func (is *IntensitySegments) Add(from, to, amount int) {
	is.updateSegments(from, to, func(prevValue int) int { return prevValue + amount })
}

// Set 方法在指定范围内设置强度值。
func (is *IntensitySegments) Set(from, to, amount int) {
	is.updateSegments(from, to, func(_ int) int { return amount })
}

// updateSegments 是一个辅助方法，通过传入的更新函数对指定范围内的强度值进行更新。
func (is *IntensitySegments) updateSegments(from, to int, updateFunc func(int) int) {
	// 确保起点和终点在 map 中已定义
	is.ensureBoundary(from)
	is.ensureBoundary(to)

	// 更新指定范围内的值
	for key := range is.segments {
		if key >= from && key < to {
			is.segments[key] = updateFunc(is.segments[key])
		}
	}

	// 合并相同强度值的连续区间
	is.mergeSegments()
}

// ensureBoundary 方法确保一个点在 map 中存在，如果不存在则添加，并继承之前的值。
func (is *IntensitySegments) ensureBoundary(point int) {
	if _, exists := is.segments[point]; !exists {
		prevValue := 0
		for k := range is.segments {
			if k < point && (prevValue == 0 || k > prevValue) {
				prevValue = k
			}
		}
		is.segments[point] = is.segments[prevValue]
	}
}

// mergeSegments 合并具有相同强度值的相邻区间。
func (is *IntensitySegments) mergeSegments() {
	keys := is.sortedKeys()
	prefixCleared := false // 清除前缀 0
	for i := 1; i < len(keys); i++ {
		// 可能存在多个连续的相同强度值
		for k := i; k < len(keys); k++ {
			prevKey, currentKey := keys[i-1], keys[k]
			if value := is.segments[prevKey]; value != 0 {
				prefixCleared = true
			}
			if prefixCleared == false {
				delete(is.segments, prevKey)
			}
			if is.segments[prevKey] == is.segments[currentKey] {
				delete(is.segments, currentKey)
				continue
			} else {
				i = k
				break
			}
		}
	}
}

// sortedKeys 返回按升序排列的 map 键。
func (is *IntensitySegments) sortedKeys() []int {
	keys := make([]int, 0, len(is.segments))
	for key := range is.segments {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func (is *IntensitySegments) ToString() {
	fmt.Printf("%v\n", is.dumps())
}

// String 方法将区间转换为字符串格式。
func (is *IntensitySegments) dumps() string {
	keys := is.sortedKeys()
	result := ""
	for _, key := range keys {
		result += fmt.Sprintf("[%d,%d],", key, is.segments[key])
	}
	if len(result) > 0 {
		result = result[:len(result)-1] // 移除末尾的逗号
	}
	return fmt.Sprintf("[%s]", result)
}
