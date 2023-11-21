package main

import (
	client "gorilla-mux/client"
	c "gorilla-mux/config"
)

func main() {
	c.HandleRequests()
	client.CallRestApi()
}
