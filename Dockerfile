# docker.io/guzmansalv/youtube_chan_stats
FROM golang as base

RUN go get -u "github.com/PuerkitoBio/goquery"
RUN go get -u "github.com/lib/pq"
RUN go get -u "github.com/imroc/req"

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s -extldflags -static" -o main src/

FROM alpine
COPY --from=base /app/main /main
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/main"]