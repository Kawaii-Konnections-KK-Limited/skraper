FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

VOLUME [ "/usr/src/app/session" ]

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app"]