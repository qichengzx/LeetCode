package main

func singleNumber(nums []int) int {
    result := 0
    length := len(nums)
    for i:=0; i< length; i++ {
        result ^= nums[i]
    }
    return result
}
