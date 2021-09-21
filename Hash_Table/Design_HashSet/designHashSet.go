package main

type MyHashSet struct {
	buckets [1000][]int
}

func Constructor() MyHashSet {
	return *(new(MyHashSet))
}

func (hs *MyHashSet) hashf(key int) int {
	return key % 1000
}

func (hs *MyHashSet) Add(key int) {
	bucket_index := hs.hashf(key)
	for i := 0; i < len(hs.buckets[bucket_index]); i++ {
		if hs.buckets[bucket_index][i] == key {
			return
		}
	}
	hs.buckets[bucket_index] = append(hs.buckets[bucket_index], key)
}

func (hs *MyHashSet) Remove(key int) {
	nums := hs.buckets[hs.hashf(key)]
	for i := 0; i < len(nums); {
		if nums[i] == key {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		}
		if len(nums) > i && nums[i] != key {
			i++
		}
	}
	hs.buckets[hs.hashf(key)] = nums
}

func (hs *MyHashSet) Contains(key int) bool {
	bucket_index := hs.hashf(key)
	for i := 0; i < len(hs.buckets[bucket_index]); i++ {
		if hs.buckets[bucket_index][i] == key {
			return true
		}
	}
	return false
}
