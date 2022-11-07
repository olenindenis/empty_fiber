package repositories

import (
	"envs/internal/core/domain"
	"envs/internal/core/ports"
	"envs/internal/dto"
	"envs/pkg/database"
	"github.com/Masterminds/squirrel"
	"time"
)

const (
	usersToRolesTableName = "user_role"
	usersTableName        = "users"
)

type UserRepository struct {
	database database.Connection
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(database database.Connection) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (sr *UserRepository) Find(id uint) (domain.User, error) {
	query, args, err := squirrel.Select(
		"id",
		"name",
		"email",
		"password",
		"created_at",
	).
		From(usersTableName).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	var item domain.User

	if err != nil {
		return item, err
	}

	conn, err := sr.database.Connection()
	if err != nil {
		return item, err
	}

	row := conn.QueryRow(query, args...)
	err = row.Scan(&item.ID, &item.Name, &item.Email, &item.Password, &item.CreatedAt)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (sr *UserRepository) FindByEmail(email string) (domain.User, error) {
	query, args, err := squirrel.Select(
		"id",
		"name",
		"email",
		"password",
		"created_at",
	).
		From(usersTableName).
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	var item domain.User

	if err != nil {
		return item, err
	}

	conn, err := sr.database.Connection()
	if err != nil {
		return item, err
	}

	row := conn.QueryRow(query, args...)
	err = row.Scan(&item.ID, &item.Name, &item.Email, &item.Password, &item.CreatedAt)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (sr *UserRepository) Store(name, email, password string) (domain.User, error) {
	query, args, err := squirrel.Insert(usersTableName).
		Columns(
			"name",
			"email",
			"password",
			"created_at",
		).
		Values(
			name,
			email,
			password,
			time.Now().UTC(),
		).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return domain.User{}, err
	}

	conn, err := sr.database.Connection()
	if err != nil {
		return domain.User{}, err
	}

	tx, err := conn.Begin()
	if err != nil {
		return domain.User{}, err
	}

	var userID int
	err = tx.QueryRow(query, args...).Scan(&userID)
	if err != nil {
		return domain.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return domain.User{}, err
	}

	//userID, err := result.LastInsertId()
	//if err != nil {
	//	return domain.User{}, err
	//}

	return domain.User{
		ID:        uint(userID),
		Name:      name,
		Email:     email,
		CreatedAt: domain.TimeStampUnix(time.Now().UTC().Second()),
	}, nil
}

func (sr *UserRepository) List(filter dto.ListFilter) ([]domain.User, error) {
	var items []domain.User

	query := squirrel.Select("id, name, email, created_at").
		From(usersTableName).
		Limit(uint64(filter.Limit)).
		Offset(uint64(filter.Offset)).
		OrderBy(filter.Order).
		PlaceholderFormat(squirrel.Dollar)

	queryString, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	conn, err := sr.database.Connection()
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item domain.User

		err = rows.Scan(&item.ID, &item.Name, &item.Email, &item.CreatedAt)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (sr *UserRepository) Update(user domain.User) error {
	query := squirrel.Update(usersTableName).
		Set("name", user.Name).
		PlaceholderFormat(squirrel.Dollar)

	queryString, args, err := query.ToSql()
	if err != nil {
		return err
	}

	conn, err := sr.database.Connection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}

func (sr *UserRepository) Delete(id uint) error {
	query := squirrel.Delete(usersTableName).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	queryString, args, err := query.ToSql()
	if err != nil {
		return err
	}

	conn, err := sr.database.Connection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
