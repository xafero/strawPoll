FROM golang:alpine AS build-env
ADD . /go/src/github.com/hibooboo2/strawpoll/
RUN ls /go/src/github.com/hibooboo2/strawpoll/ \
    && apk update \
    && apk upgrade \
    && apk add --no-cache git \
    && go get "github.com/gorilla/mux" \
    && go get "github.com/nanobox-io/golang-scribble" \
    && go get "github.com/skratchdot/open-golang/open" \
    && cd /go/src/github.com/hibooboo2/strawpoll/ \
    && go build -o goapp

FROM alpine
EXPOSE 8080
WORKDIR /app
COPY --from=build-env /go/src/github.com/hibooboo2/strawpoll/goapp /app/
ENTRYPOINT ./goapp
