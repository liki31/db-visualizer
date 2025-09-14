package models

type Column struct {
	Name string `json:"name"`
	DataType string `json:"dataType"`
	IsNull bool `json:"isNull"`
	IsPk bool `json:"isPk"`
}

type Table struct {
	Name string `json:"name"`
	Columns []Column `json:"columns"`
}