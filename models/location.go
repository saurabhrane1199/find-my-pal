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
func GetNearbyUsers(lat, lon float64, radius float64) ([]string, error) {
	locations, err := config.RedisClient.GeoRadius(config.Ctx, GEO_KEY, lon, lat, &redis.GeoRadiusQuery{
		Radius:      radius,
		Unit:        "km",
		WithDist:    false,
		WithCoord:   false,
		WithGeoHash: false,
	}).Result()

	if err != nil {
		return nil, err
	}

	usernames := make([]string, len(locations))
	for i, loc := range locations {
		usernames[i] = loc.Name
	}

	return usernames, nil
}
