package main

import "fmt"

func average(xs []float64) float64 { // Notice how type is after and return is last smh
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func main() {
    xs := []float64{98,93,77,82,83} // Averaging a slice
    fmt.Println(average(xs))

    x, y := f() // Testing returning multiple values
    fmt.Println(x, y)

    z := f2() // Testing returning a named value
    fmt.Println(z)

    z = f3() // Testing returning a named value a different way
    fmt.Println(z)

    fmt.Println(add(1,2,3)) // Testing a variadiac function that adds ints

    randomSlice := []int{1,2,3} // Testing a variadiac function that adds an int slice
    fmt.Println(add(randomSlice...)) // Pass in a slice with an ellipsis

    addClosure := func(x, y int) int { // Testing a closure that adds 2 ints
        return x + y
    }
    fmt.Println(addClosure(1,1))

    count := 0 // jesus is this js
    increment := func() int {
        count++
        return count
    }
    fmt.Println(increment())
    fmt.Println(increment())

    nextEven := makeEvenGenerator() // Make even numbers w/ closure
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4

    fmt.Println(factorial(3)) // testing factorial() function

    /*
    defer is often used when resources need to be freed in some way. For example, when
    we open a file, we need to make sure to close it later. With defer:
    f, _ := os.Open(filename)
    defer f.Close()
    This has three advantages:
    • It keeps our Close call near our Open call so it’s easier to understand.
    • If our function had multiple return statements (perhaps one in an if and one in
    an else), Close will happen before both of them.
    • Deferred functions are run even if a runtime panic occurs.
    */
    defer second()
    first()

    // Just like C/C++
    // * = pointer (value of address)
    // & = address of value
    n := 5
    zero(&n)
    fmt.Println(n) // n is 0

    m := new(int) // the `new` keyword returns an address
    zero(m)
    fmt.Println(*m) // the value at address m is 0

    //recover stops the panic and returns the value that was passed to the call to panic
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}

func f() (int, int) { // Return multiple values
    return 5, 6
}

func f2() (r int) { // Name return variable
    r = 1
    return
}

func f3() (int) { // Also works
    r := 1
    return r
}

// Variadiac parameter
func add(args ...int) int {
    total := 0
    for _, v := range args {
    total += v
    }
    return total
}

func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}

func first() {
    fmt.Println("1st")
}

func second() {
    fmt.Println("2nd")
}

func zero(nPtr *int) {
    *nPtr = 0
}
