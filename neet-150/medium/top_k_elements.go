package medium

func TopKFrequent(nums []int, k int) []int {

	hashmap := make(map[int]int, len(nums)+1)

	for _, v := range nums {
		hashmap[v] += 1
	}

	bucket := make([][]int, len(nums)+1)

	for k, val := range hashmap {
		bucket[val] = append(bucket[val], k)
	}

	res := []int{}

	for i := len(bucket) - 1; i > 0 && k > 0; i-- {
		bucketLen := len(bucket[i])
		if bucketLen > 0 {
			k -= bucketLen
			res = append(res, bucket[i]...)
		}
	}

	return res
}
