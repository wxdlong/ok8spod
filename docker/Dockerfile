FROM golang:alpine as builder
COPY ./ /home
RUN go env -w GO111MODULE="on" && cd /home && \
    go build -ldflags "-s" k8sweb.go

FROM alpine:latest
COPY  --from=builder /home/k8sweb /home/k8sweb
RUN chmod +x /home/k8sweb
EXPOSE 80
CMD ["/home/k8sweb"]


