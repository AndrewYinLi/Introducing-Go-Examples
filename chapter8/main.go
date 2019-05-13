package main

import (
	//"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"io/ioutil"
	"path/filepath"
	"container/list"
	"sort"
	"hash/crc32" // hash includes adler32, crc32, crc64, and fnv. Hashes are frequently used in programming for everything from looking up data to easily detecting changes.
	"crypto/sha1" // Cryptographic hash functions are similar to their non-cryptographic counterparts, but they have the added property of being hard to reverse
)

// To demonstrate sorting
type Person struct {
	Name string
	Age int
}
type ByName []Person // slice that is filled in main function

/*
The Sort function in sort takes a sort.Interface and sorts it. The sort.Interface
requires three methods: Len, Less, and Swap.
 */
func (ps ByName) Len() int { // Len should return the length of the thing we are sorting. For a slice, simply return len(ps).
	return len(ps)
}
func (ps ByName) Less(i, j int) bool { // Less is used to determine whether the item at position i is strictly less than the item at position j. In this case, we simply compare ps[i].Name to ps[j].Name.
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) { // swap positions in array
	ps[i], ps[j] = ps[j], ps[i]
}

// Get hash from file
func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer f.Close()
	// create a hasher
	h := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(h, f)
	// we don't care about how many bytes were written, but we do want to
	// handle the error
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func main(){

	// String manipulation with `strings` library

	fmt.Println(strings.Contains("test", "es")) // Just like Java but first char is uppercase
	fmt.Println(strings.Count("test", "t")) // Why is this a built-in function lol
	fmt.Println(strings.HasPrefix("test", "te")) // Whyyyyy
	fmt.Println(strings.HasSuffix("test", "st")) // Smh
	fmt.Println(strings.Index("teste", "e")) // Returns index of first occurrence of substr
	fmt.Println(strings.Join([]string{"a","b"}, "-")) // Takes a list of strings and joins them into a single string delimited by a different string
	fmt.Println(strings.Repeat("a", 5)) // ok
	fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // -1 to do as many times possible, else `2` in this case
	fmt.Println(strings.Split("a-b-c-d-e", "-")) // => []string{"a","b","c","d","e"}
	fmt.Println(strings.ToLower("TEST")) // str8 outta java
	fmt.Println(strings.ToUpper("test"))

	// ___________________________________________________________________

	// Read/write file with `os` and `io` libraries

	// Example of reading file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error, could not open file")
		// handle the error here
		return
	}
	defer file.Close() //We use defer file.Close() right after opening the file to make sure the file is closed as soon as the function completes.
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println("Printing test.txt: " + str)



	// Faster way to read file
	fastBs, fastErr := ioutil.ReadFile("test.txt")
	if fastErr != nil {
		return
	}
	fastStr := string(fastBs)
	fmt.Println("Printing test.txt (fast way): " + fastStr)



	// Create a file
	newFile, newFileErr := os.Create("new_file.txt")
	if newFileErr != nil {
		// handle the error here
		return
	}
	defer newFile.Close()
	newFile.WriteString("test")



	// Get contents of a directory
	dir, dirErr := os.Open(".")
	if dirErr != nil {
		return
	}
	defer dir.Close()

	fileInfos, fileInfosErr := dir.Readdir(-1) // takes a single argument that limits the number of entries returned. By passing in -1, we return all of the entries.
	if fileInfosErr != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}



	// Recurse through folder and all subfolders
	filepath.Walk(".", func(path string, info os.FileInfo, walkError error) error {
		fmt.Println(path)
		return nil
	})



	// customError := errors.New("error message") // declare custom error



	// Linked List
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}



	// Demonstrate sorting
	kids := []Person{ // kids array
		{"Jill",9},
		{"Jack",10},
	}
	sort.Sort(ByName(kids)) // fill slice with `kids` and sort
	fmt.Println(kids)


	// hashing with crc32
	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)



	// diff file contents using getHash() function implemented above main()
	h1, h1Err := getHash("test.txt")
	if h1Err != nil {
		return
	}
	h2, h2Err := getHash("test2.txt")
	if h2Err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)



	// Crypto hash
	/*
	This example is very similar to the crc32 one, because both crc32 and sha1 implement
	the hash.Hash interface. The main difference is that whereas crc32 computes a
	32-bit hash, sha1 computes a 160-bit hash. There is no native type to represent a 160-
	bit number, so we use a slice of 20 bytes instead.
	 */
	cH := sha1.New()
	cH.Write([]byte("test"))
	bs = cH.Sum([]byte{})
	fmt.Println(bs)


}