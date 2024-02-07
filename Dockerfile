FROM golang:1.21 as build

WORKDIR /usr/src/app
COPY . .

RUN go mod download && \
    CGO_ENABLED=0 go build -o /usr/bin/subscribed ./cmd/subscribed

FROM gcr.io/distroless/static-debian12

COPY --from=build /usr/bin/subscribed /usr/bin/subscribed
CMD ["/usr/bin/subscribed", "--help"]
