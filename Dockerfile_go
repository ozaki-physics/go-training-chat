FROM golang:1.15

ARG REPOSITORY

# go mod を使うために GOPATH を削除するために path を上書き
ENV GOPATH=""
WORKDIR /go/src/$REPOSITORY
RUN go mod init $REPOSITORY

CMD ["bash"]
