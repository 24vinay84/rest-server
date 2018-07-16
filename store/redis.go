// Remote Redis Server for Caching.

package store

import (
    "log"
	"strconv"
	"time"
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

func getNewClient() {
	
	log.Println("Redis server is connecting")
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("Redis Server is not connecting ******")
		panic(err)
	}
	log.Println(pong, err)
}

func incementCount(ipadd string) {

	val, err := client.Get(ipadd).Result()
	if err == redis.Nil {
		log.Println(ipadd + " key does not exists")
		err := client.Set(ipadd, "1", 5*time.Minute)   // Expire in 5 minutes
		if err != nil {
			panic(err)
		}	
	} else if err != nil {
		panic(err)
	} else {
		incr := client.Incr(ipadd)
		log.Println(ipadd, val,incr)
	}
}

func getCount(ipadd string) int64 {

	get := client.Get(ipadd)
	if get.Err() != nil {
		panic(get.Err())
	}
	i64, err := strconv.ParseInt(get.Val(), 10, 0)
	if err != nil {
		panic(get.Err())
	}
	return i64;
}

func blockIP(ipadd string) {

        blockIP :=  "block-" + ipadd  
		err := client.Set(blockIP, "1", time.Hour)   // Block for one hour
		if err != nil {
			panic(err)
		}
}

func isBlockIP(ipadd string) bool{

	blockIP :=  "block-" + ipadd  
	
	val, err := client.Get(blockIP).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		panic(err)
	} else {
		log.Println(ipadd, val)
		return true
	}
}

