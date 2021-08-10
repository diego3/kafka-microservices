FROM  golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app/

#RUN ls -la

RUN go build -o customer-service

EXPOSE 9090

ENTRYPOINT [ "customer-service" ]