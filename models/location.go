package models

import (
	"findmypal/config"

	"github.com/redis/go-redis/v9"
)

const GEO_KEY = "user_locations"

// Store user location as a geohash in Redis
func StoreUserLocation(username string, lat, lon float64) error {
	return config.RedisClient.GeoAdd(config.Ctx, GEO_KEY, &redis.GeoLocation{
		Name:      username,
		Longitude: lon,
		Latitude:  lat,
	}).Err()
}

// Get nearby users within a radius (e.g., 5km)
func GetNearbyUsers(username string, radius float64) ([]string, error) {
	locations, err := config.RedisClient.GeoSearch(config.Ctx, GEO_KEY, &redis.GeoSearchQuery{
		Member:     username,
		Radius:     radius,
		RadiusUnit: "km",
		Sort:       "ASC", // Sort by closest first
	}).Result()

	if err != nil {
		return nil, err
	}
	return locations, nil
}
