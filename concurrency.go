package main

import (
	"fmt"
	"os"
)

//imagine 2 functions
func greet(){
    fmt.Println(" General Kenobi")
}

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

func main(){

    storeData("Someone on Facebook called Kratos, 'Keratosis'"," Funny typos")

}