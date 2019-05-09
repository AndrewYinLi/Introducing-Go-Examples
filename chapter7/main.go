package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}

/*
Arguments are always copied in Go. If we attempted to modify one of the fields inside of the circleArea function,
it would not modify the original variable. Because of this, we would typically write the function using a pointer to the Circle
*/
func circleArea(c *Circle) float64 {
    return math.Pi * c.r * c.r
}

// `type` keyword introduces a new type
// followed by name of the type, `Circle` in this case
// concluded with `struct` keyword to indicate we are defining a struct
type Circle struct {
    x, y, r float64
}

// `(c *Circle)`is a receiver and means that the `Circle` struct contains this method
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

type Person struct {
    Name string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
    Person // Android IS a Person
    Model string // Android HAS a `Model` string
}

type Shape interface { // Declaring an interface
    /*
    But instead of defining fields, we define a method set.
    A method set is a list of methods that a type must have
    in order to implement the interface.
    */
    area() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
    area += s.area()
    }
    return area
}

// Interfaces can also be used as fields. Consider a MultiShape that is made up of several smaller shapes:
type MultiShape struct {
    shapes []Shape
}

// We can even turn MultiShape itself into a Shape by giving it an area method.
// Now a MultiShape can contain Circles, Rectangles, or even other MultiShapes.
When building new programs, you often won’t know what your types should look
func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

/*
Advice oon interfaces from book: 
When building new programs, you often won’t know what your types should look
like when you start—and that’s OK. In Go, you generally focus more on the behavior
of your program than on a taxonomy of types. Create a few small structs that do what
you want, add in methods that you need, and as you build your program, useful interfaces
will tend to emerge. There’s no need to have them all figured out ahead of time.
Interfaces are particularly useful as software projects grow and become more complex.
They allow us to hide the incidental details of implementation (e.g., the fields of
our struct), which makes it easier to reason about software components in isolation.
In our example, as long as the area methods we defined continue to produce the
same results, we’re free to change how a Circle or Rectangle is structured without
having to worry about whether or not the totalArea function will continue to work.
*/

func main() {

    // var c Circle // defines a new variable `c` of type `Circle` with everything initialized to 0 or nil
    // c := new(Circle) // same as above but returns a ptr
    // c := Circle{x: 0, y: 0, r: 5} // when you don't know parameter names
    c := Circle{0, 0, 5} // when you know the order
    //c := &Circle{0, 0, 5} // prev way but getting pointer
    fmt.Println(circleArea(&c)) // Testing out regular function
    fmt.Println(c.area()) // Testing out method version, notice no ptr needed

    r := Rectangle{0, 0, 10, 10} // Making a rectange
    fmt.Println(r.area()) // Testing out method, not function

    a := new(Android)
    // Both work
    a.Person.Talk()
    a.Talk()

    fmt.Println(totalArea(&c, &r)) // Testing total area

    multiShape := MultiShape{
        shapes: []Shape{
            Circle{0, 0, 5},
            Rectangle{0, 0, 10, 10},
        },
    }
}
