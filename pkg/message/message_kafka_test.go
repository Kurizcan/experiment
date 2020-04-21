package message

import (
	"encoding/json"
	"experiment/config"
	"experiment/model"
	"fmt"
	"testing"
)

func TestKafkaClient_CreateTopic(t *testing.T) {
	// init config
	if err := config.Init("G:\\experiment\\conf\\config.yaml"); err != nil {
		panic(err)
	}
	GetKafkaClient()
	fmt.Print("first")
	GetKafkaClient()
	fmt.Println("second")
}

func TestKafkaClient_Produce(t *testing.T) {
	// init config
	if err := config.Init("G:\\experiment\\conf\\config.yaml"); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	/*problem := model.ProblemModel{}
	err := problem.Detail("102")
	if err != nil {
		t.Error("read problem fail")
	}
	file, err := os.Open("G:\\experiment\\data\\9ru4VtCWg.sql")
	if err != nil {
		t.Errorf("file open tail file :%s", problem.Data)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		 t.Errorf("read data fail, err: %v", err)
		return
	}

	msg := TopicProblemMessage{
		ProblemId:  problem.ProblemId,
		DataSource: data,
		Solution:   problem.Solution,
		OutPut:     problem.Output,
	}

	realMsg, err := json.Marshal(msg)
	if err != nil {
		t.Error("json marshal fail")
	}
	*/

	answer := model.AnswerModel{}
	err := answer.DetailById(109)
	if err != nil {
		t.Error("get answer fail")
		return
	}
	msg := TopicAnswerMessage{
		AnswerId:   answer.Id,
		ProblemId:  answer.ProblemId,
		Submit:     answer.Submit,
		UpdateTime: answer.UpdateTime,
	}
	realMsg, err := json.Marshal(msg)
	if err != nil {
		t.Error("json marshal fail")
	}
	client := GetKafkaClient()
	client.Produce(TopicAnswer, realMsg)
}
