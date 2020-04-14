package constvar

const (
	Student   = 0
	Teacher   = 1
	Admin     = 2
	NEW       = "new"
	NOSUBMIT  = "no_submit"
	SUBMIT    = "submit"
	COMPLETED = "completed"
	RUNING    = "running"
	ACCEPT    = "accept"
	WRONG     = "wrong"
	EMPTY     = -1
)

var EXPERIMENT_STUDENT_STATUS = map[string]int{
	NEW:       0,
	NOSUBMIT:  1,
	SUBMIT:    2,
	COMPLETED: 3,
}

var PROBLEM_SUBMIT_STATUS = map[string]int{
	RUNING: 1,
	ACCEPT: 2,
	WRONG:  3,
}
