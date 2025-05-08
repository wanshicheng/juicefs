// Copyright 2022 Juicedata Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wasm
// +build wasm

package meta

import (
	"context"
	"database/sql"
	"io"
	"reflect"
	"syscall"
	"time"
)

type Engine struct {
}

// Rows represents an iterateable object for rows
type Rows struct{}

// Close closes the iteration
func (rs *Rows) Close() error {
	return syscall.ENOSYS
}

// Next returns true if there is another row
func (rs *Rows) Next() bool {
	return false
}

// Err returns error
func (rs *Rows) Err() error {
	return syscall.ENOSYS
}

// Scan assigns the values to the provided variables
func (rs *Rows) Scan(dest ...interface{}) error {
	return syscall.ENOSYS
}

// DriverName return the current sql driver's name
func (engine *Engine) DriverName() string {
	return ""
}

// DataSourceName return the current connection string
func (engine *Engine) DataSourceName() string {
	return ""
}

// EnableSessionID if enable session id
func (engine *Engine) EnableSessionID(enable bool) {
}

// SetDisableGlobalCache disable global cache or not
func (engine *Engine) SetDisableGlobalCache(disable bool) {
}

// ShowSQL show SQL statement or not on logger if log level is great than INFO
func (engine *Engine) ShowSQL(show ...bool) {
}

// Logger return the logger interface
func (engine *Engine) Logger() interface{} {
	return nil
}

// SetLogger set the new logger
func (engine *Engine) SetLogger(logger interface{}) {
}

// SetLogLevel sets the logger level
func (engine *Engine) SetLogLevel(level interface{}) {
}

// BufferSize sets buffer size for iterate
func (engine *Engine) BufferSize(size int) interface{} {
	return nil
}

// SetQuotePolicy sets the special quote policy
func (engine *Engine) SetQuotePolicy(quotePolicy interface{}) {
}

// Update records, bean's non-empty fields are updated contents,
// condiBean' non-empty filds are conditions
// CAUTION:
//
//	1.bool will defaultly be updated content nor conditions
//	 You should call UseBool if you have bool to use.
//	2.float32 & float64 may be not inexact as conditions
func (engine *Engine) Update(bean interface{}, condiBeans ...interface{}) (int64, error) {
	return 0, syscall.ENOSYS
}

// Delete records, bean's non-empty fields are conditions
func (engine *Engine) Delete(bean interface{}) (int64, error) {
	return 0, syscall.ENOSYS
}

// Get retrieve one record from table, bean's non-empty fields
// are conditions
func (engine *Engine) Get(bean interface{}) (bool, error) {
	return false, syscall.ENOSYS
}

// Exist returns true if the record exist otherwise return false
func (engine *Engine) Exist(bean ...interface{}) (bool, error) {
	return false, syscall.ENOSYS
}

// Find retrieve records from table, condiBeans's non-empty fields
// are conditions. beans could be []Struct, []*Struct, map[int64]Struct
// map[int64]*Struct
func (engine *Engine) Find(beans interface{}, condiBeans ...interface{}) error {
	return syscall.ENOSYS
}

// FindAndCount find the results and also return the counts
func (engine *Engine) FindAndCount(rowsSlicePtr interface{}, condiBean ...interface{}) (int64, error) {
	return 0, syscall.ENOSYS
}

// Iterate record by record handle records from table, bean's non-empty fields
// are conditions.
func (engine *Engine) Iterate(bean interface{}, fun func(idx int64, bean interface{}) error) error {
	return syscall.ENOSYS
}

// Rows return sql.Rows compatible Rows obj, as a forward Iterator object for iterating record by record, bean's non-empty fields
// are conditions.
func (engine *Engine) Rows(bean interface{}) (*Rows, error) {
	return nil, syscall.ENOSYS
}

// Count counts the records. bean's non-empty fields are conditions.
func (engine *Engine) Count(bean ...interface{}) (int64, error) {
	return 0, syscall.ENOSYS
}

// Sum sum the records by some column. bean's non-empty fields are conditions.
func (engine *Engine) Sum(bean interface{}, colName string) (float64, error) {
	return 0, syscall.ENOSYS
}

// SumInt sum the records by some column. bean's non-empty fields are conditions.
func (engine *Engine) SumInt(bean interface{}, colName string) (int64, error) {
	return 0, syscall.ENOSYS
}

// Sums sum the records by some columns. bean's non-empty fields are conditions.
func (engine *Engine) Sums(bean interface{}, colNames ...string) ([]float64, error) {
	return nil, syscall.ENOSYS
}

