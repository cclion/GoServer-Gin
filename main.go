package main

import (
	"ginserver/routers"
)

var db = make(map[string]string)

func main() {
	r := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
