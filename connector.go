package main

import (
	"fmt"
	"github.com/gocql/gocql"
//	"log"
	"time"
)

type Session struct{
	Session_id 	gocql.UUID
	Coordinator	string
	Duration	int
	//Parameters	gocql.map
	Request		string
	Started_at	time.Time
}

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("172.31.28.240")
	cluster.Keyspace = "system_traces"
	cluster.Consistency = gocql.Quorum
	session, err:= cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("error creating session: %v", err))
	}
	defer session.Close()
	var id string	
	session.Query(`SELECT session_id FROM system_traces.sessions;`).Consistency(gocql.One).Scan(&id)
	fmt.Println("Check this out: ", id)
	var query_sessions = `SELECT session_id, coordinator, duration, request, started_at FROM system_traces.sessions;`
	q:= session.Query(query_sessions)
	//if err != nil {
	//	log.Fatal(err)
	//}

	sub := Session{}
	q.Scan(&sub.Session_id, &sub.Coordinator, &sub.Duration, &sub.Request, &sub.Started_at)
	fmt.Println("Fetched sessions")
	fmt.Println(sub)
}
