package main

import (
	"geisterchor.com/gctsdb/gctsdb"

	log "github.com/Sirupsen/logrus"

	"time"
)

func example(gctsdbServer *gctsdb.GCTSDBServer) {
	bucketSize := 7 * 24 * time.Hour
	ch, err := gctsdb.NewChannel("temperature", "f32", &bucketSize)
	if err != nil {
		log.Errorf(err.Error())
	}
	if err := gctsdbServer.CreateChannel(ch); err != nil {
		log.Errorf("Could not create channel: %s", err)
	}

	channels := gctsdbServer.GetChannels("temp")
	log.Infof("I have found the following channels: %s", channels)

	points := []gctsdb.DataPoint{
		gctsdb.DataPoint{
			Timestamp: time.Now(),
			Value:     float32(4.5),
		},
	}
	if err := gctsdbServer.AddDataPoints(ch, points); err != nil {
		log.Errorf("Could not add data point: %s", err)
	}

	ps, err := gctsdbServer.GetDataPoints("temperature", time.Now().Add(-time.Hour), time.Now().Add(time.Hour))
	if err != nil {
		log.Errorf("Could not get data points: %s", err)
	}
	log.Infof("Got %d points: %s", len(ps), ps)

	if err := gctsdbServer.DeleteChannel("temperature"); err != nil {
		log.Errorf("Could not delete channel: %s", err)
	}
}
