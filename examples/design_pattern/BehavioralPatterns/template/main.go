package main

import (
	"database/sql"
	"io/ioutil"
	"log"
)

/*
the Template Method Pattern
	> the Template Method Pattern provides a template for performing an operation but
	allows the subclasses to modify certain steps of the operation without changing its overall structure.
*/

type DataReaderProcessor interface {
	ReadData() ([]byte, error)
	ProcessData(data []byte) error
	GetDataSourceName() string
}

type DataReaderProcessorTemplate struct {
	dataSource DataReaderProcessor
}

func (p *DataReaderProcessorTemplate) ReadAndProcessData() error {
	data, err := p.dataSource.ReadData()
	if err != nil {
		return err
	}
	err = p.dataSource.ProcessData(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *DataReaderProcessorTemplate) GetDataSourceName() string {
	return p.dataSource.GetDataSourceName()
}

type FileReaderProcessor struct {
	DataReaderProcessorTemplate
	fileName string
}

func (p *FileReaderProcessor) ReadData() ([]byte, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *FileReaderProcessor) ProcessData(data []byte) error {
	// process data
	return nil
}

func (p *FileReaderProcessor) GetDataSourceName() string {
	return p.fileName
}

type DatabaseReaderProcessor struct {
	DataReaderProcessorTemplate
	db *sql.DB
}

func (p *DatabaseReaderProcessor) ReadData() ([]byte, error) {
	// read data from database
	return nil, nil
}

func (p *DatabaseReaderProcessor) ProcessData(data []byte) error {
	// process data
	return nil
}

func (p *DatabaseReaderProcessor) GetDataSourceName() string {
	return "Database"
}

func main() {
	processor := &FileReaderProcessor{
		DataReaderProcessorTemplate: DataReaderProcessorTemplate{
			dataSource: &FileReaderProcessor{
				fileName: "/path/to/file",
			},
		},
	}
	err := processor.ReadAndProcessData()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
