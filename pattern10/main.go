package constructor

// Экспортируемый пакетом тип данных.
type Type struct {
	state map[int]string // поле требует инициализации
	jobs  chan int       // поле требует инициализации
}

// Конструктор.
// Упрощает использование пакета.
func New() *Type {
	t := Type{}
	t.state = make(map[int]string)
	t.jobs = make(chan int)
	return &t
}
