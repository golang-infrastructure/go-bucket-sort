package bucket_sort

import (
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-maths"
	"github.com/golang-infrastructure/go-pointer"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// BucketFunc 为元素分配一个bucket，bucket用整数表示，这个bucket可以是正数也可以是负数，都能够兼容，
// 但是bucket的范围分布应该尽量密集，分布范围比较大的数据不适合使用桶排序
type BucketFunc[T any] func(index int, value T) int

// ------------------------------------------------ ---------------------------------------------------------------------

// Sort 对整数类型的数组进行升序排序
func Sort[T gtypes.Integer](slice []T, bucketLimit ...int) error {
	return SortByFunc(slice, func(index int, value T) int {
		return int(value)
	}, bucketLimit...)
}

// SortDesc 对整数类型的数组进行倒序排序
func SortDesc[T gtypes.Integer](slice []T, bucketLimit ...int) error {
	return SortByFunc(slice, func(index int, value T) int {
		return -int(value)
	}, bucketLimit...)
}

// ------------------------------------------------ ---------------------------------------------------------------------

// SortByFunc 根据桶函数对任意类型的切片进行桶排序
// slice: 要排序的切片
// bucketFunc: 对切片中的每一个元素，如何为其分配桶
// bucketLimit: 桶大小的限制
func SortByFunc[T any](slice []T, bucketFunc BucketFunc[T], bucketLimit ...int) error {

	// 桶数量限制
	allowBucketSize := DefaultBucketLimit
	if len(bucketLimit) > 0 {
		allowBucketSize = bucketLimit[0]
	}

	// 正数和负数分别搞一个计数，但是它们是共享bucket数量限制的
	positiveMin, positiveMax, negativeMin, negativeMax := findBoundary(slice, bucketFunc)
	var positiveBucket, negativeBucket [][]T
	if positiveMax != nil && positiveMin != nil {
		bucketSize := int(*positiveMax - *positiveMin + 1)
		if allowBucketSize < bucketSize {
			return ErrBucketLimit
		}
		positiveBucket = make([][]T, bucketSize)
		allowBucketSize -= bucketSize
	}
	if negativeMax != nil {
		bucketSize := int(maths.Abs(*negativeMin) - maths.Abs(*negativeMax) + 1)
		if allowBucketSize < bucketSize {
			return ErrBucketLimit
		}
		negativeBucket = make([][]T, bucketSize)
		allowBucketSize -= bucketSize
	}

	// 开始分配bucket计数
	// 负数把最大的那边与0对齐，能够节省掉一些bucket空间
	absNegativeMax := maths.Abs(pointer.FromPointerOrDefault(negativeMax, 0))
	for index, value := range slice {
		bucket := bucketFunc(index, value)
		if bucket >= 0 {
			// 正数
			bucket := bucket - *positiveMin
			positiveBucket[bucket] = append(positiveBucket[bucket], value)
		} else {
			// 负数
			bucket := maths.Abs(bucket) - absNegativeMax
			negativeBucket[bucket] = append(negativeBucket[bucket], value)
		}
	}

	// 统计结果
	index := 0
	// 先把负数放进去，要从后往前放，因为距离原点越近的值越大，因此要从距离原点较远的那个方向往原点走
	for bucket := len(negativeBucket) - 1; bucket >= 0; bucket-- {
		valueSlice := negativeBucket[bucket]
		for _, value := range valueSlice {
			slice[index] = value
			index++
		}
	}
	// 再把正数放进去
	for _, valueSlice := range positiveBucket {
		for _, value := range valueSlice {
			slice[index] = value
			index++
		}
	}
	return nil
}

// 找到正数和负数的边界值，确定bucket的时候要使用一个最小的区间
func findBoundary[T any](slice []T, bucketFunc BucketFunc[T]) (positiveMin, positiveMax, negativeMin, negativeMax *int) {
	for index, value := range slice {
		bucket := bucketFunc(index, value)
		if bucket >= 0 {
			// 正数
			if positiveMin == nil || bucket < *positiveMin {
				positiveMin = pointer.ToPointer(bucket)
			}
			if positiveMax == nil || bucket > *positiveMax {
				positiveMax = pointer.ToPointer(bucket)
			}
		} else {
			// 负数
			if negativeMin == nil || bucket < *negativeMin {
				negativeMin = pointer.ToPointer(bucket)
			}
			if negativeMax == nil || bucket > *negativeMax {
				negativeMax = pointer.ToPointer(bucket)
			}
		}
	}
	return
}

// ------------------------------------------------ ---------------------------------------------------------------------
