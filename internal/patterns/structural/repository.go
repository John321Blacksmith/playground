// This package contains a model
// of the repository design pattern
// implementation and the corresponding
// real-life example
package structural

import (
	"fmt"
)

// MODEL

type Entity struct{}

type DBConfigs struct {
	port        string
	ssl         bool
	connTimeout int
	poolCount   int
}

type DBDriver struct {
	name    string
	configs *DBConfigs
}

func NewDriver(cfg *DBConfigs, name string) *DBDriver {
	return &DBDriver{
		name:    name,
		configs: cfg,
	}
}

func (d *DBDriver) PerformQueryMethod(query string) ([]struct{}, error) {
	fmt.Printf("Applying DML '%s' to the DB.", query)
	return []struct{}{}, nil
}

type DBRepository struct {
	driver *DBDriver
}

func NewDBRepository(driver *DBDriver) *DBRepository {
	return &DBRepository{driver: driver}
}

func (repo *DBRepository) QuerySomething(count int) ([]Entity, error) {
	var someQuery = `SELECT * FROM table;`
	var entities []Entity
	fmt.Printf("Applying client parameters '%d' with query.", count)
	results, err := repo.driver.PerformQueryMethod(someQuery)
	if err != nil {
		return nil, fmt.Errorf("Query did not execute well: %s", err)
	}
	for i := range results {
		entities = append(entities, results[i])
	}
	return entities, nil
}
