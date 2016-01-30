# remember, CGO_ENABLED=0 GOOS=linux

FROM scratch
MAINTAINER Jacob Dearing <jacob.dearing@gmail.com>

ADD havoc_server havoc_server
ENV PORT 8000
EXPOSE 80

CMD ["/havoc_server"]
