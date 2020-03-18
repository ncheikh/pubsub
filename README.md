# pubsub
HTTP Based PubSub System that uses HTTP SSE to publish messages to subscribers


## Running

`make run`

`make serve` (Hot Reloading)

Subscribe: `curl localhost:8080/subscribe`

Publish: `curl --header "Content-Type: application/json" --request POST --data '{"message":"some message"}' localhost:8080/publish`

## Testing

`make test`