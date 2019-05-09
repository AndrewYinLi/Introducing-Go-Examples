package main

import "fmt"

var outsideScope string = "I'm outside lol"

func main() {
    var x string = "Hello, World" // What the hell why can't this be `string x` instead
    fmt.Println(x)

    // Reassigning x
    x = "first"
    fmt.Println(x)
    x = "second"
    fmt.Println(x)

    // + to append strings, idk why the book has this example
    x = "first "
    fmt.Println(x)
    x += "second"
    fmt.Println(x)

    // Wtf you use == to compare strings in Go, not okay >:(
    // false, obviously
    var y string = "hello"
    var z string = "world"
    fmt.Println(y == z)
    // true, obviously
    z = "hello"
    fmt.Println(y == z)

    // Pythonic shortcut, using a colon, which is actually pretty cool
    pythonicShortcut := "Hello, World"
    fmt.Println(pythonicShortcut)

    /*
    Explaining Idiomatic Go, pertaining to the usage of :=

    Languages often have a set of informal, conventional rules. For example, in English,
    there are rules that govern the order of adjectives in a sentence. If, instead of “my
    small black cat,” you were to say, “my black small cat,” it would strike native speakers
    as very strange.
    Programming languages often exhibit the same phenomenon. Although there are
    often many ways to do something, not all of those ways are seen as a natural expression
    of the language. You may write an expression (var x int = 5) which is semantically
    valid in the language (it compiles), but which may look strange to experienced
    Go programmers.
    Go programmers (as well as programmers in other communities) often refer to this
    as the idiomatic usage of the language. Learning idiomatic Go is a worthwhile pursuit,
    but at this stage, you should probably focus on simply writing correct programs, as
    that’s challenging enough. Don’t let the best be the enemy of the good.
    */

    // Printing a variable outside of the scope, which obviously works
    fmt.Println(outsideScope)

    const constStr string = "I can't change"
    // The below would fail with `cannot assign to constStr`
    // constStr = "qwertyuiop"

    // Shorthand
    var (
        a = 5
        b = 10
        c = 15
    )
    fmt.Println(a, b, c) // Prints: 5 10 15

    // Also shorthand, sugoi
    var (d = 5; e = 10; f = 15)
    fmt.Println(d, e, f) // Prints: 5 10 15

    // you can use const too
    const (
        g = 5
        h = 10
        i = 15
    )
    fmt.Println(g, h, i) // Prints: 5 10 15

    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input) // Read in, similar to C, NOTE THAT A STRING IS 0 FOR SOME REASON, NO ERROR THROWN
    output := input * 2
    fmt.Println(output)

}
