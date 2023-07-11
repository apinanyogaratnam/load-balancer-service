start:
	go run main.go

start1:
	go run server1/main.go

start2:
	go run server2/main.go

start3:
	go run server3/main.go

up:
	cd server1 && docker-compose up --build
