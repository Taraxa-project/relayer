# builder image
FROM golang:1.21.1-alpine as gobuilder

RUN apk update && apk add curl protobuf make \
    rm -rf /var/cache/apk/*

RUN mkdir /app
COPY . /app/
WORKDIR /app

# Download Go modules
RUN go mod download

# Build
RUN make build


# executable image
FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=gobuilder /app/build/linux_amd64/relayer /relayer

# Run
CMD [ "/relayer" ]