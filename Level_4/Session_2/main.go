package main

import "idstar.com/session2/app/routers"

func main() {
	r := routers.SetupRouter()
	r.Run(":5001")
}
