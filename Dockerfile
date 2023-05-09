FROM golang:1.20 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o notification cmd/notification/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=build app/notification .

EXPOSE 80

CMD ["/notification"]
