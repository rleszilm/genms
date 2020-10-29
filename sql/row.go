package sql

// Row is the result from queries that return a single value
type Row interface {
	Scan(...interface{}) error
	Columns() ([]string, error)
	Err() error
	SliceScan() ([]interface{}, error)
	MapScan(map[string]interface{}) error
	StructScan(interface{}) error
}
