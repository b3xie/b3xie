FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
COPY ./ ./
RUN go mod download
RUN cd cmd/web/
RUN ls
RUN go build -o /b3xie
CMD ["./b3xie"]
EXPOSE 1329
ENTRYPOINT ["./cmd/web/server.go"]