// SumsInt like Sums but return slice of int64 instead of float64.
func (engine *Engine) SumsInt(bean interface{}, colNames ...string) ([]int64, error) {
	return nil, syscall.ENOSYS
}

// Query a raw sql and return records as []map[string][]byte
func (engine *Engine) Query(sqlOrArgs ...interface{}) (resultsSlice []map[string][]byte, err error) {
	return nil, syscall.ENOSYS
}

// QueryString runs a raw sql and return records as []map[string]string
func (engine *Engine) QueryString(sqlOrArgs ...interface{}) ([]map[string]string, error) {
	return nil, syscall.ENOSYS
}

// QueryInterface runs a raw sql and return records as []map[string]interface{}
func (engine *Engine) QueryInterface(sqlOrArgs ...interface{}) ([]map[string]interface{}, error) {
	return nil, syscall.ENOSYS
}

// Exec raw sql
func (engine *Engine) Exec(sqlOrArgs ...interface{}) (int64, error) {
	return 0, syscall.ENOSYS
}

// Insert one or more records
func (engine *Engine) Insert(beans ...interface{}) (int64, error) {
	return 0, syscall.ENOSYS
}

// InsertOne insert only one record
func (engine *Engine) InsertOne(bean interface{}) (int64, error) {
	return 0, syscall.ENOSYS
}

// SQL method let's you manually write raw SQL and operate
func (engine *Engine) SQL(query interface{}, args ...interface{}) interface{} {
	return nil
}

// NoAutoTime Default if your struct has "created" or "updated" filed tag, the fields
// will automatically be filled with current time when Insert or Update
// invoked. Call NoAutoTime if you dont' want to fill automatically.
func (engine *Engine) NoAutoTime() interface{} {
	return nil
}

// NoAutoCondition disable auto generate Where condition from bean or not
func (engine *Engine) NoAutoCondition(no ...bool) interface{} {
	return nil
}

// Cascade use cascade or not
func (engine *Engine) Cascade(trueOrFalse ...bool) interface{} {
	return nil
}

// Where method provide a condition query
func (engine *Engine) Where(query interface{}, args ...interface{}) interface{} {
	return nil
}

// ID method provoide a condition as (id) = ?
func (engine *Engine) ID(id interface{}) interface{} {
	return nil
}

// Before apply before Processor, affected bean is passed to closure arg
func (engine *Engine) Before(closures func(interface{})) interface{} {
	return nil
}

// After apply after insert Processor, affected bean is passed to closure arg
func (engine *Engine) After(closures func(interface{})) interface{} {
	return nil
}

// Charset set charset when create table, only support mysql now
func (engine *Engine) Charset(charset string) interface{} {
	return nil
}

// StoreEngine set store engine when create table, only support mysql now
func (engine *Engine) StoreEngine(storeEngine string) interface{} {
	return nil
}

// Distinct use for distinct columns
func (engine *Engine) Distinct(columns ...string) interface{} {
	return nil
}

// Select customerize your select columns or contents
func (engine *Engine) Select(str string) interface{} {
	return nil
}

// Cols only use the parameters as select or update columns
func (engine *Engine) Cols(columns ...string) interface{} {
	return nil
}

// AllCols indicates that all columns should be use
func (engine *Engine) AllCols() interface{} {
	return nil
}

// MustCols specify some columns must use even if they are empty
func (engine *Engine) MustCols(columns ...string) interface{} {
	return nil
}

// UseBool use bool fields
func (engine *Engine) UseBool(columns ...string) interface{} {
	return nil
}

// Omit only not use the parameters as select or update columns
func (engine *Engine) Omit(columns ...string) interface{} {
	return nil
}

// Nullable set null when column is zero-value and nullable for update
func (engine *Engine) Nullable(columns ...string) interface{} {
	return nil
}

// In will generate "column IN (?, ?)"
func (engine *Engine) In(column string, args ...interface{}) interface{} {
	return nil
}

// NotIn will generate "column NOT IN (?, ?)"
func (engine *Engine) NotIn(column string, args ...interface{}) interface{} {
	return nil
}

// Incr provides a update string like "column = column + ?"
func (engine *Engine) Incr(column string, arg ...interface{}) interface{} {
	return nil
}

// Decr provides a update string like "column = column - ?"
func (engine *Engine) Decr(column string, arg ...interface{}) interface{} {
	return nil
}

// SetExpr provides a update string like "column = {expression}"
func (engine *Engine) SetExpr(column string, expression interface{}) interface{} {
	return nil
}

