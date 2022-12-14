// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gmail-clone-backend/ent/email"
	"gmail-clone-backend/ent/predicate"
	"gmail-clone-backend/ent/user"
	"sync"
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEmail = "Email"
	TypeUser  = "User"
)

// EmailMutation represents an operation that mutates the Email nodes in the graph.
type EmailMutation struct {
	config
	op            Op
	typ           string
	id            *int
	from          *string
	subject       *string
	body          *string
	sentAt        *time.Time
	archived      *bool
	read          *bool
	clearedFields map[string]struct{}
	user          *uuid.UUID
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Email, error)
	predicates    []predicate.Email
}

var _ ent.Mutation = (*EmailMutation)(nil)

// emailOption allows management of the mutation configuration using functional options.
type emailOption func(*EmailMutation)

// newEmailMutation creates new mutation for the Email entity.
func newEmailMutation(c config, op Op, opts ...emailOption) *EmailMutation {
	m := &EmailMutation{
		config:        c,
		op:            op,
		typ:           TypeEmail,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEmailID sets the ID field of the mutation.
func withEmailID(id int) emailOption {
	return func(m *EmailMutation) {
		var (
			err   error
			once  sync.Once
			value *Email
		)
		m.oldValue = func(ctx context.Context) (*Email, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Email.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEmail sets the old Email of the mutation.
func withEmail(node *Email) emailOption {
	return func(m *EmailMutation) {
		m.oldValue = func(context.Context) (*Email, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EmailMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EmailMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EmailMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EmailMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Email.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetFrom sets the "from" field.
func (m *EmailMutation) SetFrom(s string) {
	m.from = &s
}

// From returns the value of the "from" field in the mutation.
func (m *EmailMutation) From() (r string, exists bool) {
	v := m.from
	if v == nil {
		return
	}
	return *v, true
}

// OldFrom returns the old "from" field's value of the Email entity.
// If the Email object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailMutation) OldFrom(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFrom is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFrom requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFrom: %w", err)
	}
	return oldValue.From, nil
}

// ResetFrom resets all changes to the "from" field.
func (m *EmailMutation) ResetFrom() {
	m.from = nil
}

// SetSubject sets the "subject" field.
func (m *EmailMutation) SetSubject(s string) {
	m.subject = &s
}

// Subject returns the value of the "subject" field in the mutation.
func (m *EmailMutation) Subject() (r string, exists bool) {
	v := m.subject
	if v == nil {
		return
	}
	return *v, true
}

// OldSubject returns the old "subject" field's value of the Email entity.
// If the Email object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailMutation) OldSubject(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSubject is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSubject requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSubject: %w", err)
	}
	return oldValue.Subject, nil
}

// ResetSubject resets all changes to the "subject" field.
func (m *EmailMutation) ResetSubject() {
	m.subject = nil
}

// SetBody sets the "body" field.
func (m *EmailMutation) SetBody(s string) {
	m.body = &s
}

// Body returns the value of the "body" field in the mutation.
func (m *EmailMutation) Body() (r string, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old "body" field's value of the Email entity.
// If the Email object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ResetBody resets all changes to the "body" field.
func (m *EmailMutation) ResetBody() {
	m.body = nil
}

// SetSentAt sets the "sentAt" field.
func (m *EmailMutation) SetSentAt(t time.Time) {
	m.sentAt = &t
}

// SentAt returns the value of the "sentAt" field in the mutation.
func (m *EmailMutation) SentAt() (r time.Time, exists bool) {
	v := m.sentAt
	if v == nil {
		return
	}
	return *v, true
}

// OldSentAt returns the old "sentAt" field's value of the Email entity.
// If the Email object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailMutation) OldSentAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSentAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSentAt: %w", err)
	}
	return oldValue.SentAt, nil
}

// ResetSentAt resets all changes to the "sentAt" field.
func (m *EmailMutation) ResetSentAt() {
	m.sentAt = nil
}

// SetArchived sets the "archived" field.
func (m *EmailMutation) SetArchived(b bool) {
	m.archived = &b
}

// Archived returns the value of the "archived" field in the mutation.
func (m *EmailMutation) Archived() (r bool, exists bool) {
	v := m.archived
	if v == nil {
		return
	}
	return *v, true
}

// OldArchived returns the old "archived" field's value of the Email entity.
// If the Email object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailMutation) OldArchived(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldArchived is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldArchived requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldArchived: %w", err)
	}
	return oldValue.Archived, nil
}

// ResetArchived resets all changes to the "archived" field.
func (m *EmailMutation) ResetArchived() {
	m.archived = nil
}

// SetRead sets the "read" field.
func (m *EmailMutation) SetRead(b bool) {
	m.read = &b
}

// Read returns the value of the "read" field in the mutation.
func (m *EmailMutation) Read() (r bool, exists bool) {
	v := m.read
	if v == nil {
		return
	}
	return *v, true
}

// OldRead returns the old "read" field's value of the Email entity.
// If the Email object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EmailMutation) OldRead(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRead is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRead requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRead: %w", err)
	}
	return oldValue.Read, nil
}

