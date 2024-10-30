unit-test:
	go test ./... --short

integration-test:
	# docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
	
	docker start redis-stack
	go test ./...
	docker stop redis-stack
