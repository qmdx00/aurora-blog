// Code generated by entc, DO NOT EDIT.

package ent

import (
	"aurora/blog/api/app/service/manager/internal/data/ent/theme"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Theme is the model entity for the Theme schema.
type Theme struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Theme) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case theme.FieldID:
			values[i] = new(sql.NullInt64)
		case theme.FieldCreatedAt, theme.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Theme", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Theme fields.
func (t *Theme) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case theme.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case theme.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case theme.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Theme.
// Note that you need to call Theme.Unwrap() before calling this method if this Theme
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Theme) Update() *ThemeUpdateOne {
	return (&ThemeClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Theme entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Theme) Unwrap() *Theme {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Theme is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Theme) String() string {
	var builder strings.Builder
	builder.WriteString("Theme(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Themes is a parsable slice of Theme.
type Themes []*Theme

func (t Themes) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
