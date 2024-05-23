package entity

import (
	"time"
)

/**
пользователи системы
*/

type StatusEnum int
type RoleEnum int

const (
	ActiveStatusEnum  StatusEnum = 1
	BlockedStatusEnum StatusEnum = 2
	StoppedStatusEnum StatusEnum = 2

	AdminRoleEnum    RoleEnum = 1
	SalesRoleEnum    RoleEnum = 2
	HRRoleEnum       RoleEnum = 3
	DirectorRoleEnum RoleEnum = 4
	AccountRoleEnum  RoleEnum = 5
)

type User struct {
	ID                  string // uuid
	Login               string
	Password            string
	Email               string
	EmailConfirmed      bool   // подверждени ли email
	ParentOrganithation string // родительская организация
	Phone               string
	Role                RoleEnum   // роль пользователя
	Status              StatusEnum // состание пользователя
	LicenceExpiredAt    time.Time  // срок окончания лизензии
	CreatedAt           time.Time
	DeletedAt           time.Time
}
