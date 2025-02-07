package encapsulation

// 这种命名无法通过跨包去实例化
// 即使能实例化，也无法给score赋值
// 通常这种命名方式，是为了放一些私有属性，不给外部(跨包)直接调用
type student struct {
	Name  string
	score float32
}

// 根据我定义的规范进行实例化
func NewStudent(name string, score float32) *student {
	stu := &student{
		Name:  name,
		score: score,
	}
	return stu
}

func (s *student) GetScore() float32 {
	return s.score
}

func (s *student) SetScore(score float32) {
	s.score = score
}
