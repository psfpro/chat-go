package postgres

import (
	"chatgo/internal/chatgo/domain"
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTable(ctx context.Context) error {
	query := `
CREATE TABLE IF NOT EXISTS task (
    id UUID PRIMARY KEY,
    user_id UUID,
    title VARCHAR(255),
    description VARCHAR(255),
    state VARCHAR(255),
    technical_solution VARCHAR(255),
    comments VARCHAR(255)
);
`
	_, err := r.db.ExecContext(ctx, query)

	return err
}

func (r *TaskRepository) GetAllByUserId(ctx context.Context, userID uuid.UUID) ([]*domain.Task, error) {
	var res []*domain.Task
	rows, err := r.db.QueryContext(ctx, `
SELECT id, user_id, title, description, state, technical_solution, comments 
FROM task where user_id=$1
`, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		data := domain.Task{}
		err := rows.Scan(
			&data.ID,
			&data.UserID,
			&data.Title,
			&data.Description,
			&data.State,
			&data.TechnicalSolution,
			&data.Comments,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, &data)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *TaskRepository) Save(ctx context.Context, Task *domain.Task) error {
	stm, err := r.db.PrepareContext(ctx, `
INSERT INTO task (id, user_id, title, description, state, technical_solution, comments )
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (id)
DO UPDATE SET
    user_id = $2,
    title = $3,
	description = $4,
	state = $5,
	technical_solution = $6,
	comments = $7
`)
	if err != nil {
		return err
	}

	_, err = stm.Exec(
		Task.ID,
		Task.UserID,
		Task.Title,
		Task.Description,
		Task.State,
		Task.TechnicalSolution,
		Task.Comments,
	)
	if err != nil {
		return err
	}

	return nil
}
