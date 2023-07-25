package inmemory

type MemStorage struct {
	list map[int64]map[string]bool
}

func NewStorage() *MemStorage {
	var ms MemStorage
	ms.list = make(map[int64]map[string]bool)
	return &ms
}

func (ms *MemStorage) Save(chatId int, book string) bool {
	if _, ok := ms[chatId]; !ok {
		ms[chatId] = make(map[string]bool)
	}

	if _, ok := ms[chatId][book]; !ok {
		return false
	}

	ms[chatId][book] = true
	return true
}

func (ms *MemStorage) Delete() {
	if _, ok := ms[chatId]; !ok {
		ms[chatId] = make(map[string]bool)
	}

	if _, ok := ms[chatId][book]; !ok {
		return false
	}

	delete(ms[chatId], ms[chatId][book])
	return true
}

func (ms *MemStorage) Exists() {
	if _, ok := ms[chatId]; !ok {
		ms[chatId] = make(map[string]bool)
	}

	if _, ok := ms[chatId][book]; !ok {
		return false
	}
	return true
}

func (ms *MemStorage) Rand() {

}

func (ms *MemStorage) List() {

}
