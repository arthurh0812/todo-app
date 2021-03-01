FROM golang:1.16-alpine

LABEL name=frontend

COPY . /app

WORKDIR /app

EXPOSE 8000

RUN go build -o main .

ENTRYPOINT ["/app/main", "-host=localhost", "-port=8000"]


