package redis

import (
	"github.com/aplescia-chwy/lets-go/util"
	"github.com/go-redis/redis"
	"os"
)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

//ClusterClient creates a pointer to a redis cluster client. By default looks for a redis cluster at
//localhost on port 7000. Address can be configured with the REDIS_HOST environment
//variable. A password can be configured with the REDIS_PASSWORD environment variable.
//An example for elasticache:
//	somecluster.amazonaws.com:6379
func ClusterClient() *redis.ClusterClient {

	clusterURL := util.GetEnvOrDefault("REDIS_HOST", "127.0.0.1:7000")
	password := os.Getenv("REDIS_PASSWORD")
	clientOptions := &redis.ClusterOptions{Addrs: []string{clusterURL}}
	if password != "" {
		clientOptions.Password = password
	}
	clusterClient := redis.NewClusterClient(clientOptions)

	_, err := clusterClient.Ping().Result()
	if err != nil {
		log.Error("Could not connect to redis! ", err)
		return nil
	}

	return clusterClient
}

//ClusterClientWithOpts creates a pointer to a redis cluster client. Supports passing a pointer to a redis.Options struct with
////configurable parameters. An example for elasticache:
//	somecluster.amazonaws.com:6379
func ClusterClientWithOpts(options *redis.ClusterOptions) *redis.ClusterClient {
	clusterClient := redis.NewClusterClient(options)
	_, err := clusterClient.Ping().Result()
	if err != nil {
		log.Error("Could not connect to redis!", err)
		return nil
	}

	return clusterClient
}

//StandardClient creates a pointer to a redis client. By default looks for a redis server at
//localhost on port 6379. Address can be configured with the REDIS_HOST environment
//variable. An example for elasticache:
//	someserver.amazonaws.com:6379
func StandardClient() *redis.Client {
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr: util.GetEnvOrDefault("REDIS_HOST", "127.0.0.1:6379"),
	})
	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}
	return client
}

//StandardClientWithOpts creates a pointer to a redis client. Supports passing a pointer to a redis.Options struct with
//configurable parameters. An example for elasticache:
//	someserver.amazonaws.com:6379
func StandardClientWithOpts(options *redis.Options) *redis.Client {
	var client = redis.NewClient(options)
	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}
	return client
}
