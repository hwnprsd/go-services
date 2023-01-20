# Choose whatever you want, version >= 1.16
FROM golang:1.19-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod tidy 

COPY . ./
#
CMD ["air", "-c", ".air.toml"]
