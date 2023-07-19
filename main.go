package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zupu/pkg/rest"
)

func main() {

	fmt.Print("start")
	engine := gin.Default()
	configRoute(engine)
	engine.Run()
	fmt.Print("start success.")
}

func configRoute(engine *gin.Engine) {
	rs := rest.NewRS(engine)
	v1 := rs.Group("/api/v1")
	rs.GET(v1, "/user/:id", rs.GetUser)
	rs.GET(v1, "/person/:id", rs.GetPerson)

}
