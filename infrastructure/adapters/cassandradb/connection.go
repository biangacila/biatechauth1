package cassandradb

import (
	"fmt"
	"github.com/gocql/gocql"
)

var session *gocql.Session

func InitSession(host, keyspace string) {
	cluster := gocql.NewCluster(host)
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = keyspace
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("cassandra connection err: ", err)
	}
}

func GetSession() *gocql.Session {
	return session
}
func CloseSession() {
	session.Close()
}
