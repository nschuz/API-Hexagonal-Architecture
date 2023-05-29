package postgres

import "gorm.io/gorm"

type client struct {
	db *gorm.DB
}

// NewClient returns a new insatnce to use postgres
func NewClient() *client {
	return &client{
		db: connect(),
	}
}

//db esta ailado no se deb eexponer esa es la idea de heaxagona

// Create stores a newy record in the database
func (c *client) Create(value interface{}) error {
	return c.db.Create(value).Error
}

// Firsr finds recors that math given conditions
func (c *client) First(dest interface{}, conds ...interface{}) error {
	return c.db.First(dest, conds...).Error
}
