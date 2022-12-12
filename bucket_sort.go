package bucket_sort

import (
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-maths"
	"github.com/golang-infrastructure/go-pointer"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// Sort slice: 要进行桶排序的切片
// bucketLimit: 最大桶的限制
func Sort[T gtypes.Integer](slice []T, bucketLimit ...int) error {

	// 桶数量限制
	leftBucket := DefaultBucketLimit
	if len(bucketLimit) > 0 {
		leftBucket = bucketLimit[0]
	}

	// 正数和负数分别搞一个计数
	positiveMin, positiveMax, negativeMin, negativeMax := findBoundary(slice)
	var positiveBucket, negativeBucket []int
	if positiveMax != nil {
		bucketSize := int(*positiveMax - *positiveMin + 1)
		if leftBucket < bucketSize {
			return ErrBucketLimit
		}
		positiveBucket = make([]int, bucketSize)
		leftBucket -= bucketSize
	}
	if negativeMax != nil {
		bucketSize := int(maths.Abs(*negativeMin) - maths.Abs(*negativeMax) + 1)
		if leftBucket < bucketSize {
			return ErrBucketLimit
		}
		negativeBucket = make([]int, bucketSize)
		leftBucket -= bucketSize
	}

	// 开始分配bucket计数
	absNegativeMax := maths.Abs(pointer.FromPointerOrDefault(negativeMax, 0))
	for _, value := range slice {
		if value >= 0 {
			// 正数
			bucket := value - *positiveMin
			positiveBucket[bucket]++
		} else {
			// 负数
			bucket := maths.Abs(value) - absNegativeMax
			negativeBucket[bucket]++
		}
	}

	// 统计结果
	index := 0
	// 先把负数放进去，要从后往前放，因为距离原点越近的值越大，因此要从距离原点较远的那个方向往原点走
	for value := len(negativeBucket) - 1; value >= 0; value-- {
		count := negativeBucket[value]
		realValue := T(-1 * (value + int(absNegativeMax)))
		for count > 0 {
			slice[index] = realValue
			count--
			index++
		}
	}
	// 再把正数放进去
	for value, count := range positiveBucket {
		realValue := T(value) + *positiveMin
		for count > 0 {
			slice[index] = realValue
			count--
			index++
		}
	}
	return nil
}

// 找到正数和负数的边界值，确定bucket的时候要使用一个最小的区间
func findBoundary[T gtypes.Integer](slice []T) (positiveMin, positiveMax, negativeMin, negativeMax *T) {
	for _, value := range slice {
		if value >= 0 {
			// 正数
			if positiveMin == nil || value < *positiveMin {
				positiveMin = pointer.ToPointer(value)
			}
			if positiveMax == nil || value > *positiveMax {
				positiveMax = pointer.ToPointer(value)
			}
		} else {
			// 负数
			if negativeMin == nil || value < *negativeMin {
				negativeMin = pointer.ToPointer(value)
			}
			if negativeMax == nil || value > *negativeMax {
				negativeMax = pointer.ToPointer(value)
			}
		}
	}
	return
}

// ------------------------------------------------ ---------------------------------------------------------------------

//// SortByBucketFunc 根据自定义的bucket进行排序
//func SortByBucketFunc[T any](slice []T, bucketFunc func(value T) int) {
//
//}

// ------------------------------------------------ ---------------------------------------------------------------------
