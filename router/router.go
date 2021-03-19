package router

import (
	routing "github.com/buaazp/fasthttprouter"
	"github.com/visaruttt/simple-fasthttp/handler"
)

func New() *routing.Router {
	return routing.New()
}

func Mount(r *routing.Router){
	r.GET("/helloworld", handler.Helloworld)
}
