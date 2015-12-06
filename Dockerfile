FROM ubuntu

RUN adduser --system --home /opt/geisterchor/gcTSDB gctsdb
ADD target/gcTSDB /opt/geisterchor/gcTSDB/gcTSDB
ADD static /opt/geisterchor/gcTSDB/static/

WORKDIR /opt/geisterchor/gcTSDB
USER gctsdb
CMD ["/opt/geisterchor/gcTSDB/gcTSDB"]
