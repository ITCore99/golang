package model

type worker struct {
	Name string
}

func (w worker) GetName() string { // 这里的首字母也要大写不让无法挎包
	return w.Name
}

func CreateWroker(name string) *worker {
	return &worker{Name: name}
}