FROM golang:1.18 as development

WORKDIR /app
COPY ./go.mod ./go.sum /app/
RUN go mod download
COPY . /app

EXPOSE 8080

CMD cd cmd/webhookListener && go run main.go