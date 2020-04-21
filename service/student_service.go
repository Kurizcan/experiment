package service

import (
	"experiment/model"
	"experiment/pkg/constvar"
	"experiment/pkg/errno"
	"experiment/pkg/redis"
	"fmt"
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
	//如果用户未逐题提交，则在这里是默认没有提交的
	// TODO 需要判断每道题目是否允许完成，如果允许完成，返回结果，需要进行一个判断，使用缓存实现判断
	// gid_sid: {pid, status} hash
	res, err := getExperimentStatus(groupId, userId)
	return res, err
}

func getExperimentStatus(groupId, studentId int) (model.ExperimentResult, error) {
	// 获取该学生的所有 answer 情况
	experimentKey := redis.GetGroupSidKey(groupId, studentId)
	detail, err := redis.Client.HGetAll(experimentKey)
	if err == nil && isRunning(detail) {
		return model.ExperimentResult{}, errno.ErrSubmitRunning
	} else {
		// 缓存失效
		dataForRedis := make(map[string]interface{})
		answer := model.AnswerModel{}
		problemList, err := answer.GetProblemScore(groupId, studentId)
		if err != nil || len(problemList) == 0 {
			return model.ExperimentResult{}, errno.ErrDatabase
		}

		list := make([]model.ProblemResultList, len(problemList))
		score, isRunning := 0, false
		wg := sync.WaitGroup{}
		for i, problem := range problemList {
			score += problem.Score
			dataForRedis[fmt.Sprintf("%d", problem.ProblemId)] = problem.Status
			if isRunning || problem.Status == constvar.ProblemSubmitStatus[constvar.RUNNING] {
				isRunning = true
				continue
			}
			wg.Add(1)
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

		// 更新 redis
		err = redis.Client.HSetAll(experimentKey, dataForRedis)

		return res, errno.OK
	}
}

func isRunning(data map[string]string) bool {
	for _, status := range data {
		statusNum, _ := strconv.Atoi(status)
		if statusNum == constvar.ProblemSubmitStatus[constvar.RUNNING] {
			return true
		}
	}
	return false
}
