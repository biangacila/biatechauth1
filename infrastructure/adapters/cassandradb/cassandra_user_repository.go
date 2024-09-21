package cassandradb

import (
	"fmt"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/biangacila/biatechauth1/domain/entities"
	"github.com/gocql/gocql"
	"strings"
	"time"
)

type CassandraUserRepository struct {
	CassandraGenericRepository[entities.User]
	session *gocql.Session
}

// NewCassandraUserRepository creates a new instance of CassandraUserRepository
func NewCassandraUserRepository() *CassandraUserRepository {
	cassGeneric := NewCassandraGenericRepository(entities.User{})
	return &CassandraUserRepository{
		CassandraGenericRepository: *cassGeneric, // Initialize generic repo
		session:                    GetSession(),
	}
}

// Lock locks a user account by updating the status and logger info
func (c *CassandraUserRepository) Lock(code string) error {
	query := fmt.Sprintf("UPDATE %v.users SET status='locked',UpdatedAt=? WHERE  AND email='%v'",
		constants.DbName, code)
	return c.session.Query(query, time.Now()).Exec()
}
func (c *CassandraUserRepository) ResetPassword(email, password string) error {
	query := fmt.Sprintf("UPDATE %v.users SET password='%v',UpdatedAt=? WHERE  AND email='%v'",
		constants.DbName, password, email)
	return c.session.Query(query, time.Now()).Exec()
}

// UnLock unlocks a user account by updating the status and logger info
func (c *CassandraUserRepository) UnLock(code string) error {
	query := fmt.Sprintf("UPDATE %v.users SET status='active',UpdatedAt=? WHERE  AND email='%v'",
		constants.DbName, code)
	return c.session.Query(query, time.Now()).Exec()
}

func (c *CassandraUserRepository) FindByEmail(email string) (user entities.User, err error) {
	email = strings.ToLower(email)
	email = strings.TrimSpace(email)
	query := fmt.Sprintf("SELECT * FROM users WHERE email='%v'", email)
	user, err = FetchRecord(c.session, query, entities.User{})
	return user, err
}
