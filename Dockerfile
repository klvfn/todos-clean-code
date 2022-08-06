# Builder
FROM golang:1.18-alpine as builder
RUN apk update && apk upgrade
WORKDIR /app
COPY . .
RUN go build -o api cmd/*.go

# Distribution
FROM alpine
ARG TZ="Asia/Jakarta"
RUN apk update && apk upgrade && apk --update --no-cache add tzdata \
    && cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /app 
EXPOSE 3030
COPY config.json .
COPY --from=builder /app/api /app
CMD ["/app/api"]