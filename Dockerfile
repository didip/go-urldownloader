FROM ubuntu:14.04

RUN apt-get update -q

RUN DEBIAN_FRONTEND=noninteractive apt-get install -qy build-essential curl git mercurial bzr

RUN curl -s https://storage.googleapis.com/golang/go1.3.src.tar.gz | tar -v -C /usr/local -xz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1

ENV GOPATH /go
ENV PATH   /usr/local/go/bin:$PATH

ADD . /opt/go-urldownloader
RUN cd /opt/go-urldownloader && sh ./go-get-dependencies.sh && go build .

CMD ["/opt/go-urldownloader/bin/go-urldownloader"]