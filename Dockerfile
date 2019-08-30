FROM golang:latest as builder
WORKDIR /go/src/github.com/feiyuWeb/ginApi
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . .
RUN go get
# Build the Go app
#RUN CGO_ENABLED=0 go build -a -ldflags="-w -s"
RUN go build -o main .

FROM scratch as final
MAINTAINER feiyu <709463253@qq.com>
WORKDIR /ginApi
COPY --from=builder /go/src/github.com/feiyuWeb/ginApi/config /ginApi/config
COPY --from=builder /go/src/github.com/feiyuWeb/ginApi/controls /ginApi/controls
COPY --from=builder /go/src/github.com/feiyuWeb/ginApi/docs /ginApi/docs
COPY --from=builder /go/src/github.com/feiyuWeb/ginApi/models /ginApi/models
EXPOSE 8080
ENTRYPOINT ["/ginApi"]
