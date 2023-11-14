package broker

type Broker interface {
	CreateTopic(topicName string, partCount int, repCount int) error
}
