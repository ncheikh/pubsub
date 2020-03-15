all:
	make build.server

clean:
	make clean.server

run:
	make run.server

test:
	go test ./...

serve:
	make serve.server


# Server Commands
clean.server:
	rm ./out/server

build.server:
	go build -race -o ./out/server ./cmd/server

run.server:
	make build.server

	./out/server

# Run with hot reload
serve.server:
	~/.air -c build/air/server.conf 
