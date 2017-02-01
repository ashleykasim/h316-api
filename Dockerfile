FROM golang:1.7.3
ADD . /go/src/recipes-api
WORKDIR /go/src/recipes-api
RUN go get && go build -o bin/recipes-api
ENTRYPOINT bin/recipes-api
EXPOSE 8080
