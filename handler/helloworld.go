package handler

import (
	"github.com/valyala/fasthttp"
	"log"
)

func Helloworld(ctx *fasthttp.RequestCtx){
	log.Println("helloworld")
}
