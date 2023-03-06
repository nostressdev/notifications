FROM golang:1.17 as builder

RUN wget --no-check-certificate https://github.com/apple/foundationdb/releases/download/6.3.22/foundationdb-clients_6.3.22-1_amd64.deb
RUN dpkg -i foundationdb-clients_6.3.22-1_amd64.deb
RUN apt update
RUN apt install -y dnsutils

#RUN apt update
#RUN apt install -y protobuf-compiler
#RUN go get -u github.com/golang/protobuf/protoc-gen-go@latest
#RUN go get -u google.golang.org/grpc@latest
#RUN GO111MODULE=off go get -u github.com/envoyproxy/protoc-gen-validate
#RUN cd ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate && make

WORKDIR /service

COPY ./go.mod /service
COPY ./firebase_creds.json /service
RUN go mod download

COPY . /service

RUN make bin/notifications

FROM ubuntu:20.04
RUN apt update
RUN apt install -y wget
RUN apt install -y dnsutils
RUN wget --no-check-certificate https://github.com/apple/foundationdb/releases/download/6.3.22/foundationdb-clients_6.3.22-1_amd64.deb
RUN dpkg -i foundationdb-clients_6.3.22-1_amd64.deb

COPY --from=builder /service/docker-entrypoint.sh /
COPY --from=builder /service/bin/notifications /bin/service

CMD [ "bash", "-c", "./bin/service" ]