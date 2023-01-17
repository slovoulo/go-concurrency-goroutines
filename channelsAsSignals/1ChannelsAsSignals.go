//Channels as signals for complete code execution
package channelsassignals

import (
	"fmt"
	"os"
)

//1.imagine 2 functions

//Stores data to a file
func storeData(storableText string, fileName string){
    file, err:=os.OpenFile(
        fileName, //Desired file name
        os.O_CREATE| //Create the file if it doesnt exist
        os.O_APPEND| //Append file contents to existing contents without overwrite
        os.O_WRONLY, //Open the file in write mode
        0666,
    )

    if err!=nil{
        fmt.Println("An error occured. Exiting")
        return
    }
    defer file.Close()

    _,err=file.WriteString(storableText) //Actual writing of text to the file

    if err!=nil{
        fmt.Println("Writing to the file failed")
    }
}

//Repeatedly writes data to a file
func repeatedWrite(lines int, fileName string, c chan int){

    //Takes a number of lines and loops upto them
    //Calls the storeData function for each iteration
    
    for i:=1; i<lines; i++{
        text :=fmt.Sprintf("For %v Mordor\n",i)
        storeData(text,fileName)
    }

    fmt.Printf("--Done storing %v lines of text --\n", lines)

    //3)
    //This function is going to be a goroutine. 
    //To ensure it's completion before stoppin the program we add a channel as its parameter:  c chan int
    //To use the channel: Call the channel followed by a (go specific) back arrow operator symbolizing the direction of data flow followed by any value matching the channel type

    c<-1
}
func main(){

   // storeData("Someone on Facebook called Kratos, 'Keratosis'"," Funny typos")

   //Create a channel and define its datatype
   channel:=make(chan int)

   // 1)
   //Writing 5500 lines to a file might take a while especially when running multiple write operations
   //Instead of waiting for each operation to finish before starting the next we can use goroutines to run them concurrently
   //a goroutine is defined by adding the keyword "go" before a function

            // go repeatedWrite(5500,"Middle Earth")   //Before channels
            // go repeatedWrite(5500,"Middle Mars")    //Before channels

    //2)
    //By default goroutines start the next line of code before the routine is finished. If there is no more code the program stops
    //In the above example code execution stops immediately after starting the second routine (Before it's fully complete)
    //To prevent the above, channels can be used as defined by "channel:=make(chan int)"
    //After defining a channel pass it to the function as a parameter
        go repeatedWrite(5500,"Middle Earth",channel)   
        go repeatedWrite(5500,"Middle Mars",channel)   
    
    //4
    //Since the routine here have a channel param, go will start them simultaneously but wont close the progrum until data is received from the channel
    //To fetch data from a channel use the back arrow operator to denote data flow direction
    //For each routine, call the channel individually
    <- channel
    <- channel
}