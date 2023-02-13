FROM golang:1.19

RUN mkdir /api
WORKDIR /api

COPY ./api/go.* ./

RUN go mod tidy

COPY ./api* ./

EXPOSE 3000

CMD ["go", "run", "main.go"]