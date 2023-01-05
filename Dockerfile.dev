FROM golang:alpine as BUILD

ENV CGO_ENABLED=0

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o /demo

CMD ["/demo"]