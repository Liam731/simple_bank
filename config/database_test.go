package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	fmt.Println("test database connection start...")
	db := DatabaseConnection()
	assert.NotNil(t, db, "failed to connect to database")
}
