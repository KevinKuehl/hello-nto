#
# Copyright 2019 Kevin Kuehl
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.
#
FROM golang:1.11 AS build-env

COPY . /go/src/github.com/kevinkuehl/hello-nto
WORKDIR /go/src/github.com/kevinkuehl/hello-nto

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/hello-nto-app

FROM scratch
COPY --from=build-env /go/bin/hello-nto-app /go/bin/

ENTRYPOINT ["/go/bin/hello-nto-app"]

EXPOSE 8080
