package main

import (
	"geisterchor.com/gctsdb/cassandra"

	log "github.com/Sirupsen/logrus"

	"os"
	"strings"
	"time"
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

	gctsdbClient, err := CreateGCTSDBClient(cClient)
	if err != nil {
		log.Fatalf("Could not create gcTSDB client: %s", err)
	}

	bucketSize := 7 * 24 * time.Hour
	ch, err := NewChannel("temperature", "f32", &bucketSize)
	if err != nil {
		log.Errorf(err.Error())
	}
	if err := gctsdbClient.CreateChannel(ch); err != nil {
		log.Errorf("Could not create channel: %s", err)
	}

	channels := gctsdbClient.GetChannels("temp")
	log.Infof("I have found the following channels: %s", channels)

	points := []DataPoint{
		DataPoint{
			Timestamp: time.Now(),
			Value:     float32(4.5),
		},
	}
	if err := gctsdbClient.AddDataPoints(ch, points); err != nil {
		log.Errorf("Could not add data point: %s", err)
	}

	if err := gctsdbClient.DeleteChannel("temperature"); err != nil {
		log.Errorf("Could not delete channel: %s", err)
	}

}
