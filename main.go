package main

import (
	"fmt"
	"io"
	"os"
)
func main(){
  fmt.Println("Hello go!")
  if(len(os.Args) > 1){
	f,e :=os.OpenFile(os.Args[1],os.O_RDONLY,0666)
	if(e != nil){
		os.Exit(1)
	}
	defer f.Close()
	bs,err :=io.ReadAll(f)
	if err != nil{
		os.Exit(1)
	}
	fmt.Println(string(bs))
  }	
}