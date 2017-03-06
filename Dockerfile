FROM golang:1.7.3
ADD . /go/src/h316-api
WORKDIR /go/src/h316-api
RUN go get && go build -o bin/h316-api
ENTRYPOINT bin/h316-api
EXPOSE 8080
