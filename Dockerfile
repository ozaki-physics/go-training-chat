FROM golang:1.15

ARG REPOSITORY
ENV GOPATH="/go"
WORKDIR /go/src/$REPOSITORY
COPY ./go.mod .
RUN go mod download
CMD ["bash"]
