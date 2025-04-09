FROM golang:1.24

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY . .

RUN go mod tidy

CMD [ "air", "-c", ".air.toml" ]