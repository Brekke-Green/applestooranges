FROM golang:1.20

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o main .

CMD ./main
