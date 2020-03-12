package controller

import (
	"fmt"
	"log"
	"net/http"
)

type Service struct {
	port string
}

func NewService(port string) *Service {
	return &Service{port:port}
}

func (s *Service)Run()  {
	handler := NewServiceHandler()
	http.Handle("/payment", handler)
	log.Printf("Server starting on port %v\n", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil))
}