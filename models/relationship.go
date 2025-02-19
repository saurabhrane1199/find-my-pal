package models

import (
	"findmypal/config"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Send a friend request
func SendFriendRequest(sender, receiver string) error {
	session := config.Neo4jDriver.NewSession(config.Ctx, neo4j.SessionConfig{})
	defer session.Close(config.Ctx)

	_, err := session.ExecuteWrite(config.Ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MERGE (a:User {username: $sender}) MERGE (b:User {username: $receiver})
		          MERGE (a)-[:SENT_REQUEST]->(b)`
		_, err := tx.Run(config.Ctx, query, map[string]interface{}{
			"sender":   sender,
			"receiver": receiver,
		})
		return nil, err
	})

	if err != nil {
		return fmt.Errorf("failed to send friend request: %v", err)
	}
	return nil
}

// Accept friend request
func AcceptFriendRequest(sender, receiver string) error {
	session := config.Neo4jDriver.NewSession(config.Ctx, neo4j.SessionConfig{})
	defer session.Close(config.Ctx)

	_, err := session.ExecuteWrite(config.Ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MATCH (a:User {username: $sender})-[r:SENT_REQUEST]->(b:User {username: $receiver})
		          DELETE r
		          MERGE (a)-[:FRIENDS_WITH]->(b)
		          MERGE (b)-[:FRIENDS_WITH]->(a)`
		_, err := tx.Run(config.Ctx, query, map[string]interface{}{
			"sender":   sender,
			"receiver": receiver,
		})
		return nil, err
	})

	if err != nil {
		return fmt.Errorf("failed to accept friend request: %v", err)
	}
	return nil
}

// Get list of friends
func GetFriends(username string) ([]string, error) {
	session := config.Neo4jDriver.NewSession(config.Ctx, neo4j.SessionConfig{})
	defer session.Close(config.Ctx)

	friends := []string{}
	_, err := session.ExecuteRead(config.Ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `MATCH (a:User {username: $username})-[:FRIENDS_WITH]-(b:User) RETURN b.username`
		result, err := tx.Run(config.Ctx, query, map[string]interface{}{
			"username": username,
		})
		if err != nil {
			return nil, err
		}

		for result.Next(config.Ctx) {
			record := result.Record()
			if friend, ok := record.Get("b.username"); ok {
				friends = append(friends, friend.(string))
			}
		}
		return nil, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to fetch friends: %v", err)
	}
	return friends, nil
}
