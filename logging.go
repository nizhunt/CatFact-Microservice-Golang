package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

// this is how you make a middleware, take a service, inject your methods and return that service back
func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s\n error=%v\n took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())
	return s.next.GetCatFact(ctx)
}
