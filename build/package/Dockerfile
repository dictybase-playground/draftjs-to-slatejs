FROM golang:1.13.8-alpine3.10
LABEL maintainer="Eric Hartline <eric.hartline@northwestern.edu>"
ENV GOPROXY https://proxy.golang.org
RUN apk add --no-cache git build-base
RUN mkdir -p /draftjs-to-slatejs
WORKDIR /draftjs-to-slatejs
# set up and build go app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD cmd cmd
ADD internal internal
RUN go build -o app cmd/main.go
# set up nodejs
RUN apk add --update nodejs nodejs-npm
COPY package*.json ./
RUN npm install
ADD nodejs nodejs
ENTRYPOINT [ "./app" ]