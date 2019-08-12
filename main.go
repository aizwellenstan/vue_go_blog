package main

import "github.com/aizwellenstan/vue-go-blog/router"

func main() {
	r := router.GetRouter()
	r.Run(":2000")
}
