FROM golang:latest AS builder
WORKDIR /app
COPY . /app/
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /app
ENV GIN_MODE=release
EXPOSE 8080
COPY --from=builder /app/main /app
CMD [ "/app/main" ]