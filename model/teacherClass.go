package model

type TeacherClassModel struct {
	Id        int    `json:"id" gorm:"column:id;primary_key;"`
	TeacherId int    `json:"teacher_id" gorm:"column:teacherId;"`
	ClassId   string `json:"class_id" gorm:"column:classId;"`
}

type Classes struct {
	ClassId string `json:"class_id" gorm:"column:classId;"`
}

func (t *TeacherClassModel) TableName() string {
	return "teacher_class"
}

func (t *TeacherClassModel) GetClassByTeacher(teacherId int) []string {
	c := make([]Classes, 0)
	DB.Self.Table(t.TableName()).Select("classId").Where("teacherId = ?", teacherId).Scan(&c)
	res := make([]string, len(c))
	for index, cla := range c {
		res[index] = cla.ClassId
	}
	return res
}
