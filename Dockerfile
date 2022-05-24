FROM golang:1.18-alpine3.15
RUN mkdir /micro
WORKDIR /micro
RUN apk --no-cache add ca-certificates
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on
RUN go build -o gw
EXPOSE 8000
CMD [ "/micro/gw" ]