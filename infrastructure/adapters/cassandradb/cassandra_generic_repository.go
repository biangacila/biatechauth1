package cassandradb

import (
	"errors"
	"fmt"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/gocql/gocql"
)

type CassandraGenericRepository[T any] struct {
	session *gocql.Session
}

func (c *CassandraGenericRepository[T]) Find(entity string, fieldValues map[string]interface{}, t T) (T, error) {
	records, err := c.Get(entity, fieldValues, t)
	if err != nil {
		return t, err
	}
	if len(records) == 0 {
		return t, errors.New("no records found")
	}
	return records[(len(records) - 1)], nil
}

func (c *CassandraGenericRepository[T]) Get(entity string, fieldValues map[string]interface{}, t T) ([]T, error) {
	return FetchRecordWithConditions(c.session, constants.DbName, entity, fieldValues, t, " ALLOW FILTERING ")
}

func (c *CassandraGenericRepository[T]) Update(entity string, conditions, fieldValues map[string]interface{}, t T) error {
	valuesWhere, err := WhereClauseBuilder(conditions)
	if err != nil {
		return err
	}
	valuesToUpdate, err := UpdateClauseBuilder(fieldValues)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", constants.DbName, entity, valuesToUpdate, valuesWhere)

	return c.session.Query(query, fieldValues).Exec()
}

func (c *CassandraGenericRepository[T]) Delete(entity string, fieldValues map[string]interface{}, t T) error {
	valuesWhere, err := WhereClauseBuilder(fieldValues)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s.%s SET %s WHERE %s", constants.DbName, entity, valuesWhere)
	return c.session.Query(query, fieldValues).Exec()
}

func NewCassandraGenericRepository[T any](t T) *CassandraGenericRepository[T] {
	return &CassandraGenericRepository[T]{
		session: GetSession(),
	}
}
func (c *CassandraGenericRepository[T]) Save(entity string, record any, t T) error {
	return InsertRecord(c.session, constants.DbName, entity, record)
}
