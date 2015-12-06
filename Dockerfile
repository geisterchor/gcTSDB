FROM ubuntu

RUN adduser --system --home /opt/geisterchor geisterchor
ADD target/gcTSDB /opt/geisterchor/gcTSDB

WORKDIR /opt/geisterchor
USER geisterchor
CMD ["/opt/geisterchor/gcTSDB"]
