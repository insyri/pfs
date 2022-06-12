FROM golang

WORKDIR /go/app/backend

COPY . .

RUN go mod download
RUN PATH=$PATH:$(go env GOPATH)/bin
RUN go get github.com/cosmtrek/air
RUN go install github.com/cosmtrek/air

CMD air
# CMD go run main.go
EXPOSE 8080