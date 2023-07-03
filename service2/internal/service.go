package internal

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Service struct {
	BaseUrl string
	Client  *http.Client
}

func NewService() *Service {
	return &Service{
		BaseUrl: "http://localhost:8080",
		// BaseUrl: "http://app1:8080", // docker-compose
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (s *Service) FindAll(ctx context.Context) (products []Product, err error) {
	resp, err := s.Client.Get(s.BaseUrl + "/api/v1/products")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&products)
	return
}
