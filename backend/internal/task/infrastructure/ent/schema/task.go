package schema

import (
	"ToDo/internal/task/application/domain"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				id, err := uuid.NewV7()
				if err != nil {
					panic("failed to generate UUIDv7")
				}
				return id
			}),
		field.String("title").MaxLen(domain.MaxTitleLength).NotEmpty(),
		field.String("description").MaxLen(domain.MaxDescriptionLength).NotEmpty(),
		field.Enum("priority").Values("LOW", "MIDDLE", "HIGH").Default("MIDDLE"),
		field.Enum("stage").Values("TODO", "IN_PROGRESS", "DONE"),
		field.Time("updated_at").Default(time.Now),
	}
}