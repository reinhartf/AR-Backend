package service

import (
	"database/sql"
	"github.com/reinhartf/AR-Backend/model"
	"github.com/jackc/pgx"
	"github.com/op/go-logging"
)

type RoleService struct {
	db  *pgx.Conn
	log *logging.Logger
}

func NewRoleService(db *pgx.Conn, log *logging.Logger) *RoleService {
	return &RoleService{db: db, log: log}
}

func (r *RoleService) FindByUserId(userId *string) ([]*model.Role, error) {
	roles := make([]*model.Role, 0)

	roleSQL := `SELECT role.*
	FROM roles role
	INNER JOIN rel_users_roles ur ON role.id = ur.role_id
	WHERE ur.user_id = $1 `
	rows, err := r.db.Query(roleSQL, userId)
	if err == sql.ErrNoRows {
		return roles, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var role model.Role
		err = rows.Scan(&role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}
