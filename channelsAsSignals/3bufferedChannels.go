package channelsassignals

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
//Sometimes we may want to execute channels inchucks instead of all at the same time
//For the above will use a new channel (limiter) which takes a second parameter of buffer size limiter:=make(chan int, 3)
//This means that go will execute the first three channel operations then pause as it awaits comnpletion of one or more before proceeding

//2)
//Remember to add the limitter as a parameter to the function being channeled

func main(){
c:=make(chan int)
limiter:=make(chan int, 3)

//Generate two values, get their sum and print the sum
 go generateValue(c,limiter)
 go generateValue(c,limiter)
 go generateValue(c,limiter)
 go generateValue(c,limiter)

 //2)
 //When creating the generateValue() function, we've set it to store its return data in a channel
 //Here we now extract the stored data into x and y variables
 //x:= <- c means that the data will flow from channel C into variable x


        // x:= <- c    //Before goroutine looping     
        // y:= <- c    //Before goroutine looping
  
        //sum:=x+y    //Before goroutine looping
        //fmt.Println(sum)

//3)
//Its also possible to loop through goroutines in a situation where we execute multiple routines concurrently
sum:=0
i:=0

//Here the goal is to loop throught the return values of each channel as they finish executing
//To prevent a deadlock we can check for the number of iterations
//As we expect 4 iterations, we will close the operation after 4 iterations
for num:=range c{
    sum=sum+num
    i++
    if i==4{
        close(c)
    }
   
//4)
//
    

}
fmt.Println(sum)
 
}