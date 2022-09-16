// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cubdesign/haikara/ent/article"
	"github.com/cubdesign/haikara/ent/site"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt time.Time `json:"published_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleQuery when eager-loading is set.
	Edges         ArticleEdges `json:"edges"`
	site_articles *int
}

// ArticleEdges holds the relations/edges for other nodes in the graph.
type ArticleEdges struct {
	// Site holds the value of the site edge.
	Site *Site `json:"site,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SiteOrErr returns the Site value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleEdges) SiteOrErr() (*Site, error) {
	if e.loadedTypes[0] {
		if e.Site == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: site.Label}
		}
		return e.Site, nil
	}
	return nil, &NotLoadedError{edge: "site"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			values[i] = new(sql.NullInt64)
		case article.FieldTitle, article.FieldURL:
			values[i] = new(sql.NullString)
		case article.FieldCreatedAt, article.FieldUpdatedAt, article.FieldPublishedAt:
			values[i] = new(sql.NullTime)
		case article.ForeignKeys[0]: // site_articles
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Article", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case article.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case article.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case article.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field published_at", values[i])
			} else if value.Valid {
				a.PublishedAt = value.Time
			}
		case article.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field site_articles", value)
			} else if value.Valid {
				a.site_articles = new(int)
				*a.site_articles = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySite queries the "site" edge of the Article entity.
func (a *Article) QuerySite() *SiteQuery {
	return (&ArticleClient{config: a.config}).QuerySite(a)
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return (&ArticleClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(a.URL)
	builder.WriteString(", ")
	builder.WriteString("published_at=")
	builder.WriteString(a.PublishedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article

func (a Articles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
