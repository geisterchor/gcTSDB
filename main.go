package main

import (
	"geisterchor.com/gctsdb/cassandra"
	"geisterchor.com/gctsdb/gctsdb"

	log "github.com/Sirupsen/logrus"

	"os"
	"strings"
)

func main() {
	initializeLogger()
	log.Infof("gcTSDB - geisterchor Time Series Database")

	cassandraHosts := strings.Split(os.Getenv("CASSANDRA_HOSTS"), ",")
	cClient, err := cassandra.CreateCassandraClient(
		cassandraHosts,
		os.Getenv("CASSANDRA_USER"),
		os.Getenv("CASSANDRA_PASSWORD"),
		os.Getenv("CASSANDRA_KEYSPACE"),
	)
	if err != nil {
		log.Fatalf("Could not connect to Cassandra: %s", err)
	}

	if err := cassandra.InitializeSchema(cClient); err != nil {
		log.Fatalf("Could not create Cassandra schema: %s", err)
	}

	gctsdbClient, err := gctsdb.CreateGCTSDBClient(cClient)
	if err != nil {
		log.Fatalf("Could not create gcTSDB client: %s", err)
	}

	example(gctsdbClient)
}
