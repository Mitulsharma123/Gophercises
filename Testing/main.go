package main

import "fmt"
import m "golang-book/chapter11/math"

func main(){
 xs := []float64{1,2,3,4}
 avg := m.Average(xs)
 fmt.Println(avg)
}


// m is the alias

/*You may have noticed that every function in the packages we've seen start with a capital letter. In Go if something starts with a capital letter that means other packages (and programs) are able to see it. If we had named the function average instead of Average our main program would not have been able to see it.*/
