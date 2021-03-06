//This program will print “ping” forever (hit enter to stop it)

package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) { //channel type represented by keyword chan followed by type
  for i := 0; ; i++ {
    c <- "ping"     // <- operator to send & receive messages
  }			// c<-"ping" means send ping
}

func ponger(c chan string){
  for i:=0; ; i++ {
    c <- "pong"
  }
}

func printer(c chan string) {
  for {
    msg := <- c			// receive a message and store it in msg
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string 
  fmt.Scanln(&input)

}
