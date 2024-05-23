package http_server

import (
	"gitlab.com/bat.ggl/bgDigital/internal/core/entity"
	"time"
)

type LoginUserResponse struct {
	ID                  string // uuid
	Login               string
	EmailConfirmed      bool              // подверждени ли email
	ParentOrganithation string            // родительская организация
	Role                entity.RoleEnum   // роль пользователя
	Status              entity.StatusEnum // состание пользователя
	LicenceExpiredAt    time.Time         // срок окончания лизензии
}

type LogoutResponse struct {
	Status bool
}
