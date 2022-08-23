FROM golang:1.18 as development

RUN apt update && apt install sqlite3 libsqlite3-dev -y

WORKDIR /app
COPY ./go.mod ./go.sum /app/
RUN go mod download
COPY . /app

RUN /usr/bin/sqlite3 /app/officio_assassinorum.db

EXPOSE 8080

CMD cd cmd/webhookListener && go run main.go