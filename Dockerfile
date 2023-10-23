FROM golang:1.18-alpine3.17
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .
CMD ./out/dist
EXPOSE 80
EXPOSE 8080