//Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
package main

import "fmt"

func main(){
   var input_fahr float32
   fmt.Print("Enter the temperature in Fahrenheit: ")
   fmt.Scanf("%f", &input_fahr)
   
   output_celc := float32(input_fahr-32)*5/9
   
   fmt.Println("Temperature in Celsius:", output_celc)
}
