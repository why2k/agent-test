FROM golang:1.4
ADD . /go/
RUN go build -o agentcmd

CMD agentcmd
