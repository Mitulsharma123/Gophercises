package main

import ("fmt"; "math")

type Shape interface{
  area() float64
  perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

/*func rectangleArea(x1, y1, x2, y2 float64) float64 {
  l := distance(x1, y1, x1,y2)
  w := distance(x1, y1, x2, y1)
  return l * w
}

func circleArea(x, y, r float64) float64 {
  return math.Pi * r*r
}
*/
type Circle struct{
  x, y, r float64
}

type Rectangle struct {
  x1,y1,x2,y2 float64
} 

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (c *Circle) area() float64{
  return math.Pi * c.r * c.r
}

func main() {
  //var rx1, ry1 float64 = 0, 0
  //var rx2, ry2 float64 = 10, 10
  //var cx, cy, cr float64 = 0, 0, 5
  
  //fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
  //fmt.Println(circleArea(cx, cy, cr))
  
  c := Circle{1,1,5}
  r := Rectangle{1, 1, 10, 10}

  fmt.Println(c.area())
  fmt.Println(r.area())

}

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
    return 2 * (l + w)
}

func (c *Circle) perimeter() float64{
  return 2 * math.Pi * c.r
}


