//This program will print “ping” forever (hit enter to stop it)

package main

import (
  "fmt"
  "time"
)

func pinger(c chan float64) { //channel type represented by keyword chan followed by type
  for i := 0; ; i++ {
    c <- 100     // <- operator to send & receive messages
  }			// c<-"ping" means send ping
}

func ponger(c chan float64){
  for i:=0; ; i++ {
    c <- 99
  }
}

func printer(c chan float64) {
  for {
    msg := <- c			// receive a message and store it in msg
    fmt.Println("%f", msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan float64 = make(chan float64)

  go pinger(c)
  go ponger(c)
  go printer(c)

  //var input string 
  //fmt.Scanln(&input)

}
