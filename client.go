package main

import (
	"geisterchor.com/gctsdb/cassandra"

	"github.com/gocql/gocql"
)

type GCTSDBClient struct {
	cClient *cassandra.CassandraClient
}

func (c *GCTSDBClient) GetCSession() *gocql.Session {
	return c.cClient.GetSession()
}

func CreateGCTSDBClient(cassandraClient *cassandra.CassandraClient) (*GCTSDBClient, error) {
	c := GCTSDBClient{
		cClient: cassandraClient,
	}
	return &c, nil
}
