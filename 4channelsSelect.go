package main

import (
	"fmt"
	"math/rand"
	"time"
)

//random number generator
var source=rand.NewSource(time.Now().Unix())
var randN=rand.New(source)

func generateValue(c chan int, limit chan int ){
    limit<-1 //Send a random value to the limit channel at the start of the function execution
    sleepTime:=randN.Intn(3)

    time.Sleep(time.Duration(sleepTime)*time.Second)// This line is here to simulate that the number generation will take some time 
    
    //instead of returning a value, we send it into a channel
    

    c <- randN.Intn(10)
    <- limit //Free up the limit channel at the end of the function execution
   // return randN.Intn(10)   - We are not returning a value here but it's also possible
}
//1)
//There are instances when we might want to perform operations concurrently but stop once the first one completes
//Since it's impossible to know which goroutine will finish first, we can set up a checking mechanism
//In such a scenario we'll use a special channels control structure known as select
//We'll create 2 separate channels and use them for different routines
//we then create two variables a and b to store the values returned by each channel
//With the select statement we'll be able to check for the channel that returns a value first then proceed with its subsequent code execution




func main(){
x:=make(chan int)
y:=make(chan int)
limiter:=make(chan int, 3)

var a int
var b int

//Generate two values, get their sum and print the sum
 go generateValue(x,limiter)
 go generateValue(y,limiter)

 select{
 //If channel x finishes first save its return value in a then do something else   
 case a= <-x :
    fmt.Printf("x finished first and returned %v", a)
 
 case b= <-y :
    fmt.Printf("y finished first and returned %v", b)
 }



    


 
}