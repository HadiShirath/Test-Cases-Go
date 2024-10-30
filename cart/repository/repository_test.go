package repository_test

import (
	"context"
	"test_cases/cart/repository"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// integration test cases
func TestRepository_AddToCart_Success(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// connect to real caching system - redis
	redisOptions := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	redisClient := redis.NewClient(redisOptions)

	// initiate cache repository
	repo := repository.New(redisClient)

	// execute repository methods
	err := repo.AddToCart(context.TODO(), "userID", "productID")

	// assert the error should be nil
	assert.NoError(t, err, "it should not return error, redis container is there")

	// assert value of store cache
	result, err := redisClient.HGet(context.Background(), "cart-userID", "name").Result()
	assert.NoError(t, err, "it should not return error")
	assert.Equal(t, result, "Sepatu lokal dari UMKM", "it should return the expected cached product item name")

}
