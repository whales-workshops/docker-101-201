package main
// 🚀 >>> go test -v -p 1
import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	redisModule "github.com/testcontainers/testcontainers-go/modules/redis"
)

// go test -v -run TestGetRestaurant
func TestGetRestaurant(t *testing.T) {

	ctx := context.Background()

	redisContainer, err := redisModule.RunContainer(ctx, testcontainers.WithImage("redis:7.2.4"))
	require.NoError(t, err)

	endpoint, err := redisContainer.Endpoint(ctx, "")
	require.NoError(t, err)

	rdb := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	restaurants := []Restaurant{
		{ID: "1", Name: "Restaurant A", Address: "123 Main St", Website: "https://a.com", Phone: "123-456-7890", Tags: "Italian"},
		{ID: "2", Name: "Restaurant B", Address: "456 Elm St", Website: "https://b.com", Phone: "987-654-3210", Tags: "Mexican"},
	}

	for _, r := range restaurants {
		rdb.HSet(ctx, "restaurant:"+r.ID, map[string]interface{}{
			"name":    r.Name,
			"address": r.Address,
			"website": r.Website,
			"phone":   r.Phone,
			"tags":    r.Tags,
		})
	}

	restaurant, err := getRestaurant(ctx, rdb, "2")

	fmt.Println("🍱", restaurant)

	assert.NoError(t, err)
	assert.Equal(t, "Restaurant B", restaurant.Name)
}

// go test -v -run TestGetAllRestaurants
func TestGetAllRestaurants(t *testing.T) {
	ctx := context.Background()

	redisContainer, err := redisModule.RunContainer(ctx, testcontainers.WithImage("redis:7.2.4"))
	require.NoError(t, err)

	defer func() {
		if err := redisContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}()

	endpoint, err := redisContainer.Endpoint(ctx, "")
	require.NoError(t, err)

	rdb := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	defer rdb.Close()

	restaurants := []Restaurant{
		{ID: "1", Name: "Restaurant A", Address: "123 Main St", Website: "https://a.com", Phone: "123-456-7890", Tags: "Italian"},
		{ID: "2", Name: "Restaurant B", Address: "456 Elm St", Website: "https://b.com", Phone: "987-654-3210", Tags: "Mexican"},
	}

	for _, r := range restaurants {
		rdb.HSet(ctx, "restaurant:"+r.ID, map[string]interface{}{
			"name":    r.Name,
			"address": r.Address,
			"website": r.Website,
			"phone":   r.Phone,
			"tags":    r.Tags,
		})
		rdb.SAdd(ctx, "restaurants", r.ID)
	}

	restaurants, err = getAllRestaurants(ctx, rdb)

	fmt.Println("🥘", restaurants)

	assert.NoError(t, err)
	assert.Len(t, restaurants, 2)
}
