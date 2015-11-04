FROM golang:1.4
RUN apt-get update && apt-get install -y autoconf file g++ gcc libc-dev make pkg-config git
RUN wget https://github.com/libgit2/libgit2/archive/v0.23.2.tar.gz && tar xzf v0.23.2.tar.gz && cd libgit2-0.23.2/ && cmake . && make && make install

RUN go get github.com/why2k/agent-test

CMD agent-test
