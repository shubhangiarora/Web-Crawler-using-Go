//// To run this file run command go run crawl.go website name 


package main

import (
  "flag"          
  "fmt"
  "net/http"      // 'http' will retrieve pages for us
  "io/ioutil"     // 'ioutil' will help us print pages to the screen
  "os"
)

func main() {
  flag.Parse()

  args := flag.Args()
  fmt.Println(args)
  if len(args) < 1 {
    fmt.Println("Please specify start page")
    os.Exit(1)
  }
  retrieve(args[0])  // The reason we can call 'retrieve' is because
                     // it's defined in the same package as the calling function.
}

func retrieve(uri string) {  // This func(tion) takes a parameter and the
                             // format for a function parameter definition is
                             // to say what the name of the parameter is and then
                             // the type.
                             // So here we're expecting to be given a
                             // string that we'll refer to as 'uri'
  resp, err := http.Get(uri)
  if err != nil {            //To handle the error typically works in Go.
    return                   
  }
  defer resp.Body.Close()  // to close the resource we opened
                          

  body, _ := ioutil.ReadAll(resp.Body) 
  fmt.Println(string(body))             
}
