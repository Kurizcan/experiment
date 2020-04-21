package service

import (
	"experiment/model"
	"experiment/pkg/constvar"
	"experiment/pkg/errno"
	"github.com/lexkong/log"
	"sync"
)

type Teacher struct{}

func NewTeacher() *Teacher {
	return &Teacher{}
}

// 分发实验
func (t *Teacher) Distributed(groupId int, classList []string) error {
	ch := make(chan string, len(classList))
	res := make(chan error)
	go t.updateExperimentClass(ch, res, groupId, classList)
	go t.updateExperimentStudent(ch, groupId)
	// 还有一种处理方式，专门用一个 channel 记录失败的 classId，另起一个协程不断的重试处理
	// 这里为了简单处理，就暂时直接返回错误信息，不做重试处理
	return <-res
}

func (t *Teacher) updateExperimentClass(class chan string, res chan error, groupId int, classList []string) {
	// fail case record
	cnt := 0
	// get group name
	em := model.ExperimentModel{}
	if err := em.GetDetail(groupId); err != nil {
		log.Errorf(err, "get group name fail, id: %d", groupId)
		res <- errno.ErrDatabase
	}
	wg := sync.WaitGroup{}
	for _, classId := range classList {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			ec := model.ClassExperimentModel{
				GroupId:   groupId,
				ClassId:   id,
				GroupName: em.GroupName,
			}
			if err := ec.Create(); err != nil {
				log.Errorf(err, "Distributed experiment to class:%s fail", id)
				return
			}
			class <- id
			cnt++
		}(classId)
	}
	wg.Wait()
	// 确保像 channel 发送完数据后才关闭 channel
	if cnt != len(classList) {
		// 上述操作有出现错误
		res <- errno.InternalServerError
	}
	close(class)
	res <- nil
	close(res)
}

func (t *Teacher) updateExperimentStudent(class chan string, groupId int) {
	for id := range class {
		log.Infof("get classId from chan class, id: %s", id)
		go func(classId string) {
			// 获取该班级所有学生列表
			sc := model.StudentClassModel{}
			studentList, err := sc.GetStudent(classId)
			if err != nil || len(studentList) == 0 {
				log.Errorf(err, "Due to class number is zero Distributed experiment to class:%s fail", id)
				return
			}
			// 更新目标表
			for _, sid := range studentList {
				se := model.StudentExperimentModel{
					GroupId:   groupId,
					StudentId: sid,
					Status:    constvar.ExperimentStudentStatus[constvar.NEW],
				}
				if err := se.Create(); err != nil {
					log.Errorf(err, "Distributed experiment to student:%s fail", sid)
					return
				}
			}
		}(id)
	}
}
