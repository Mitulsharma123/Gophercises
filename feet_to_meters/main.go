//Write another program that converts from feet into meters. (1 ft = 0.3048 m)

package main

import "fmt"

func main(){
  fmt.Print("Enter the measurement in inches:")
  
  var input_inch int
  fmt.Scanf("%d",&input_inch)
  
  output_met := 0.2048*float32(input_inch)
  fmt.Println("Converted in meters:", output_met)
}
