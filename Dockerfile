FROM golang:1.19 AS BUILD
ARG Release=dev

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV APP_ENV production

COPY . .
RUN go build -ldflags="-X 'main.Release=${Release}'" -v -o bot .

# CMD ["/app/bot"]

FROM gcr.io/distroless/base-debian10 as DEPLOY

WORKDIR /app

COPY --from=BUILD /app/bot /app/bot

USER nonroot:nonroot

ENV APP_ENV production

CMD ["/app/bot"]
