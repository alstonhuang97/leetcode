func singleNumber(nums []int) int {
    m := map[int]struct{}{}
    for _, num := range nums {
        if _, ok := m[num]; ok {
            delete(m, num)
        } else {
            m[num] = struct{}{}
        }
    }
    for k, _ := range m {
        return k
    } 
    return -1
}