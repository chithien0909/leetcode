package main

func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    // Assign i = 1
    i := 1
    // Keep incrementing i until i * i > x
    // When x = 8
    // i = 1 => i * i = 1 < 8
    // i = 2 => i * i = 4 > 8
    // i = 3 => i * i = 9 > 8
    // return i - 1 => 2
    for i * i <= x {
        i++
    }
    return i - 1
}