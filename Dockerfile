FROM golang:1.4

RUN go build -o agentcmd

CMD agentcmd