// Table temporarily change the Get, Find, Update's table
func (engine *Engine) Table(tableNameOrBean interface{}) interface{} {
	return nil
}

// Alias set the table alias
func (engine *Engine) Alias(alias string) interface{} {
	return nil
}

// Limit will generate "LIMIT start, limit"
func (engine *Engine) Limit(limit int, start ...int) interface{} {
	return nil
}

// Desc will generate "ORDER BY column1 DESC, column2 DESC"
func (engine *Engine) Desc(colNames ...string) interface{} {
	return nil
}

// Asc will generate "ORDER BY column1,column2 Asc"
func (engine *Engine) Asc(colNames ...string) interface{} {
	return nil
}

// OrderBy will generate "ORDER BY order"
func (engine *Engine) OrderBy(order string) interface{} {
	return nil
}

// Join the join_operator should be one of INNER, LEFT OUTER, CROSS etc
func (engine *Engine) Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) interface{} {
	return nil
}

// GroupBy generate group by statement
func (engine *Engine) GroupBy(keys string) interface{} {
	return nil
}

// Having generate having statement
func (engine *Engine) Having(conditions string) interface{} {
	return nil
}

// NewSession New a session
func (engine *Engine) NewSession() interface{} {
	return nil
}

// Close the engine
func (engine *Engine) Close() error {
	return syscall.ENOSYS
}

// Ping tests if database is alive
func (engine *Engine) Ping() error {
	return syscall.ENOSYS
}

// NewDB provides an interface to operate database directly
func (engine *Engine) NewDB() (interface{}, error) {
	return nil, syscall.ENOSYS
}

// DB return the wrapper of sql.DB
func (engine *Engine) DB() interface{} {
	return nil
}

// Dialect return database dialect
func (engine *Engine) Dialect() interface{} {
	return nil
}

// SetCacher sets cacher for the table
func (engine *Engine) SetCacher(tableName string, cacher interface{}) {
}

// GetCacher returns the cachher of the special table
func (engine *Engine) GetCacher(tableName string) interface{} {
	return nil
}

// SetMapper set the name mapping rules
func (engine *Engine) SetMapper(mapper interface{}) {
}

// SetTableMapper set the table name mapping rule
func (engine *Engine) SetTableMapper(mapper interface{}) {
}

// SetColumnMapper set the column name mapping rule
func (engine *Engine) SetColumnMapper(mapper interface{}) {
}

// GetColumnMapper returns the column name mapper
func (engine *Engine) GetColumnMapper() interface{} {
	return nil
}

// GetTableMapper returns the table name mapper
func (engine *Engine) GetTableMapper() interface{} {
	return nil
}

// GetTZLocation returns time zone of the application
func (engine *Engine) GetTZLocation() *time.Location {
	return time.UTC
}

// SetTZLocation sets time zone of the application
func (engine *Engine) SetTZLocation(tz *time.Location) {
}

// GetTZDatabase returns time zone of the database
func (engine *Engine) GetTZDatabase() *time.Location {
	return time.UTC
}

// SetTZDatabase sets time zone of the database
func (engine *Engine) SetTZDatabase(tz *time.Location) {
}

// Quote Use QuoteStr quote the string sql
func (engine *Engine) Quote(value string) string {
	return ""
}

// QuoteTo quotes string and writes into the buffer
func (engine *Engine) QuoteTo(buf interface{}, value string) {
}

// SQLType A simple wrapper to dialect's core.SqlType method
func (engine *Engine) SQLType(c interface{}) string {
	return ""
}

// AutoIncrStr Database's autoincrement statement
func (engine *Engine) AutoIncrStr() string {
	return ""
}

// SetMaxOpenConns is only available for go 1.2+
func (engine *Engine) SetMaxOpenConns(conns int) {
}

