FROM alpine:latest
RUN apk add go
WORKDIR /app
COPY /webapp.go .
RUN go build -o main  webapp.go
CMD ["/app/main"]
