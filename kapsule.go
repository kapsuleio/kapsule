package main

import (
        "fmt"
         "flag"
    )

func main () {
	servPtr := flag.String("serv", "main", "a string")
	
	flag.Parse()
	
	fmt.Println("serv:" , *servPtr)
	
	fmt.Println("tail:", flag.Args())
}
