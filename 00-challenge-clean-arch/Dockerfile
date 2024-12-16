FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o challenge cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/challenge .
COPY --from=builder /app/cmd/ordersystem/.env .


RUN apk --no-cache add ca-certificates

EXPOSE 8080 8000 50051
CMD ["./challenge"]