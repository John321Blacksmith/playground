package structural

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

// EXAMPLE

// Some business entities
type Log struct {
	level     string
	message   string
	timestamp string
}

type LogQuerysets []map[string][]Log

///////////

// Some DB driver
type FireBirdDriver struct {
	db string
}

func NewFBDriver() *FireBirdDriver {
	return &FireBirdDriver{db: "firebird"}
}

func (d *FireBirdDriver) PerformSelectQuery(query string) []byte {
	return []byte{}
}

///////////

// Log reading behavior
type LogReader interface {
	GetLogsByServiceName(service string) ([]Log, error)
}

// Some Repository
type LogRepository struct {
	driver *FireBirdDriver
}

func NewLogRepository(driver *FireBirdDriver) *LogRepository {
	return &LogRepository{driver: driver}
}

func (repo *LogRepository) GetLogsByServiceName(service string) ([]Log, error) {
	var logs []Log
	getLogsByServiceName := fmt.Sprintf(`
					SELECT
						id,
						s.name,
						level,
						message,
						timestamp,
						ip_address,
						name_process
					FROM
						logs
					JOIN
						services AS s
					ON
						service_id = s.id
					WHERE
						s.name = %s
					`, service)
	rawData := repo.driver.PerformSelectQuery(getLogsByServiceName)
	reader := bytes.NewReader(rawData)
	err := binary.Read(reader, binary.NativeEndian, logs)
	if err != nil {
		return nil, fmt.Errorf("Transaction failed, because %s", err)
	}
	return logs, nil
}

////////////

// Some Usecase
type LogExtractionUseCase struct {
	repo *LogRepository
}

func NewLogExtractionUseCase(repo LogReader) *LogExtractionUseCase {
	r := repo.(*LogRepository)
	return &LogExtractionUseCase{repo: r}
}

func (usecase *LogExtractionUseCase) GetLogsListByService(services []string) (LogQuerysets, error) {
	querySets := LogQuerysets{}
	for _, service := range services {
		logs, err := usecase.repo.GetLogsByServiceName(service)
		if err != nil {
			return nil, err
		}
		querySet := map[string][]Log{service: logs}
		querySets = append(querySets, querySet)
	}
	return querySets, nil
}

// Apply this pattern
var services []string = []string{"broker_auth", "broker_admin", "broker_terminals", "broker_client"}

func main() {
	driver := NewFBDriver()
	repository := NewLogRepository(driver)
	usecase := NewLogExtractionUseCase(repository)
	querysets, err := usecase.GetLogsListByService(services)
	if err != nil {
		log.Fatalf("Cannot fetch logs from the DB because %s", err)
	}
	if len(querysets) != 0 {
		for service, logs := range querysets {
			fmt.Println(service, logs)
		}
	}
}
