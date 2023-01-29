## Build
FROM golang:1.19-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /binary

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /binary /binary

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/binary"]
