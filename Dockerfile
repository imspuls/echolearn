FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

# Expose the port 3000
EXPOSE 3000:3000

ENTRYPOINT [ "/app/binary" ]