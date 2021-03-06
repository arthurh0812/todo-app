FROM golang:1.15-alpine

COPY . /app

WORKDIR /app

RUN go get firebase.google.com/go/v4

ENV GOOGLE_APPLICATION_CREDENTIALS "C:/Users/Caroline/Coding/todo-app/backend/service-account-key.json"

RUN go build -o main .

ENTRYPOINT ["/app/main"]
