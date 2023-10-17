FROM golang:latest

LABEL MICRO-SERVICE = "BANK EMPLOYEE API"

WORKDIR /app

COPY go.mod /app/

COPY go.sum /app/

RUN go mod download

COPY ./src /app/

RUN go build -o /bank-employee-microservice

EXPOSE 8084

CMD ["/bank-employee-microservice"]