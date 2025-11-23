FROM golang:1.24-alpine AS build

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o app .


FROM alpine:3.19

RUN apk add --no-cache tzdata
ENV TZ=Asia/Bangkok

WORKDIR /app

COPY --from=build /app/app .

COPY --from=build /app/config.yaml /app/config.yaml

EXPOSE 8000

CMD ["/app/app"]
