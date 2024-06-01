package domain

type Storage interface {
	SaveFiles(files map[string]*File)
}
