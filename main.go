package main

import(
	"fmt"
  "log"
	"net/http"

)

func(){
	fileServer := http.FileServer( http.Dir("./static")


	)
}
