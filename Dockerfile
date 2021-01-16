# Builder
FROM golang:1.15-alpine as builder
RUN apk update && apk upgrade
WORKDIR /app
COPY . .
RUN go build -o api ./*.go

# Distribution
FROM alpine
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata
WORKDIR /app 
EXPOSE 3030
COPY .env .
COPY --from=builder /app/api /app
CMD /app/api