package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	wmkafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/cikupin/learn-firehose/config"
	"github.com/cikupin/learn-firehose/payload"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var grpcProducerCmd = &cobra.Command{
	Use:   "grpc-producer",
	Short: "produce GRPC message to Kafka in every 5 seconds",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv()
		doGRPC()
	},
}

func init() {
	rootCmd.AddCommand(grpcProducerCmd)
}

func doGRPC() {
	kafka0Host := fmt.Sprintf("%s:%d", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)
	pubConfig := wmkafka.PublisherConfig{
		Brokers:               []string{kafka0Host},
		Marshaler:             wmkafka.DefaultMarshaler{},
		OverwriteSaramaConfig: wmkafka.DefaultSaramaSyncPublisherConfig(),
	}

	publisher, err := wmkafka.NewPublisher(pubConfig, watermill.NewStdLogger(true, false))
	if err != nil {
		log.Fatalln(err.Error())
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					publishGRPC(t, publisher)
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

func publishGRPC(t time.Time, publisher *wmkafka.Publisher) {
	newID, _ := uuid.NewRandom()
	newPayload := &payload.Flag{
		Id: newID.String(),
	}

	bytePayload, err := proto.Marshal(newPayload)
	if err != nil {
		log.Printf("[ERROR - marshal proto] %s\n", err.Error())
		return
	}

	newMessage := message.NewMessage(newID.String(), bytePayload)
	err = publisher.Publish(config.Cfg.Kafka.TopicGRPC, newMessage)
	if err != nil {
		log.Printf("[ERROR - publish message] %s\n", err.Error())
		return
	}

	fmt.Printf("---> [%v] Publishing :%s\n", t.Format(time.Kitchen), newID.String())
}