// ResetRead resets all changes to the "read" field.
func (m *EmailMutation) ResetRead() {
	m.read = nil
}

// SetUserID sets the "user" edge to the User entity by id.
func (m *EmailMutation) SetUserID(id uuid.UUID) {
	m.user = &id
}

// ClearUser clears the "user" edge to the User entity.
func (m *EmailMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *EmailMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *EmailMutation) UserID() (id uuid.UUID, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *EmailMutation) UserIDs() (ids []uuid.UUID) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *EmailMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the EmailMutation builder.
func (m *EmailMutation) Where(ps ...predicate.Email) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *EmailMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Email).
func (m *EmailMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EmailMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.from != nil {
		fields = append(fields, email.FieldFrom)
	}
	if m.subject != nil {
		fields = append(fields, email.FieldSubject)
	}
	if m.body != nil {
		fields = append(fields, email.FieldBody)
	}
	if m.sentAt != nil {
		fields = append(fields, email.FieldSentAt)
	}
	if m.archived != nil {
		fields = append(fields, email.FieldArchived)
	}
	if m.read != nil {
		fields = append(fields, email.FieldRead)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EmailMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case email.FieldFrom:
		return m.From()
	case email.FieldSubject:
		return m.Subject()
	case email.FieldBody:
		return m.Body()
	case email.FieldSentAt:
		return m.SentAt()
	case email.FieldArchived:
		return m.Archived()
	case email.FieldRead:
		return m.Read()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EmailMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case email.FieldFrom:
		return m.OldFrom(ctx)
	case email.FieldSubject:
		return m.OldSubject(ctx)
	case email.FieldBody:
		return m.OldBody(ctx)
	case email.FieldSentAt:
		return m.OldSentAt(ctx)
	case email.FieldArchived:
		return m.OldArchived(ctx)
	case email.FieldRead:
		return m.OldRead(ctx)
	}
	return nil, fmt.Errorf("unknown Email field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailMutation) SetField(name string, value ent.Value) error {
	switch name {
	case email.FieldFrom:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFrom(v)
		return nil
	case email.FieldSubject:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSubject(v)
		return nil
	case email.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	case email.FieldSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSentAt(v)
		return nil
	case email.FieldArchived:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetArchived(v)
		return nil
	case email.FieldRead:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRead(v)
		return nil
	}
	return fmt.Errorf("unknown Email field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EmailMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EmailMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EmailMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Email numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EmailMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EmailMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EmailMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Email nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EmailMutation) ResetField(name string) error {
	switch name {
	case email.FieldFrom:
		m.ResetFrom()
		return nil
	case email.FieldSubject:
		m.ResetSubject()
		return nil
	case email.FieldBody:
		m.ResetBody()
		return nil
	case email.FieldSentAt:
		m.ResetSentAt()
		return nil
	case email.FieldArchived:
		m.ResetArchived()
		return nil
	case email.FieldRead:
		m.ResetRead()
		return nil
	}
	return fmt.Errorf("unknown Email field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EmailMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, email.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EmailMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case email.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EmailMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EmailMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EmailMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, email.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EmailMutation) EdgeCleared(name string) bool {
	switch name {
	case email.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EmailMutation) ClearEdge(name string) error {
	switch name {
	case email.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Email unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EmailMutation) ResetEdge(name string) error {
	switch name {
	case email.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Email edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	username      *string
	first_name    *string
	last_name     *string
	password      *string
	email         *string
	clearedFields map[string]struct{}
	emails        map[int]struct{}
	removedemails map[int]struct{}
	clearedemails bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetFirstName sets the "first_name" field.
func (m *UserMutation) SetFirstName(s string) {
	m.first_name = &s
}

// FirstName returns the value of the "first_name" field in the mutation.
func (m *UserMutation) FirstName() (r string, exists bool) {
	v := m.first_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "first_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "first_name" field.
func (m *UserMutation) ResetFirstName() {
	m.first_name = nil
}

// SetLastName sets the "last_name" field.
func (m *UserMutation) SetLastName(s string) {
	m.last_name = &s
}

// LastName returns the value of the "last_name" field in the mutation.
func (m *UserMutation) LastName() (r string, exists bool) {
	v := m.last_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "last_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "last_name" field.
func (m *UserMutation) ResetLastName() {
	m.last_name = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// AddEmailIDs adds the "emails" edge to the Email entity by ids.
func (m *UserMutation) AddEmailIDs(ids ...int) {
	if m.emails == nil {
		m.emails = make(map[int]struct{})
	}
	for i := range ids {
		m.emails[ids[i]] = struct{}{}
	}
}

// ClearEmails clears the "emails" edge to the Email entity.
func (m *UserMutation) ClearEmails() {
	m.clearedemails = true
}

// EmailsCleared reports if the "emails" edge to the Email entity was cleared.
func (m *UserMutation) EmailsCleared() bool {
	return m.clearedemails
}

// RemoveEmailIDs removes the "emails" edge to the Email entity by IDs.
func (m *UserMutation) RemoveEmailIDs(ids ...int) {
	if m.removedemails == nil {
		m.removedemails = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.emails, ids[i])
		m.removedemails[ids[i]] = struct{}{}
	}
}

// RemovedEmails returns the removed IDs of the "emails" edge to the Email entity.
func (m *UserMutation) RemovedEmailsIDs() (ids []int) {
	for id := range m.removedemails {
		ids = append(ids, id)
	}
	return
}

// EmailsIDs returns the "emails" edge IDs in the mutation.
func (m *UserMutation) EmailsIDs() (ids []int) {
	for id := range m.emails {
		ids = append(ids, id)
	}
	return
}

// ResetEmails resets all changes to the "emails" edge.
func (m *UserMutation) ResetEmails() {
	m.emails = nil
	m.clearedemails = false
	m.removedemails = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.first_name != nil {
		fields = append(fields, user.FieldFirstName)
	}
	if m.last_name != nil {
		fields = append(fields, user.FieldLastName)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUsername:
		return m.Username()
	case user.FieldFirstName:
		return m.FirstName()
	case user.FieldLastName:
		return m.LastName()
	case user.FieldPassword:
		return m.Password()
	case user.FieldEmail:
		return m.Email()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldFirstName:
		return m.OldFirstName(ctx)
	case user.FieldLastName:
		return m.OldLastName(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case user.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldFirstName:
		m.ResetFirstName()
		return nil
	case user.FieldLastName:
		m.ResetLastName()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.emails != nil {
		edges = append(edges, user.EdgeEmails)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeEmails:
		ids := make([]ent.Value, 0, len(m.emails))
		for id := range m.emails {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedemails != nil {
		edges = append(edges, user.EdgeEmails)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeEmails:
		ids := make([]ent.Value, 0, len(m.removedemails))
		for id := range m.removedemails {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedemails {
		edges = append(edges, user.EdgeEmails)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeEmails:
		return m.clearedemails
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeEmails:
		m.ResetEmails()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
