// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

// IndexColumn represents index column info.
type IndexColumn struct {
	SeqNo      int    // seq_no
	Cid        int    // cid
	ColumnName string // column_name
}

// PgIndexColumns runs a custom query, returning results as IndexColumn.
func PgIndexColumns(db XODB, schema string, index string) ([]*IndexColumn, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`(row_number() over()), ` + // ::integer AS seq_no
		`a.attnum, ` + // ::integer AS cid
		`a.attname ` + // ::varchar AS column_name
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`LEFT JOIN pg_attribute a ON i.indrelid = a.attrelid AND a.attnum = ANY(i.indkey) AND a.attisdropped = false ` +
		`WHERE i.indkey <> '0' AND n.nspname = $1 AND ic.relname = $2`

	// run query
	XOLog(sqlstr, schema, index)
	q, err := db.Query(sqlstr, schema, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IndexColumn{}
	for q.Next() {
		ic := IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}

	return res, nil
}

// MyIndexColumns runs a custom query, returning results as IndexColumn.
func MyIndexColumns(db XODB, schema string, table string, index string) ([]*IndexColumn, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`seq_in_index AS seq_no, ` +
		`column_name ` +
		`FROM information_schema.statistics ` +
		`WHERE index_schema = ? AND table_name = ? AND index_name = ? ` +
		`ORDER BY seq_in_index`

	// run query
	XOLog(sqlstr, schema, table, index)
	q, err := db.Query(sqlstr, schema, table, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IndexColumn{}
	for q.Next() {
		ic := IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}

	return res, nil
}

// SqIndexColumns runs a custom query, returning results as IndexColumn.
func SqIndexColumns(db XODB, index string) ([]*IndexColumn, error) {
	var err error

	// sql query
	var sqlstr = `PRAGMA index_info(` + index + `)`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IndexColumn{}
	for q.Next() {
		ic := IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}

	return res, nil
}

// OrIndexColumns runs a custom query, returning results as IndexColumn.
func OrIndexColumns(db XODB, schema string, table string, index string) ([]*IndexColumn, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`column_position AS seq_no, ` +
		`LOWER(column_name) AS column_name ` +
		`FROM all_ind_columns ` +
		`WHERE index_owner = UPPER(:1) AND table_name = UPPER(:2) AND index_name = UPPER(:3) ` +
		`ORDER BY column_position`

	// run query
	XOLog(sqlstr, schema, table, index)
	q, err := db.Query(sqlstr, schema, table, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IndexColumn{}
	for q.Next() {
		ic := IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}

	return res, nil
}

// MsIndexColumns runs a custom query, returning results as IndexColumn.
func MsIndexColumns(db XODB, schema string, table string, index string) ([]*IndexColumn, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`k.keyno AS seq_no, ` +
		`k.colid AS cid, ` +
		`c.name AS column_name ` +
		`FROM sysindexes i ` +
		`INNER JOIN sysobjects o ON i.id = o.id ` +
		`INNER JOIN sysindexkeys k ON k.id = o.id AND k.indid = i.indid ` +
		`INNER JOIN syscolumns c ON c.id = o.id AND c.colid = k.colid ` +
		`WHERE o.type = 'U' AND SCHEMA_NAME(o.uid) = $1 AND o.name = $2 AND i.name = $3 ` +
		`ORDER BY k.keyno`

	// run query
	XOLog(sqlstr, schema, table, index)
	q, err := db.Query(sqlstr, schema, table, index)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*IndexColumn{}
	for q.Next() {
		ic := IndexColumn{}

		// scan
		err = q.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ic)
	}

	return res, nil
}
