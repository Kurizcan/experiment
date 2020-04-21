package message

import (
	"github.com/Shopify/sarama"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"sync"
)

type KafkaClient struct {
	Cluster sarama.ClusterAdmin
	Config  *sarama.Config
}

var singleton *KafkaClient
var once sync.Once

func GetKafkaClient() *KafkaClient {
	once.Do(func() {
		config := sarama.NewConfig()
		config.Version = sarama.V0_10_1_0
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true
		config.Producer.Retry.Max = 5
		// 选择使用的压缩算法，这里没有使用
		//config.Producer.Compression = sarama.CompressionGZIP
		cluster, err := sarama.NewClusterAdmin(viper.GetStringSlice("mq.kafka.brokers"), config)
		if err != nil {
			log.Fatal("cluster create fail ", err)
		}
		singleton = &KafkaClient{
			Cluster: cluster,
			Config:  config,
		}
	})
	return singleton
}

func (client *KafkaClient) GetProducer() (sarama.AsyncProducer, error) {
	producer, err := sarama.NewAsyncProducer(viper.GetStringSlice("mq.kafka.brokers"), client.Config)
	if err != nil {
		log.Fatal("producer create fail ", err)
		return nil, err
	}
	return producer, nil
}

// 该方法不向外提供，并且只会执行一次
func (client *KafkaClient) createTopic() error {
	// create topic
	// bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 1 --topic problem
	cluster := client.Cluster
	defer cluster.Close()

	topics, _ := cluster.ListTopics()
	_, topicProblemExist := topics[TopicProblem]
	_, topicAnswerExist := topics[TopicAnswer]
	var err error
	if !topicProblemExist {
		err = createTopic(cluster, TopicProblem, PartitionsNum, ReplicationFactor)
		if err != nil {
			log.Error("create TOPIC_PROBLEM fail ", err)
		}
	}
	if !topicAnswerExist {
		err = createTopic(cluster, TopicAnswer, PartitionsNum, ReplicationFactor)
		if err != nil {
			log.Error("create TOPIC_ANSWER fail ", err)
		}
	}
	return err
}

func (client *KafkaClient) Produce(topic string, data []byte) error {
	producer, err := client.GetProducer()
	if err != nil {
		log.Error("get producer fail", err)
		return err
	}
	defer producer.AsyncClose()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}
	producer.Input() <- msg
	log.Infof("send message topic:%s, value:%v", topic, data)
	select {
	case err := <-producer.Errors():
		log.Errorf(err.Err, "send message fail topic:%s, value:%v, msg: %v", topic, data, err.Msg)
		return err
	case ok := <-producer.Successes():
		log.Infof("send message success topic:%s, value:%v", topic, ok.Value)
		return nil
	}
}

func createTopic(cluster sarama.ClusterAdmin, topic string, numPartitions int32, replicationFactor int16) error {
	err := cluster.CreateTopic(topic, &sarama.TopicDetail{
		NumPartitions:     numPartitions,
		ReplicationFactor: replicationFactor,
		ReplicaAssignment: nil,
		ConfigEntries:     nil,
	}, false)
	if err != nil {
		return err
	}
	return nil
}
