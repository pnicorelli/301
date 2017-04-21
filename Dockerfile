FROM golang

RUN mkdir -p /usr/local/301

ENV GOPATH /usr/local/301
ENV GOBIN /usr/local/301


COPY . /usr/local/301
WORKDIR /usr/local/301

RUN go get ./
RUN go build 301.go

CMD ["./301"]

EXPOSE 8080
