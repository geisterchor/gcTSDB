package gctsdb

import (
	"geisterchor.com/gcTSDB/cassandra"

	"github.com/gocql/gocql"
)

type GCTSDBServer struct {
	cClient *cassandra.CassandraClient
}

func (c *GCTSDBServer) GetCSession() *gocql.Session {
	return c.cClient.GetSession()
}

func CreateGCTSDBServer(cassandraClient *cassandra.CassandraClient) (*GCTSDBServer, error) {
	c := GCTSDBServer{
		cClient: cassandraClient,
	}
	return &c, nil
}
