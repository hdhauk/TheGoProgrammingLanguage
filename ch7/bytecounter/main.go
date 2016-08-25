package main

import "fmt"

// ByteCounter merly counts written bytes before discarding them
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // Converts int to ByteCounter
	return len(p), nil
}

func main() {
	var bc ByteCounter
	bc.Write([]byte("hello"))
	fmt.Println(bc)

	bc = 0
	var name = "Dolly Duck"
	fmt.Fprintf(&bc, "hello, %s", name)
	fmt.Println(bc) // "17", = len("hello, Dolly Duck")
}
