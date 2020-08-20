FROM golang:1.15 AS builder
 
RUN mkdir -p /app
 
WORKDIR /app
 
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env . 

EXPOSE 8080

CMD ["./main"]