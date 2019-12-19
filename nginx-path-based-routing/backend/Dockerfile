# [multi stage build] the build stage
FROM golang:1.13-alpine3.10 as builder
LABEL maintainer="kenju <ken901waga@gmail.com>"

# install prerequisisite packages
RUN apk update && apk upgrade && apk add git gcc musl-dev

# build go binary
COPY . /app
WORKDIR /app
RUN touch REVISION
RUN go build -o backend-service


# [multi stage build] the runtime stage
FROM alpine:3.10

RUN apk update && apk upgrade && apk add ca-certificates

WORKDIR /go

# copy the minimum necessary failes from the build stage
COPY --from=builder /app/backend-service /app/REVISION ./

EXPOSE 8080
CMD ["/go/backend-service"]
