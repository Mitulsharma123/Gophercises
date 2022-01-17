//Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.

package main

import ("fmt"; "math")

type Shape interface{  //interface Shape 
  area() float64
  perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {  //function distance 
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

type Circle struct{  //defining a struct type of circle
  x, y, r float64    //list of field
}

type Rectangle struct { ////defining a struct type of rectangle
  x1,y1,x2,y2 float64
} 

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}
//In between the keyword func and the name of the function we've added a “receiver”. 
//The receiver is like a parameter – it has a name and a type –
//but by creating the function in this way it allows us to call the function using the . operator

func (c *Circle) area() float64{
  return math.Pi * c.r * c.r
}

func main() {
  
  c := Circle{1,1,5}
  r := Rectangle{1, 1, 10, 10}

  fmt.Println(c.area())
  fmt.Println(r.area())

}

//perimeter method to calculate perimeter of shape

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
    return 2 * (l + w)
}

func (c *Circle) perimeter() float64{
  return 2 * math.Pi * c.r
}


