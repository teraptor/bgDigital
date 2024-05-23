package entity

/*
*
сущность компании
*/
type Company struct {
	ID               string
	OrganizationName string // название организации
	INN              string // ИНН
	KPP              string // КПП
	Email            string // при создании, автоматически создается пользователь для компании.
	Phone            string
	Password         string
}
