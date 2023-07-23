## Getting started
```bash
$ git pull git@github.com:sumimayo/restful-api-with-gin.git
```

## Compile and run the code
```bash
$ go run .
```

## From another terminal, make a request to GET album list
```bash
$ curl http://localhost:8080/albums
```

## Or make a request to GET specific album
```bash
$ curl http://localhost:8080/albums/2
```

## Make a request to POST album
```bash
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "5","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

## Future version
- [ ] Authentication
- [ ] List book reviews
- [ ] Review of book read
- [ ] Comments
