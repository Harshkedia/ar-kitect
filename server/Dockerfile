# build server
FROM golang:1.15-rc-alpine AS gobuilder
RUN apk update && apk add --no-cache git make build-base

WORKDIR $GOPATH/src/app

COPY go.mod .
COPY server.go .
COPY ./haikunator ./haikunator

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/server

# build web viewer
FROM node:stretch-slim AS webbuilder
WORKDIR /app
COPY ./web  .
RUN npm install
RUN npm run build

# get usdz container
FROM leon/usd-from-gltf

RUN apt-get update && apt-get install npm -y && npm install -g obj2gltf
# put it together
COPY --from=webbuilder /app/dist ./static
COPY --from=gobuilder /go/bin/server .

COPY ./FBX2glTF .
COPY ./FBX2glTF .
COPY ./sample-models ./models

# ENV CERT_DOMAIN = ar.protfo.io
# ENV APP_STATIC_PATH = usr/app/static
# ENV MODELS_PATH     = usr/app/models

ENTRYPOINT [ "/usr/bin/env" ]

RUN chmod +x ./server
RUN chmod +x ./FBX2glTF

CMD [ "./server" ]
