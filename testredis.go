package main

import (

	"github.com/go-redis/redis"
	"fmt"
	"bufio"
	"os"
)

var RedisClient *redis.Client
var RedisAddress string = "localhost:6379"
var RedisPassword string = ""
var PubSubChannel *redis.PubSub = nil

func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisAddress,
		Password: RedisPassword, // no password set
		DB:       0,             // use default DB
	})

	fmt.Println( RedisClient.Ping().Result() )

}

func ServerRedisPublish(redisChannel string, message string, channelId int32) {
	/*_, ok := PubSubChannel[redisChannel]
	if !ok {
		ServerRedisSubscribe(redisChannel, channelId)
	}*/

	fmt.Println("To publish " + message + " to: " + redisChannel)

	err := RedisClient.Publish(redisChannel, message).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func ServerRedisReceive() {
	defer PubSubChannel.Close()

	for {
		//for _, ps := range PubSubChannel {
		msg, err := PubSubChannel.ReceiveMessage()
		if err != nil {

			continue
		}

		fmt.Println("received", msg.Payload, "from", msg.Channel)


		//}
	}
}

func main(){
	InitRedisClient()
	PubSubChannel = RedisClient.Subscribe("test")
	go ServerRedisReceive()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		ServerRedisPublish("test" ,text , 0 )
	}
}