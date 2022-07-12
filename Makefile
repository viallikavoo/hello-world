
APP_NAME := "hello_world"

clean:
	rm -rf bin || true
	rm Dockerfile || true

generate-dockerfile:
	go build -o bin/ template_generator.go
	bin/template_generator $(APP_NAME)

test: 
	docker build --no-cache -t $(APP_NAME) .
	docker run $(APP_NAME) | grep "hello world"