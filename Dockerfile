FROM golang:1.4
RUN apt-get update && apt-get install -y autoconf file g++ gcc libc-dev make pkg-config git

RUN go get github.com/why2k/agent-test

CMD agent-test
