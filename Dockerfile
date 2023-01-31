FROM golang:1.19

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV APP_ENV production

COPY . .
RUN go build -v -o bot .

CMD ["/app/bot"]
