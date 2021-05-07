FROM golang:1.16.3-alpine3.13 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
## Pull in any dependencies
RUN go mod download
RUN go mod verify
## Build with the necessary go libraries included.
RUN go build -o main ./cmd/agartha/agartha.go
## Start created binary executable
FROM scratch
COPY --from=builder /app/main /bin/agartha
ENTRYPOINT ["/bin/agartha"]
