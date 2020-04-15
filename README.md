HTTP over TLS demo in Go. Stolen from [Liz Rice](https://github.com/lizrice/secure-connections) :-).

server:

```
minica -ca-common-name localhost
go run server/main.go
```

client:

```
curl https://localhost:8080 --cacert cacert.crt
```