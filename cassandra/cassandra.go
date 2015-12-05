package cassandra

import (
	"github.com/gocql/gocql"
)

type CassandraClient struct {
	session *gocql.Session
}

func (c *CassandraClient) GetSession() *gocql.Session {
	return c.session
}

func CreateCassandraClient(hosts []string, username, password, keyspace string) (*CassandraClient, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Consistency = gocql.One

	if len(keyspace) == 0 {
		keyspace = "gctsdb"
	}
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	c := CassandraClient{
		session: session,
	}
	return &c, nil
}
