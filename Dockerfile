FROM golang:latest

WORKDIR /go/src/github.com/WardrobeGenerator
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

CMD ["app"]