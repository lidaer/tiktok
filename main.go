package main

import (
	"tiktok/router"
)

func main() {
	r := router.InitDouyinRouter()
	err := r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
