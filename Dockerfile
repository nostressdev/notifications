FROM foundationdb/foundationdb:6.0.18 as fdb
FROM golang:1.16 as builder

RUN wget --no-check-certificate https://www.foundationdb.org/downloads/6.0.18/ubuntu/installers/foundationdb-clients_6.0.18-1_amd64.deb
RUN dpkg -i foundationdb-clients_6.0.18-1_amd64.deb
RUN apt update
RUN apt install -y dnsutils

COPY --from=fdb /var/fdb/scripts/create_cluster_file.bash /

RUN apt update
RUN apt install -y protobuf-compiler
RUN go get -u github.com/golang/protobuf/protoc-gen-go@latest
RUN go get -u google.golang.org/grpc@latest
RUN GO111MODULE=off go get -u github.com/envoyproxy/protoc-gen-validate
RUN cd ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate && make

WORKDIR /service

COPY ./.libs /service/.libs
COPY ./proto /service/proto
COPY ./go.mod /service
RUN go mod download

COPY . /service

RUN make

CMD [ "bash", "-c", "./docker-entrypoint.sh" ]