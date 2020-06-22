package main

import (
	"gin_demo02/opt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	req := &opt.Req{}

	r.GET("/add_register", req.AddRegisterData)
	_ = r.Run(":8082")
}
