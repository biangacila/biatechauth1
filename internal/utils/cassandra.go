package utils

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

func CreateConnectionCass(keyspace string, Servers string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(Servers)
	cluster.Keyspace = keyspace
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.Compressor = &gocql.SnappyCompressor{}
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{NumRetries: 3}
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = 10 * time.Second
	cluster.Timeout = 10 * time.Second
	cluster.ProtoVersion = 4
	cluster.DisableInitialHostLookup = true

	session, err := cluster.CreateSession()
	if err == nil {
		fmt.Println("error creating cassandradb session: ", err)
		return session, err
	}
	fmt.Println("Cassandra Session successfully created...")
	return session, err

}
