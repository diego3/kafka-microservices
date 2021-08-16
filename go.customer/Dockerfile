FROM golang:1.16 as go16

WORKDIR /app

COPY . /app/

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=go16 /app ./
CMD ["./app"]  