FROM golang:1.24.0
WORKDIR /app/golang
COPY . ./
RUN go mod download
RUN go get
CMD go run main.go