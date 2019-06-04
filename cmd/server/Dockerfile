FROM golang:alpine3.9 as builder

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    go get github.com/bcmmbaga/microsvc/pkg/taskservice && \
    go get google.golang.org/grpc && \
    go get github.com/bcmmbaga/microsvc/pkg/taskpb

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo -o /go/bin/server .

FROM alpine:3.9

WORKDIR /server

COPY --from=builder /go/bin/server .

EXPOSE 3000

ENTRYPOINT [ "./server" ]