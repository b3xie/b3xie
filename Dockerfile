FROM alpine as base
WORKDIR /
COPY . .
RUN apk add go
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /cmd/web/server ./cmd/web/server.go
COPY /cmd/web/ ./
EXPOSE 1329
ENTRYPOINT ./cmd/web/server