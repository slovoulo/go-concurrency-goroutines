package channelsassignals

import (
	"fmt"
	"math/rand"
	"time"
)

//random number generator
var source=rand.NewSource(time.Now().Unix())
var randN=rand.New(source)

func generateValue(c chan int ){
    sleepTime:=randN.Intn(3)

    time.Sleep(time.Duration(sleepTime)*time.Second)// This line is here to simulate that the number generation will take some time 
    
    //instead of returning a value, we send it into a channel
    

    c <- randN.Intn(10)
   // return randN.Intn(10)   - We are not returning a value here but it's also possible
}
//1)
//Channels can also be used to store values
//In the below code we expect some integers from generateValue() but since they are go routines we'll rely on channels to store their returns


func main(){
c:=make(chan int)

//Generate two values, get their sum and print the sum
 go generateValue(c)
 go generateValue(c)
 go generateValue(c)
 go generateValue(c)

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