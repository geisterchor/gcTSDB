package cassandra

import (
	"github.com/gocql/gocql"
)

func InitializeSchema(cClient *CassandraClient) error {
	if err := cClient.GetSession().Query(`CREATE TABLE IF NOT EXISTS gctsdb_index (
      category text,
      channel text,
      type text static,
      bucket_size bigint static,
      buckets bigint,
      PRIMARY KEY (category, channel, buckets)
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
