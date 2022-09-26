FROM golang:1.19-buster

# General setup.
EXPOSE 8000

# File structure.
RUN mkdir /app
WORKDIR /app
COPY . .

# Compile.
RUN go install

CMD [ "micro-go" ]
