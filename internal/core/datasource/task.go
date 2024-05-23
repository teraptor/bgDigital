package datasource

/*
таски с напоминанем о каком-либо действии
*/
type TaskRepository struct {
}

func (t *TaskRepository) CreateTaskRepository() *TaskRepository {
	return &TaskRepository{}
}
