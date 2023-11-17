package db

import (
	"context"
	"docker/warehouse"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisState struct {
	redis *redis.Client
}

var productsIdIncrement = 1

func NewRedisState() *RedisState {

	return &RedisState{
		redis.NewClient(&redis.Options{
			Addr: "redis:6379",
			DB:   0,
		})}
}

func (s *RedisState) SaveProduct(p warehouse.Product) warehouse.Product {

	p.ID = productsIdIncrement

	productString, err := json.Marshal(p)

	if err != nil {
		log.Println(err)
		return p
	}

	err = s.redis.Set(context.TODO(), fmt.Sprintf("products:%d", productsIdIncrement), string(productString), 0).Err()

	if err != nil {
		log.Println(err)
		return p
	}

	productsIdIncrement++

	return p
}

// TODO: redo implementation, use LIST Redis structure instead of STRING
func (s *RedisState) ListProducts() []warehouse.Product {
	keys := s.redis.Keys(context.TODO(), "products:*").Val()

	log.Println(keys)

	productStrings := s.redis.MGet(context.TODO(), keys...).Val()

	log.Println(productStrings)

	products := []warehouse.Product{}

	for _, p := range productStrings {
		product := warehouse.Product{}
		err := json.Unmarshal([]byte(p.(string)), &product)

		if err != nil {
			log.Println(err)
			return products
		}

		products = append(products, product)
	}

	return products
}
