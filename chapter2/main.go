package main

import "fmt"

func main(){
    // Numbers
    fmt.Println("1 + 1 =", 1 + 1) // Note that there automatically is a space inserted between args
    fmt.Println("1 + 1 =", 1.0 + 1.0) // Same result as above... -_-

    // Strings
    fmt.Println(len("Hello, World")) // 12
    fmt.Println("Hello, World"[1]) // wtf why the ascii number instead
    fmt.Println("Hello, " + "World") // Yes

    // Booleans ALL OF THESE ARE JUST LIKE JAVA + THE C FAM YAY
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
