FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /workdir/genshincalendar
COPY . .
RUN go build cmd/main/main.go

# runtime image
FROM alpine
COPY --from=builder /workdir/genshincalendar /app

CMD /app/main $PORT