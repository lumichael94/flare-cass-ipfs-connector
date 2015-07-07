package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"reflect"
	"time"
)

type sessions struct {
	Session_id  gocql.UUID
	Coordinator string
	Duration    int
	Parameters  map[string]string
	Request     string
	Started_at  time.Time
}
type sessions struct {
	Session_id  gocql.UUID
	Coordinator string
	Duration    int
	Parameters  map[string]string
	Request     string
	Started_at  time.Time
}

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("172.31.28.240")
	cluster.Keyspace = "system_traces"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("error creating session: %v", err))
	}
	defer session.Close()
	//var id string

	s := sessions{}
	var test_query = `SELECT session_id FROM system_traces.sessions;`
	session.Query(test_query).Consistency(gocql.One).Scan(&s.Session_id)
	fmt.Println("Check this out: ", s.Session_id)
	fmt.Println("The type of parameter is: ", reflect.TypeOf(s.Session_id))

	var query_sessions = `SELECT session_id, coordinator, duration, parameters, request, started_at FROM system_traces.sessions;`
	iter := session.Query(query_sessions).Consistency(gocql.One).Iter()

	//session.Query(query_sessions).Consistency(gocql.One).Scan(&s.Session_id, &s.Coordinator, &s.Duration, &s.Parameters, &s.Request, &s.Started_at)
	fmt.Println("Fetched sessions")
	//fmt.Println("Session: ", &s.Session_id, s.Coordinator, s.Duration, s.Parameters, s.Request, s.Started_at)
	for iter.Scan(&s.Session_id, &s.Coordinator, &s.Duration, &s.Parameters, &s.Request, &s.Started_at) {
		fmt.Println("Session: ", &s.Session_id, s.Coordinator, s.Duration, s.Parameters, s.Request, s.Started_at)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
