package student

import (
	. "experiment/handler"
	"experiment/model"
	"experiment/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

func MyExperiments(c *gin.Context) {
	studentId, exists := c.Get("userId")
	if !exists {
		SendResponse(c, errno.ErrParam, nil)
		return
	}
	fmt.Println(studentId)

	se := model.StudentExperimentModel{}
	groups, err := se.GetGroups(int(studentId.(float64)))
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	res := make([]model.StudentExperiment, len(groups))

	wg := sync.WaitGroup{}
	for i, group := range groups {
		wg.Add(1)
		go func(index int, se model.StudentExperimentModel) {
			defer wg.Done()
			experiment := model.ExperimentModel{}
			if err := experiment.GetDetail(se.GroupId); err != nil {
				SendResponse(c, errno.ErrDatabase, nil)
				return
			}
			res[i] = model.StudentExperiment{
				GroupId:   se.GroupId,
				GroupName: experiment.GroupName,
				Status:    se.Status,
				Score:     se.Score,
			}
		}(i, group)
	}
	wg.Wait()

	SendResponse(c, errno.OK, res)
}
