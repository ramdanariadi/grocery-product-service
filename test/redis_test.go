package test

import (
	"encoding/json"
	"fmt"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/proto/product"
	"github.com/ramdanariadi/grocery-be-golang/main/utils"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func Test_write_redis(t *testing.T) {
	redisClient := utils.NewRedisClient()
	product := product.Product{
		Id:          "s8s978",
		Name:        "Brocolly",
		CategoryId:  "hgu8023",
		ImageUrl:    "",
		Description: "Good for health",
		PerUnit:     100,
		Weight:      10,
		Price:       1000,
		Category:    "Vegetables",
	}
	bytes, err := json.Marshal(product)
	helpers.LogIfError(err)
	ctx := context.Background()
	err = redisClient.Set(ctx, "s8s978", bytes, 1*time.Hour).Err()
	helpers.LogIfError(err)
}

func Test_read_redis(t *testing.T) {
	redistClient := utils.NewRedisClient()
	ctx := context.Background()
	since := time.Now()
	result, err := redistClient.Get(ctx, "s8s978").Result()
	helpers.LogIfError(err)
	after := time.Since(since)
	fmt.Printf("done on %d miliseconds \n", after.Milliseconds())
	fmt.Println("productId", result)
}
