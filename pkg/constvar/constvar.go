package constvar

const (
	Student   = 0
	Teacher   = 1
	Admin     = 2
	NEW       = "new"
	NOSUBMIT  = "no_submit"
	SUBMIT    = "submit"
	COMPLETED = "completed"
	RUNNING   = "running"
	ACCEPT    = "accept"
	WRONG     = "wrong"
	EMPTY     = -1
)

var ExperimentStudentStatus = map[string]int{
	NEW:       0,
	NOSUBMIT:  1,
	SUBMIT:    2,
	COMPLETED: 3,
}

var ProblemSubmitStatus = map[string]int{
	RUNNING: 1,
	ACCEPT:  2,
	WRONG:   3,
}
