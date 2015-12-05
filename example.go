package main

import (
	log "github.com/Sirupsen/logrus"

	"time"
)

func example(gctsdbClient *GCTSDBClient) {
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

	ps, err := gctsdbClient.GetDataPoints("temperature", time.Now().Add(-time.Hour), time.Now().Add(time.Hour))
	if err != nil {
		log.Errorf("Could not get data points: %s", err)
	}
	log.Infof("Got %d points: %s", len(ps), ps)

	if err := gctsdbClient.DeleteChannel("temperature"); err != nil {
		log.Errorf("Could not delete channel: %s", err)
	}
}
