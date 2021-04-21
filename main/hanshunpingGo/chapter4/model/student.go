package model

// 如果首字母大写，外包可以正常访问，但是如果首字母小写，外包就不能直接访问
type student struct {
	Name  string
	score float64
}

// 通过工厂模式来获取student的实例
func NewStudent(n string, s float64) *student {
	return &student{Name: n, score: s}
}

func NewStudent2(n string, s float64) student {
	return student{Name: n, score: s}
}

func (s *student) GetScore() float64 {
	return s.score
}
