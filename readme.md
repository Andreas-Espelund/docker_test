# Testing docker for building and deploying and API


## Start the server

```bash
cd golang-server

docker build -t server

docker run -p 8080:8080 server
```

## Start the client

```bash
cd python-client

docker build -t client

docker run --network="host" client
```