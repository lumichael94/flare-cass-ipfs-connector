package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("172.31.28.240")
	cluster.Keyspace = "system_traces"
	session, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("error creating session: %v", err))
	}
	defer session.Close()

	var query_sessions = `SELECT session_id, coordinator, duration, parameters, request, started_at FROM system_traces.sessions;`
	q, err := session.Query(query_sessions)
	if err != nil {
		log.Fatal(err)
	}
	q.Scan(&sub.Session_id, &sub.Coordinator, &sub.Duration, &sub.Parameters, &sub.Request, &sub.Started_at)
	fmt.Println("Fetched sessions")
	fmt.Println(sub)
}
