package service

import (
	"experiment/model"
	"experiment/pkg/constvar"
	"experiment/pkg/errno"
	"github.com/lexkong/log"
	"strconv"
	"sync"
)

type Student struct{}

func NewStudent() *Student {
	return &Student{}
}

func (s *Student) Submit(groupId, userId int) (model.ExperimentResult, error) {
	/*
			1. 判断 student_experiment 对应 status 是否为已完成，当所有完成后才可以返回结果
		  2. 搜索 answer 获取所有答题详情
		  3. 更新 student_experiment 总分
	*/

	se := model.StudentExperimentModel{}
	if err := se.Detail(groupId, userId); err != nil {
		return model.ExperimentResult{}, errno.ErrDatabase
	}

	//TODO 如果用户未逐题提交，则在这里将所有涉及这个同学的题目提交到判题队列中去

	if se.Status != constvar.EXPERIMENT_STUDENT_STATUS[constvar.COMPLETED] {
		return model.ExperimentResult{}, errno.ErrSubmitRunning
	}

	answer := model.AnswerModel{}
	problemList, err := answer.GetProblemScore(groupId, userId)
	if err != nil || len(problemList) == 0 {
		return model.ExperimentResult{}, errno.ErrDatabase
	}

	list := make([]model.ProblemResultList, len(problemList))
	score := 0
	wg := sync.WaitGroup{}
	for i, problem := range problemList {
		wg.Add(1)
		score += problem.Score
		// get title
		go func(index int, pro model.ProblemScore) {
			defer wg.Done()
			p := model.ProblemModel{}
			if err := p.Detail(strconv.Itoa(pro.ProblemId)); err != nil {
				log.Error("get title fail", err)
				return
			}
			list[index].Title = p.Title
			list[index].Score = pro.Score
			list[index].ProblemId = pro.ProblemId
		}(i, problem)
	}
	wg.Wait()

	res := model.ExperimentResult{
		Score: score,
		List:  list,
	}

	//fmt.Println(res)

	return res, errno.OK
}
