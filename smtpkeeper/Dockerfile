# Build binary
FROM golang:1.15 AS builder
COPY . /code
WORKDIR /code
RUN CGO_ENABLED=0 go build

# Build final image with a binary and necessary resources
FROM alpine
COPY --from=builder /code/smtpkeeper /opt/smtpkeeper/
WORKDIR /opt/smtpkeeper
EXPOSE 8080
ENTRYPOINT ["/opt/smtpkeeper/smtpkeeper"]
