package main

import (
	"tutorial/server"
)

// var (
// 	strContentType     = []byte("Content-Type")
// 	strApplicationJSON = []byte("application/json")
// )

// func Introduce(ctx *fasthttp.RequestCtx) {
// 	if string(ctx.Method()) != "POST" {
// 		ctx.Error("Wrong Method", fasthttp.StatusNotFound)
// 		return
// 	}
// 	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	fmt.Println(string(ctx.PostBody()))

// }

// func Create(ctx *fasthttp.RequestCtx) {
// 	if string(ctx.Method()) != "POST" {
// 		ctx.Error("Wrong Method", fasthttp.StatusNotFound)
// 		return
// 	}
// 	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	fmt.Println(string(ctx.PostBody()))
// }
// func Read(ctx *fasthttp.RequestCtx) {
// 	if string(ctx.Method()) != "GET" {
// 		ctx.Error("Wrong Method", fasthttp.StatusNotFound)
// 		return
// 	}
// 	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	fmt.Println(string(ctx.PostBody()))
// }

// func Update(ctx *fasthttp.RequestCtx) {
// 	if string(ctx.Method()) != "PUT" {
// 		ctx.Error("Wrong Method", fasthttp.StatusNotFound)
// 		return
// 	}
// 	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	fmt.Println(string(ctx.PostBody()))
// }

// func Delete(ctx *fasthttp.RequestCtx) {
// 	if string(ctx.Method()) != "DELETE" {
// 		ctx.Error("Wrong Method", fasthttp.StatusNotFound)
// 		return
// 	}
// 	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	fmt.Println(string(ctx.PostBody()))
// }

// IN FUNC MAIN
// the corresponding fasthttp request handler
// requestHandler := func(ctx *fasthttp.RequestCtx) {
// 	switch string(ctx.Path()) {
// 	case "/introduce":
// 		Introduce(ctx)
// 	case "/create":
// 		Create(ctx)
// 	case "/read":
// 		Read(ctx)
// 	case "/update":
// 		Update(ctx)
// 	case "/delete":
// 		Delete(ctx)
// 	default:
// 		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
// 	}
// }
// fasthttp.ListenAndServe(":8080", requestHandler)

func main() {
	server.NewServer()
}
