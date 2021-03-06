gcTSDB - geisterchor Time Series Database
=========================================

Compiling
---------

You need to have Go and Make installed.
Clone this repo, `cd` into the directory and then execute `make` to compile and
create an executable called `gcTSDB` in the `target` directory.


Configuration
-------------

For now the configuration is passed via environment variables to the software.

| Variable           | Description                                                                   |
|--------------------|-------------------------------------------------------------------------------|
| CASSANDRA_HOSTS    | comma separated list of Cassandra nodes                                       |
| CASSANDRA_USER     | Cassandra user name                                                           |
| CASSANDRA_PASSWORD | Cassandra password                                                            |
| CASSANDRA_KEYSPACE | Cassandra keyspace (default: gctsdb)                                          |
| LOG_FORMAT         | you may set this to `JSON` to output all log messages in one-line JSON format |


Running
-------

You can execute this application directly from your console like this:

    CASSANDRA_USER=dev CASSANDRA_PASSWORD=dev ./target/gcTSDB

Or via Docker:

    docker run -d --name gctsdb -p 3000:3000 -e CASSANDRA_HOSTS=cas1,cas2,cas3 \
        -e CASSANDRA_USER=dev -e CASSANDRA_PASSWORD=dev geisterchor/gctsdb


Contributing
------------
### Dependencies
All Go dependencies are managed by [govendor](https://github.com/kardianos/govendor).
The dependency code is copied to the `vendor` directory.


License
-------
gcTSDB is licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE.md) for the full license text.

Copyright (c) 2015 the gcTSDB authors.
