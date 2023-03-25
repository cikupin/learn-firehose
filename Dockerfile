FROM scratch

COPY bin/learn-firehose /usr/bin/learn-firehose
COPY payload/ /usr/bin/payload/

EXPOSE 3000
ENTRYPOINT ["/usr/bin/learn-firehose"]
