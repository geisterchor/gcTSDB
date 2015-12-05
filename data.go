package main

import (
	//	"github.com/shopspring/decimal"
	//log "github.com/Sirupsen/logrus"

	"fmt"
	"time"
)

type DataPoint struct {
	Timestamp time.Time
	Value     interface{}
}

func (c *GCTSDBClient) AddDataPoints(channel *Channel, points []DataPoint) error {
	buckets := map[int64]bool{}

	for _, p := range points {
		bucket := p.Timestamp.UnixNano() - (p.Timestamp.UnixNano() % int64(*channel.BucketSize))
		buckets[bucket] = true

		s := c.GetCSession()
		qTempl := fmt.Sprintf("INSERT INTO gctsdb_data (channel, bucket, unixnano, %s) VALUES (?, ?, ?, ?);", channel.DataType)
		q := c.GetCSession().Query("")

		succ := false
		if val, ok := p.Value.(float32); ok {
			q = s.Query(qTempl, channel.Name, bucket, p.Timestamp.UnixNano(), val)
			succ = true
		}
		if val, ok := p.Value.(float64); ok {
			q = s.Query(qTempl, channel.Name, bucket, p.Timestamp.UnixNano(), val)
			succ = true
		}

		if !succ {
			return fmt.Errorf("Could not detect value type")
		}

		if err := q.Exec(); err != nil {
			return err
		}
	}

	for bucket, _ := range buckets {
		if err := c.GetCSession().Query("INSERT INTO gctsdb_buckets (channel, bucket) VALUES (?, ?)", channel.Name, bucket).Exec(); err != nil {
			return err
		}
	}

	return nil
}
