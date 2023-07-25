package tracker

type Tracker struct {
	Status map[int64]string
}

func New() *Tracker {
	status := make(map[int64]string)
	return &Tracker{status}
}

func (t *Tracker) IsSet(chatId int64, command string) bool {
	return t.Status[chatId] == command
}

func (t *Tracker) UnSet(chatId int64) {
	t.Status[chatId] = ""
}

func (t *Tracker) NotSet(chatId int64) bool {
	return t.Status[chatId] == ""
}

func (t *Tracker) Update(chatId int64, command string) {
	t.Status[chatId] = command
}
