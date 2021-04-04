FROM golang:1.16-alpine3.13 AS builder

RUN apk update 
WORKDIR /go/src/mockWorker/
COPY ./mockWorker.go /go/src/mockWorker/
COPY ./go.mod /go/src/mockWorker/
RUN go build -o /bin/mockWorker

#############################################
FROM yobasystems/alpine-docker
RUN apk update

COPY --from=builder /bin/mockWorker /bin/mockWorker
ENTRYPOINT [ "/bin/mockWorker" ]
