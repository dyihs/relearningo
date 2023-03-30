package main

import "net/http"

type Server interface {
    Route(pattern string, handleFunc http.HandlerFunc)
    Start(address string) error
}

type sdkHttpServer struct {
    Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
    // TODO implement me
    http.HandleFunc(pattern, handleFunc)
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
