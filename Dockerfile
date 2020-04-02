FROM golang as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /go/src/github.com && mkdir /go/src/github.com/racoon63 && mkdir /go/src/github.com/racoon63/go-bandwidth-monitor-api

COPY app/ /go/src/github.com/racoon63/go-bandwidth-monitor-api/app
COPY internal/ /go/src/github.com/racoon63/go-bandwidth-monitor-api/internal

RUN ls -l /go/src/github.com/racoon63/go-bandwidth-monitor-api/

WORKDIR /go/src/github.com/racoon63/go-bandwidth-monitor-api/app/go-bandwidth-monitor-api

RUN go build main.go

FROM golang

COPY --chown=bwm --from=build /go/src/github.com/racoon63/go-bandwidth-monitor-api/app/go-bandwidth-monitor-api/go-bandwidth-monitor-api /app/api

WORKDIR /app/go-bandwidth-monitor-api

RUN useradd --system --no-create-home --no-user-group bwm

USER bwm

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl -f http://localhost:8080/api/v1/" ]

ENTRYPOINT [ "/app/api" ]