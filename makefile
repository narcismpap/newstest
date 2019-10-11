build: 
	docker build -t gonews .

run:
	docker run -p 9000:9000 -it gonews

up: build run
