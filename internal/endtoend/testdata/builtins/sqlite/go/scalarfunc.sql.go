// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: scalarfunc.sql

package querytest

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const getAbs = `-- name: GetAbs :one
SELECT abs(int_val) FROM test
`

func (q *Queries) GetAbs(ctx context.Context) (float64, error) {
	row := q.db.QueryRowContext(ctx, getAbs)
	var abs float64
	err := row.Scan(&abs)
	return abs, err
}

const getChanges = `-- name: GetChanges :one
SELECT changes()
`

func (q *Queries) GetChanges(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getChanges)
	var changes int64
	err := row.Scan(&changes)
	return changes, err
}

const getChar1 = `-- name: GetChar1 :one
SELECT char(65)
`

func (q *Queries) GetChar1(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getChar1)
	var char string
	err := row.Scan(&char)
	return char, err
}

const getChar3 = `-- name: GetChar3 :one
SELECT char(65, 66, 67)
`

func (q *Queries) GetChar3(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getChar3)
	var char string
	err := row.Scan(&char)
	return char, err
}

const getCoalesce = `-- name: GetCoalesce :one
SELECT coalesce(NULL, 1, 'test')
`

func (q *Queries) GetCoalesce(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getCoalesce)
	var coalesce interface{}
	err := row.Scan(&coalesce)
	return coalesce, err
}

const getFormat = `-- name: GetFormat :one
SELECT format('Hello %s', 'world')
`

func (q *Queries) GetFormat(ctx context.Context) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getFormat)
	var format sql.NullString
	err := row.Scan(&format)
	return format, err
}

const getGlob = `-- name: GetGlob :one
SELECT glob('a*c', 'abc')
`

func (q *Queries) GetGlob(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getGlob)
	var glob int64
	err := row.Scan(&glob)
	return glob, err
}

const getHex = `-- name: GetHex :one
SELECT hex(123456)
`

func (q *Queries) GetHex(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getHex)
	var hex string
	err := row.Scan(&hex)
	return hex, err
}

const getIfnull = `-- name: GetIfnull :one
SELECT ifnull(1, 2)
`

func (q *Queries) GetIfnull(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getIfnull)
	var ifnull interface{}
	err := row.Scan(&ifnull)
	return ifnull, err
}

const getIif = `-- name: GetIif :one
SELECT iif(1, 2, 3)
`

func (q *Queries) GetIif(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getIif)
	var iif interface{}
	err := row.Scan(&iif)
	return iif, err
}

const getInstr = `-- name: GetInstr :one
SELECT instr('hello', 'l')
`

func (q *Queries) GetInstr(ctx context.Context) (sql.NullInt64, error) {
	row := q.db.QueryRowContext(ctx, getInstr)
	var instr sql.NullInt64
	err := row.Scan(&instr)
	return instr, err
}

const getLastInsertRowID = `-- name: GetLastInsertRowID :one
SELECT last_insert_rowid()
`

func (q *Queries) GetLastInsertRowID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLastInsertRowID)
	var last_insert_rowid int64
	err := row.Scan(&last_insert_rowid)
	return last_insert_rowid, err
}

const getLength = `-- name: GetLength :one
SELECT length('12345')
`

func (q *Queries) GetLength(ctx context.Context) (sql.NullInt64, error) {
	row := q.db.QueryRowContext(ctx, getLength)
	var length sql.NullInt64
	err := row.Scan(&length)
	return length, err
}

const getLike2 = `-- name: GetLike2 :one
SELECT like('%bc%', 'abcd')
`

func (q *Queries) GetLike2(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLike2)
	var like int64
	err := row.Scan(&like)
	return like, err
}

const getLike3 = `-- name: GetLike3 :one
SELECT like('$%1%', '%100', '$')
`

func (q *Queries) GetLike3(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLike3)
	var like int64
	err := row.Scan(&like)
	return like, err
}

const getLikelihood = `-- name: GetLikelihood :one
SELECT likelihood('12345', 0.5)
`

func (q *Queries) GetLikelihood(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getLikelihood)
	var likelihood interface{}
	err := row.Scan(&likelihood)
	return likelihood, err
}

const getLikely = `-- name: GetLikely :one
SELECT likely('12345')
`

func (q *Queries) GetLikely(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getLikely)
	var likely interface{}
	err := row.Scan(&likely)
	return likely, err
}

const getLower = `-- name: GetLower :one
SELECT lower('ABCDE')
`

