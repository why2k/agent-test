FROM golang:1.4
ADD .
RUN go build -o agentcmd

CMD agentcmd
