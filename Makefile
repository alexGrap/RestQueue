testing:
	go run cmd/main.go 2 &
	go run test/testing.go
	lsof -ti tcp:3000 | xargs kill


