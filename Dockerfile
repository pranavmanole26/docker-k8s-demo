FROM golang:alpine

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o /demo

EXPOSE 8080

# FROM alpine:latest
# COPY --from=BUILD /app/demo /bin

CMD ["/demo"]
