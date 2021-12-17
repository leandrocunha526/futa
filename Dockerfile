FROM golang:1.17-bookworm

WORKDIR /src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
