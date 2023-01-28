FROM golang:latest as builder
WORKDIR ./workspace
COPY ./pkg/ ./pkg/ 
COPY ./config/ ./config/
COPY ./cmd/main.go ./cmd/main.go
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod tidy 
EXPOSE 9092
CMD ["bash", "-c", "go run ./cmd/main.go"]
