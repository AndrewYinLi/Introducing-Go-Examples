package main

import "fmt"

func main(){

    // Arrays

    // Calculate averages
    var x [5]float64 // float64 not float
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83
    var total float64 = 0
    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    fmt.Println(total / float64(len(x))) // Can't divide float64 by int, so cast

    total = 0
    for _, value := range x { // Foreach loop, underscore says to skip iterator
        total += value
    }
    fmt.Println(total / float64(len(x))) // Can't divide float64 by int

    // Other way to declare an array
    x = [5]float64{ 98, 93, 77, 82, 83 }

    // Another way as well
    // Notice the extra trailing , after 83. This is required by Go and it allows us to easily
    // remove an element from the array by commenting out the line (stupid af)
    x = [5]float64{
        98,
        93,
        77,
        82,
        83,
    }

    // Slices (Go lists)
    // Typically used to represent lists of items, particularly when you need to access
    // the nth item quickly—for example, player #33, or the 10th most popular search query.

    // A slice is a segment of an array. Like arrays, slices are indexable and have a length.
    // Unlike arrays, this length is allowed to change. Example of a slice:
    var y []float64

    y = make([]float64, 5) // Different way to make slice w/ length 5
    y = make([]float64, 5, 10) // Slice w/ length 5, capacity 10

    // Making a slice from an array
    arr := [5]float64{1,2,3,4,5}
    y = arr[0:5]
    y = arr[0:] // Pythonic
    y = arr[:len(arr)]

    // append
    y = append(arr[0:5], 6, 7) // Make sure to append a slice not an array
    fmt.Println(arr[0:5], y)

    // copy
    // If the lengths of the two slices are not the same, the smaller of the two will be used.
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)

    // Maps (basically dictionaries from Python)
    m := make(map[string]int)
    m["key"] = 10
    fmt.Println(m["key"])
    delete(m, "key")

    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"
    fmt.Println(elements["Li"])
    fmt.Println(elements["Un"]) // Returns ""

    name, ok := elements["Un"] // Returns two values
    fmt.Println(name, ok) // name is value of lookup, ok is whether it was successful

    //First, we try to get the value from the map. Then, if it’s successful, we run the code
    // inside of the block. The stuff inside the `ok` block won't run.
    if name, ok = elements["Un"]; ok {
        fmt.Println(name, ok)
        fmt.Println("SUCCESSFUL")
    }

    // faster way to declare map
    // elements2 := map[string]string{
    //     "H": "Hydrogen",
    //     "He": "Helium",
    //     "Li": "Lithium",
    //     "Be": "Beryllium",
    //     "B": "Boron",
    //     "C": "Carbon",
    //     "N": "Nitrogen",
    //     "O": "Oxygen",
    //     "F": "Fluorine",
    //     "Ne": "Neon",
    // }

    // Nested
    elements3 := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium",
            "state":"solid",
        },
        "B": map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C": map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N": map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O": map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F": map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne": map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }
    if el, ok := elements3["Li"]; ok {
    fmt.Println(el["name"], el["state"])
    }
}
