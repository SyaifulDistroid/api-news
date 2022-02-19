FROM golang:1.16-alpine

WORKDIR /api-news/

COPY go.mod go.sum ./

RUN go mod download

COPY COPY . .

COPY api-news/main.go .

WORKDIR /api-news

RUN go build -o main .

CMD ["./api-news/main"]

EXPOSE 8585