package cassandradb

import (
	"github.com/biangacila/biatechauth1/constants"
	"github.com/biangacila/biatechauth1/domain/entities"
	"github.com/gocql/gocql"
)

type CassandraLoginRepository struct {
	CassandraGenericRepository[entities.Login]
	session *gocql.Session
}

func NewCassandraLoginRepository() *CassandraLoginRepository {
	cassGeneric := NewCassandraGenericRepository(entities.Login{})
	return &CassandraLoginRepository{
		CassandraGenericRepository: *cassGeneric, // Initialize generic repo
		session:                    GetSession(),
	}
}

func (c CassandraLoginRepository) New(login *entities.Login) (*entities.Login, error) {
	err := InsertRecord(c.session, constants.DbName, "logins", login)
	return login, err
}

func (c CassandraLoginRepository) HasLogin(username string) (*entities.Login, error) {
	//TODO implement me
	panic("implement me")
}
