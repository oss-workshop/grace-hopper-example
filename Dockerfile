FROM golang:1.14 as builder


COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /dist/main ./main.go

FROM scratch

COPY --from=builder /dist/main /app/main
ENTRYPOINT ["/app/main"]