package student

import "github.com/pkg/errors"

type (
	Student struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Age   int64  `json:"age"`
		Score int64  `json:"score"`
	}

	studentManager struct {
		info map[string]*Student
	}
)

var Manager *studentManager

func init() {
	Manager = new(studentManager)
	Manager.info = make(map[string]*Student)
	Manager.info["1"] = &Student{
		Id:    "1",
		Name:  "A",
		Age:   20,
		Score: 84,
	}

	Manager.info["2"] = &Student{
		Id:    "2",
		Name:  "B",
		Age:   20,
		Score: 70,
	}
}

func (m *studentManager) FindAll() []Student {
	list := make([]Student, 0, len(m.info))
	for _, student := range m.info {
		list = append(list, *student)
	}
	return list
}

func (m *studentManager) Find(id string) (Student, error) {
	student, exists := m.info[id]
	if exists {
		return *student, nil
	}

	return Student{}, errors.Errorf("student id: %s not found ", id)
}

func (m *studentManager) Insert(student *Student) error {
	_, err := m.Find(student.Id)
	if err == nil {
		return errors.Errorf("duplicate student id: %s", student.Id)
	}
	m.info[student.Id] = student
	return nil
}

func (m *studentManager) Update(student *Student) error {
	oldStudent, err := m.Find(student.Id)
	if err != nil {
		return err
	}

	if len(student.Name) > 0 {
		oldStudent.Name = student.Name
	}

	if student.Age > 0 && oldStudent.Age != student.Age {
		oldStudent.Age = student.Age
	}

	if student.Score > 0 && oldStudent.Score != student.Score {
		oldStudent.Score = student.Score
	}
	m.info[student.Id] = &oldStudent
	return nil
}

func (m *studentManager) Delete(id string) error {
	_, err := m.Find(id)
	if err != nil {
		return err
	}
	delete(m.info, id)
	return nil
}