func (q *Queries) GetLower(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getLower)
	var lower string
	err := row.Scan(&lower)
	return lower, err
}

const getLtrim = `-- name: GetLtrim :one
SELECT ltrim(' ABCDE')
`

func (q *Queries) GetLtrim(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getLtrim)
	var ltrim string
	err := row.Scan(&ltrim)
	return ltrim, err
}

const getLtrim2 = `-- name: GetLtrim2 :one
SELECT ltrim(':ABCDE', ':')
`

func (q *Queries) GetLtrim2(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getLtrim2)
	var ltrim string
	err := row.Scan(&ltrim)
	return ltrim, err
}

const getMax3 = `-- name: GetMax3 :one
SELECT max(1, 3, 2)
`

func (q *Queries) GetMax3(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getMax3)
	var max interface{}
	err := row.Scan(&max)
	return max, err
}

const getMin3 = `-- name: GetMin3 :one
SELECT min(1, 3, 2)
`

func (q *Queries) GetMin3(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getMin3)
	var min interface{}
	err := row.Scan(&min)
	return min, err
}

const getNullif = `-- name: GetNullif :one
SELECT nullif(1, 2)
`

func (q *Queries) GetNullif(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getNullif)
	var nullif interface{}
	err := row.Scan(&nullif)
	return nullif, err
}

const getPrintf = `-- name: GetPrintf :one
SELECT printf('Hello %s', 'world')
`

func (q *Queries) GetPrintf(ctx context.Context) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getPrintf)
	var printf sql.NullString
	err := row.Scan(&printf)
	return printf, err
}

const getQuote = `-- name: GetQuote :one
SELECT quote(1)
`

func (q *Queries) GetQuote(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getQuote)
	var quote string
	err := row.Scan(&quote)
	return quote, err
}

const getRandom = `-- name: GetRandom :one
SELECT random()
`

func (q *Queries) GetRandom(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getRandom)
	var random interface{}
	err := row.Scan(&random)
	return random, err
}

const getRandomBlob = `-- name: GetRandomBlob :one
SELECT randomblob(16)
`

func (q *Queries) GetRandomBlob(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getRandomBlob)
	var randomblob interface{}
	err := row.Scan(&randomblob)
	return randomblob, err
}

const getReplace = `-- name: GetReplace :one
SELECT replace('abc', 'bc', 'df')
`

func (q *Queries) GetReplace(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getReplace)
	var replace string
	err := row.Scan(&replace)
	return replace, err
}

const getRound = `-- name: GetRound :one
SELECT round(1.1)
`

func (q *Queries) GetRound(ctx context.Context) (float64, error) {
	row := q.db.QueryRowContext(ctx, getRound)
	var round float64
	err := row.Scan(&round)
	return round, err
}

const getRound2 = `-- name: GetRound2 :one
SELECT round(1.1, 2)
`

func (q *Queries) GetRound2(ctx context.Context) (float64, error) {
	row := q.db.QueryRowContext(ctx, getRound2)
	var round float64
	err := row.Scan(&round)
	return round, err
}

const getRtrim = `-- name: GetRtrim :one
SELECT rtrim('ABCDE ')
`

func (q *Queries) GetRtrim(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getRtrim)
	var rtrim string
	err := row.Scan(&rtrim)
	return rtrim, err
}

const getRtrim2 = `-- name: GetRtrim2 :one
SELECT rtrim('ABCDE:', ':')
`

func (q *Queries) GetRtrim2(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getRtrim2)
	var rtrim string
	err := row.Scan(&rtrim)
	return rtrim, err
}

const getSQLiteCompileOptionGet = `-- name: GetSQLiteCompileOptionGet :one
SELECT sqlite_compileoption_get(1)
`

func (q *Queries) GetSQLiteCompileOptionGet(ctx context.Context) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getSQLiteCompileOptionGet)
	var sqlite_compileoption_get sql.NullString
	err := row.Scan(&sqlite_compileoption_get)
	return sqlite_compileoption_get, err
}

const getSQLiteCompileOptionUsed = `-- name: GetSQLiteCompileOptionUsed :one
SELECT sqlite_compileoption_used(1)
`

func (q *Queries) GetSQLiteCompileOptionUsed(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getSQLiteCompileOptionUsed)
	var sqlite_compileoption_used int64
	err := row.Scan(&sqlite_compileoption_used)
	return sqlite_compileoption_used, err
}

const getSQLiteOffset = `-- name: GetSQLiteOffset :one
SELECT sqlite_offset(1)
`

