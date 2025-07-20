FROM golang:1.24.4

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 3000

CMD ["air", "serve"]
