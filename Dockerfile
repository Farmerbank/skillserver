FROM golang:1.8-alpine

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
EXPOSE 3000
CMD ["go-wrapper", "run"] # ["app"]
