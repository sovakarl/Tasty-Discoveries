FROM golang:1.25 AS build 

WORKDIR /app
COPY go.mod .
RUN go mod download 

COPY . .
RUN go build -o main src/cmd/main.go

FROM alpine:latest 
WORKDIR /root 
COPY --from=build app/main .
# CMD [ "./main" ]