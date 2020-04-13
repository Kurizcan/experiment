package constvar

const (
	Student  = 0
	Teacher  = 1
	Admin    = 2
	NEW      = "new"
	NOSUBMIT = "no_submit"
	SUBMITED = "submit"
)

var EXPERIMENT_STUDENT_STATUS = map[string]int{
	NEW:      0,
	NOSUBMIT: 1,
	SUBMITED: 2,
}
