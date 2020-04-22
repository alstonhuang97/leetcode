func missingNumber(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    n := len(nums)
    return (1 + n) * n / 2 - sum
}