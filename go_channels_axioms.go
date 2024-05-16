package main

//import "fmt"


     func main(){
        //1. close nil channel = panic
        /* 
        var ch chan int
        close(ch) 
        */
        
        //2. close closed channel = panic
        /* 
        ch:=make(chan int)
        close(ch)
        close(ch)
        */
        
        //3. read nil channel = block
        /* 
        var ch chan int
        <-ch
        */
        
        //4. read closed channel = default value
        /* 
        ch:=make(chan int)
        close(ch)
        x:=<-ch
        fmt.Println(x)
        */
        
        //5. write nil channel = block
        /* 
        var ch chan int
        ch<-0
        */
        
        //6. write closed channel = panic
        /* 
        ch:=make(chan int)
        close(ch)
        ch<-0
        */
}
    
  
