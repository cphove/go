# go

### How to use

1. use go get to add dependency

```shell
cd agent
go get .
```

2. use go build 

```shell
go build -o . main.go
```

3. run server

```shell
./main
```

4. test

```shell
curl http://localhost:8080/run --include --header "Content-Type: application/json" --request "POST" --data '{"args": ["python", "--version"]}'
```

If you have python installed on your computer, you will see the following:

```json
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 24 Mar 2023 08:39:48 GMT
Content-Length: 71

{
    "code": 200,
    "message": null,
    "data": "Python 3.10.4\n"
}
``