// SetMaxIdleConns set the max idle connections on pool, default is 2
func (engine *Engine) SetMaxIdleConns(conns int) {
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused
func (engine *Engine) SetConnMaxLifetime(d time.Duration) {
}

// SetDefaultCacher set the default cacher
func (engine *Engine) SetDefaultCacher(cacher interface{}) {
}

// GetDefaultCacher returns the default cacher
func (engine *Engine) GetDefaultCacher() interface{} {
	return nil
}

// NoCache if you has set default cacher, and you want temporilly stop use cache
func (engine *Engine) NoCache() interface{} {
	return nil
}

// NoCascade If you do not want to auto cascade load object
func (engine *Engine) NoCascade() interface{} {
	return nil
}

// MapCacher Set a table use a special cacher
func (engine *Engine) MapCacher(bean interface{}, cacher interface{}) error {
	return syscall.ENOSYS
}

// CreateIndexes create indexes
func (engine *Engine) CreateIndexes(bean interface{}) error {
	return syscall.ENOSYS
}

// CreateUniques create uniques
func (engine *Engine) CreateUniques(bean interface{}) error {
	return syscall.ENOSYS
}

// ClearCacheBean if enabled cache, clear the cache bean
func (engine *Engine) ClearCacheBean(bean interface{}, id string) error {
	return syscall.ENOSYS
}

// ClearCache if enabled cache, clear some tables' cache
func (engine *Engine) ClearCache(beans ...interface{}) error {
	return syscall.ENOSYS
}

// UnMapType remove table from tables cache
func (engine *Engine) UnMapType(t reflect.Type) {
}

// Sync the new struct changes to database
func (engine *Engine) Sync(beans ...interface{}) error {
	return syscall.ENOSYS
}

// Sync2 synchronize structs to database tables
func (engine *Engine) Sync2(beans ...interface{}) error {
	return syscall.ENOSYS
}

// CreateTables create tables according bean
func (engine *Engine) CreateTables(beans ...interface{}) error {
	return syscall.ENOSYS
}

// DropTables drop specify tables
func (engine *Engine) DropTables(beans ...interface{}) error {
	return syscall.ENOSYS
}

// DropIndexes drop indexes of a table
func (engine *Engine) DropIndexes(bean interface{}) error {
	return syscall.ENOSYS
}

// ImportFile SQL DDL file
func (engine *Engine) ImportFile(ddlPath string) ([]interface{}, error) {
	return nil, syscall.ENOSYS
}

// Import SQL DDL from io.Reader
func (engine *Engine) Import(r io.Reader) ([]interface{}, error) {
	return nil, syscall.ENOSYS
}

// TableInfo get table info according to bean's content
func (engine *Engine) TableInfo(bean interface{}) (interface{}, error) {
	return nil, syscall.ENOSYS
}

// IsTableEmpty if a table has any record
func (engine *Engine) IsTableEmpty(bean interface{}) (bool, error) {
	return false, syscall.ENOSYS
}

// IsTableExist if a table is exist
func (engine *Engine) IsTableExist(beanOrTableName interface{}) (bool, error) {
	return false, syscall.ENOSYS
}

// TableName returns table name with schema prefix if has
func (engine *Engine) TableName(bean interface{}, includeSchema ...bool) string {
	return ""
}

// tbNameWithSchema returns table name with schema
func (engine *Engine) tbNameWithSchema(v string) string {
	return ""
}

// DBMetas Retrieve all tables, columns, indexes' informations from database
func (engine *Engine) DBMetas() ([]interface{}, error) {
	return nil, syscall.ENOSYS
}

// DumpAllToFile dump database all table structs and data to a file
func (engine *Engine) DumpAllToFile(fp string, tp ...interface{}) error {
	return syscall.ENOSYS
}

// DumpAll dump database all table structs and data to w
func (engine *Engine) DumpAll(w io.Writer, tp ...interface{}) error {
	return syscall.ENOSYS
}

// DumpTablesToFile dump specified tables to SQL file
func (engine *Engine) DumpTablesToFile(tables []interface{}, fp string, tp ...interface{}) error {
	return syscall.ENOSYS
}

// DumpTables dump specify tables to io.Writer
func (engine *Engine) DumpTables(tables []interface{}, w io.Writer, tp ...interface{}) error {
	return syscall.ENOSYS
}

// SetSchema sets the schema of database
func (engine *Engine) SetSchema(schema string) {
}

// AddHook adds a context Hook
func (engine *Engine) AddHook(hook interface{}) {
}

// Unscoped always disable struct tag "deleted"
func (engine *Engine) Unscoped() interface{} {
	return nil
}

// Context creates a session with the context
func (engine *Engine) Context(ctx context.Context) interface{} {
	return nil
}

// SetDefaultContext set the default context
func (engine *Engine) SetDefaultContext(ctx context.Context) {
}

// PingContext tests if database is alive
func (engine *Engine) PingContext(ctx context.Context) error {
	return syscall.ENOSYS
}

// Transaction Execute sql wrapped in a transaction, tx will automatic commit if no errors occurred
func (engine *Engine) Transaction(f func(interface{}) (interface{}, error), opts ...*sql.TxOptions) (interface{}, error) {
	return nil, syscall.ENOSYS
}
