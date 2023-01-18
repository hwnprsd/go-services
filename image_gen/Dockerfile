FROM golang:1.19-alpine as build
WORKDIR /src 
# Copy the two important files
COPY go.mod /src 
COPY go.sum /src 
# Download all the deps
RUN go mod download 
COPY *.go ./
# Build the executable
RUN go build -o /exe
# # Run the executable
# CMD [ "/exe" ]

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /exe /exe

USER nonroot:nonroot

ENTRYPOINT ["/exe"]

