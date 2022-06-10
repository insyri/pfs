FROM golang

WORKDIR /go/app/backend

COPY go.mod ./
COPY go.sum ./
COPY ./*.go ./
COPY database.env ./

RUN go mod download
RUN PATH=$PATH:$(go env GOPATH)/bin
RUN go get github.com/cosmtrek/air
RUN go install github.com/cosmtrek/air

CMD air
EXPOSE 8080