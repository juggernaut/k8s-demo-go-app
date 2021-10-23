FROM golang:1.17

WORKDIR /go/src/app
COPY . .
RUN go install -v
EXPOSE 8080

CMD ["k8s-demo-go-app"]
