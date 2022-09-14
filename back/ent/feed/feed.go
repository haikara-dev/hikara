// Code generated by ent, DO NOT EDIT.

package feed

import (
	"time"
)

const (
	// Label holds the string label denoting the feed type in the database.
	Label = "feed"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldContents holds the string denoting the contents field in the database.
	FieldContents = "contents"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// EdgeSite holds the string denoting the site edge name in mutations.
	EdgeSite = "site"
	// Table holds the table name of the feed in the database.
	Table = "feeds"
	// SiteTable is the table that holds the site relation/edge.
	SiteTable = "feeds"
	// SiteInverseTable is the table name for the Site entity.
	// It exists in this package in order to avoid circular dependency with the "site" package.
	SiteInverseTable = "sites"
	// SiteColumn is the table column denoting the site relation/edge.
	SiteColumn = "site_feeds"
)

// Columns holds all SQL columns for feed fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldContents,
	FieldCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "feeds"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"site_feeds",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// ContentsValidator is a validator for the "contents" field. It is called by the builders before save.
	ContentsValidator func(string) error
	// DefaultCount holds the default value on creation for the "count" field.
	DefaultCount int
)
