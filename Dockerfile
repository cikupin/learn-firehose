FROM scratch

COPY bin/learn-firehose /usr/bin/learn-firehose

EXPOSE 3000
ENTRYPOINT ["/usr/bin/learn-firehose"]
