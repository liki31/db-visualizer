package db

import (
	"fmt"

	"github.com/liki31/db-visualizer/backend/models"
)

func (c *DBConnection) GetTables() ([]models.Table, error) {
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE';`

	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, models.Table{Name: name})
	}
	return tables, nil
}

func (c *DBConnection) GetColumns(tableName string) ([]models.Column, error) {
	query := `
		SELECT column_name, data_type, is_nullable,
		       (SELECT EXISTS (
		           SELECT 1
		           FROM information_schema.table_constraints tc
		           JOIN information_schema.key_column_usage kcu
		             ON tc.constraint_name = kcu.constraint_name
		           WHERE tc.table_name = $1 AND tc.constraint_type = 'PRIMARY KEY'
		             AND kcu.column_name = c.column_name
		       ))
		FROM information_schema.columns c
		WHERE table_name = $1;`

	rows, err := c.DB.Query(query, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []models.Column
	for rows.Next() {
		var col models.Column
		var nullable string
		if err := rows.Scan(&col.Name, &col.DataType, &nullable, &col.IsPk); err != nil {
			return nil, err
		}
		col.IsNull = nullable == "YES"
		columns = append(columns, col)
	}
	return columns, nil
}

func (c *DBConnection) GetRows(tableName string, limit, offset int) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", tableName)

	rows, err := c.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, col := range cols {
			rowMap[col] = values[i]
		}
		results = append(results, rowMap)
	}
	return results, nil
}
