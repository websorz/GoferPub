package consumer

type Consumer struct {
	BootstrapServer string
	ConsumerGroup   string
	Offset          string
	Topics          []string
}
