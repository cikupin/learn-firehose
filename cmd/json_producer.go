package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	wmkafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/cikupin/learn-firehose/config"
	"github.com/cikupin/learn-firehose/factory"
	"github.com/cikupin/learn-firehose/payload"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var jsonProducerCmd = &cobra.Command{
	Use:   "json-producer",
	Short: "produce JSON message to Kafka in every 5 seconds",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv()
		doJSON()
	},
}

func init() {
	rootCmd.AddCommand(jsonProducerCmd)
}

func doJSON() {
	publisher := factory.NewKafkaPublisher()

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					publishJSON(t, publisher)
				}
			}
		}()
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	publisher.Close()
	fmt.Println("application stopped...")
}

func publishJSON(t time.Time, publisher *wmkafka.Publisher) {
	newID, _ := uuid.NewRandom()
	createdAt := time.Now()

	newPayload := &payload.Identity{
		ID:        newID.String(),
		CreatedAt: createdAt,
	}

	err := faker.FakeData(&newPayload)
	if err != nil {
		log.Printf("[ERROR - faking data] %s\n", err.Error())
		return
	}

	bytePayload, err := json.Marshal(newPayload)
	if err != nil {
		log.Printf("[ERROR - marshal payload] %s\n", err.Error())
		return
	}

	newMessage := message.NewMessage(newID.String(), bytePayload)
	err = publisher.Publish(config.Cfg.Kafka.TopicJSON, newMessage)
	if err != nil {
		log.Printf("[ERROR - publish message] %s\n", err.Error())
		return
	}

	fmt.Printf("---> [%v] Publishing :%s\n", t.Format(time.Kitchen), newID.String())
}
