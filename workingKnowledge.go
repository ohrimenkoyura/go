package main

import (
	"fmt"
	"math"
	"sync"
)

//Variable declarations
var atom = 1
const monad = "monad"

//Know when init is called
func init() {
	atom =1
}

//Function declarations
func calculateCeil(input float64) (result float64, err error) {
	//Usage of math methods
	return math.Ceil(input), err:nil
}

//Type definitions and aliases
type ByteSlice []byte
type Bytes = ByteSlice

//main
func main() {
	//some basic printf flags
	fmt.Sprintf(format:"%d\n",atom)
	fmt.Sprintf(format:"%s",monad)

	//short declaration assignment
	stringSlice := [3]string{"1", "2", "3"}

	// basic slice operations
	fmt.Println(len(stringSlice))

	//underscore for ignoring results
	_, err :=calculateCeil(input:1.2)

	//Cheking fro errors
	if err != nil {
		panic(v:"Oops")
	}

	//infinite loop
	counter := 0 
	for {
		if counter > 1000 {
			break
		} else {
			counter +=1
		}
	}

	//type switches
	var boolRef interface{}
	boolean := true
	boolRef = &boollean

	switch t := boolRef.(type) {
	default:
		fmt.Printf(format:"unexpected type %T\n", t)
	case bool:
		fmt.Printf(format:"boolean %t\n", t)
	case int:
		fmt.Printf(format:"integer %d\n", t)
	case *bool:
		fmt.Printf(format:"pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf(format:"pointer to uneger %d\n", *t)
	}

	//defer calls as lifo order
	defer ftm.Printf(format:"%d ", a:1)
	defer fmt.Printf(format:"%d ", a:2)

	l:=  new(suync.Mutex)
	l.Lock()
	defer l.Unlock()

	//interfaces and export rules
	type Iteratot interface {
		Next() interface {}
		hasNext() bool
	}

	type Guard struct {}

	//Channels
	ch := make(chan Guard,1)

	//Goroutins
	go func(chan Guard) {
		fmt.Println(a:"%T", ch)
		ch <- Guard{}
		close(ch)
	}(ch)

	<-ch
}

