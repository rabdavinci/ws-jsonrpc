## DESCRIPTION
This example shows how to set up a gorilla/websocket server to process JSON-RPC calls
Also implements minimal Ping/Pong handling
## HOW 2RUN

### INSTALL MODULES
1. `go mod init github.com/rabdavinci/ws-jsonrpc`
2. `go mod download`
3. `go mod vendor`
4. `go mod verify`

### RUN SERVER
1. `cd server && go run *.go`
### RUN CLIENTS
1. `cd client_first && go run main.go`
2. `cd client_second && go run main.go`

## TODO
1. Add tests
2. Dockerize
