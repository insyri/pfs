FROM golang

WORKDIR /go/app/backend

COPY go.mod ./
COPY ./*.go ./

RUN go mod download
RUN go build -o ./pfs

CMD "./pfs"
EXPOSE 8080