FROM golang:latest-alpine as BUILD

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN go build -o demo

FROM alpine:latest
COPY --from=BUILD /app/demo /bin

CMD ["/demo"]
