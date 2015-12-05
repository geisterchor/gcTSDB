package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"

	"fmt"
	"time"
)

type Channel struct {
	Name       string
	DataType   string
	BucketSize *time.Duration
}

func NewChannel(name, datatype string, bucketSize *time.Duration) (*Channel, error) {
	types := []string{"i32", "i64", "f32", "f64", "dec", "str"}

	if !stringInSlice(datatype, types) {
		return nil, fmt.Errorf("'%s' is not a valid gcTSDB datatype", datatype)
	}

	ch := Channel{
		Name:       name,
		DataType:   datatype,
		BucketSize: bucketSize,
	}
	return &ch, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (c *GCTSDBClient) CreateChannel(channel *Channel) error {
	log.Infof("Creating channel '%s' (type: '%s', bucketSize: %s)", channel.Name, channel.DataType, channel.BucketSize)

	var tmp string
	err := c.GetCSession().Query(`
        SELECT channel FROM gctsdb_index WHERE pk = 0 AND channel = ?;
      `, channel.Name).Scan(&tmp)
	if err == nil {
		return fmt.Errorf("Channel '%s' already exists.", channel.Name)
	}
	if err != gocql.ErrNotFound {
		return err
	}

	if err := c.GetCSession().Query(`
      INSERT INTO gctsdb_index (pk, channel, type, bucket_size) VALUES (0, ?, ?, ?);
    `, channel.Name, channel.DataType, channel.BucketSize).Exec(); err != nil {
		return fmt.Errorf("Could not create channel '%s': %s", channel.Name, err)
	}

	return nil
}

func (c *GCTSDBClient) DeleteChannel(channel string) error {
	log.Infof("Deleting channel '%s'", channel)
	var tmp string
	err := c.GetCSession().Query(`
          SELECT channel FROM gctsdb_index WHERE pk = 0 AND channel = ?;
        `, channel).Scan(&tmp)
	if err != nil {
		return err
	}

	iter := c.GetCSession().Query("SELECT bucket FROM gctsdb_buckets WHERE channel = ?;", channel).Iter()

	var bucket int64
	for iter.Scan(&bucket) {
		if err := c.GetCSession().Query("DELETE FROM gctsdb_data WHERE channel = ? AND bucket = ?;", channel, bucket).Exec(); err != nil {
			return err
		}
	}

	if err := c.GetCSession().Query("DELETE FROM gctsdb_buckets WHERE channel = ?;", channel).Exec(); err != nil {
		return err
	}

	if err := c.GetCSession().Query("DELETE FROM gctsdb_index WHERE pk = 0 AND channel = ?;", channel).Exec(); err != nil {
		return err
	}

	return nil
}

func (c *GCTSDBClient) GetChannel(name string) (*Channel, error) {
	var ch Channel
	if err := c.GetCSession().Query("SELECT channel, type, bucket_size FROM gctsdb_index WHERE pk = 0 AND channel = ?;",
		name).Scan(&ch.Name, &ch.DataType, &ch.BucketSize); err != nil {
		return nil, err
	}

	return &ch, nil
}

func (c *GCTSDBClient) GetChannels(prefix string) []Channel {
	incrPrefix := []byte(prefix)
	incrPrefix[len(incrPrefix)-1] += 1
	iter := c.GetCSession().Query("SELECT channel, type, bucket_size FROM gctsdb_index WHERE pk = 0 AND channel >= ? AND channel < ?;", prefix, string(incrPrefix)).Iter()

	channels := []Channel{}
	var ch Channel
	for iter.Scan(&ch.Name, &ch.DataType, &ch.BucketSize) {
		channels = append(channels, ch)
	}

	return channels
}
