FROM golang:1.16

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./
CMD ["go","run","main.go"]