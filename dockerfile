FROM golang:1.15 as builder


RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o sshforever .

FROM scratch
COPY --from=builder /build/sshforever /app/
WORKDIR /app

ENTRYPOINT [ "./sshforever" ]