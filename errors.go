package bucket_sort

import "errors"

// ErrBucketLimit 达到最大bucket数量的限制，为了防止溢出，可以指定桶数量的限制
var ErrBucketLimit = errors.New("arrive max bucket limit")
