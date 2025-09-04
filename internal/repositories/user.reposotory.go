package repositories

import (
	"context"
	"log"

	"github.com/cristian-yw/Minitask_W10D2/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, body models.User) (models.User, error) {
	sql := "INSERT INTO users (email, role, password_hash) VALUES ($1, $2, $3) returning id, email, role"
	values := []any{body.Email, body.Role, body.Password_hash}
	var newuser models.User
	if err := r.db.QueryRow(ctx, sql, values...).Scan(&newuser.Id, &newuser.Email, &newuser.Role); err != nil {
		log.Println("Internal Server Error", err.Error())
		return models.User{}, err
	}
	return newuser, nil
}

func (r *UserRepository) InsertNewUser(ctx context.Context, body models.User) (pgconn.CommandTag, error) {
	sql := "INSERT INTO users (email, role, password_hash) VALUES ($1, $2, $3)"
	values := []any{body.Email, body.Role, body.Password_hash}
	return r.db.Exec(ctx, sql, values...)
}
func (r *UserRepository) GetAllUsers(ctx context.Context, offset, limit int) ([]models.User, error) {
	sql := "SELECT id, email, role, created_at from users LIMIT $1 OFFSET $2"
	values := []any{limit, offset}
	rows, err := r.db.Query(ctx, sql, values...)
	if err != nil {
		log.Println("Error executing query: ", err.Error())
		return []models.User{}, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Role); err != nil {
			log.Println("Error scanning row: ", err.Error())
			return []models.User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, body models.User) (models.User, error) {
	sql := "UPDATE users SET email =coalesce(nullif($1,''), email), role = coalesce(nullif($2, '')::type_role, role), password_hash = coalesce(nullif($3, ''), password_hash), updated_at = now() WHERE id = $4 returning id, email, role"
	values := []any{body.Email, body.Role, body.Password_hash, id}
	var UpdatedUser models.User
	if err := r.db.QueryRow(ctx, sql, values...).Scan(&UpdatedUser.Id, &UpdatedUser.Email, &UpdatedUser.Role); err != nil {
		log.Printf("Error executing query: %v", err)
		return models.User{}, err
	}
	return UpdatedUser, nil
}
