FROM debian:bullseye-slim

RUN apt-get update \
    && apt-get install -y -q --no-install-recommends \
    ca-certificates \
    && apt-get clean \
    && rm -r /var/lib/apt/lists/*
RUN mkdir -p /data/hcm/etc/certs
RUN mkdir -p /data/hcm/logs

COPY bk-hcm-hcservice /data/hcm/

CMD ["/data/hcm/bk-hcm-hcservice", "--config-file", "/data/hcm/etc/config.yaml"]
