ARG GOLANG_IMAGE_TAG

##### Base
FROM golang:$GOLANG_IMAGE_TAG as base

RUN apk update

##### Local
FROM base as local

WORKDIR /app

# for extension
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install -v github.com/stamblerre/gocode@latest

# formatter
RUN go install golang.org/x/tools/cmd/goimports@latest

# lint
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# task runner
RUN go install github.com/go-task/task/v3/cmd/task@latest