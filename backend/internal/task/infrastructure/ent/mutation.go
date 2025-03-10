// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ToDo/internal/task/infrastructure/ent/predicate"
	"ToDo/internal/task/infrastructure/ent/task"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTask = "Task"
)

// TaskMutation represents an operation that mutates the Task nodes in the graph.
type TaskMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	title         *string
	description   *string
	priority      *task.Priority
	stage         *task.Stage
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Task, error)
	predicates    []predicate.Task
}

var _ ent.Mutation = (*TaskMutation)(nil)

// taskOption allows management of the mutation configuration using functional options.
type taskOption func(*TaskMutation)

// newTaskMutation creates new mutation for the Task entity.
func newTaskMutation(c config, op Op, opts ...taskOption) *TaskMutation {
	m := &TaskMutation{
		config:        c,
		op:            op,
		typ:           TypeTask,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTaskID sets the ID field of the mutation.
func withTaskID(id uuid.UUID) taskOption {
	return func(m *TaskMutation) {
		var (
			err   error
			once  sync.Once
			value *Task
		)
		m.oldValue = func(ctx context.Context) (*Task, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Task.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTask sets the old Task of the mutation.
func withTask(node *Task) taskOption {
	return func(m *TaskMutation) {
		m.oldValue = func(context.Context) (*Task, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TaskMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TaskMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Task entities.
func (m *TaskMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TaskMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TaskMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Task.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *TaskMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *TaskMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *TaskMutation) ResetTitle() {
	m.title = nil
}

// SetDescription sets the "description" field.
func (m *TaskMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *TaskMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *TaskMutation) ResetDescription() {
	m.description = nil
}

// SetPriority sets the "priority" field.
func (m *TaskMutation) SetPriority(t task.Priority) {
	m.priority = &t
}

// Priority returns the value of the "priority" field in the mutation.
func (m *TaskMutation) Priority() (r task.Priority, exists bool) {
	v := m.priority
	if v == nil {
		return
	}
	return *v, true
}

// OldPriority returns the old "priority" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldPriority(ctx context.Context) (v task.Priority, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPriority is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPriority requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPriority: %w", err)
	}
	return oldValue.Priority, nil
}

// ResetPriority resets all changes to the "priority" field.
func (m *TaskMutation) ResetPriority() {
	m.priority = nil
}

// SetStage sets the "stage" field.
func (m *TaskMutation) SetStage(t task.Stage) {
	m.stage = &t
}

// Stage returns the value of the "stage" field in the mutation.
func (m *TaskMutation) Stage() (r task.Stage, exists bool) {
	v := m.stage
	if v == nil {
		return
	}
	return *v, true
}

// OldStage returns the old "stage" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldStage(ctx context.Context) (v task.Stage, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStage: %w", err)
	}
	return oldValue.Stage, nil
}

// ResetStage resets all changes to the "stage" field.
func (m *TaskMutation) ResetStage() {
	m.stage = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TaskMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TaskMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TaskMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the TaskMutation builder.
func (m *TaskMutation) Where(ps ...predicate.Task) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TaskMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TaskMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Task, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TaskMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TaskMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Task).
func (m *TaskMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TaskMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.title != nil {
		fields = append(fields, task.FieldTitle)
	}
	if m.description != nil {
		fields = append(fields, task.FieldDescription)
	}
	if m.priority != nil {
		fields = append(fields, task.FieldPriority)
	}
	if m.stage != nil {
		fields = append(fields, task.FieldStage)
	}
	if m.updated_at != nil {
		fields = append(fields, task.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TaskMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case task.FieldTitle:
		return m.Title()
	case task.FieldDescription:
		return m.Description()
	case task.FieldPriority:
		return m.Priority()
	case task.FieldStage:
		return m.Stage()
	case task.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TaskMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case task.FieldTitle:
		return m.OldTitle(ctx)
	case task.FieldDescription:
		return m.OldDescription(ctx)
	case task.FieldPriority:
		return m.OldPriority(ctx)
	case task.FieldStage:
		return m.OldStage(ctx)
	case task.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Task field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TaskMutation) SetField(name string, value ent.Value) error {
	switch name {
	case task.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case task.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case task.FieldPriority:
		v, ok := value.(task.Priority)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPriority(v)
		return nil
	case task.FieldStage:
		v, ok := value.(task.Stage)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStage(v)
		return nil
	case task.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TaskMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TaskMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TaskMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Task numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TaskMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TaskMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TaskMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Task nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TaskMutation) ResetField(name string) error {
	switch name {
	case task.FieldTitle:
		m.ResetTitle()
		return nil
	case task.FieldDescription:
		m.ResetDescription()
		return nil
	case task.FieldPriority:
		m.ResetPriority()
		return nil
	case task.FieldStage:
		m.ResetStage()
		return nil
	case task.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TaskMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TaskMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TaskMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TaskMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TaskMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TaskMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TaskMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Task unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TaskMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Task edge %s", name)
}
