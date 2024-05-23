package datasource

/*
репозиторий для сбора всех постов с телеграмм каналов/чатов/ботов
- посты о новых вакансиях
- посты о кандидатах
*/

const (
	collectionname = "Posts"
)

type PostRepository struct {
}

/*
**
Функция создания постов
*/
func CreatePostReposiroty() *PostRepository {
	return &PostRepository{}
}

func (p *PostRepository) CreatePost() {

}
