package main

func moveZeroes(nums []int)  {
	j := 0
	for i, k := range nums {
		if k != 0 {
			if(i > j){
				nums[j] = nums[i];
				nums[i] = 0;
			}
			j++;
		}
	}
}