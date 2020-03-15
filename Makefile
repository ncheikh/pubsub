all:
	make build.server

clean:
	make clean.storage-ledger

run:
	make run.server

test:
	make test.server

serve:
	make serve.server

build.server:
	go build -race -o ./out/server ./cmd/server

build.client:
	go build -race -o ./out/client ./cmd/client


test.server:
	go test ./cmd/server


run.server:
	make build.server

	./out/server

run.client:
	make build.client

	./out/client

# Run with hot reload
serve.server:
	~/.air -c build/air/server.conf 