func (q *Queries) GetSQLiteOffset(ctx context.Context) (sql.NullInt64, error) {
	row := q.db.QueryRowContext(ctx, getSQLiteOffset)
	var sqlite_offset sql.NullInt64
	err := row.Scan(&sqlite_offset)
	return sqlite_offset, err
}

const getSQLiteSourceID = `-- name: GetSQLiteSourceID :one
SELECT sqlite_source_id()
`

func (q *Queries) GetSQLiteSourceID(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getSQLiteSourceID)
	var sqlite_source_id string
	err := row.Scan(&sqlite_source_id)
	return sqlite_source_id, err
}

const getSQLiteVersion = `-- name: GetSQLiteVersion :one
SELECT sqlite_version()
`

func (q *Queries) GetSQLiteVersion(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getSQLiteVersion)
	var sqlite_version string
	err := row.Scan(&sqlite_version)
	return sqlite_version, err
}

const getSign = `-- name: GetSign :one
SELECT sign(1)
`

func (q *Queries) GetSign(ctx context.Context) (sql.NullInt64, error) {
	row := q.db.QueryRowContext(ctx, getSign)
	var sign sql.NullInt64
	err := row.Scan(&sign)
	return sign, err
}

const getSoundex = `-- name: GetSoundex :one
SELECT soundex('abc')
`

func (q *Queries) GetSoundex(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getSoundex)
	var soundex string
	err := row.Scan(&soundex)
	return soundex, err
}

const getSubstr2 = `-- name: GetSubstr2 :one
SELECT substr('abcdef', 2)
`

func (q *Queries) GetSubstr2(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getSubstr2)
	var substr string
	err := row.Scan(&substr)
	return substr, err
}

const getSubstr3 = `-- name: GetSubstr3 :one
SELECT substr('abcdef', 1, 2)
`

func (q *Queries) GetSubstr3(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getSubstr3)
	var substr string
	err := row.Scan(&substr)
	return substr, err
}

const getSubstring2 = `-- name: GetSubstring2 :one
SELECT substring('abcdef', 1)
`

func (q *Queries) GetSubstring2(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getSubstring2)
	var substring string
	err := row.Scan(&substring)
	return substring, err
}

const getSusbstring3 = `-- name: GetSusbstring3 :one
SELECT substring('abcdef', 1, 2)
`

func (q *Queries) GetSusbstring3(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getSusbstring3)
	var substring string
	err := row.Scan(&substring)
	return substring, err
}

const getTotalChanges = `-- name: GetTotalChanges :one
SELECT total_changes()
`

func (q *Queries) GetTotalChanges(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalChanges)
	var total_changes int64
	err := row.Scan(&total_changes)
	return total_changes, err
}

const getTrim = `-- name: GetTrim :one
SELECT trim(' ABCDE ')
`

func (q *Queries) GetTrim(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getTrim)
	var trim string
	err := row.Scan(&trim)
	return trim, err
}

const getTrim2 = `-- name: GetTrim2 :one
SELECT trim(':ABCDE:', ':')
`

func (q *Queries) GetTrim2(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getTrim2)
	var trim string
	err := row.Scan(&trim)
	return trim, err
}

const getTypeof = `-- name: GetTypeof :one
SELECT typeof('ABCDE')
`

func (q *Queries) GetTypeof(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getTypeof)
	var typeof string
	err := row.Scan(&typeof)
	return typeof, err
}

const getUnicode = `-- name: GetUnicode :one
SELECT unicode('A')
`

func (q *Queries) GetUnicode(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUnicode)
	var unicode int64
	err := row.Scan(&unicode)
	return unicode, err
}

const getUnlikely = `-- name: GetUnlikely :one
SELECT unlikely('12345')
`

func (q *Queries) GetUnlikely(ctx context.Context) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getUnlikely)
	var unlikely interface{}
	err := row.Scan(&unlikely)
	return unlikely, err
}

const getUpper = `-- name: GetUpper :one
SELECT upper('abcde')
`

func (q *Queries) GetUpper(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, getUpper)
	var upper string
	err := row.Scan(&upper)
	return upper, err
}

const getZeroblob = `-- name: GetZeroblob :one
SELECT zeroblob(16)
`

func (q *Queries) GetZeroblob(ctx context.Context) ([]uint8, error) {
	row := q.db.QueryRowContext(ctx, getZeroblob)
	var zeroblob []uint8
	err := row.Scan(pq.Array(&zeroblob))
	return zeroblob, err
}
