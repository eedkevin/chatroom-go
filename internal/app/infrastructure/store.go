package infrastructure

type IStorage interface {
	Save(id string, obj interface{}) error
	Get(id string) (interface{}, error)
	List() (map[string]interface{}, error)
	Delete(id string) error
}
