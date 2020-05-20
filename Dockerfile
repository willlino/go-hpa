FROM golang:alpine as build

WORKDIR /src
ADD ./src /src

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch
WORKDIR /usr/bin
COPY --from=build /src /usr/bin

CMD ["./src"]