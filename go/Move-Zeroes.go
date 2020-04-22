func moveZeroes(nums []int)  {
    j := 0
    for i := 0; i <len(nums);i++ {
        if nums[j] == 0 {
            nums = append(nums[:j], nums[j+1:]...)
            nums = append(nums, 0)
            j--
        }
        j++
    }
}