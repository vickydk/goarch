package user

type Repository interface {
	Save(entity *Entity) (err error)
	FindByEmail(email string) (entity Entity, err error)
	FindById(id int64) (entity Entity, err error)
}
