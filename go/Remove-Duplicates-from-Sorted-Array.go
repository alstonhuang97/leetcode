func removeDuplicates(nums []int) int {
    m := map[int]bool{}
    for i := 0; i < len(nums); i++ {
        if _, ok := m[nums[i]]; ok {
            nums = append(nums[:i], nums[i+1:]...)
            i--
        } else {
            m[nums[i]] = true
        }
    }
    return len(nums)
}