FROM golang:1.18-alpine as BUILD

WORKDIR $GOPATH/src/github.com/ramdanariadi/grocery-product-service

COPY . .
RUN go mod download
RUN go build -o /app

EXPOSE 9090
CMD [ "/app" ]