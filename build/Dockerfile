FROM golang:1.18 as development

WORKDIR /app
COPY ./go.mod ./go.sum /app/
RUN go mod download
COPY . /app

CMD cd cmd/webhook && run main.go