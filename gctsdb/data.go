package gctsdb

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/shopspring/decimal"

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

func (c *GCTSDBClient) GetDataPoints(channelName string, t1, t2 time.Time) ([]DataPoint, error) {
	channel, err := c.GetChannel(channelName)
	if err != nil {
		return nil, err
	}

	bucket1 := t1.UnixNano() - (t1.UnixNano() % int64(*channel.BucketSize))
	bucket2 := t2.UnixNano() - (t2.UnixNano() % int64(*channel.BucketSize))

	nBuckets := (bucket2-bucket1)/int64(*channel.BucketSize) + 1

	var points []DataPoint
	for i := int64(0); i < nBuckets; i++ {
		bucket := bucket1 + i*int64(*channel.BucketSize)

		q := fmt.Sprintf("SELECT unixnano, %s FROM gctsdb_data WHERE channel = ? AND bucket = ? AND unixnano >= ? AND unixnano < ? ORDER BY unixnano ASC;", channel.DataType)

		iter := c.GetCSession().Query(q, channel.Name, bucket, t1.UnixNano(), t2.UnixNano()).Iter()

		p := DataPoint{}
		var t int64
		if channel.DataType == "i32" {
			var val int32
			for iter.Scan(&t, &val) {
				p.Value = val
				p.Timestamp = time.Unix(0, t)
				points = append(points, p)
			}
		}
		if channel.DataType == "i64" {
			var val int64
			for iter.Scan(&t, &val) {
				p.Value = val
				p.Timestamp = time.Unix(0, t)
				points = append(points, p)
			}
		}
		if channel.DataType == "f32" {
			var val float32
			for iter.Scan(&t, &val) {
				p.Value = val
				p.Timestamp = time.Unix(0, t)
				points = append(points, p)
			}
		}
		if channel.DataType == "f64" {
			var val float64
			for iter.Scan(&t, &val) {
				p.Value = val
				p.Timestamp = time.Unix(0, t)
				points = append(points, p)
			}
		}
		if channel.DataType == "dec" {
			var val decimal.Decimal
			for iter.Scan(&t, &val) {
				p.Value = val
				p.Timestamp = time.Unix(0, t)
				points = append(points, p)
			}
		}
		if channel.DataType == "str" {
			var val string
			for iter.Scan(&t, &val) {
				p.Value = val
				p.Timestamp = time.Unix(0, t)
				points = append(points, p)
			}
		}
	}

	return points, nil
}
