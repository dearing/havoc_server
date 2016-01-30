# remember, CGO_ENABLED=0 GOOS=linux

FROM scratch
MAINTAINER Jacob Dearing <jacob.dearing@gmail.com>

ADD havoc_server havoc_server
EXPOSE 8080 8081

CMD ["/havoc_server"]
