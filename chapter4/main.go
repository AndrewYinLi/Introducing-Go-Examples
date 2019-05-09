package main

import "fmt"

func main(){

    // Pretty standard minus the colon
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "even")
        } else { // ELSE CANT BE ON A NEW LINE WTF
            fmt.Println(i, "odd")
        }
    }

    // Case-switch, no need for break statements which is nice
    for i := 1; i <= 10; i++ {
        switch i {
            case 0: // Can format like so
                fmt.Println("Zero")
            case 1: fmt.Println("One") // Or like this, very clean
            case 2: fmt.Println("Two")
            case 3: fmt.Println("Three")
            case 4: fmt.Println("Four")
            case 5: fmt.Println("Five")
            default: fmt.Println("Greater than five, much big such wow")
        }
    }

}
