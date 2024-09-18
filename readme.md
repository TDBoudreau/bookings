# Bookings and Reservations

This is the repository for the project I am following on Udemy. Learning web app design in Go and, more specifically, working the html/template package. https://pkg.go.dev/html/template

- Built using Go version 1.22.6
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [Alex Edwards SCS session manager](https://github.com/alexedwards/scs)
- Uses [nosurf](https://github.com/justinas/nosurf)

# Testing
- go test ./cmd/web/
- go test ./cmd/web/ -v
- go test ./cmd/web/ -coverprofile=coverage.out && go tool cover -html=coverage.out