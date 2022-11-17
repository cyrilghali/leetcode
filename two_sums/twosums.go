package main
import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    visited := make(map[int] int)
    for i:= 0; i < len(nums); i++ {
        diff := target - nums[i]
        if val, exist := visited[diff]; exist {
            return []int{val, i}
       }
        visited[nums[i]] = i
    }
    return nil;
}

func main(){
    nums := []int{2, 7, 9, 11} // init
    res := twoSum(nums, 9)
    fmt.Println(res)
}
