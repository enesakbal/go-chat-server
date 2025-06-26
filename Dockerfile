FROM golang
RUN mkdir /go-chat-server
WORKDIR /go-chat-server

COPY . .

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]

EXPOSE 8080