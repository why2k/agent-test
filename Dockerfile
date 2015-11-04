FROM golang:1.4
RUN go get github.com/why2k/agent-test

CMD agent-test
