HTTP over TLS demo in Go. Stolen from [Liz Rice](https://github.com/lizrice/secure-connections) :-).

server:

```
minica -ca-common-name localhost # generate CA cert and private key
go run server/main.go
```

client:

```
curl https://localhost:8080 # observe the client and server error
curl https://localhost:8080 --cacert cacert.crt
```