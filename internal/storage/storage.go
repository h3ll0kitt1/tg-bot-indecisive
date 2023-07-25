package storage

type Storage interface {
	Save(chatId int64, in string) (bool, error)
	Delete(chatId int64, in string) (bool, error)
	Exists(chatId int64, in string) (bool, error)
	Rand(chatId int64) (string, error)
	List(chatId int64) ([]string, error)
}
