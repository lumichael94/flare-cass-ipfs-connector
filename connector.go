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
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	/* Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */
	var query_sessions = "COPY sessions (session_id, coordinator, duration, parameters, request, started_at) TO '~/tracing_session.csv' WITH DELIMITER = '|' AND NULL = '<null>';"

	var query_events = "COPY events (session_id, event_id, activity, source, source_elapsed, thread) TO '~/tracing_events.csv' WITH DELIMITER = '|' AND NULL = '<null>';"
	if err := session.Query(query_sessions).Consistency(gocql.One); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched sessions")

	if err := session.Query(query_events).Consistency(gocql.One); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched events")
}
