package config

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var Neo4jDriver neo4j.DriverWithContext
var ctx = context.Background()

func InitNeo4j() {
	var err error
	Neo4jDriver, err = neo4j.NewDriverWithContext(
		"bolt://localhost:7687",
		neo4j.BasicAuth("neo4j", "Saurabh99#", ""),
	)

	if err != nil {
		log.Fatal("Error connecting to Neo4j:", err)
	}

	err = Neo4jDriver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal("Neo4j connectivity check failed:", err)
	}
}
