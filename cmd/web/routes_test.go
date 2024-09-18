package main

import (
	"fmt"
	"github.com/TDBoudreau/bookings/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("Expected type of *chi.Mux, instead got %T", v))
	}
}
