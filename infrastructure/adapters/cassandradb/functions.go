package cassandradb

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/biangacila/biatechauth1/constants"
	"github.com/gocql/gocql"
	"sort"
	"strconv"
	"strings"
)

func InsertRecord(session *gocql.Session, dbName, table string, record any) error {
	str, err := json.Marshal(record)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %v.%v  JSON '%v' ", dbName, table, string(str))
	return session.Query(query).Exec()
}
func FindRecord[T any](session *gocql.Session, query string, outs T) (T, error) {
	records, err := FetchData2(session, query, outs)
	if err != nil {
		fmt.Println("::error FetchData> ", err)
		return outs, err
	}
	var rec T
	for _, record := range records {
		rec = record
	}
	return rec, nil
}
func FetchRecord[T any](session *gocql.Session, query string, outs T) (T, error) {
	records, err := FetchData2(session, query, outs)
	if err != nil {
		fmt.Println("::error FetchData> ", err)
		return outs, err
	}
	if len(records) == 0 {
		return outs, fmt.Errorf("no records found")
	}
	var rec T
	for _, record := range records {
		rec = record
	}
	return rec, nil
}
func FetchData[T any](session *gocql.Session, query string, outs T) ([]T, error) {
	iter := session.Query(query, constants.AppName).Iter()
	rows, err := iter.SliceMap()
	if err != nil {
		return []T{}, err
	}
	b, _ := json.Marshal(rows)
	var data []T
	_ = json.Unmarshal(b, &data)
	return data, nil
}
func FetchData2[T any](session *gocql.Session, query string, outs T) ([]T, error) {
	iter := session.Query(query).Iter()
	rows, err := iter.SliceMap()
	if err != nil {
		return []T{}, err
	}
	b, _ := json.Marshal(rows)
	var data []T
	_ = json.Unmarshal(b, &data)
	return data, nil
}
func FetchRecordWithConditions[T any](session *gocql.Session, dbName, table string, conditions map[string]interface{}, outs T, allowFiltering ...string) ([]T, error) {
	qry := fmt.Sprintf("SELECT * FROM %v.%s ", dbName, table)
	qry = qry + " where "
	if conditions != nil {
		for k, v := range conditions {
			if _, ok := v.(string); ok {
				qry = qry + fmt.Sprintf(" %s ='%v' and", k, v)
			} else {
				qry = qry + fmt.Sprintf(" %s =%v and", k, v)
			}
		}
		qry = strings.TrimSuffix(qry, "and")
	}

	if len(allowFiltering) > 0 {
		qry = qry + allowFiltering[0]
	}

	data, err := FetchData2(session, qry, outs)
	return data, err
}
func ExecuteQuery(session *gocql.Session, query string) error {
	return session.Query(query).Exec()
}
func GenerateSequenceNumber(session *gocql.Session, table, fieldName, prefixCode string, prefixStart int) (nexCode string, err error) {
	query := fmt.Sprintf(`select %v as code from %v.%v  `, fieldName, constants.DbName, table)
	iter := session.Query(query).Iter()
	rows, err := iter.SliceMap()
	if err != nil {
		return "", err
	}
	b, _ := json.Marshal(rows)
	type PrefixConf struct {
		Code string
	}

	var codes []PrefixConf
	_ = json.Unmarshal(b, &codes)

	var ls []int
	for _, code := range codes {
		arr := strings.Split(code.Code, prefixCode)
		value, _ := strconv.Atoi(arr[1])
		ls = append(ls, value)
	}
	sort.Ints(ls)

	var value = int64(prefixStart)
	if len(ls) > 0 {
		lastIndex := len(ls) - 1
		value = int64(ls[lastIndex])
	}
	nextValue := fmt.Sprintf("%v%v", prefixCode, value+1)
	return nextValue, nil
}
func WhereClauseBuilder(conditions map[string]interface{}) (string, error) {
	qry := ""
	if conditions == nil {
		return qry, errors.New("conditions is nil")
	}
	if len(conditions) == 0 {
		return qry, errors.New("conditions is empty")
	}
	qry = qry + " where "
	for k, v := range conditions {
		if _, ok := v.(string); ok {
			qry = qry + fmt.Sprintf(" %s ='%v' and", k, v)
		} else {
			qry = qry + fmt.Sprintf(" %s =%v and", k, v)
		}
	}
	qry = strings.TrimSuffix(qry, "and")
	return qry, nil
}
func UpdateClauseBuilder(conditions map[string]interface{}) (string, error) {
	qry := ""
	if conditions == nil {
		return qry, errors.New("update conditions is nil")
	}
	if len(conditions) == 0 {
		return qry, errors.New("update conditions is empty")
	}
	qry = qry + " "
	for k, v := range conditions {
		if _, ok := v.(string); ok {
			qry = qry + fmt.Sprintf(" %s ='%v' ,", k, v)
		} else {
			qry = qry + fmt.Sprintf(" %s =%v ,", k, v)
		}
	}
	qry = strings.TrimSuffix(qry, ",")
	return qry, nil
}
