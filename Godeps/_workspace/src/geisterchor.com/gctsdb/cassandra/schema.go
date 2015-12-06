package cassandra

import (
	"github.com/gocql/gocql"
)

func InitializeSchema(cClient *CassandraClient) error {
	if err := cClient.GetSession().Query(`CREATE TABLE IF NOT EXISTS gctsdb_index (
      pk int,
      channel text,
      type text,
      bucket_size bigint,
      PRIMARY KEY (pk, channel)
    );`).Consistency(gocql.Quorum).Exec(); err != nil {
		return err
	}

	if err := cClient.GetSession().Query(`CREATE TABLE IF NOT EXISTS gctsdb_buckets (
      channel text,
      bucket bigint,
      PRIMARY KEY (channel, bucket)
    );`).Consistency(gocql.Quorum).Exec(); err != nil {
		return err
	}

	if err := cClient.GetSession().Query(`CREATE TABLE IF NOT EXISTS gctsdb_data (
      channel text,
      bucket bigint,
      unixnano bigint,
      i32 int,
      i64 bigint,
      f32 float,
      f64 double,
      dec decimal,
      str text,
      PRIMARY KEY ((channel, bucket), unixnano)
    );`).Consistency(gocql.Quorum).Exec(); err != nil {
		return err
	}
	return nil
}
