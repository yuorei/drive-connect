FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build  -o drive-connect
CMD [ "./drive-connect" ]

EXPOSE 50051