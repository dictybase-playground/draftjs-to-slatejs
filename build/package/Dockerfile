FROM golang:1.13.8-alpine3.10
LABEL maintainer="Eric Hartline <eric.hartline@northwestern.edu>"
ENV GOPROXY https://proxy.golang.org
RUN apk add --no-cache git build-base
RUN mkdir -p /draftjs-to-slatejs
WORKDIR /draftjs-to-slatejs
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
# set up binary command (node-cli)
RUN npm link

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=0 /draftjs-to-slatejs/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]