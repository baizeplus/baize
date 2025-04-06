package sliceUtils

import (
	"baize/app/baize"
)

// HasDuplicates 查看是否有重复的元素
func HasDuplicates[T comparable](slice []T) bool {
	s := baize.Set[T]{}
	for _, v := range slice {
		if s.Contains(v) {
			return true
		}
		s.Add(v)
	}
	return false
}

// Difference 获取两个数组的差异数据
func Difference[T comparable](slice1, slice2 []T) (s1, s2 []T) {
	set1 := baize.NewSet(slice1)
	set1.RemoveAll(slice2...)
	set2 := baize.NewSet(slice2)
	set2.RemoveAll(slice1...)
	return set1.ToSlice(), set2.ToSlice()
}

// Intersection 获取两个数组的交集
func Intersection[T comparable](a, b []T) []T {
	sa := baize.NewSet(a)
	intersect := make([]T, 0)
	for _, t := range b {
		if sa.Contains(t) {
			intersect = append(intersect, t)
		}
	}
	return intersect
}

// Union 函数返回两个整数数组的并集，并去除重复元素。
func Union[T comparable](a, b []T) []T {
	sa := baize.NewSet(a)
	sa.AddAll(b...)
	return sa.ToSlice()
}

// AddUnique 函数负责将新元素添加到切片中，并确保没有重复元素。
func AddUnique[T comparable](a []T, b ...T) []T {
	sa := baize.NewSet(a)
	sa.AddAll(b...)
	return sa.ToSlice()
}

func FindIndex[T comparable](a T, b []T) int {
	for i, t := range b {
		if t == a {
			return i
		}
	}
	return -1
}
