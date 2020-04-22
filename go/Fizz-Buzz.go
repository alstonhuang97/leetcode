func fizzBuzz(n int) []string {
    f := []string{"1"}
    for i := 2; i <= n; i++ {
        if i % 15 == 0 {
            f = append(f, "FizzBuzz")
        } else if i % 5 == 0 {
            f = append(f, "Buzz")
        } else if i % 3 == 0 {
            f = append(f, "Fizz")
        } else {
             f = append(f, strconv.Itoa(i))
        }
    }
    return f
}