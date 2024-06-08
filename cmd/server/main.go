package main

import "github.com/nk-hung/go-ecommerce-backend-api/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run(":3001")
}
