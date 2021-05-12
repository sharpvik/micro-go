FROM golang:1.16-alpine3.12 AS micro_builder
RUN mkdir /micro
WORKDIR /micro
RUN apk --no-cache add ca-certificates
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on
RUN go build -o bin
EXPOSE 8000
CMD [ "/micro/bin" ]