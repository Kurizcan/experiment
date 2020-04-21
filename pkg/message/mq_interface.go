package message

type MQ interface {
	Produce(topic string, data []byte) error
	createTopic() error
}

const (
	TopicProblem      = "problem_experiment" // 2 个副本， 3 个分区
	TopicAnswer       = "answer_experiment"  // 2 个副本， 3 个分区
	PartitionsNum     = 3
	ReplicationFactor = 2
)

type TopicProblemMessage struct {
	ProblemId  int
	DataSource []byte
	Solution   string
	OutPut     []byte
}

type TopicAnswerMessage struct {
	AnswerId   int
	ProblemId  int
	Submit     string
	UpdateTime int64
}
