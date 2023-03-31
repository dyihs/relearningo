package main

import (
    "fmt"
    "net/http"
)

type Server interface {
    Route(pattern string, handleFunc func(ctx *Context))
    Start(address string) error
}

type sdkHttpServer struct {
    Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *Context)) {
    // 自己控制 Context 的创建
    http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
        ctx := NewContext(writer, request)
        handleFunc(ctx)
    })
}

func (s *sdkHttpServer) Start(address string) error {
    // TODO implement me
    return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
    return &sdkHttpServer{
        Name: name,
    }
}

// SignUp 在没有 context 抽象的情况下，是长这样的
func SignUp(ctx Context) {
    req := &signUpReq{}

    err := ctx.ReadJson(req)
    if err != nil {
        ctx.BadRequestJson(err)
        return
    }

    resp := &commonResponse{
        Data: 123,
    }
    err = ctx.WriteJson(http.StatusOK, resp)
    if err != nil {
        fmt.Printf("写入响应失败：%v", err)
        return
    }
}

type signUpReq struct {
    Email             string `json:"email"`
    Password          string `json:"password"`
    ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
    BizCode int         `json:"biz_code"`
    Msg     string      `json:"msg"`
    Data    interface{} `json:"data"`
}
