// Code generated from SQLiteParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package sqlgrammar // SQLiteParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SQLiteParser struct {
	*antlr.BaseParser
}

var sqliteparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqliteparserParserInit() {
	staticData := &sqliteparserParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'.'", "'('", "')'", "','", "'='", "'*'", "'+'", "'-'", "'~'",
		"'||'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'|'", "'<'", "'<='", "'>'",
		"'>='", "'=='", "'!='", "'<>'", "'::'", "'ABORT'", "'ADD'", "'ALL'",
		"'AND'", "'ASC'", "'AS'", "'BETWEEN'", "'BY'", "'CASE'", "'COLLATE'",
		"'COMMIT'", "'CONFLICT'", "'CREATE'", "'CROSS'", "'DEFAULT'", "'DELETE'",
		"'DESC'", "'DISTINCT'", "'DO'", "'ELSE'", "'END'", "'ESCAPE'", "'EXCEPT'",
		"'EXISTS'", "'FAIL'", "'FALSE'", "'FILTER'", "'FIRST'", "'FROM'", "'FULL'",
		"'GLOB'", "'GROUPS'", "'GROUP'", "'HAVING'", "'IGNORE'", "'INDEXED'",
		"'INNER'", "'INSERT'", "'INTERSECT'", "'INTO'", "'IN'", "'ISNULL'",
		"'IS'", "'JOIN'", "'LAST'", "'LEFT'", "'LIKE'", "'LIMIT'", "'MATCH'",
		"'NATURAL'", "'NOTHING'", "'NOTNULL'", "'NOT'", "'NULLS'", "'NULL'",
		"'OFFSET'", "'OF'", "'ON'", "'ORDER'", "'OR'", "'OUTER'", "'RAISE'",
		"'REGEXP'", "'REPLACE'", "'RETURNING'", "'RIGHT'", "'ROLLBACK'", "'SELECT'",
		"'SET'", "'THEN'", "'TRUE'", "'UNION'", "'UPDATE'", "'USING'", "'VALUES'",
		"'WHEN'", "'WHERE'", "'WITH'",
	}
	staticData.symbolicNames = []string{
		"", "SCOL", "DOT", "OPEN_PAR", "CLOSE_PAR", "COMMA", "ASSIGN", "STAR",
		"PLUS", "MINUS", "TILDE", "PIPE2", "DIV", "MOD", "LT2", "GT2", "AMP",
		"PIPE", "LT", "LT_EQ", "GT", "GT_EQ", "EQ", "NOT_EQ1", "NOT_EQ2", "TYPE_CAST",
		"ABORT_", "ADD_", "ALL_", "AND_", "ASC_", "AS_", "BETWEEN_", "BY_",
		"CASE_", "COLLATE_", "COMMIT_", "CONFLICT_", "CREATE_", "CROSS_", "DEFAULT_",
		"DELETE_", "DESC_", "DISTINCT_", "DO_", "ELSE_", "END_", "ESCAPE_",
		"EXCEPT_", "EXISTS_", "FAIL_", "FALSE_", "FILTER_", "FIRST_", "FROM_",
		"FULL_", "GLOB_", "GROUPS_", "GROUP_", "HAVING_", "IGNORE_", "INDEXED_",
		"INNER_", "INSERT_", "INTERSECT_", "INTO_", "IN_", "ISNULL_", "IS_",
		"JOIN_", "LAST_", "LEFT_", "LIKE_", "LIMIT_", "MATCH_", "NATURAL_",
		"NOTHING_", "NOTNULL_", "NOT_", "NULLS_", "NULL_", "OFFSET_", "OF_",
		"ON_", "ORDER_", "OR_", "OUTER_", "RAISE_", "REGEXP_", "REPLACE_", "RETURNING_",
		"RIGHT_", "ROLLBACK_", "SELECT_", "SET_", "THEN_", "TRUE_", "UNION_",
		"UPDATE_", "USING_", "VALUES_", "WHEN_", "WHERE_", "WITH_", "IDENTIFIER",
		"NUMERIC_LITERAL", "BIND_PARAMETER", "STRING_LITERAL", "BLOB_LITERAL",
		"SINGLE_LINE_COMMENT", "MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	}
	staticData.ruleNames = []string{
		"parse", "sql_stmt_list", "sql_stmt", "indexed_column", "cte_table_name",
		"common_table_expression", "common_table_stmt", "delete_stmt", "expr",
		"cast_type", "type_cast", "literal_value", "value_row", "values_clause",
		"insert_stmt", "returning_clause", "upsert_update", "upsert_clause",
		"select_stmt_core", "select_stmt", "join_clause", "select_core", "table_or_subquery",
		"result_column", "returning_clause_result_column", "join_operator",
		"join_constraint", "compound_operator", "update_set_subclause", "update_stmt",
		"column_name_list", "qualified_table_name", "order_by_stmt", "limit_stmt",
		"ordering_term", "asc_desc", "function_keyword", "function_name", "table_name",
		"table_alias", "column_name", "column_alias", "collation_name", "index_name",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 112, 689, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 1, 0, 5, 0, 90, 8, 0, 10, 0, 12, 0, 93, 9, 0,
		1, 0, 1, 0, 1, 1, 5, 1, 98, 8, 1, 10, 1, 12, 1, 101, 9, 1, 1, 1, 1, 1,
		4, 1, 105, 8, 1, 11, 1, 12, 1, 106, 1, 1, 5, 1, 110, 8, 1, 10, 1, 12, 1,
		113, 9, 1, 1, 1, 5, 1, 116, 8, 1, 10, 1, 12, 1, 119, 9, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 125, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 134, 8, 4, 10, 4, 12, 4, 137, 9, 4, 1, 4, 1, 4, 3, 4, 141, 8, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 153,
		8, 6, 10, 6, 12, 6, 156, 9, 6, 1, 7, 3, 7, 159, 8, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 166, 8, 7, 1, 7, 3, 7, 169, 8, 7, 1, 8, 1, 8, 1, 8, 3,
		8, 174, 8, 8, 1, 8, 1, 8, 3, 8, 178, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 183,
		8, 8, 1, 8, 1, 8, 3, 8, 187, 8, 8, 1, 8, 3, 8, 190, 8, 8, 1, 8, 3, 8, 193,
		8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 203, 8, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 213, 8, 8, 10, 8,
		12, 8, 216, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 223, 8, 8, 1, 8,
		1, 8, 1, 8, 5, 8, 228, 8, 8, 10, 8, 12, 8, 231, 9, 8, 1, 8, 3, 8, 234,
		8, 8, 1, 8, 1, 8, 3, 8, 238, 8, 8, 1, 8, 1, 8, 3, 8, 242, 8, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 4, 8, 249, 8, 8, 11, 8, 12, 8, 250, 1, 8, 1, 8, 3,
		8, 255, 8, 8, 1, 8, 1, 8, 3, 8, 259, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 283, 8, 8, 1, 8, 1, 8, 3, 8, 287, 8,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 292, 8, 8, 1, 8, 3, 8, 295, 8, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 300, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 318, 8, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 324, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		331, 8, 8, 5, 8, 333, 8, 8, 10, 8, 12, 8, 336, 9, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 349, 8,
		12, 10, 12, 12, 12, 352, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		5, 13, 360, 8, 13, 10, 13, 12, 13, 363, 9, 13, 1, 14, 3, 14, 366, 8, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 373, 8, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 379, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 385, 8,
		14, 10, 14, 12, 14, 388, 9, 14, 1, 14, 1, 14, 3, 14, 392, 8, 14, 1, 14,
		1, 14, 3, 14, 396, 8, 14, 1, 14, 3, 14, 399, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 5, 15, 405, 8, 15, 10, 15, 12, 15, 408, 9, 15, 1, 16, 1, 16, 3,
		16, 412, 8, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 5, 17, 423, 8, 17, 10, 17, 12, 17, 426, 9, 17, 1, 17, 1, 17, 1,
		17, 3, 17, 431, 8, 17, 3, 17, 433, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 5, 17, 442, 8, 17, 10, 17, 12, 17, 445, 9, 17, 1, 17,
		1, 17, 3, 17, 449, 8, 17, 3, 17, 451, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		5, 18, 457, 8, 18, 10, 18, 12, 18, 460, 9, 18, 1, 18, 3, 18, 463, 8, 18,
		1, 18, 3, 18, 466, 8, 18, 1, 19, 3, 19, 469, 8, 19, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 478, 8, 20, 10, 20, 12, 20, 481, 9,
		20, 1, 21, 1, 21, 3, 21, 485, 8, 21, 1, 21, 1, 21, 1, 21, 5, 21, 490, 8,
		21, 10, 21, 12, 21, 493, 9, 21, 1, 21, 1, 21, 1, 21, 3, 21, 498, 8, 21,
		3, 21, 500, 8, 21, 1, 21, 1, 21, 3, 21, 504, 8, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 5, 21, 511, 8, 21, 10, 21, 12, 21, 514, 9, 21, 1, 21, 1,
		21, 3, 21, 518, 8, 21, 3, 21, 520, 8, 21, 1, 22, 1, 22, 1, 22, 3, 22, 525,
		8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 532, 8, 22, 3, 22, 534,
		8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 544,
		8, 23, 3, 23, 546, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 552, 8, 24,
		3, 24, 554, 8, 24, 1, 25, 3, 25, 557, 8, 25, 1, 25, 1, 25, 3, 25, 561,
		8, 25, 1, 25, 3, 25, 564, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 3, 27, 573, 8, 27, 1, 27, 1, 27, 3, 27, 577, 8, 27, 1, 28, 1,
		28, 3, 28, 581, 8, 28, 1, 28, 1, 28, 1, 28, 1, 29, 3, 29, 587, 8, 29, 1,
		29, 1, 29, 1, 29, 3, 29, 592, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		5, 29, 599, 8, 29, 10, 29, 12, 29, 602, 9, 29, 1, 29, 1, 29, 1, 29, 3,
		29, 607, 8, 29, 3, 29, 609, 8, 29, 1, 29, 1, 29, 3, 29, 613, 8, 29, 1,
		29, 3, 29, 616, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 622, 8, 30, 10,
		30, 12, 30, 625, 9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 3, 31, 632,
		8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 639, 8, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 5, 32, 646, 8, 32, 10, 32, 12, 32, 649, 9, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 655, 8, 33, 1, 34, 1, 34, 1, 34, 3,
		34, 660, 8, 34, 1, 34, 3, 34, 663, 8, 34, 1, 34, 1, 34, 3, 34, 667, 8,
		34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 3, 37, 675, 8, 37, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 43, 0, 1, 16, 44, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 0, 13, 1, 0, 8, 10,
		2, 0, 7, 7, 12, 13, 1, 0, 8, 9, 1, 0, 14, 17, 1, 0, 18, 21, 4, 0, 56, 56,
		66, 66, 74, 74, 88, 88, 5, 0, 51, 51, 80, 80, 96, 96, 105, 105, 107, 107,
		3, 0, 55, 55, 71, 71, 91, 91, 5, 0, 26, 26, 50, 50, 60, 60, 89, 89, 92,
		92, 2, 0, 5, 5, 81, 81, 2, 0, 53, 53, 70, 70, 2, 0, 30, 30, 42, 42, 3,
		0, 56, 56, 72, 72, 89, 89, 773, 0, 91, 1, 0, 0, 0, 2, 99, 1, 0, 0, 0, 4,
		124, 1, 0, 0, 0, 6, 126, 1, 0, 0, 0, 8, 128, 1, 0, 0, 0, 10, 142, 1, 0,
		0, 0, 12, 148, 1, 0, 0, 0, 14, 158, 1, 0, 0, 0, 16, 258, 1, 0, 0, 0, 18,
		337, 1, 0, 0, 0, 20, 339, 1, 0, 0, 0, 22, 342, 1, 0, 0, 0, 24, 344, 1,
		0, 0, 0, 26, 355, 1, 0, 0, 0, 28, 365, 1, 0, 0, 0, 30, 400, 1, 0, 0, 0,
		32, 411, 1, 0, 0, 0, 34, 416, 1, 0, 0, 0, 36, 452, 1, 0, 0, 0, 38, 468,
		1, 0, 0, 0, 40, 472, 1, 0, 0, 0, 42, 482, 1, 0, 0, 0, 44, 533, 1, 0, 0,
		0, 46, 545, 1, 0, 0, 0, 48, 553, 1, 0, 0, 0, 50, 556, 1, 0, 0, 0, 52, 567,
		1, 0, 0, 0, 54, 576, 1, 0, 0, 0, 56, 580, 1, 0, 0, 0, 58, 586, 1, 0, 0,
		0, 60, 617, 1, 0, 0, 0, 62, 628, 1, 0, 0, 0, 64, 640, 1, 0, 0, 0, 66, 650,
		1, 0, 0, 0, 68, 656, 1, 0, 0, 0, 70, 668, 1, 0, 0, 0, 72, 670, 1, 0, 0,
		0, 74, 674, 1, 0, 0, 0, 76, 676, 1, 0, 0, 0, 78, 678, 1, 0, 0, 0, 80, 680,
		1, 0, 0, 0, 82, 682, 1, 0, 0, 0, 84, 684, 1, 0, 0, 0, 86, 686, 1, 0, 0,
		0, 88, 90, 3, 2, 1, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89,
		1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0,
		94, 95, 5, 0, 0, 1, 95, 1, 1, 0, 0, 0, 96, 98, 5, 1, 0, 0, 97, 96, 1, 0,
		0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100,
		102, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 111, 3, 4, 2, 0, 103, 105, 5,
		1, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 104, 1, 0, 0,
		0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 3, 4, 2, 0, 109,
		104, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 117, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 116, 5, 1,
		0, 0, 115, 114, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0,
		117, 118, 1, 0, 0, 0, 118, 3, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 125,
		3, 14, 7, 0, 121, 125, 3, 28, 14, 0, 122, 125, 3, 38, 19, 0, 123, 125,
		3, 58, 29, 0, 124, 120, 1, 0, 0, 0, 124, 121, 1, 0, 0, 0, 124, 122, 1,
		0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 5, 1, 0, 0, 0, 126, 127, 3, 80, 40,
		0, 127, 7, 1, 0, 0, 0, 128, 140, 3, 76, 38, 0, 129, 130, 5, 3, 0, 0, 130,
		135, 3, 80, 40, 0, 131, 132, 5, 5, 0, 0, 132, 134, 3, 80, 40, 0, 133, 131,
		1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0,
		0, 0, 136, 138, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 5, 4, 0, 0,
		139, 141, 1, 0, 0, 0, 140, 129, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		9, 1, 0, 0, 0, 142, 143, 3, 8, 4, 0, 143, 144, 5, 31, 0, 0, 144, 145, 5,
		3, 0, 0, 145, 146, 3, 36, 18, 0, 146, 147, 5, 4, 0, 0, 147, 11, 1, 0, 0,
		0, 148, 149, 5, 103, 0, 0, 149, 154, 3, 10, 5, 0, 150, 151, 5, 5, 0, 0,
		151, 153, 3, 10, 5, 0, 152, 150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 13, 1, 0, 0, 0, 156, 154, 1,
		0, 0, 0, 157, 159, 3, 12, 6, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0,
		0, 159, 160, 1, 0, 0, 0, 160, 161, 5, 41, 0, 0, 161, 162, 5, 54, 0, 0,
		162, 165, 3, 62, 31, 0, 163, 164, 5, 102, 0, 0, 164, 166, 3, 16, 8, 0,
		165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167,
		169, 3, 30, 15, 0, 168, 167, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 15,
		1, 0, 0, 0, 170, 171, 6, 8, -1, 0, 171, 173, 3, 22, 11, 0, 172, 174, 3,
		20, 10, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 259, 1, 0,
		0, 0, 175, 177, 5, 106, 0, 0, 176, 178, 3, 20, 10, 0, 177, 176, 1, 0, 0,
		0, 177, 178, 1, 0, 0, 0, 178, 259, 1, 0, 0, 0, 179, 180, 3, 76, 38, 0,
		180, 181, 5, 2, 0, 0, 181, 183, 1, 0, 0, 0, 182, 179, 1, 0, 0, 0, 182,
		183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 186, 3, 80, 40, 0, 185, 187,
		3, 20, 10, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 259, 1,
		0, 0, 0, 188, 190, 5, 78, 0, 0, 189, 188, 1, 0, 0, 0, 189, 190, 1, 0, 0,
		0, 190, 191, 1, 0, 0, 0, 191, 193, 5, 49, 0, 0, 192, 189, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 5, 3, 0, 0, 195, 196,
		3, 36, 18, 0, 196, 197, 5, 4, 0, 0, 197, 259, 1, 0, 0, 0, 198, 199, 5,
		3, 0, 0, 199, 200, 3, 16, 8, 0, 200, 202, 5, 4, 0, 0, 201, 203, 3, 20,
		10, 0, 202, 201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 259, 1, 0, 0, 0,
		204, 205, 7, 0, 0, 0, 205, 259, 3, 16, 8, 17, 206, 207, 5, 78, 0, 0, 207,
		259, 3, 16, 8, 6, 208, 209, 5, 3, 0, 0, 209, 214, 3, 16, 8, 0, 210, 211,
		5, 5, 0, 0, 211, 213, 3, 16, 8, 0, 212, 210, 1, 0, 0, 0, 213, 216, 1, 0,
		0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 1, 0, 0, 0,
		216, 214, 1, 0, 0, 0, 217, 218, 5, 4, 0, 0, 218, 259, 1, 0, 0, 0, 219,
		220, 3, 74, 37, 0, 220, 233, 5, 3, 0, 0, 221, 223, 5, 43, 0, 0, 222, 221,
		1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 229, 3, 16,
		8, 0, 225, 226, 5, 5, 0, 0, 226, 228, 3, 16, 8, 0, 227, 225, 1, 0, 0, 0,
		228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230,
		234, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 234, 5, 7, 0, 0, 233, 222,
		1, 0, 0, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0,
		0, 0, 235, 237, 5, 4, 0, 0, 236, 238, 3, 20, 10, 0, 237, 236, 1, 0, 0,
		0, 237, 238, 1, 0, 0, 0, 238, 259, 1, 0, 0, 0, 239, 241, 5, 34, 0, 0, 240,
		242, 3, 16, 8, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 248,
		1, 0, 0, 0, 243, 244, 5, 101, 0, 0, 244, 245, 3, 16, 8, 0, 245, 246, 5,
		95, 0, 0, 246, 247, 3, 16, 8, 0, 247, 249, 1, 0, 0, 0, 248, 243, 1, 0,
		0, 0, 249, 250, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0,
		251, 254, 1, 0, 0, 0, 252, 253, 5, 45, 0, 0, 253, 255, 3, 16, 8, 0, 254,
		252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257,
		5, 46, 0, 0, 257, 259, 1, 0, 0, 0, 258, 170, 1, 0, 0, 0, 258, 175, 1, 0,
		0, 0, 258, 182, 1, 0, 0, 0, 258, 192, 1, 0, 0, 0, 258, 198, 1, 0, 0, 0,
		258, 204, 1, 0, 0, 0, 258, 206, 1, 0, 0, 0, 258, 208, 1, 0, 0, 0, 258,
		219, 1, 0, 0, 0, 258, 239, 1, 0, 0, 0, 259, 334, 1, 0, 0, 0, 260, 261,
		10, 15, 0, 0, 261, 262, 5, 11, 0, 0, 262, 333, 3, 16, 8, 16, 263, 264,
		10, 14, 0, 0, 264, 265, 7, 1, 0, 0, 265, 333, 3, 16, 8, 15, 266, 267, 10,
		13, 0, 0, 267, 268, 7, 2, 0, 0, 268, 333, 3, 16, 8, 14, 269, 270, 10, 12,
		0, 0, 270, 271, 7, 3, 0, 0, 271, 333, 3, 16, 8, 13, 272, 273, 10, 11, 0,
		0, 273, 274, 7, 4, 0, 0, 274, 333, 3, 16, 8, 12, 275, 294, 10, 10, 0, 0,
		276, 295, 5, 6, 0, 0, 277, 295, 5, 22, 0, 0, 278, 295, 5, 23, 0, 0, 279,
		295, 5, 24, 0, 0, 280, 282, 5, 68, 0, 0, 281, 283, 5, 78, 0, 0, 282, 281,
		1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 295, 1, 0, 0, 0, 284, 286, 5, 68,
		0, 0, 285, 287, 5, 78, 0, 0, 286, 285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0,
		287, 288, 1, 0, 0, 0, 288, 289, 5, 43, 0, 0, 289, 295, 5, 54, 0, 0, 290,
		292, 5, 78, 0, 0, 291, 290, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293,
		1, 0, 0, 0, 293, 295, 7, 5, 0, 0, 294, 276, 1, 0, 0, 0, 294, 277, 1, 0,
		0, 0, 294, 278, 1, 0, 0, 0, 294, 279, 1, 0, 0, 0, 294, 280, 1, 0, 0, 0,
		294, 284, 1, 0, 0, 0, 294, 291, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296,
		333, 3, 16, 8, 11, 297, 299, 10, 8, 0, 0, 298, 300, 5, 78, 0, 0, 299, 298,
		1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302, 5, 32,
		0, 0, 302, 303, 3, 16, 8, 0, 303, 304, 5, 29, 0, 0, 304, 305, 3, 16, 8,
		9, 305, 333, 1, 0, 0, 0, 306, 307, 10, 5, 0, 0, 307, 308, 5, 29, 0, 0,
		308, 333, 3, 16, 8, 6, 309, 310, 10, 4, 0, 0, 310, 311, 5, 85, 0, 0, 311,
		333, 3, 16, 8, 5, 312, 313, 10, 16, 0, 0, 313, 314, 5, 35, 0, 0, 314, 333,
		3, 84, 42, 0, 315, 317, 10, 9, 0, 0, 316, 318, 5, 78, 0, 0, 317, 316, 1,
		0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 5, 72, 0,
		0, 320, 323, 3, 16, 8, 0, 321, 322, 5, 47, 0, 0, 322, 324, 3, 16, 8, 0,
		323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 333, 1, 0, 0, 0, 325,
		330, 10, 7, 0, 0, 326, 331, 5, 67, 0, 0, 327, 331, 5, 77, 0, 0, 328, 329,
		5, 78, 0, 0, 329, 331, 5, 80, 0, 0, 330, 326, 1, 0, 0, 0, 330, 327, 1,
		0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 333, 1, 0, 0, 0, 332, 260, 1, 0, 0,
		0, 332, 263, 1, 0, 0, 0, 332, 266, 1, 0, 0, 0, 332, 269, 1, 0, 0, 0, 332,
		272, 1, 0, 0, 0, 332, 275, 1, 0, 0, 0, 332, 297, 1, 0, 0, 0, 332, 306,
		1, 0, 0, 0, 332, 309, 1, 0, 0, 0, 332, 312, 1, 0, 0, 0, 332, 315, 1, 0,
		0, 0, 332, 325, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0,
		334, 335, 1, 0, 0, 0, 335, 17, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 337, 338,
		5, 104, 0, 0, 338, 19, 1, 0, 0, 0, 339, 340, 5, 25, 0, 0, 340, 341, 3,
		18, 9, 0, 341, 21, 1, 0, 0, 0, 342, 343, 7, 6, 0, 0, 343, 23, 1, 0, 0,
		0, 344, 345, 5, 3, 0, 0, 345, 350, 3, 16, 8, 0, 346, 347, 5, 5, 0, 0, 347,
		349, 3, 16, 8, 0, 348, 346, 1, 0, 0, 0, 349, 352, 1, 0, 0, 0, 350, 348,
		1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 353, 1, 0, 0, 0, 352, 350, 1, 0,
		0, 0, 353, 354, 5, 4, 0, 0, 354, 25, 1, 0, 0, 0, 355, 356, 5, 100, 0, 0,
		356, 361, 3, 24, 12, 0, 357, 358, 5, 5, 0, 0, 358, 360, 3, 24, 12, 0, 359,
		357, 1, 0, 0, 0, 360, 363, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 362,
		1, 0, 0, 0, 362, 27, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 364, 366, 3, 12,
		6, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 372, 1, 0, 0, 0,
		367, 373, 5, 89, 0, 0, 368, 373, 5, 63, 0, 0, 369, 370, 5, 63, 0, 0, 370,
		371, 5, 85, 0, 0, 371, 373, 5, 89, 0, 0, 372, 367, 1, 0, 0, 0, 372, 368,
		1, 0, 0, 0, 372, 369, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 375, 5, 65,
		0, 0, 375, 378, 3, 76, 38, 0, 376, 377, 5, 31, 0, 0, 377, 379, 3, 78, 39,
		0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 391, 1, 0, 0, 0, 380,
		381, 5, 3, 0, 0, 381, 386, 3, 80, 40, 0, 382, 383, 5, 5, 0, 0, 383, 385,
		3, 80, 40, 0, 384, 382, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386, 384, 1,
		0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 389, 1, 0, 0, 0, 388, 386, 1, 0, 0,
		0, 389, 390, 5, 4, 0, 0, 390, 392, 1, 0, 0, 0, 391, 380, 1, 0, 0, 0, 391,
		392, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 395, 3, 26, 13, 0, 394, 396,
		3, 34, 17, 0, 395, 394, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 398, 1,
		0, 0, 0, 397, 399, 3, 30, 15, 0, 398, 397, 1, 0, 0, 0, 398, 399, 1, 0,
		0, 0, 399, 29, 1, 0, 0, 0, 400, 401, 5, 90, 0, 0, 401, 406, 3, 48, 24,
		0, 402, 403, 5, 5, 0, 0, 403, 405, 3, 48, 24, 0, 404, 402, 1, 0, 0, 0,
		405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407,
		31, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 412, 3, 80, 40, 0, 410, 412,
		3, 60, 30, 0, 411, 409, 1, 0, 0, 0, 411, 410, 1, 0, 0, 0, 412, 413, 1,
		0, 0, 0, 413, 414, 5, 6, 0, 0, 414, 415, 3, 16, 8, 0, 415, 33, 1, 0, 0,
		0, 416, 417, 5, 83, 0, 0, 417, 432, 5, 37, 0, 0, 418, 419, 5, 3, 0, 0,
		419, 424, 3, 6, 3, 0, 420, 421, 5, 5, 0, 0, 421, 423, 3, 6, 3, 0, 422,
		420, 1, 0, 0, 0, 423, 426, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 424, 425,
		1, 0, 0, 0, 425, 427, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 427, 430, 5, 4,
		0, 0, 428, 429, 5, 102, 0, 0, 429, 431, 3, 16, 8, 0, 430, 428, 1, 0, 0,
		0, 430, 431, 1, 0, 0, 0, 431, 433, 1, 0, 0, 0, 432, 418, 1, 0, 0, 0, 432,
		433, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 450, 5, 44, 0, 0, 435, 451,
		5, 76, 0, 0, 436, 437, 5, 98, 0, 0, 437, 438, 5, 94, 0, 0, 438, 443, 3,
		32, 16, 0, 439, 440, 5, 5, 0, 0, 440, 442, 3, 32, 16, 0, 441, 439, 1, 0,
		0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0,
		444, 448, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 447, 5, 102, 0, 0, 447,
		449, 3, 16, 8, 0, 448, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 451,
		1, 0, 0, 0, 450, 435, 1, 0, 0, 0, 450, 436, 1, 0, 0, 0, 451, 35, 1, 0,
		0, 0, 452, 458, 3, 42, 21, 0, 453, 454, 3, 54, 27, 0, 454, 455, 3, 42,
		21, 0, 455, 457, 1, 0, 0, 0, 456, 453, 1, 0, 0, 0, 457, 460, 1, 0, 0, 0,
		458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 462, 1, 0, 0, 0, 460,
		458, 1, 0, 0, 0, 461, 463, 3, 64, 32, 0, 462, 461, 1, 0, 0, 0, 462, 463,
		1, 0, 0, 0, 463, 465, 1, 0, 0, 0, 464, 466, 3, 66, 33, 0, 465, 464, 1,
		0, 0, 0, 465, 466, 1, 0, 0, 0, 466, 37, 1, 0, 0, 0, 467, 469, 3, 12, 6,
		0, 468, 467, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470,
		471, 3, 36, 18, 0, 471, 39, 1, 0, 0, 0, 472, 479, 3, 44, 22, 0, 473, 474,
		3, 50, 25, 0, 474, 475, 3, 44, 22, 0, 475, 476, 3, 52, 26, 0, 476, 478,
		1, 0, 0, 0, 477, 473, 1, 0, 0, 0, 478, 481, 1, 0, 0, 0, 479, 477, 1, 0,
		0, 0, 479, 480, 1, 0, 0, 0, 480, 41, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0,
		482, 484, 5, 93, 0, 0, 483, 485, 5, 43, 0, 0, 484, 483, 1, 0, 0, 0, 484,
		485, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 491, 3, 46, 23, 0, 487, 488,
		5, 5, 0, 0, 488, 490, 3, 46, 23, 0, 489, 487, 1, 0, 0, 0, 490, 493, 1,
		0, 0, 0, 491, 489, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0, 492, 499, 1, 0, 0,
		0, 493, 491, 1, 0, 0, 0, 494, 497, 5, 54, 0, 0, 495, 498, 3, 44, 22, 0,
		496, 498, 3, 40, 20, 0, 497, 495, 1, 0, 0, 0, 497, 496, 1, 0, 0, 0, 498,
		500, 1, 0, 0, 0, 499, 494, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 503,
		1, 0, 0, 0, 501, 502, 5, 102, 0, 0, 502, 504, 3, 16, 8, 0, 503, 501, 1,
		0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 519, 1, 0, 0, 0, 505, 506, 5, 58, 0,
		0, 506, 507, 5, 33, 0, 0, 507, 512, 3, 16, 8, 0, 508, 509, 5, 5, 0, 0,
		509, 511, 3, 16, 8, 0, 510, 508, 1, 0, 0, 0, 511, 514, 1, 0, 0, 0, 512,
		510, 1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 517, 1, 0, 0, 0, 514, 512,
		1, 0, 0, 0, 515, 516, 5, 59, 0, 0, 516, 518, 3, 16, 8, 0, 517, 515, 1,
		0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 520, 1, 0, 0, 0, 519, 505, 1, 0, 0,
		0, 519, 520, 1, 0, 0, 0, 520, 43, 1, 0, 0, 0, 521, 524, 3, 76, 38, 0, 522,
		523, 5, 31, 0, 0, 523, 525, 3, 78, 39, 0, 524, 522, 1, 0, 0, 0, 524, 525,
		1, 0, 0, 0, 525, 534, 1, 0, 0, 0, 526, 527, 5, 3, 0, 0, 527, 528, 3, 36,
		18, 0, 528, 531, 5, 4, 0, 0, 529, 530, 5, 31, 0, 0, 530, 532, 3, 78, 39,
		0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 534, 1, 0, 0, 0, 533,
		521, 1, 0, 0, 0, 533, 526, 1, 0, 0, 0, 534, 45, 1, 0, 0, 0, 535, 546, 5,
		7, 0, 0, 536, 537, 3, 76, 38, 0, 537, 538, 5, 2, 0, 0, 538, 539, 5, 7,
		0, 0, 539, 546, 1, 0, 0, 0, 540, 543, 3, 16, 8, 0, 541, 542, 5, 31, 0,
		0, 542, 544, 3, 82, 41, 0, 543, 541, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0,
		544, 546, 1, 0, 0, 0, 545, 535, 1, 0, 0, 0, 545, 536, 1, 0, 0, 0, 545,
		540, 1, 0, 0, 0, 546, 47, 1, 0, 0, 0, 547, 554, 5, 7, 0, 0, 548, 551, 3,
		16, 8, 0, 549, 550, 5, 31, 0, 0, 550, 552, 3, 82, 41, 0, 551, 549, 1, 0,
		0, 0, 551, 552, 1, 0, 0, 0, 552, 554, 1, 0, 0, 0, 553, 547, 1, 0, 0, 0,
		553, 548, 1, 0, 0, 0, 554, 49, 1, 0, 0, 0, 555, 557, 5, 75, 0, 0, 556,
		555, 1, 0, 0, 0, 556, 557, 1, 0, 0, 0, 557, 563, 1, 0, 0, 0, 558, 560,
		7, 7, 0, 0, 559, 561, 5, 86, 0, 0, 560, 559, 1, 0, 0, 0, 560, 561, 1, 0,
		0, 0, 561, 564, 1, 0, 0, 0, 562, 564, 5, 62, 0, 0, 563, 558, 1, 0, 0, 0,
		563, 562, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564, 565, 1, 0, 0, 0, 565,
		566, 5, 69, 0, 0, 566, 51, 1, 0, 0, 0, 567, 568, 5, 83, 0, 0, 568, 569,
		3, 16, 8, 0, 569, 53, 1, 0, 0, 0, 570, 572, 5, 97, 0, 0, 571, 573, 5, 28,
		0, 0, 572, 571, 1, 0, 0, 0, 572, 573, 1, 0, 0, 0, 573, 577, 1, 0, 0, 0,
		574, 577, 5, 64, 0, 0, 575, 577, 5, 48, 0, 0, 576, 570, 1, 0, 0, 0, 576,
		574, 1, 0, 0, 0, 576, 575, 1, 0, 0, 0, 577, 55, 1, 0, 0, 0, 578, 581, 3,
		80, 40, 0, 579, 581, 3, 60, 30, 0, 580, 578, 1, 0, 0, 0, 580, 579, 1, 0,
		0, 0, 581, 582, 1, 0, 0, 0, 582, 583, 5, 6, 0, 0, 583, 584, 3, 16, 8, 0,
		584, 57, 1, 0, 0, 0, 585, 587, 3, 12, 6, 0, 586, 585, 1, 0, 0, 0, 586,
		587, 1, 0, 0, 0, 587, 588, 1, 0, 0, 0, 588, 591, 5, 98, 0, 0, 589, 590,
		5, 85, 0, 0, 590, 592, 7, 8, 0, 0, 591, 589, 1, 0, 0, 0, 591, 592, 1, 0,
		0, 0, 592, 593, 1, 0, 0, 0, 593, 594, 3, 62, 31, 0, 594, 595, 5, 94, 0,
		0, 595, 600, 3, 56, 28, 0, 596, 597, 5, 5, 0, 0, 597, 599, 3, 56, 28, 0,
		598, 596, 1, 0, 0, 0, 599, 602, 1, 0, 0, 0, 600, 598, 1, 0, 0, 0, 600,
		601, 1, 0, 0, 0, 601, 608, 1, 0, 0, 0, 602, 600, 1, 0, 0, 0, 603, 606,
		5, 54, 0, 0, 604, 607, 3, 44, 22, 0, 605, 607, 3, 40, 20, 0, 606, 604,
		1, 0, 0, 0, 606, 605, 1, 0, 0, 0, 607, 609, 1, 0, 0, 0, 608, 603, 1, 0,
		0, 0, 608, 609, 1, 0, 0, 0, 609, 612, 1, 0, 0, 0, 610, 611, 5, 102, 0,
		0, 611, 613, 3, 16, 8, 0, 612, 610, 1, 0, 0, 0, 612, 613, 1, 0, 0, 0, 613,
		615, 1, 0, 0, 0, 614, 616, 3, 30, 15, 0, 615, 614, 1, 0, 0, 0, 615, 616,
		1, 0, 0, 0, 616, 59, 1, 0, 0, 0, 617, 618, 5, 3, 0, 0, 618, 623, 3, 80,
		40, 0, 619, 620, 5, 5, 0, 0, 620, 622, 3, 80, 40, 0, 621, 619, 1, 0, 0,
		0, 622, 625, 1, 0, 0, 0, 623, 621, 1, 0, 0, 0, 623, 624, 1, 0, 0, 0, 624,
		626, 1, 0, 0, 0, 625, 623, 1, 0, 0, 0, 626, 627, 5, 4, 0, 0, 627, 61, 1,
		0, 0, 0, 628, 631, 3, 76, 38, 0, 629, 630, 5, 31, 0, 0, 630, 632, 3, 78,
		39, 0, 631, 629, 1, 0, 0, 0, 631, 632, 1, 0, 0, 0, 632, 638, 1, 0, 0, 0,
		633, 634, 5, 61, 0, 0, 634, 635, 5, 33, 0, 0, 635, 639, 3, 86, 43, 0, 636,
		637, 5, 78, 0, 0, 637, 639, 5, 61, 0, 0, 638, 633, 1, 0, 0, 0, 638, 636,
		1, 0, 0, 0, 638, 639, 1, 0, 0, 0, 639, 63, 1, 0, 0, 0, 640, 641, 5, 84,
		0, 0, 641, 642, 5, 33, 0, 0, 642, 647, 3, 68, 34, 0, 643, 644, 5, 5, 0,
		0, 644, 646, 3, 68, 34, 0, 645, 643, 1, 0, 0, 0, 646, 649, 1, 0, 0, 0,
		647, 645, 1, 0, 0, 0, 647, 648, 1, 0, 0, 0, 648, 65, 1, 0, 0, 0, 649, 647,
		1, 0, 0, 0, 650, 651, 5, 73, 0, 0, 651, 654, 3, 16, 8, 0, 652, 653, 7,
		9, 0, 0, 653, 655, 3, 16, 8, 0, 654, 652, 1, 0, 0, 0, 654, 655, 1, 0, 0,
		0, 655, 67, 1, 0, 0, 0, 656, 659, 3, 16, 8, 0, 657, 658, 5, 35, 0, 0, 658,
		660, 3, 84, 42, 0, 659, 657, 1, 0, 0, 0, 659, 660, 1, 0, 0, 0, 660, 662,
		1, 0, 0, 0, 661, 663, 3, 70, 35, 0, 662, 661, 1, 0, 0, 0, 662, 663, 1,
		0, 0, 0, 663, 666, 1, 0, 0, 0, 664, 665, 5, 79, 0, 0, 665, 667, 7, 10,
		0, 0, 666, 664, 1, 0, 0, 0, 666, 667, 1, 0, 0, 0, 667, 69, 1, 0, 0, 0,
		668, 669, 7, 11, 0, 0, 669, 71, 1, 0, 0, 0, 670, 671, 7, 12, 0, 0, 671,
		73, 1, 0, 0, 0, 672, 675, 5, 104, 0, 0, 673, 675, 3, 72, 36, 0, 674, 672,
		1, 0, 0, 0, 674, 673, 1, 0, 0, 0, 675, 75, 1, 0, 0, 0, 676, 677, 5, 104,
		0, 0, 677, 77, 1, 0, 0, 0, 678, 679, 5, 104, 0, 0, 679, 79, 1, 0, 0, 0,
		680, 681, 5, 104, 0, 0, 681, 81, 1, 0, 0, 0, 682, 683, 5, 104, 0, 0, 683,
		83, 1, 0, 0, 0, 684, 685, 5, 104, 0, 0, 685, 85, 1, 0, 0, 0, 686, 687,
		5, 104, 0, 0, 687, 87, 1, 0, 0, 0, 97, 91, 99, 106, 111, 117, 124, 135,
		140, 154, 158, 165, 168, 173, 177, 182, 186, 189, 192, 202, 214, 222, 229,
		233, 237, 241, 250, 254, 258, 282, 286, 291, 294, 299, 317, 323, 330, 332,
		334, 350, 361, 365, 372, 378, 386, 391, 395, 398, 406, 411, 424, 430, 432,
		443, 448, 450, 458, 462, 465, 468, 479, 484, 491, 497, 499, 503, 512, 517,
		519, 524, 531, 533, 543, 545, 551, 553, 556, 560, 563, 572, 576, 580, 586,
		591, 600, 606, 608, 612, 615, 623, 631, 638, 647, 654, 659, 662, 666, 674,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SQLiteParserInit initializes any static state used to implement SQLiteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSQLiteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SQLiteParserInit() {
	staticData := &sqliteparserParserStaticData
	staticData.once.Do(sqliteparserParserInit)
}

// NewSQLiteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSQLiteParser(input antlr.TokenStream) *SQLiteParser {
	SQLiteParserInit()
	this := new(SQLiteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sqliteparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SQLiteParser.g4"

	return this
}

// SQLiteParser tokens.
const (
	SQLiteParserEOF                 = antlr.TokenEOF
	SQLiteParserSCOL                = 1
	SQLiteParserDOT                 = 2
	SQLiteParserOPEN_PAR            = 3
	SQLiteParserCLOSE_PAR           = 4
	SQLiteParserCOMMA               = 5
	SQLiteParserASSIGN              = 6
	SQLiteParserSTAR                = 7
	SQLiteParserPLUS                = 8
	SQLiteParserMINUS               = 9
	SQLiteParserTILDE               = 10
	SQLiteParserPIPE2               = 11
	SQLiteParserDIV                 = 12
	SQLiteParserMOD                 = 13
	SQLiteParserLT2                 = 14
	SQLiteParserGT2                 = 15
	SQLiteParserAMP                 = 16
	SQLiteParserPIPE                = 17
	SQLiteParserLT                  = 18
	SQLiteParserLT_EQ               = 19
	SQLiteParserGT                  = 20
	SQLiteParserGT_EQ               = 21
	SQLiteParserEQ                  = 22
	SQLiteParserNOT_EQ1             = 23
	SQLiteParserNOT_EQ2             = 24
	SQLiteParserTYPE_CAST           = 25
	SQLiteParserABORT_              = 26
	SQLiteParserADD_                = 27
	SQLiteParserALL_                = 28
	SQLiteParserAND_                = 29
	SQLiteParserASC_                = 30
	SQLiteParserAS_                 = 31
	SQLiteParserBETWEEN_            = 32
	SQLiteParserBY_                 = 33
	SQLiteParserCASE_               = 34
	SQLiteParserCOLLATE_            = 35
	SQLiteParserCOMMIT_             = 36
	SQLiteParserCONFLICT_           = 37
	SQLiteParserCREATE_             = 38
	SQLiteParserCROSS_              = 39
	SQLiteParserDEFAULT_            = 40
	SQLiteParserDELETE_             = 41
	SQLiteParserDESC_               = 42
	SQLiteParserDISTINCT_           = 43
	SQLiteParserDO_                 = 44
	SQLiteParserELSE_               = 45
	SQLiteParserEND_                = 46
	SQLiteParserESCAPE_             = 47
	SQLiteParserEXCEPT_             = 48
	SQLiteParserEXISTS_             = 49
	SQLiteParserFAIL_               = 50
	SQLiteParserFALSE_              = 51
	SQLiteParserFILTER_             = 52
	SQLiteParserFIRST_              = 53
	SQLiteParserFROM_               = 54
	SQLiteParserFULL_               = 55
	SQLiteParserGLOB_               = 56
	SQLiteParserGROUPS_             = 57
	SQLiteParserGROUP_              = 58
	SQLiteParserHAVING_             = 59
	SQLiteParserIGNORE_             = 60
	SQLiteParserINDEXED_            = 61
	SQLiteParserINNER_              = 62
	SQLiteParserINSERT_             = 63
	SQLiteParserINTERSECT_          = 64
	SQLiteParserINTO_               = 65
	SQLiteParserIN_                 = 66
	SQLiteParserISNULL_             = 67
	SQLiteParserIS_                 = 68
	SQLiteParserJOIN_               = 69
	SQLiteParserLAST_               = 70
	SQLiteParserLEFT_               = 71
	SQLiteParserLIKE_               = 72
	SQLiteParserLIMIT_              = 73
	SQLiteParserMATCH_              = 74
	SQLiteParserNATURAL_            = 75
	SQLiteParserNOTHING_            = 76
	SQLiteParserNOTNULL_            = 77
	SQLiteParserNOT_                = 78
	SQLiteParserNULLS_              = 79
	SQLiteParserNULL_               = 80
	SQLiteParserOFFSET_             = 81
	SQLiteParserOF_                 = 82
	SQLiteParserON_                 = 83
	SQLiteParserORDER_              = 84
	SQLiteParserOR_                 = 85
	SQLiteParserOUTER_              = 86
	SQLiteParserRAISE_              = 87
	SQLiteParserREGEXP_             = 88
	SQLiteParserREPLACE_            = 89
	SQLiteParserRETURNING_          = 90
	SQLiteParserRIGHT_              = 91
	SQLiteParserROLLBACK_           = 92
	SQLiteParserSELECT_             = 93
	SQLiteParserSET_                = 94
	SQLiteParserTHEN_               = 95
	SQLiteParserTRUE_               = 96
	SQLiteParserUNION_              = 97
	SQLiteParserUPDATE_             = 98
	SQLiteParserUSING_              = 99
	SQLiteParserVALUES_             = 100
	SQLiteParserWHEN_               = 101
	SQLiteParserWHERE_              = 102
	SQLiteParserWITH_               = 103
	SQLiteParserIDENTIFIER          = 104
	SQLiteParserNUMERIC_LITERAL     = 105
	SQLiteParserBIND_PARAMETER      = 106
	SQLiteParserSTRING_LITERAL      = 107
	SQLiteParserBLOB_LITERAL        = 108
	SQLiteParserSINGLE_LINE_COMMENT = 109
	SQLiteParserMULTILINE_COMMENT   = 110
	SQLiteParserSPACES              = 111
	SQLiteParserUNEXPECTED_CHAR     = 112
)

// SQLiteParser rules.
const (
	SQLiteParserRULE_parse                          = 0
	SQLiteParserRULE_sql_stmt_list                  = 1
	SQLiteParserRULE_sql_stmt                       = 2
	SQLiteParserRULE_indexed_column                 = 3
	SQLiteParserRULE_cte_table_name                 = 4
	SQLiteParserRULE_common_table_expression        = 5
	SQLiteParserRULE_common_table_stmt              = 6
	SQLiteParserRULE_delete_stmt                    = 7
	SQLiteParserRULE_expr                           = 8
	SQLiteParserRULE_cast_type                      = 9
	SQLiteParserRULE_type_cast                      = 10
	SQLiteParserRULE_literal_value                  = 11
	SQLiteParserRULE_value_row                      = 12
	SQLiteParserRULE_values_clause                  = 13
	SQLiteParserRULE_insert_stmt                    = 14
	SQLiteParserRULE_returning_clause               = 15
	SQLiteParserRULE_upsert_update                  = 16
	SQLiteParserRULE_upsert_clause                  = 17
	SQLiteParserRULE_select_stmt_core               = 18
	SQLiteParserRULE_select_stmt                    = 19
	SQLiteParserRULE_join_clause                    = 20
	SQLiteParserRULE_select_core                    = 21
	SQLiteParserRULE_table_or_subquery              = 22
	SQLiteParserRULE_result_column                  = 23
	SQLiteParserRULE_returning_clause_result_column = 24
	SQLiteParserRULE_join_operator                  = 25
	SQLiteParserRULE_join_constraint                = 26
	SQLiteParserRULE_compound_operator              = 27
	SQLiteParserRULE_update_set_subclause           = 28
	SQLiteParserRULE_update_stmt                    = 29
	SQLiteParserRULE_column_name_list               = 30
	SQLiteParserRULE_qualified_table_name           = 31
	SQLiteParserRULE_order_by_stmt                  = 32
	SQLiteParserRULE_limit_stmt                     = 33
	SQLiteParserRULE_ordering_term                  = 34
	SQLiteParserRULE_asc_desc                       = 35
	SQLiteParserRULE_function_keyword               = 36
	SQLiteParserRULE_function_name                  = 37
	SQLiteParserRULE_table_name                     = 38
	SQLiteParserRULE_table_alias                    = 39
	SQLiteParserRULE_column_name                    = 40
	SQLiteParserRULE_column_alias                   = 41
	SQLiteParserRULE_collation_name                 = 42
	SQLiteParserRULE_index_name                     = 43
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllSql_stmt_list() []ISql_stmt_listContext
	Sql_stmt_list(i int) ISql_stmt_listContext

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(SQLiteParserEOF, 0)
}

func (s *ParseContext) AllSql_stmt_list() []ISql_stmt_listContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISql_stmt_listContext); ok {
			len++
		}
	}

	tst := make([]ISql_stmt_listContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISql_stmt_listContext); ok {
			tst[i] = t.(ISql_stmt_listContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Sql_stmt_list(i int) ISql_stmt_listContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISql_stmt_listContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISql_stmt_listContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Parse() (localctx IParseContext) {
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLiteParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9223369837831520254) != 0) || ((int64((_la-89)) & ^0x3f) == 0 && ((int64(1)<<(_la-89))&16913) != 0) {
		{
			p.SetState(88)
			p.Sql_stmt_list()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(SQLiteParserEOF)
	}

	return localctx
}

// ISql_stmt_listContext is an interface to support dynamic dispatch.
type ISql_stmt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSql_stmt() []ISql_stmtContext
	Sql_stmt(i int) ISql_stmtContext
	AllSCOL() []antlr.TerminalNode
	SCOL(i int) antlr.TerminalNode

	// IsSql_stmt_listContext differentiates from other interfaces.
	IsSql_stmt_listContext()
}

type Sql_stmt_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmt_listContext() *Sql_stmt_listContext {
	var p = new(Sql_stmt_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_sql_stmt_list
	return p
}

func (*Sql_stmt_listContext) IsSql_stmt_listContext() {}

func NewSql_stmt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmt_listContext {
	var p = new(Sql_stmt_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_sql_stmt_list

	return p
}

func (s *Sql_stmt_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_stmt_listContext) AllSql_stmt() []ISql_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISql_stmtContext); ok {
			len++
		}
	}

	tst := make([]ISql_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISql_stmtContext); ok {
			tst[i] = t.(ISql_stmtContext)
			i++
		}
	}

	return tst
}

func (s *Sql_stmt_listContext) Sql_stmt(i int) ISql_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISql_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISql_stmtContext)
}

func (s *Sql_stmt_listContext) AllSCOL() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserSCOL)
}

func (s *Sql_stmt_listContext) SCOL(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserSCOL, i)
}

func (s *Sql_stmt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_stmt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_stmt_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitSql_stmt_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Sql_stmt_list() (localctx ISql_stmt_listContext) {
	this := p
	_ = this

	localctx = NewSql_stmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLiteParserRULE_sql_stmt_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserSCOL {
		{
			p.SetState(96)
			p.Match(SQLiteParserSCOL)
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(102)
		p.Sql_stmt()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == SQLiteParserSCOL {
				{
					p.SetState(103)
					p.Match(SQLiteParserSCOL)
				}

				p.SetState(106)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(108)
				p.Sql_stmt()
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(114)
				p.Match(SQLiteParserSCOL)
			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ISql_stmtContext is an interface to support dynamic dispatch.
type ISql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Delete_stmt() IDelete_stmtContext
	Insert_stmt() IInsert_stmtContext
	Select_stmt() ISelect_stmtContext
	Update_stmt() IUpdate_stmtContext

	// IsSql_stmtContext differentiates from other interfaces.
	IsSql_stmtContext()
}

type Sql_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmtContext() *Sql_stmtContext {
	var p = new(Sql_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_sql_stmt
	return p
}

func (*Sql_stmtContext) IsSql_stmtContext() {}

func NewSql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmtContext {
	var p = new(Sql_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_sql_stmt

	return p
}

func (s *Sql_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_stmtContext) Delete_stmt() IDelete_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelete_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelete_stmtContext)
}

func (s *Sql_stmtContext) Insert_stmt() IInsert_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsert_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsert_stmtContext)
}

func (s *Sql_stmtContext) Select_stmt() ISelect_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_stmtContext)
}

func (s *Sql_stmtContext) Update_stmt() IUpdate_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdate_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdate_stmtContext)
}

func (s *Sql_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitSql_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Sql_stmt() (localctx ISql_stmtContext) {
	this := p
	_ = this

	localctx = NewSql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLiteParserRULE_sql_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(120)
			p.Delete_stmt()
		}

	case 2:
		{
			p.SetState(121)
			p.Insert_stmt()
		}

	case 3:
		{
			p.SetState(122)
			p.Select_stmt()
		}

	case 4:
		{
			p.SetState(123)
			p.Update_stmt()
		}

	}

	return localctx
}

// IIndexed_columnContext is an interface to support dynamic dispatch.
type IIndexed_columnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Column_name() IColumn_nameContext

	// IsIndexed_columnContext differentiates from other interfaces.
	IsIndexed_columnContext()
}

type Indexed_columnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexed_columnContext() *Indexed_columnContext {
	var p = new(Indexed_columnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_indexed_column
	return p
}

func (*Indexed_columnContext) IsIndexed_columnContext() {}

func NewIndexed_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Indexed_columnContext {
	var p = new(Indexed_columnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_indexed_column

	return p
}

func (s *Indexed_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Indexed_columnContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Indexed_columnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Indexed_columnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Indexed_columnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitIndexed_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Indexed_column() (localctx IIndexed_columnContext) {
	this := p
	_ = this

	localctx = NewIndexed_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLiteParserRULE_indexed_column)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Column_name()
	}

	return localctx
}

// ICte_table_nameContext is an interface to support dynamic dispatch.
type ICte_table_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Table_name() ITable_nameContext
	OPEN_PAR() antlr.TerminalNode
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext
	CLOSE_PAR() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCte_table_nameContext differentiates from other interfaces.
	IsCte_table_nameContext()
}

type Cte_table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCte_table_nameContext() *Cte_table_nameContext {
	var p = new(Cte_table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_cte_table_name
	return p
}

func (*Cte_table_nameContext) IsCte_table_nameContext() {}

func NewCte_table_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_table_nameContext {
	var p = new(Cte_table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_cte_table_name

	return p
}

func (s *Cte_table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Cte_table_nameContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Cte_table_nameContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *Cte_table_nameContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Cte_table_nameContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Cte_table_nameContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *Cte_table_nameContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Cte_table_nameContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Cte_table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitCte_table_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Cte_table_name() (localctx ICte_table_nameContext) {
	this := p
	_ = this

	localctx = NewCte_table_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLiteParserRULE_cte_table_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Table_name()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOPEN_PAR {
		{
			p.SetState(129)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(130)
			p.Column_name()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(131)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(132)
				p.Column_name()
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(138)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	}

	return localctx
}

// ICommon_table_expressionContext is an interface to support dynamic dispatch.
type ICommon_table_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Cte_table_name() ICte_table_nameContext
	AS_() antlr.TerminalNode
	OPEN_PAR() antlr.TerminalNode
	Select_stmt_core() ISelect_stmt_coreContext
	CLOSE_PAR() antlr.TerminalNode

	// IsCommon_table_expressionContext differentiates from other interfaces.
	IsCommon_table_expressionContext()
}

type Common_table_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommon_table_expressionContext() *Common_table_expressionContext {
	var p = new(Common_table_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_common_table_expression
	return p
}

func (*Common_table_expressionContext) IsCommon_table_expressionContext() {}

func NewCommon_table_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Common_table_expressionContext {
	var p = new(Common_table_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_common_table_expression

	return p
}

func (s *Common_table_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Common_table_expressionContext) Cte_table_name() ICte_table_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICte_table_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICte_table_nameContext)
}

func (s *Common_table_expressionContext) AS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAS_, 0)
}

func (s *Common_table_expressionContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *Common_table_expressionContext) Select_stmt_core() ISelect_stmt_coreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_stmt_coreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_stmt_coreContext)
}

func (s *Common_table_expressionContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *Common_table_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Common_table_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Common_table_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitCommon_table_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Common_table_expression() (localctx ICommon_table_expressionContext) {
	this := p
	_ = this

	localctx = NewCommon_table_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLiteParserRULE_common_table_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Cte_table_name()
	}
	{
		p.SetState(143)
		p.Match(SQLiteParserAS_)
	}
	{
		p.SetState(144)
		p.Match(SQLiteParserOPEN_PAR)
	}
	{
		p.SetState(145)
		p.Select_stmt_core()
	}
	{
		p.SetState(146)
		p.Match(SQLiteParserCLOSE_PAR)
	}

	return localctx
}

// ICommon_table_stmtContext is an interface to support dynamic dispatch.
type ICommon_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WITH_() antlr.TerminalNode
	AllCommon_table_expression() []ICommon_table_expressionContext
	Common_table_expression(i int) ICommon_table_expressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCommon_table_stmtContext differentiates from other interfaces.
	IsCommon_table_stmtContext()
}

type Common_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommon_table_stmtContext() *Common_table_stmtContext {
	var p = new(Common_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_common_table_stmt
	return p
}

func (*Common_table_stmtContext) IsCommon_table_stmtContext() {}

func NewCommon_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Common_table_stmtContext {
	var p = new(Common_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_common_table_stmt

	return p
}

func (s *Common_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Common_table_stmtContext) WITH_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserWITH_, 0)
}

func (s *Common_table_stmtContext) AllCommon_table_expression() []ICommon_table_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommon_table_expressionContext); ok {
			len++
		}
	}

	tst := make([]ICommon_table_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommon_table_expressionContext); ok {
			tst[i] = t.(ICommon_table_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Common_table_stmtContext) Common_table_expression(i int) ICommon_table_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommon_table_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommon_table_expressionContext)
}

func (s *Common_table_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Common_table_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Common_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Common_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Common_table_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitCommon_table_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Common_table_stmt() (localctx ICommon_table_stmtContext) {
	this := p
	_ = this

	localctx = NewCommon_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLiteParserRULE_common_table_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(SQLiteParserWITH_)
	}
	{
		p.SetState(149)
		p.Common_table_expression()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(150)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(151)
			p.Common_table_expression()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDelete_stmtContext is an interface to support dynamic dispatch.
type IDelete_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE_() antlr.TerminalNode
	FROM_() antlr.TerminalNode
	Qualified_table_name() IQualified_table_nameContext
	Common_table_stmt() ICommon_table_stmtContext
	WHERE_() antlr.TerminalNode
	Expr() IExprContext
	Returning_clause() IReturning_clauseContext

	// IsDelete_stmtContext differentiates from other interfaces.
	IsDelete_stmtContext()
}

type Delete_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_stmtContext() *Delete_stmtContext {
	var p = new(Delete_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_delete_stmt
	return p
}

func (*Delete_stmtContext) IsDelete_stmtContext() {}

func NewDelete_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_stmtContext {
	var p = new(Delete_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_delete_stmt

	return p
}

func (s *Delete_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_stmtContext) DELETE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDELETE_, 0)
}

func (s *Delete_stmtContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFROM_, 0)
}

func (s *Delete_stmtContext) Qualified_table_name() IQualified_table_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualified_table_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualified_table_nameContext)
}

func (s *Delete_stmtContext) Common_table_stmt() ICommon_table_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommon_table_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommon_table_stmtContext)
}

func (s *Delete_stmtContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserWHERE_, 0)
}

func (s *Delete_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Delete_stmtContext) Returning_clause() IReturning_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturning_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturning_clauseContext)
}

func (s *Delete_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitDelete_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Delete_stmt() (localctx IDelete_stmtContext) {
	this := p
	_ = this

	localctx = NewDelete_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLiteParserRULE_delete_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(157)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(160)
		p.Match(SQLiteParserDELETE_)
	}
	{
		p.SetState(161)
		p.Match(SQLiteParserFROM_)
	}
	{
		p.SetState(162)
		p.Qualified_table_name()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWHERE_ {
		{
			p.SetState(163)
			p.Match(SQLiteParserWHERE_)
		}
		{
			p.SetState(164)
			p.expr(0)
		}

	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserRETURNING_ {
		{
			p.SetState(167)
			p.Returning_clause()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetElevate_expr returns the elevate_expr rule contexts.
	GetElevate_expr() IExprContext

	// GetUnary_expr returns the unary_expr rule contexts.
	GetUnary_expr() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetCase_expr returns the case_expr rule contexts.
	GetCase_expr() IExprContext

	// GetElse_expr returns the else_expr rule contexts.
	GetElse_expr() IExprContext

	// SetElevate_expr sets the elevate_expr rule contexts.
	SetElevate_expr(IExprContext)

	// SetUnary_expr sets the unary_expr rule contexts.
	SetUnary_expr(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetCase_expr sets the case_expr rule contexts.
	SetCase_expr(IExprContext)

	// SetElse_expr sets the else_expr rule contexts.
	SetElse_expr(IExprContext)

	// GetExpr_list returns the expr_list rule context list.
	GetExpr_list() []IExprContext

	// GetWhen_expr returns the when_expr rule context list.
	GetWhen_expr() []IExprContext

	// GetThen_expr returns the then_expr rule context list.
	GetThen_expr() []IExprContext

	// SetExpr_list sets the expr_list rule context list.
	SetExpr_list([]IExprContext)

	// SetWhen_expr sets the when_expr rule context list.
	SetWhen_expr([]IExprContext)

	// SetThen_expr sets the then_expr rule context list.
	SetThen_expr([]IExprContext)

	// Getter signatures
	Literal_value() ILiteral_valueContext
	Type_cast() IType_castContext
	BIND_PARAMETER() antlr.TerminalNode
	Column_name() IColumn_nameContext
	Table_name() ITable_nameContext
	DOT() antlr.TerminalNode
	OPEN_PAR() antlr.TerminalNode
	Select_stmt_core() ISelect_stmt_coreContext
	CLOSE_PAR() antlr.TerminalNode
	EXISTS_() antlr.TerminalNode
	NOT_() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	TILDE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Function_name() IFunction_nameContext
	STAR() antlr.TerminalNode
	DISTINCT_() antlr.TerminalNode
	CASE_() antlr.TerminalNode
	END_() antlr.TerminalNode
	AllWHEN_() []antlr.TerminalNode
	WHEN_(i int) antlr.TerminalNode
	AllTHEN_() []antlr.TerminalNode
	THEN_(i int) antlr.TerminalNode
	ELSE_() antlr.TerminalNode
	PIPE2() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LT2() antlr.TerminalNode
	GT2() antlr.TerminalNode
	AMP() antlr.TerminalNode
	PIPE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LT_EQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	GT_EQ() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NOT_EQ1() antlr.TerminalNode
	NOT_EQ2() antlr.TerminalNode
	IS_() antlr.TerminalNode
	FROM_() antlr.TerminalNode
	IN_() antlr.TerminalNode
	GLOB_() antlr.TerminalNode
	MATCH_() antlr.TerminalNode
	REGEXP_() antlr.TerminalNode
	BETWEEN_() antlr.TerminalNode
	AND_() antlr.TerminalNode
	OR_() antlr.TerminalNode
	COLLATE_() antlr.TerminalNode
	Collation_name() ICollation_nameContext
	LIKE_() antlr.TerminalNode
	ESCAPE_() antlr.TerminalNode
	ISNULL_() antlr.TerminalNode
	NOTNULL_() antlr.TerminalNode
	NULL_() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	elevate_expr IExprContext
	unary_expr   IExprContext
	_expr        IExprContext
	expr_list    []IExprContext
	case_expr    IExprContext
	when_expr    []IExprContext
	then_expr    []IExprContext
	else_expr    IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetElevate_expr() IExprContext { return s.elevate_expr }

func (s *ExprContext) GetUnary_expr() IExprContext { return s.unary_expr }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) GetCase_expr() IExprContext { return s.case_expr }

func (s *ExprContext) GetElse_expr() IExprContext { return s.else_expr }

func (s *ExprContext) SetElevate_expr(v IExprContext) { s.elevate_expr = v }

func (s *ExprContext) SetUnary_expr(v IExprContext) { s.unary_expr = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) SetCase_expr(v IExprContext) { s.case_expr = v }

func (s *ExprContext) SetElse_expr(v IExprContext) { s.else_expr = v }

func (s *ExprContext) GetExpr_list() []IExprContext { return s.expr_list }

func (s *ExprContext) GetWhen_expr() []IExprContext { return s.when_expr }

func (s *ExprContext) GetThen_expr() []IExprContext { return s.then_expr }

func (s *ExprContext) SetExpr_list(v []IExprContext) { s.expr_list = v }

func (s *ExprContext) SetWhen_expr(v []IExprContext) { s.when_expr = v }

func (s *ExprContext) SetThen_expr(v []IExprContext) { s.then_expr = v }

func (s *ExprContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *ExprContext) Type_cast() IType_castContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_castContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_castContext)
}

func (s *ExprContext) BIND_PARAMETER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserBIND_PARAMETER, 0)
}

func (s *ExprContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *ExprContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *ExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDOT, 0)
}

func (s *ExprContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *ExprContext) Select_stmt_core() ISelect_stmt_coreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_stmt_coreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_stmt_coreContext)
}

func (s *ExprContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *ExprContext) EXISTS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserEXISTS_, 0)
}

func (s *ExprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNOT_, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SQLiteParserMINUS, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SQLiteParserPLUS, 0)
}

func (s *ExprContext) TILDE() antlr.TerminalNode {
	return s.GetToken(SQLiteParserTILDE, 0)
}

func (s *ExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *ExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *ExprContext) Function_name() IFunction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *ExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserSTAR, 0)
}

func (s *ExprContext) DISTINCT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDISTINCT_, 0)
}

func (s *ExprContext) CASE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCASE_, 0)
}

func (s *ExprContext) END_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserEND_, 0)
}

func (s *ExprContext) AllWHEN_() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserWHEN_)
}

func (s *ExprContext) WHEN_(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserWHEN_, i)
}

func (s *ExprContext) AllTHEN_() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserTHEN_)
}

func (s *ExprContext) THEN_(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserTHEN_, i)
}

func (s *ExprContext) ELSE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserELSE_, 0)
}

func (s *ExprContext) PIPE2() antlr.TerminalNode {
	return s.GetToken(SQLiteParserPIPE2, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SQLiteParserMOD, 0)
}

func (s *ExprContext) LT2() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLT2, 0)
}

func (s *ExprContext) GT2() antlr.TerminalNode {
	return s.GetToken(SQLiteParserGT2, 0)
}

func (s *ExprContext) AMP() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAMP, 0)
}

func (s *ExprContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SQLiteParserPIPE, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLT, 0)
}

func (s *ExprContext) LT_EQ() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLT_EQ, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(SQLiteParserGT, 0)
}

func (s *ExprContext) GT_EQ() antlr.TerminalNode {
	return s.GetToken(SQLiteParserGT_EQ, 0)
}

func (s *ExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SQLiteParserASSIGN, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(SQLiteParserEQ, 0)
}

func (s *ExprContext) NOT_EQ1() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNOT_EQ1, 0)
}

func (s *ExprContext) NOT_EQ2() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNOT_EQ2, 0)
}

func (s *ExprContext) IS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIS_, 0)
}

func (s *ExprContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFROM_, 0)
}

func (s *ExprContext) IN_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIN_, 0)
}

func (s *ExprContext) GLOB_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserGLOB_, 0)
}

func (s *ExprContext) MATCH_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserMATCH_, 0)
}

func (s *ExprContext) REGEXP_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserREGEXP_, 0)
}

func (s *ExprContext) BETWEEN_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserBETWEEN_, 0)
}

func (s *ExprContext) AND_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAND_, 0)
}

func (s *ExprContext) OR_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOR_, 0)
}

func (s *ExprContext) COLLATE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOLLATE_, 0)
}

func (s *ExprContext) Collation_name() ICollation_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollation_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollation_nameContext)
}

func (s *ExprContext) LIKE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLIKE_, 0)
}

func (s *ExprContext) ESCAPE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserESCAPE_, 0)
}

func (s *ExprContext) ISNULL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserISNULL_, 0)
}

func (s *ExprContext) NOTNULL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNOTNULL_, 0)
}

func (s *ExprContext) NULL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNULL_, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SQLiteParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, SQLiteParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(171)
			p.Literal_value()
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(172)
				p.Type_cast()
			}

		}

	case 2:
		{
			p.SetState(175)
			p.Match(SQLiteParserBIND_PARAMETER)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(176)
				p.Type_cast()
			}

		}

	case 3:
		p.SetState(182)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(179)
				p.Table_name()
			}
			{
				p.SetState(180)
				p.Match(SQLiteParserDOT)
			}

		}
		{
			p.SetState(184)
			p.Column_name()
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(185)
				p.Type_cast()
			}

		}

	case 4:
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserEXISTS_ || _la == SQLiteParserNOT_ {
			p.SetState(189)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == SQLiteParserNOT_ {
				{
					p.SetState(188)
					p.Match(SQLiteParserNOT_)
				}

			}
			{
				p.SetState(191)
				p.Match(SQLiteParserEXISTS_)
			}

		}
		{
			p.SetState(194)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(195)
			p.Select_stmt_core()
		}
		{
			p.SetState(196)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	case 5:
		{
			p.SetState(198)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(199)

			var _x = p.expr(0)

			localctx.(*ExprContext).elevate_expr = _x
		}
		{
			p.SetState(200)
			p.Match(SQLiteParserCLOSE_PAR)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(201)
				p.Type_cast()
			}

		}

	case 6:
		{
			p.SetState(204)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1792) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(205)

			var _x = p.expr(17)

			localctx.(*ExprContext).unary_expr = _x
		}

	case 7:
		{
			p.SetState(206)
			p.Match(SQLiteParserNOT_)
		}
		{
			p.SetState(207)

			var _x = p.expr(6)

			localctx.(*ExprContext).unary_expr = _x
		}

	case 8:
		{
			p.SetState(208)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(209)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).expr_list = append(localctx.(*ExprContext).expr_list, localctx.(*ExprContext)._expr)
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(210)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(211)

				var _x = p.expr(0)

				localctx.(*ExprContext)._expr = _x
			}
			localctx.(*ExprContext).expr_list = append(localctx.(*ExprContext).expr_list, localctx.(*ExprContext)._expr)

			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(217)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	case 9:
		{
			p.SetState(219)
			p.Function_name()
		}
		{
			p.SetState(220)
			p.Match(SQLiteParserOPEN_PAR)
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SQLiteParserOPEN_PAR, SQLiteParserPLUS, SQLiteParserMINUS, SQLiteParserTILDE, SQLiteParserCASE_, SQLiteParserDISTINCT_, SQLiteParserEXISTS_, SQLiteParserFALSE_, SQLiteParserGLOB_, SQLiteParserLIKE_, SQLiteParserNOT_, SQLiteParserNULL_, SQLiteParserREPLACE_, SQLiteParserTRUE_, SQLiteParserIDENTIFIER, SQLiteParserNUMERIC_LITERAL, SQLiteParserBIND_PARAMETER, SQLiteParserSTRING_LITERAL:
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == SQLiteParserDISTINCT_ {
				{
					p.SetState(221)
					p.Match(SQLiteParserDISTINCT_)
				}

			}
			{
				p.SetState(224)
				p.expr(0)
			}
			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == SQLiteParserCOMMA {
				{
					p.SetState(225)
					p.Match(SQLiteParserCOMMA)
				}
				{
					p.SetState(226)
					p.expr(0)
				}

				p.SetState(231)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		case SQLiteParserSTAR:
			{
				p.SetState(232)
				p.Match(SQLiteParserSTAR)
			}

		case SQLiteParserCLOSE_PAR:

		default:
		}
		{
			p.SetState(235)
			p.Match(SQLiteParserCLOSE_PAR)
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(236)
				p.Type_cast()
			}

		}

	case 10:
		{
			p.SetState(239)
			p.Match(SQLiteParserCASE_)
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&74872360984905480) != 0) || ((int64((_la-72)) & ^0x3f) == 0 && ((int64(1)<<(_la-72))&64441418049) != 0) {
			{
				p.SetState(240)

				var _x = p.expr(0)

				localctx.(*ExprContext).case_expr = _x
			}

		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SQLiteParserWHEN_ {
			{
				p.SetState(243)
				p.Match(SQLiteParserWHEN_)
			}
			{
				p.SetState(244)

				var _x = p.expr(0)

				localctx.(*ExprContext)._expr = _x
			}
			localctx.(*ExprContext).when_expr = append(localctx.(*ExprContext).when_expr, localctx.(*ExprContext)._expr)
			{
				p.SetState(245)
				p.Match(SQLiteParserTHEN_)
			}
			{
				p.SetState(246)

				var _x = p.expr(0)

				localctx.(*ExprContext)._expr = _x
			}
			localctx.(*ExprContext).then_expr = append(localctx.(*ExprContext).then_expr, localctx.(*ExprContext)._expr)

			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserELSE_ {
			{
				p.SetState(252)
				p.Match(SQLiteParserELSE_)
			}
			{
				p.SetState(253)

				var _x = p.expr(0)

				localctx.(*ExprContext).else_expr = _x
			}

		}
		{
			p.SetState(256)
			p.Match(SQLiteParserEND_)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(261)
					p.Match(SQLiteParserPIPE2)
				}
				{
					p.SetState(262)
					p.expr(16)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(264)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12416) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(265)
					p.expr(15)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(267)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SQLiteParserPLUS || _la == SQLiteParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(268)
					p.expr(14)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(270)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(271)
					p.expr(13)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(273)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3932160) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(274)
					p.expr(12)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(294)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(276)
						p.Match(SQLiteParserASSIGN)
					}

				case 2:
					{
						p.SetState(277)
						p.Match(SQLiteParserEQ)
					}

				case 3:
					{
						p.SetState(278)
						p.Match(SQLiteParserNOT_EQ1)
					}

				case 4:
					{
						p.SetState(279)
						p.Match(SQLiteParserNOT_EQ2)
					}

				case 5:
					{
						p.SetState(280)
						p.Match(SQLiteParserIS_)
					}
					p.SetState(282)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
						{
							p.SetState(281)
							p.Match(SQLiteParserNOT_)
						}

					}

				case 6:
					{
						p.SetState(284)
						p.Match(SQLiteParserIS_)
					}
					p.SetState(286)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == SQLiteParserNOT_ {
						{
							p.SetState(285)
							p.Match(SQLiteParserNOT_)
						}

					}
					{
						p.SetState(288)
						p.Match(SQLiteParserDISTINCT_)
					}
					{
						p.SetState(289)
						p.Match(SQLiteParserFROM_)
					}

				case 7:
					p.SetState(291)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == SQLiteParserNOT_ {
						{
							p.SetState(290)
							p.Match(SQLiteParserNOT_)
						}

					}
					{
						p.SetState(293)
						_la = p.GetTokenStream().LA(1)

						if !((int64((_la-56)) & ^0x3f) == 0 && ((int64(1)<<(_la-56))&4295230465) != 0) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				}
				{
					p.SetState(296)
					p.expr(11)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				p.SetState(299)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SQLiteParserNOT_ {
					{
						p.SetState(298)
						p.Match(SQLiteParserNOT_)
					}

				}
				{
					p.SetState(301)
					p.Match(SQLiteParserBETWEEN_)
				}
				{
					p.SetState(302)
					p.expr(0)
				}
				{
					p.SetState(303)
					p.Match(SQLiteParserAND_)
				}
				{
					p.SetState(304)
					p.expr(9)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(306)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(307)
					p.Match(SQLiteParserAND_)
				}
				{
					p.SetState(308)
					p.expr(6)
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(309)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(310)
					p.Match(SQLiteParserOR_)
				}
				{
					p.SetState(311)
					p.expr(5)
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(312)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(313)
					p.Match(SQLiteParserCOLLATE_)
				}
				{
					p.SetState(314)
					p.Collation_name()
				}

			case 11:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(315)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(317)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SQLiteParserNOT_ {
					{
						p.SetState(316)
						p.Match(SQLiteParserNOT_)
					}

				}
				{
					p.SetState(319)
					p.Match(SQLiteParserLIKE_)
				}
				{
					p.SetState(320)
					p.expr(0)
				}
				p.SetState(323)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(321)
						p.Match(SQLiteParserESCAPE_)
					}
					{
						p.SetState(322)
						p.expr(0)
					}

				}

			case 12:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				p.SetState(330)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case SQLiteParserISNULL_:
					{
						p.SetState(326)
						p.Match(SQLiteParserISNULL_)
					}

				case SQLiteParserNOTNULL_:
					{
						p.SetState(327)
						p.Match(SQLiteParserNOTNULL_)
					}

				case SQLiteParserNOT_:
					{
						p.SetState(328)
						p.Match(SQLiteParserNOT_)
					}
					{
						p.SetState(329)
						p.Match(SQLiteParserNULL_)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}

		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// ICast_typeContext is an interface to support dynamic dispatch.
type ICast_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsCast_typeContext differentiates from other interfaces.
	IsCast_typeContext()
}

type Cast_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCast_typeContext() *Cast_typeContext {
	var p = new(Cast_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_cast_type
	return p
}

func (*Cast_typeContext) IsCast_typeContext() {}

func NewCast_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cast_typeContext {
	var p = new(Cast_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_cast_type

	return p
}

func (s *Cast_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Cast_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Cast_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cast_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cast_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitCast_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Cast_type() (localctx ICast_typeContext) {
	this := p
	_ = this

	localctx = NewCast_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLiteParserRULE_cast_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(SQLiteParserIDENTIFIER)
	}

	return localctx
}

// IType_castContext is an interface to support dynamic dispatch.
type IType_castContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_CAST() antlr.TerminalNode
	Cast_type() ICast_typeContext

	// IsType_castContext differentiates from other interfaces.
	IsType_castContext()
}

type Type_castContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_castContext() *Type_castContext {
	var p = new(Type_castContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_type_cast
	return p
}

func (*Type_castContext) IsType_castContext() {}

func NewType_castContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_castContext {
	var p = new(Type_castContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_type_cast

	return p
}

func (s *Type_castContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_castContext) TYPE_CAST() antlr.TerminalNode {
	return s.GetToken(SQLiteParserTYPE_CAST, 0)
}

func (s *Type_castContext) Cast_type() ICast_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICast_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICast_typeContext)
}

func (s *Type_castContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_castContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_castContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitType_cast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Type_cast() (localctx IType_castContext) {
	this := p
	_ = this

	localctx = NewType_castContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SQLiteParserRULE_type_cast)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(SQLiteParserTYPE_CAST)
	}
	{
		p.SetState(340)
		p.Cast_type()
	}

	return localctx
}

// ILiteral_valueContext is an interface to support dynamic dispatch.
type ILiteral_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	NULL_() antlr.TerminalNode
	TRUE_() antlr.TerminalNode
	FALSE_() antlr.TerminalNode

	// IsLiteral_valueContext differentiates from other interfaces.
	IsLiteral_valueContext()
}

type Literal_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_valueContext() *Literal_valueContext {
	var p = new(Literal_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_literal_value
	return p
}

func (*Literal_valueContext) IsLiteral_valueContext() {}

func NewLiteral_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_valueContext {
	var p = new(Literal_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_literal_value

	return p
}

func (s *Literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_valueContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNUMERIC_LITERAL, 0)
}

func (s *Literal_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLiteParserSTRING_LITERAL, 0)
}

func (s *Literal_valueContext) NULL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNULL_, 0)
}

func (s *Literal_valueContext) TRUE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserTRUE_, 0)
}

func (s *Literal_valueContext) FALSE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFALSE_, 0)
}

func (s *Literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitLiteral_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Literal_value() (localctx ILiteral_valueContext) {
	this := p
	_ = this

	localctx = NewLiteral_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SQLiteParserRULE_literal_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-51)) & ^0x3f) == 0 && ((int64(1)<<(_la-51))&90107177456369665) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValue_rowContext is an interface to support dynamic dispatch.
type IValue_rowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAR() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	CLOSE_PAR() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValue_rowContext differentiates from other interfaces.
	IsValue_rowContext()
}

type Value_rowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_rowContext() *Value_rowContext {
	var p = new(Value_rowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_value_row
	return p
}

func (*Value_rowContext) IsValue_rowContext() {}

func NewValue_rowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_rowContext {
	var p = new(Value_rowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_value_row

	return p
}

func (s *Value_rowContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_rowContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *Value_rowContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Value_rowContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Value_rowContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *Value_rowContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Value_rowContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Value_rowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_rowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_rowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitValue_row(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Value_row() (localctx IValue_rowContext) {
	this := p
	_ = this

	localctx = NewValue_rowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SQLiteParserRULE_value_row)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(SQLiteParserOPEN_PAR)
	}
	{
		p.SetState(345)
		p.expr(0)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(346)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(347)
			p.expr(0)
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(353)
		p.Match(SQLiteParserCLOSE_PAR)
	}

	return localctx
}

// IValues_clauseContext is an interface to support dynamic dispatch.
type IValues_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUES_() antlr.TerminalNode
	AllValue_row() []IValue_rowContext
	Value_row(i int) IValue_rowContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValues_clauseContext differentiates from other interfaces.
	IsValues_clauseContext()
}

type Values_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValues_clauseContext() *Values_clauseContext {
	var p = new(Values_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_values_clause
	return p
}

func (*Values_clauseContext) IsValues_clauseContext() {}

func NewValues_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Values_clauseContext {
	var p = new(Values_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_values_clause

	return p
}

func (s *Values_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Values_clauseContext) VALUES_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserVALUES_, 0)
}

func (s *Values_clauseContext) AllValue_row() []IValue_rowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValue_rowContext); ok {
			len++
		}
	}

	tst := make([]IValue_rowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValue_rowContext); ok {
			tst[i] = t.(IValue_rowContext)
			i++
		}
	}

	return tst
}

func (s *Values_clauseContext) Value_row(i int) IValue_rowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_rowContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_rowContext)
}

func (s *Values_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Values_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Values_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Values_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Values_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitValues_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Values_clause() (localctx IValues_clauseContext) {
	this := p
	_ = this

	localctx = NewValues_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SQLiteParserRULE_values_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(SQLiteParserVALUES_)
	}
	{
		p.SetState(356)
		p.Value_row()
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(357)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(358)
			p.Value_row()
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInsert_stmtContext is an interface to support dynamic dispatch.
type IInsert_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTO_() antlr.TerminalNode
	Table_name() ITable_nameContext
	Values_clause() IValues_clauseContext
	REPLACE_() antlr.TerminalNode
	INSERT_() antlr.TerminalNode
	OR_() antlr.TerminalNode
	Common_table_stmt() ICommon_table_stmtContext
	AS_() antlr.TerminalNode
	Table_alias() ITable_aliasContext
	OPEN_PAR() antlr.TerminalNode
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext
	CLOSE_PAR() antlr.TerminalNode
	Upsert_clause() IUpsert_clauseContext
	Returning_clause() IReturning_clauseContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsInsert_stmtContext differentiates from other interfaces.
	IsInsert_stmtContext()
}

type Insert_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_stmtContext() *Insert_stmtContext {
	var p = new(Insert_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_insert_stmt
	return p
}

func (*Insert_stmtContext) IsInsert_stmtContext() {}

func NewInsert_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_stmtContext {
	var p = new(Insert_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_insert_stmt

	return p
}

func (s *Insert_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_stmtContext) INTO_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserINTO_, 0)
}

func (s *Insert_stmtContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Insert_stmtContext) Values_clause() IValues_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValues_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValues_clauseContext)
}

func (s *Insert_stmtContext) REPLACE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserREPLACE_, 0)
}

func (s *Insert_stmtContext) INSERT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserINSERT_, 0)
}

func (s *Insert_stmtContext) OR_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOR_, 0)
}

func (s *Insert_stmtContext) Common_table_stmt() ICommon_table_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommon_table_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommon_table_stmtContext)
}

func (s *Insert_stmtContext) AS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAS_, 0)
}

func (s *Insert_stmtContext) Table_alias() ITable_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_aliasContext)
}

func (s *Insert_stmtContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *Insert_stmtContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Insert_stmtContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Insert_stmtContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *Insert_stmtContext) Upsert_clause() IUpsert_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpsert_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpsert_clauseContext)
}

func (s *Insert_stmtContext) Returning_clause() IReturning_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturning_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturning_clauseContext)
}

func (s *Insert_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Insert_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Insert_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitInsert_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Insert_stmt() (localctx IInsert_stmtContext) {
	this := p
	_ = this

	localctx = NewInsert_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SQLiteParserRULE_insert_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(364)
			p.Common_table_stmt()
		}

	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(367)
			p.Match(SQLiteParserREPLACE_)
		}

	case 2:
		{
			p.SetState(368)
			p.Match(SQLiteParserINSERT_)
		}

	case 3:
		{
			p.SetState(369)
			p.Match(SQLiteParserINSERT_)
		}
		{
			p.SetState(370)
			p.Match(SQLiteParserOR_)
		}
		{
			p.SetState(371)
			p.Match(SQLiteParserREPLACE_)
		}

	}
	{
		p.SetState(374)
		p.Match(SQLiteParserINTO_)
	}
	{
		p.SetState(375)
		p.Table_name()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserAS_ {
		{
			p.SetState(376)
			p.Match(SQLiteParserAS_)
		}
		{
			p.SetState(377)
			p.Table_alias()
		}

	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOPEN_PAR {
		{
			p.SetState(380)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(381)
			p.Column_name()
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(382)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(383)
				p.Column_name()
			}

			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(389)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	}
	{
		p.SetState(393)
		p.Values_clause()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserON_ {
		{
			p.SetState(394)
			p.Upsert_clause()
		}

	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserRETURNING_ {
		{
			p.SetState(397)
			p.Returning_clause()
		}

	}

	return localctx
}

// IReturning_clauseContext is an interface to support dynamic dispatch.
type IReturning_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURNING_() antlr.TerminalNode
	AllReturning_clause_result_column() []IReturning_clause_result_columnContext
	Returning_clause_result_column(i int) IReturning_clause_result_columnContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReturning_clauseContext differentiates from other interfaces.
	IsReturning_clauseContext()
}

type Returning_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturning_clauseContext() *Returning_clauseContext {
	var p = new(Returning_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_returning_clause
	return p
}

func (*Returning_clauseContext) IsReturning_clauseContext() {}

func NewReturning_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Returning_clauseContext {
	var p = new(Returning_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_returning_clause

	return p
}

func (s *Returning_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Returning_clauseContext) RETURNING_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserRETURNING_, 0)
}

func (s *Returning_clauseContext) AllReturning_clause_result_column() []IReturning_clause_result_columnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReturning_clause_result_columnContext); ok {
			len++
		}
	}

	tst := make([]IReturning_clause_result_columnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReturning_clause_result_columnContext); ok {
			tst[i] = t.(IReturning_clause_result_columnContext)
			i++
		}
	}

	return tst
}

func (s *Returning_clauseContext) Returning_clause_result_column(i int) IReturning_clause_result_columnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturning_clause_result_columnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturning_clause_result_columnContext)
}

func (s *Returning_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Returning_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Returning_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Returning_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Returning_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitReturning_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Returning_clause() (localctx IReturning_clauseContext) {
	this := p
	_ = this

	localctx = NewReturning_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SQLiteParserRULE_returning_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(SQLiteParserRETURNING_)
	}
	{
		p.SetState(401)
		p.Returning_clause_result_column()
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(402)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(403)
			p.Returning_clause_result_column()
		}

		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUpsert_updateContext is an interface to support dynamic dispatch.
type IUpsert_updateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	Column_name() IColumn_nameContext
	Column_name_list() IColumn_name_listContext

	// IsUpsert_updateContext differentiates from other interfaces.
	IsUpsert_updateContext()
}

type Upsert_updateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpsert_updateContext() *Upsert_updateContext {
	var p = new(Upsert_updateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_upsert_update
	return p
}

func (*Upsert_updateContext) IsUpsert_updateContext() {}

func NewUpsert_updateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upsert_updateContext {
	var p = new(Upsert_updateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_upsert_update

	return p
}

func (s *Upsert_updateContext) GetParser() antlr.Parser { return s.parser }

func (s *Upsert_updateContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SQLiteParserASSIGN, 0)
}

func (s *Upsert_updateContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Upsert_updateContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Upsert_updateContext) Column_name_list() IColumn_name_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_name_listContext)
}

func (s *Upsert_updateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upsert_updateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upsert_updateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitUpsert_update(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Upsert_update() (localctx IUpsert_updateContext) {
	this := p
	_ = this

	localctx = NewUpsert_updateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SQLiteParserRULE_upsert_update)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(411)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		{
			p.SetState(409)
			p.Column_name()
		}

	case SQLiteParserOPEN_PAR:
		{
			p.SetState(410)
			p.Column_name_list()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(413)
		p.Match(SQLiteParserASSIGN)
	}
	{
		p.SetState(414)
		p.expr(0)
	}

	return localctx
}

// IUpsert_clauseContext is an interface to support dynamic dispatch.
type IUpsert_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTarget_expr returns the target_expr rule contexts.
	GetTarget_expr() IExprContext

	// GetUpdate_expr returns the update_expr rule contexts.
	GetUpdate_expr() IExprContext

	// SetTarget_expr sets the target_expr rule contexts.
	SetTarget_expr(IExprContext)

	// SetUpdate_expr sets the update_expr rule contexts.
	SetUpdate_expr(IExprContext)

	// Getter signatures
	ON_() antlr.TerminalNode
	CONFLICT_() antlr.TerminalNode
	DO_() antlr.TerminalNode
	NOTHING_() antlr.TerminalNode
	UPDATE_() antlr.TerminalNode
	SET_() antlr.TerminalNode
	OPEN_PAR() antlr.TerminalNode
	AllIndexed_column() []IIndexed_columnContext
	Indexed_column(i int) IIndexed_columnContext
	CLOSE_PAR() antlr.TerminalNode
	AllUpsert_update() []IUpsert_updateContext
	Upsert_update(i int) IUpsert_updateContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllWHERE_() []antlr.TerminalNode
	WHERE_(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsUpsert_clauseContext differentiates from other interfaces.
	IsUpsert_clauseContext()
}

type Upsert_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	target_expr IExprContext
	update_expr IExprContext
}

func NewEmptyUpsert_clauseContext() *Upsert_clauseContext {
	var p = new(Upsert_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_upsert_clause
	return p
}

func (*Upsert_clauseContext) IsUpsert_clauseContext() {}

func NewUpsert_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upsert_clauseContext {
	var p = new(Upsert_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_upsert_clause

	return p
}

func (s *Upsert_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Upsert_clauseContext) GetTarget_expr() IExprContext { return s.target_expr }

func (s *Upsert_clauseContext) GetUpdate_expr() IExprContext { return s.update_expr }

func (s *Upsert_clauseContext) SetTarget_expr(v IExprContext) { s.target_expr = v }

func (s *Upsert_clauseContext) SetUpdate_expr(v IExprContext) { s.update_expr = v }

func (s *Upsert_clauseContext) ON_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserON_, 0)
}

func (s *Upsert_clauseContext) CONFLICT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCONFLICT_, 0)
}

func (s *Upsert_clauseContext) DO_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDO_, 0)
}

func (s *Upsert_clauseContext) NOTHING_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNOTHING_, 0)
}

func (s *Upsert_clauseContext) UPDATE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserUPDATE_, 0)
}

func (s *Upsert_clauseContext) SET_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserSET_, 0)
}

func (s *Upsert_clauseContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *Upsert_clauseContext) AllIndexed_column() []IIndexed_columnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexed_columnContext); ok {
			len++
		}
	}

	tst := make([]IIndexed_columnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexed_columnContext); ok {
			tst[i] = t.(IIndexed_columnContext)
			i++
		}
	}

	return tst
}

func (s *Upsert_clauseContext) Indexed_column(i int) IIndexed_columnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexed_columnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexed_columnContext)
}

func (s *Upsert_clauseContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *Upsert_clauseContext) AllUpsert_update() []IUpsert_updateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUpsert_updateContext); ok {
			len++
		}
	}

	tst := make([]IUpsert_updateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUpsert_updateContext); ok {
			tst[i] = t.(IUpsert_updateContext)
			i++
		}
	}

	return tst
}

func (s *Upsert_clauseContext) Upsert_update(i int) IUpsert_updateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpsert_updateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpsert_updateContext)
}

func (s *Upsert_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Upsert_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Upsert_clauseContext) AllWHERE_() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserWHERE_)
}

func (s *Upsert_clauseContext) WHERE_(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserWHERE_, i)
}

func (s *Upsert_clauseContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Upsert_clauseContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Upsert_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upsert_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upsert_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitUpsert_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Upsert_clause() (localctx IUpsert_clauseContext) {
	this := p
	_ = this

	localctx = NewUpsert_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SQLiteParserRULE_upsert_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.Match(SQLiteParserON_)
	}
	{
		p.SetState(417)
		p.Match(SQLiteParserCONFLICT_)
	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOPEN_PAR {
		{
			p.SetState(418)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(419)
			p.Indexed_column()
		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(420)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(421)
				p.Indexed_column()
			}

			p.SetState(426)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(427)
			p.Match(SQLiteParserCLOSE_PAR)
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserWHERE_ {
			{
				p.SetState(428)
				p.Match(SQLiteParserWHERE_)
			}
			{
				p.SetState(429)

				var _x = p.expr(0)

				localctx.(*Upsert_clauseContext).target_expr = _x
			}

		}

	}
	{
		p.SetState(434)
		p.Match(SQLiteParserDO_)
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserNOTHING_:
		{
			p.SetState(435)
			p.Match(SQLiteParserNOTHING_)
		}

	case SQLiteParserUPDATE_:
		{
			p.SetState(436)
			p.Match(SQLiteParserUPDATE_)
		}
		{
			p.SetState(437)
			p.Match(SQLiteParserSET_)
		}

		{
			p.SetState(438)
			p.Upsert_update()
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(439)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(440)
				p.Upsert_update()
			}

			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserWHERE_ {
			{
				p.SetState(446)
				p.Match(SQLiteParserWHERE_)
			}
			{
				p.SetState(447)

				var _x = p.expr(0)

				localctx.(*Upsert_clauseContext).update_expr = _x
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelect_stmt_coreContext is an interface to support dynamic dispatch.
type ISelect_stmt_coreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSelect_core() []ISelect_coreContext
	Select_core(i int) ISelect_coreContext
	AllCompound_operator() []ICompound_operatorContext
	Compound_operator(i int) ICompound_operatorContext
	Order_by_stmt() IOrder_by_stmtContext
	Limit_stmt() ILimit_stmtContext

	// IsSelect_stmt_coreContext differentiates from other interfaces.
	IsSelect_stmt_coreContext()
}

type Select_stmt_coreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_stmt_coreContext() *Select_stmt_coreContext {
	var p = new(Select_stmt_coreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_select_stmt_core
	return p
}

func (*Select_stmt_coreContext) IsSelect_stmt_coreContext() {}

func NewSelect_stmt_coreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_stmt_coreContext {
	var p = new(Select_stmt_coreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_select_stmt_core

	return p
}

func (s *Select_stmt_coreContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_stmt_coreContext) AllSelect_core() []ISelect_coreContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelect_coreContext); ok {
			len++
		}
	}

	tst := make([]ISelect_coreContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelect_coreContext); ok {
			tst[i] = t.(ISelect_coreContext)
			i++
		}
	}

	return tst
}

func (s *Select_stmt_coreContext) Select_core(i int) ISelect_coreContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_coreContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_coreContext)
}

func (s *Select_stmt_coreContext) AllCompound_operator() []ICompound_operatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICompound_operatorContext); ok {
			len++
		}
	}

	tst := make([]ICompound_operatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICompound_operatorContext); ok {
			tst[i] = t.(ICompound_operatorContext)
			i++
		}
	}

	return tst
}

func (s *Select_stmt_coreContext) Compound_operator(i int) ICompound_operatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompound_operatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompound_operatorContext)
}

func (s *Select_stmt_coreContext) Order_by_stmt() IOrder_by_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrder_by_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrder_by_stmtContext)
}

func (s *Select_stmt_coreContext) Limit_stmt() ILimit_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimit_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimit_stmtContext)
}

func (s *Select_stmt_coreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_stmt_coreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_stmt_coreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitSelect_stmt_core(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Select_stmt_core() (localctx ISelect_stmt_coreContext) {
	this := p
	_ = this

	localctx = NewSelect_stmt_coreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SQLiteParserRULE_select_stmt_core)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Select_core()
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-48)) & ^0x3f) == 0 && ((int64(1)<<(_la-48))&562949953486849) != 0 {
		{
			p.SetState(453)
			p.Compound_operator()
		}
		{
			p.SetState(454)
			p.Select_core()
		}

		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserORDER_ {
		{
			p.SetState(461)
			p.Order_by_stmt()
		}

	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserLIMIT_ {
		{
			p.SetState(464)
			p.Limit_stmt()
		}

	}

	return localctx
}

// ISelect_stmtContext is an interface to support dynamic dispatch.
type ISelect_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_stmt_core() ISelect_stmt_coreContext
	Common_table_stmt() ICommon_table_stmtContext

	// IsSelect_stmtContext differentiates from other interfaces.
	IsSelect_stmtContext()
}

type Select_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_stmtContext() *Select_stmtContext {
	var p = new(Select_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_select_stmt
	return p
}

func (*Select_stmtContext) IsSelect_stmtContext() {}

func NewSelect_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_stmtContext {
	var p = new(Select_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_select_stmt

	return p
}

func (s *Select_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_stmtContext) Select_stmt_core() ISelect_stmt_coreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_stmt_coreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_stmt_coreContext)
}

func (s *Select_stmtContext) Common_table_stmt() ICommon_table_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommon_table_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommon_table_stmtContext)
}

func (s *Select_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitSelect_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Select_stmt() (localctx ISelect_stmtContext) {
	this := p
	_ = this

	localctx = NewSelect_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SQLiteParserRULE_select_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(467)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(470)
		p.Select_stmt_core()
	}

	return localctx
}

// IJoin_clauseContext is an interface to support dynamic dispatch.
type IJoin_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTable_or_subquery() []ITable_or_subqueryContext
	Table_or_subquery(i int) ITable_or_subqueryContext
	AllJoin_operator() []IJoin_operatorContext
	Join_operator(i int) IJoin_operatorContext
	AllJoin_constraint() []IJoin_constraintContext
	Join_constraint(i int) IJoin_constraintContext

	// IsJoin_clauseContext differentiates from other interfaces.
	IsJoin_clauseContext()
}

type Join_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_clauseContext() *Join_clauseContext {
	var p = new(Join_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_join_clause
	return p
}

func (*Join_clauseContext) IsJoin_clauseContext() {}

func NewJoin_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_clauseContext {
	var p = new(Join_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_join_clause

	return p
}

func (s *Join_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Join_clauseContext) AllTable_or_subquery() []ITable_or_subqueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITable_or_subqueryContext); ok {
			len++
		}
	}

	tst := make([]ITable_or_subqueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITable_or_subqueryContext); ok {
			tst[i] = t.(ITable_or_subqueryContext)
			i++
		}
	}

	return tst
}

func (s *Join_clauseContext) Table_or_subquery(i int) ITable_or_subqueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_or_subqueryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_or_subqueryContext)
}

func (s *Join_clauseContext) AllJoin_operator() []IJoin_operatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoin_operatorContext); ok {
			len++
		}
	}

	tst := make([]IJoin_operatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoin_operatorContext); ok {
			tst[i] = t.(IJoin_operatorContext)
			i++
		}
	}

	return tst
}

func (s *Join_clauseContext) Join_operator(i int) IJoin_operatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoin_operatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoin_operatorContext)
}

func (s *Join_clauseContext) AllJoin_constraint() []IJoin_constraintContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoin_constraintContext); ok {
			len++
		}
	}

	tst := make([]IJoin_constraintContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoin_constraintContext); ok {
			tst[i] = t.(IJoin_constraintContext)
			i++
		}
	}

	return tst
}

func (s *Join_clauseContext) Join_constraint(i int) IJoin_constraintContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoin_constraintContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoin_constraintContext)
}

func (s *Join_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Join_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Join_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitJoin_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Join_clause() (localctx IJoin_clauseContext) {
	this := p
	_ = this

	localctx = NewJoin_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SQLiteParserRULE_join_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.Table_or_subquery()
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&68720607361) != 0 {
		{
			p.SetState(473)
			p.Join_operator()
		}
		{
			p.SetState(474)
			p.Table_or_subquery()
		}
		{
			p.SetState(475)
			p.Join_constraint()
		}

		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelect_coreContext is an interface to support dynamic dispatch.
type ISelect_coreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWhereExpr returns the whereExpr rule contexts.
	GetWhereExpr() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetHavingExpr returns the havingExpr rule contexts.
	GetHavingExpr() IExprContext

	// SetWhereExpr sets the whereExpr rule contexts.
	SetWhereExpr(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetHavingExpr sets the havingExpr rule contexts.
	SetHavingExpr(IExprContext)

	// GetGroupByExpr returns the groupByExpr rule context list.
	GetGroupByExpr() []IExprContext

	// SetGroupByExpr sets the groupByExpr rule context list.
	SetGroupByExpr([]IExprContext)

	// Getter signatures
	SELECT_() antlr.TerminalNode
	AllResult_column() []IResult_columnContext
	Result_column(i int) IResult_columnContext
	DISTINCT_() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	FROM_() antlr.TerminalNode
	WHERE_() antlr.TerminalNode
	GROUP_() antlr.TerminalNode
	BY_() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Table_or_subquery() ITable_or_subqueryContext
	Join_clause() IJoin_clauseContext
	HAVING_() antlr.TerminalNode

	// IsSelect_coreContext differentiates from other interfaces.
	IsSelect_coreContext()
}

type Select_coreContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	whereExpr   IExprContext
	_expr       IExprContext
	groupByExpr []IExprContext
	havingExpr  IExprContext
}

func NewEmptySelect_coreContext() *Select_coreContext {
	var p = new(Select_coreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_select_core
	return p
}

func (*Select_coreContext) IsSelect_coreContext() {}

func NewSelect_coreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_coreContext {
	var p = new(Select_coreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_select_core

	return p
}

func (s *Select_coreContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_coreContext) GetWhereExpr() IExprContext { return s.whereExpr }

func (s *Select_coreContext) Get_expr() IExprContext { return s._expr }

func (s *Select_coreContext) GetHavingExpr() IExprContext { return s.havingExpr }

func (s *Select_coreContext) SetWhereExpr(v IExprContext) { s.whereExpr = v }

func (s *Select_coreContext) Set_expr(v IExprContext) { s._expr = v }

func (s *Select_coreContext) SetHavingExpr(v IExprContext) { s.havingExpr = v }

func (s *Select_coreContext) GetGroupByExpr() []IExprContext { return s.groupByExpr }

func (s *Select_coreContext) SetGroupByExpr(v []IExprContext) { s.groupByExpr = v }

func (s *Select_coreContext) SELECT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserSELECT_, 0)
}

func (s *Select_coreContext) AllResult_column() []IResult_columnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResult_columnContext); ok {
			len++
		}
	}

	tst := make([]IResult_columnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResult_columnContext); ok {
			tst[i] = t.(IResult_columnContext)
			i++
		}
	}

	return tst
}

func (s *Select_coreContext) Result_column(i int) IResult_columnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResult_columnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResult_columnContext)
}

func (s *Select_coreContext) DISTINCT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDISTINCT_, 0)
}

func (s *Select_coreContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Select_coreContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Select_coreContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFROM_, 0)
}

func (s *Select_coreContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserWHERE_, 0)
}

func (s *Select_coreContext) GROUP_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserGROUP_, 0)
}

func (s *Select_coreContext) BY_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserBY_, 0)
}

func (s *Select_coreContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Select_coreContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Select_coreContext) Table_or_subquery() ITable_or_subqueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_or_subqueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_or_subqueryContext)
}

func (s *Select_coreContext) Join_clause() IJoin_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoin_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoin_clauseContext)
}

func (s *Select_coreContext) HAVING_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserHAVING_, 0)
}

func (s *Select_coreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_coreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_coreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitSelect_core(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Select_core() (localctx ISelect_coreContext) {
	this := p
	_ = this

	localctx = NewSelect_coreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SQLiteParserRULE_select_core)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(482)
		p.Match(SQLiteParserSELECT_)
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserDISTINCT_ {
		{
			p.SetState(483)
			p.Match(SQLiteParserDISTINCT_)
		}

	}
	{
		p.SetState(486)
		p.Result_column()
	}
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(487)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(488)
			p.Result_column()
		}

		p.SetState(493)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserFROM_ {
		{
			p.SetState(494)
			p.Match(SQLiteParserFROM_)
		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(495)
				p.Table_or_subquery()
			}

		case 2:
			{
				p.SetState(496)
				p.Join_clause()
			}

		}

	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWHERE_ {
		{
			p.SetState(501)
			p.Match(SQLiteParserWHERE_)
		}
		{
			p.SetState(502)

			var _x = p.expr(0)

			localctx.(*Select_coreContext).whereExpr = _x
		}

	}
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserGROUP_ {
		{
			p.SetState(505)
			p.Match(SQLiteParserGROUP_)
		}
		{
			p.SetState(506)
			p.Match(SQLiteParserBY_)
		}
		{
			p.SetState(507)

			var _x = p.expr(0)

			localctx.(*Select_coreContext)._expr = _x
		}
		localctx.(*Select_coreContext).groupByExpr = append(localctx.(*Select_coreContext).groupByExpr, localctx.(*Select_coreContext)._expr)
		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(508)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(509)

				var _x = p.expr(0)

				localctx.(*Select_coreContext)._expr = _x
			}
			localctx.(*Select_coreContext).groupByExpr = append(localctx.(*Select_coreContext).groupByExpr, localctx.(*Select_coreContext)._expr)

			p.SetState(514)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(517)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserHAVING_ {
			{
				p.SetState(515)
				p.Match(SQLiteParserHAVING_)
			}
			{
				p.SetState(516)

				var _x = p.expr(0)

				localctx.(*Select_coreContext).havingExpr = _x
			}

		}

	}

	return localctx
}

// ITable_or_subqueryContext is an interface to support dynamic dispatch.
type ITable_or_subqueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Table_name() ITable_nameContext
	AS_() antlr.TerminalNode
	Table_alias() ITable_aliasContext
	OPEN_PAR() antlr.TerminalNode
	Select_stmt_core() ISelect_stmt_coreContext
	CLOSE_PAR() antlr.TerminalNode

	// IsTable_or_subqueryContext differentiates from other interfaces.
	IsTable_or_subqueryContext()
}

type Table_or_subqueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_or_subqueryContext() *Table_or_subqueryContext {
	var p = new(Table_or_subqueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_table_or_subquery
	return p
}

func (*Table_or_subqueryContext) IsTable_or_subqueryContext() {}

func NewTable_or_subqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_or_subqueryContext {
	var p = new(Table_or_subqueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_table_or_subquery

	return p
}

func (s *Table_or_subqueryContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_or_subqueryContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Table_or_subqueryContext) AS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAS_, 0)
}

func (s *Table_or_subqueryContext) Table_alias() ITable_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_aliasContext)
}

func (s *Table_or_subqueryContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *Table_or_subqueryContext) Select_stmt_core() ISelect_stmt_coreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelect_stmt_coreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelect_stmt_coreContext)
}

func (s *Table_or_subqueryContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *Table_or_subqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_or_subqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_or_subqueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitTable_or_subquery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Table_or_subquery() (localctx ITable_or_subqueryContext) {
	this := p
	_ = this

	localctx = NewTable_or_subqueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SQLiteParserRULE_table_or_subquery)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(533)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)
			p.Table_name()
		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(522)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(523)
				p.Table_alias()
			}

		}

	case SQLiteParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(526)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(527)
			p.Select_stmt_core()
		}
		{
			p.SetState(528)
			p.Match(SQLiteParserCLOSE_PAR)
		}
		p.SetState(531)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(529)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(530)
				p.Table_alias()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IResult_columnContext is an interface to support dynamic dispatch.
type IResult_columnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	Table_name() ITable_nameContext
	DOT() antlr.TerminalNode
	Expr() IExprContext
	AS_() antlr.TerminalNode
	Column_alias() IColumn_aliasContext

	// IsResult_columnContext differentiates from other interfaces.
	IsResult_columnContext()
}

type Result_columnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResult_columnContext() *Result_columnContext {
	var p = new(Result_columnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_result_column
	return p
}

func (*Result_columnContext) IsResult_columnContext() {}

func NewResult_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Result_columnContext {
	var p = new(Result_columnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_result_column

	return p
}

func (s *Result_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Result_columnContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserSTAR, 0)
}

func (s *Result_columnContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Result_columnContext) DOT() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDOT, 0)
}

func (s *Result_columnContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Result_columnContext) AS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAS_, 0)
}

func (s *Result_columnContext) Column_alias() IColumn_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_aliasContext)
}

func (s *Result_columnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Result_columnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Result_columnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitResult_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Result_column() (localctx IResult_columnContext) {
	this := p
	_ = this

	localctx = NewResult_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SQLiteParserRULE_result_column)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(535)
			p.Match(SQLiteParserSTAR)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(536)
			p.Table_name()
		}
		{
			p.SetState(537)
			p.Match(SQLiteParserDOT)
		}
		{
			p.SetState(538)
			p.Match(SQLiteParserSTAR)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(540)
			p.expr(0)
		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(541)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(542)
				p.Column_alias()
			}

		}

	}

	return localctx
}

// IReturning_clause_result_columnContext is an interface to support dynamic dispatch.
type IReturning_clause_result_columnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	Expr() IExprContext
	AS_() antlr.TerminalNode
	Column_alias() IColumn_aliasContext

	// IsReturning_clause_result_columnContext differentiates from other interfaces.
	IsReturning_clause_result_columnContext()
}

type Returning_clause_result_columnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturning_clause_result_columnContext() *Returning_clause_result_columnContext {
	var p = new(Returning_clause_result_columnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_returning_clause_result_column
	return p
}

func (*Returning_clause_result_columnContext) IsReturning_clause_result_columnContext() {}

func NewReturning_clause_result_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Returning_clause_result_columnContext {
	var p = new(Returning_clause_result_columnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_returning_clause_result_column

	return p
}

func (s *Returning_clause_result_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Returning_clause_result_columnContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserSTAR, 0)
}

func (s *Returning_clause_result_columnContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Returning_clause_result_columnContext) AS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAS_, 0)
}

func (s *Returning_clause_result_columnContext) Column_alias() IColumn_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_aliasContext)
}

func (s *Returning_clause_result_columnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Returning_clause_result_columnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Returning_clause_result_columnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitReturning_clause_result_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Returning_clause_result_column() (localctx IReturning_clause_result_columnContext) {
	this := p
	_ = this

	localctx = NewReturning_clause_result_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SQLiteParserRULE_returning_clause_result_column)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(553)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(547)
			p.Match(SQLiteParserSTAR)
		}

	case SQLiteParserOPEN_PAR, SQLiteParserPLUS, SQLiteParserMINUS, SQLiteParserTILDE, SQLiteParserCASE_, SQLiteParserEXISTS_, SQLiteParserFALSE_, SQLiteParserGLOB_, SQLiteParserLIKE_, SQLiteParserNOT_, SQLiteParserNULL_, SQLiteParserREPLACE_, SQLiteParserTRUE_, SQLiteParserIDENTIFIER, SQLiteParserNUMERIC_LITERAL, SQLiteParserBIND_PARAMETER, SQLiteParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(548)
			p.expr(0)
		}
		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(549)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(550)
				p.Column_alias()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJoin_operatorContext is an interface to support dynamic dispatch.
type IJoin_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JOIN_() antlr.TerminalNode
	NATURAL_() antlr.TerminalNode
	INNER_() antlr.TerminalNode
	LEFT_() antlr.TerminalNode
	RIGHT_() antlr.TerminalNode
	FULL_() antlr.TerminalNode
	OUTER_() antlr.TerminalNode

	// IsJoin_operatorContext differentiates from other interfaces.
	IsJoin_operatorContext()
}

type Join_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_operatorContext() *Join_operatorContext {
	var p = new(Join_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_join_operator
	return p
}

func (*Join_operatorContext) IsJoin_operatorContext() {}

func NewJoin_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_operatorContext {
	var p = new(Join_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_join_operator

	return p
}

func (s *Join_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Join_operatorContext) JOIN_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserJOIN_, 0)
}

func (s *Join_operatorContext) NATURAL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNATURAL_, 0)
}

func (s *Join_operatorContext) INNER_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserINNER_, 0)
}

func (s *Join_operatorContext) LEFT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLEFT_, 0)
}

func (s *Join_operatorContext) RIGHT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserRIGHT_, 0)
}

func (s *Join_operatorContext) FULL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFULL_, 0)
}

func (s *Join_operatorContext) OUTER_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOUTER_, 0)
}

func (s *Join_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Join_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Join_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitJoin_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Join_operator() (localctx IJoin_operatorContext) {
	this := p
	_ = this

	localctx = NewJoin_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SQLiteParserRULE_join_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserNATURAL_ {
		{
			p.SetState(555)
			p.Match(SQLiteParserNATURAL_)
		}

	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserFULL_, SQLiteParserLEFT_, SQLiteParserRIGHT_:
		{
			p.SetState(558)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&68719542273) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserOUTER_ {
			{
				p.SetState(559)
				p.Match(SQLiteParserOUTER_)
			}

		}

	case SQLiteParserINNER_:
		{
			p.SetState(562)
			p.Match(SQLiteParserINNER_)
		}

	case SQLiteParserJOIN_:

	default:
	}
	{
		p.SetState(565)
		p.Match(SQLiteParserJOIN_)
	}

	return localctx
}

// IJoin_constraintContext is an interface to support dynamic dispatch.
type IJoin_constraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ON_() antlr.TerminalNode
	Expr() IExprContext

	// IsJoin_constraintContext differentiates from other interfaces.
	IsJoin_constraintContext()
}

type Join_constraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_constraintContext() *Join_constraintContext {
	var p = new(Join_constraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_join_constraint
	return p
}

func (*Join_constraintContext) IsJoin_constraintContext() {}

func NewJoin_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_constraintContext {
	var p = new(Join_constraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_join_constraint

	return p
}

func (s *Join_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *Join_constraintContext) ON_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserON_, 0)
}

func (s *Join_constraintContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Join_constraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Join_constraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Join_constraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitJoin_constraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Join_constraint() (localctx IJoin_constraintContext) {
	this := p
	_ = this

	localctx = NewJoin_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SQLiteParserRULE_join_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		p.Match(SQLiteParserON_)
	}
	{
		p.SetState(568)
		p.expr(0)
	}

	return localctx
}

// ICompound_operatorContext is an interface to support dynamic dispatch.
type ICompound_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNION_() antlr.TerminalNode
	ALL_() antlr.TerminalNode
	INTERSECT_() antlr.TerminalNode
	EXCEPT_() antlr.TerminalNode

	// IsCompound_operatorContext differentiates from other interfaces.
	IsCompound_operatorContext()
}

type Compound_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_operatorContext() *Compound_operatorContext {
	var p = new(Compound_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_compound_operator
	return p
}

func (*Compound_operatorContext) IsCompound_operatorContext() {}

func NewCompound_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_operatorContext {
	var p = new(Compound_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_compound_operator

	return p
}

func (s *Compound_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_operatorContext) UNION_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserUNION_, 0)
}

func (s *Compound_operatorContext) ALL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserALL_, 0)
}

func (s *Compound_operatorContext) INTERSECT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserINTERSECT_, 0)
}

func (s *Compound_operatorContext) EXCEPT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserEXCEPT_, 0)
}

func (s *Compound_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitCompound_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Compound_operator() (localctx ICompound_operatorContext) {
	this := p
	_ = this

	localctx = NewCompound_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SQLiteParserRULE_compound_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(576)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserUNION_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(570)
			p.Match(SQLiteParserUNION_)
		}
		p.SetState(572)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserALL_ {
			{
				p.SetState(571)
				p.Match(SQLiteParserALL_)
			}

		}

	case SQLiteParserINTERSECT_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(574)
			p.Match(SQLiteParserINTERSECT_)
		}

	case SQLiteParserEXCEPT_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(575)
			p.Match(SQLiteParserEXCEPT_)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUpdate_set_subclauseContext is an interface to support dynamic dispatch.
type IUpdate_set_subclauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	Column_name() IColumn_nameContext
	Column_name_list() IColumn_name_listContext

	// IsUpdate_set_subclauseContext differentiates from other interfaces.
	IsUpdate_set_subclauseContext()
}

type Update_set_subclauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_set_subclauseContext() *Update_set_subclauseContext {
	var p = new(Update_set_subclauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_update_set_subclause
	return p
}

func (*Update_set_subclauseContext) IsUpdate_set_subclauseContext() {}

func NewUpdate_set_subclauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_set_subclauseContext {
	var p = new(Update_set_subclauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_update_set_subclause

	return p
}

func (s *Update_set_subclauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_set_subclauseContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SQLiteParserASSIGN, 0)
}

func (s *Update_set_subclauseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Update_set_subclauseContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Update_set_subclauseContext) Column_name_list() IColumn_name_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_name_listContext)
}

func (s *Update_set_subclauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_set_subclauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_set_subclauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitUpdate_set_subclause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Update_set_subclause() (localctx IUpdate_set_subclauseContext) {
	this := p
	_ = this

	localctx = NewUpdate_set_subclauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SQLiteParserRULE_update_set_subclause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(580)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		{
			p.SetState(578)
			p.Column_name()
		}

	case SQLiteParserOPEN_PAR:
		{
			p.SetState(579)
			p.Column_name_list()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(582)
		p.Match(SQLiteParserASSIGN)
	}
	{
		p.SetState(583)
		p.expr(0)
	}

	return localctx
}

// IUpdate_stmtContext is an interface to support dynamic dispatch.
type IUpdate_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UPDATE_() antlr.TerminalNode
	Qualified_table_name() IQualified_table_nameContext
	SET_() antlr.TerminalNode
	AllUpdate_set_subclause() []IUpdate_set_subclauseContext
	Update_set_subclause(i int) IUpdate_set_subclauseContext
	Common_table_stmt() ICommon_table_stmtContext
	OR_() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	FROM_() antlr.TerminalNode
	WHERE_() antlr.TerminalNode
	Expr() IExprContext
	Returning_clause() IReturning_clauseContext
	ROLLBACK_() antlr.TerminalNode
	ABORT_() antlr.TerminalNode
	REPLACE_() antlr.TerminalNode
	FAIL_() antlr.TerminalNode
	IGNORE_() antlr.TerminalNode
	Table_or_subquery() ITable_or_subqueryContext
	Join_clause() IJoin_clauseContext

	// IsUpdate_stmtContext differentiates from other interfaces.
	IsUpdate_stmtContext()
}

type Update_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_stmtContext() *Update_stmtContext {
	var p = new(Update_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_update_stmt
	return p
}

func (*Update_stmtContext) IsUpdate_stmtContext() {}

func NewUpdate_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_stmtContext {
	var p = new(Update_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_update_stmt

	return p
}

func (s *Update_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_stmtContext) UPDATE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserUPDATE_, 0)
}

func (s *Update_stmtContext) Qualified_table_name() IQualified_table_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualified_table_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualified_table_nameContext)
}

func (s *Update_stmtContext) SET_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserSET_, 0)
}

func (s *Update_stmtContext) AllUpdate_set_subclause() []IUpdate_set_subclauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUpdate_set_subclauseContext); ok {
			len++
		}
	}

	tst := make([]IUpdate_set_subclauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUpdate_set_subclauseContext); ok {
			tst[i] = t.(IUpdate_set_subclauseContext)
			i++
		}
	}

	return tst
}

func (s *Update_stmtContext) Update_set_subclause(i int) IUpdate_set_subclauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdate_set_subclauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdate_set_subclauseContext)
}

func (s *Update_stmtContext) Common_table_stmt() ICommon_table_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommon_table_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommon_table_stmtContext)
}

func (s *Update_stmtContext) OR_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOR_, 0)
}

func (s *Update_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Update_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Update_stmtContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFROM_, 0)
}

func (s *Update_stmtContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserWHERE_, 0)
}

func (s *Update_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Update_stmtContext) Returning_clause() IReturning_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturning_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturning_clauseContext)
}

func (s *Update_stmtContext) ROLLBACK_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserROLLBACK_, 0)
}

func (s *Update_stmtContext) ABORT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserABORT_, 0)
}

func (s *Update_stmtContext) REPLACE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserREPLACE_, 0)
}

func (s *Update_stmtContext) FAIL_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFAIL_, 0)
}

func (s *Update_stmtContext) IGNORE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIGNORE_, 0)
}

func (s *Update_stmtContext) Table_or_subquery() ITable_or_subqueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_or_subqueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_or_subqueryContext)
}

func (s *Update_stmtContext) Join_clause() IJoin_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoin_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoin_clauseContext)
}

func (s *Update_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitUpdate_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Update_stmt() (localctx IUpdate_stmtContext) {
	this := p
	_ = this

	localctx = NewUpdate_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SQLiteParserRULE_update_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(585)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(588)
		p.Match(SQLiteParserUPDATE_)
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOR_ {
		{
			p.SetState(589)
			p.Match(SQLiteParserOR_)
		}
		{
			p.SetState(590)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1154047404580798464) != 0) || _la == SQLiteParserREPLACE_ || _la == SQLiteParserROLLBACK_) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(593)
		p.Qualified_table_name()
	}
	{
		p.SetState(594)
		p.Match(SQLiteParserSET_)
	}
	{
		p.SetState(595)
		p.Update_set_subclause()
	}
	p.SetState(600)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(596)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(597)
			p.Update_set_subclause()
		}

		p.SetState(602)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(608)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserFROM_ {
		{
			p.SetState(603)
			p.Match(SQLiteParserFROM_)
		}
		p.SetState(606)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 84, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(604)
				p.Table_or_subquery()
			}

		case 2:
			{
				p.SetState(605)
				p.Join_clause()
			}

		}

	}
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWHERE_ {
		{
			p.SetState(610)
			p.Match(SQLiteParserWHERE_)
		}
		{
			p.SetState(611)
			p.expr(0)
		}

	}
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserRETURNING_ {
		{
			p.SetState(614)
			p.Returning_clause()
		}

	}

	return localctx
}

// IColumn_name_listContext is an interface to support dynamic dispatch.
type IColumn_name_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAR() antlr.TerminalNode
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext
	CLOSE_PAR() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumn_name_listContext differentiates from other interfaces.
	IsColumn_name_listContext()
}

type Column_name_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_name_listContext() *Column_name_listContext {
	var p = new(Column_name_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_column_name_list
	return p
}

func (*Column_name_listContext) IsColumn_name_listContext() {}

func NewColumn_name_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_name_listContext {
	var p = new(Column_name_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_column_name_list

	return p
}

func (s *Column_name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_name_listContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOPEN_PAR, 0)
}

func (s *Column_name_listContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Column_name_listContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_name_listContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCLOSE_PAR, 0)
}

func (s *Column_name_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Column_name_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Column_name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_name_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitColumn_name_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Column_name_list() (localctx IColumn_name_listContext) {
	this := p
	_ = this

	localctx = NewColumn_name_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SQLiteParserRULE_column_name_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
		p.Match(SQLiteParserOPEN_PAR)
	}
	{
		p.SetState(618)
		p.Column_name()
	}
	p.SetState(623)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(619)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(620)
			p.Column_name()
		}

		p.SetState(625)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(626)
		p.Match(SQLiteParserCLOSE_PAR)
	}

	return localctx
}

// IQualified_table_nameContext is an interface to support dynamic dispatch.
type IQualified_table_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Table_name() ITable_nameContext
	AS_() antlr.TerminalNode
	Table_alias() ITable_aliasContext
	INDEXED_() antlr.TerminalNode
	BY_() antlr.TerminalNode
	Index_name() IIndex_nameContext
	NOT_() antlr.TerminalNode

	// IsQualified_table_nameContext differentiates from other interfaces.
	IsQualified_table_nameContext()
}

type Qualified_table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualified_table_nameContext() *Qualified_table_nameContext {
	var p = new(Qualified_table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_qualified_table_name
	return p
}

func (*Qualified_table_nameContext) IsQualified_table_nameContext() {}

func NewQualified_table_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Qualified_table_nameContext {
	var p = new(Qualified_table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_qualified_table_name

	return p
}

func (s *Qualified_table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Qualified_table_nameContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Qualified_table_nameContext) AS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserAS_, 0)
}

func (s *Qualified_table_nameContext) Table_alias() ITable_aliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_aliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_aliasContext)
}

func (s *Qualified_table_nameContext) INDEXED_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserINDEXED_, 0)
}

func (s *Qualified_table_nameContext) BY_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserBY_, 0)
}

func (s *Qualified_table_nameContext) Index_name() IIndex_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_nameContext)
}

func (s *Qualified_table_nameContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNOT_, 0)
}

func (s *Qualified_table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qualified_table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Qualified_table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitQualified_table_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Qualified_table_name() (localctx IQualified_table_nameContext) {
	this := p
	_ = this

	localctx = NewQualified_table_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SQLiteParserRULE_qualified_table_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(628)
		p.Table_name()
	}
	p.SetState(631)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserAS_ {
		{
			p.SetState(629)
			p.Match(SQLiteParserAS_)
		}
		{
			p.SetState(630)
			p.Table_alias()
		}

	}
	p.SetState(638)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserINDEXED_:
		{
			p.SetState(633)
			p.Match(SQLiteParserINDEXED_)
		}
		{
			p.SetState(634)
			p.Match(SQLiteParserBY_)
		}
		{
			p.SetState(635)
			p.Index_name()
		}

	case SQLiteParserNOT_:
		{
			p.SetState(636)
			p.Match(SQLiteParserNOT_)
		}
		{
			p.SetState(637)
			p.Match(SQLiteParserINDEXED_)
		}

	case SQLiteParserEOF, SQLiteParserSCOL, SQLiteParserDELETE_, SQLiteParserINSERT_, SQLiteParserREPLACE_, SQLiteParserRETURNING_, SQLiteParserSELECT_, SQLiteParserSET_, SQLiteParserUPDATE_, SQLiteParserWHERE_, SQLiteParserWITH_:

	default:
	}

	return localctx
}

// IOrder_by_stmtContext is an interface to support dynamic dispatch.
type IOrder_by_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER_() antlr.TerminalNode
	BY_() antlr.TerminalNode
	AllOrdering_term() []IOrdering_termContext
	Ordering_term(i int) IOrdering_termContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsOrder_by_stmtContext differentiates from other interfaces.
	IsOrder_by_stmtContext()
}

type Order_by_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_stmtContext() *Order_by_stmtContext {
	var p = new(Order_by_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_order_by_stmt
	return p
}

func (*Order_by_stmtContext) IsOrder_by_stmtContext() {}

func NewOrder_by_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_stmtContext {
	var p = new(Order_by_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_order_by_stmt

	return p
}

func (s *Order_by_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_stmtContext) ORDER_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserORDER_, 0)
}

func (s *Order_by_stmtContext) BY_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserBY_, 0)
}

func (s *Order_by_stmtContext) AllOrdering_term() []IOrdering_termContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrdering_termContext); ok {
			len++
		}
	}

	tst := make([]IOrdering_termContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrdering_termContext); ok {
			tst[i] = t.(IOrdering_termContext)
			i++
		}
	}

	return tst
}

func (s *Order_by_stmtContext) Ordering_term(i int) IOrdering_termContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrdering_termContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrdering_termContext)
}

func (s *Order_by_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLiteParserCOMMA)
}

func (s *Order_by_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, i)
}

func (s *Order_by_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitOrder_by_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Order_by_stmt() (localctx IOrder_by_stmtContext) {
	this := p
	_ = this

	localctx = NewOrder_by_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SQLiteParserRULE_order_by_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(640)
		p.Match(SQLiteParserORDER_)
	}
	{
		p.SetState(641)
		p.Match(SQLiteParserBY_)
	}
	{
		p.SetState(642)
		p.Ordering_term()
	}
	p.SetState(647)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(643)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(644)
			p.Ordering_term()
		}

		p.SetState(649)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILimit_stmtContext is an interface to support dynamic dispatch.
type ILimit_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT_() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	OFFSET_() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsLimit_stmtContext differentiates from other interfaces.
	IsLimit_stmtContext()
}

type Limit_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimit_stmtContext() *Limit_stmtContext {
	var p = new(Limit_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_limit_stmt
	return p
}

func (*Limit_stmtContext) IsLimit_stmtContext() {}

func NewLimit_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Limit_stmtContext {
	var p = new(Limit_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_limit_stmt

	return p
}

func (s *Limit_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Limit_stmtContext) LIMIT_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLIMIT_, 0)
}

func (s *Limit_stmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Limit_stmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Limit_stmtContext) OFFSET_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserOFFSET_, 0)
}

func (s *Limit_stmtContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOMMA, 0)
}

func (s *Limit_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Limit_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Limit_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitLimit_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Limit_stmt() (localctx ILimit_stmtContext) {
	this := p
	_ = this

	localctx = NewLimit_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SQLiteParserRULE_limit_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(650)
		p.Match(SQLiteParserLIMIT_)
	}
	{
		p.SetState(651)
		p.expr(0)
	}
	p.SetState(654)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserCOMMA || _la == SQLiteParserOFFSET_ {
		{
			p.SetState(652)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLiteParserCOMMA || _la == SQLiteParserOFFSET_) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(653)
			p.expr(0)
		}

	}

	return localctx
}

// IOrdering_termContext is an interface to support dynamic dispatch.
type IOrdering_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	COLLATE_() antlr.TerminalNode
	Collation_name() ICollation_nameContext
	Asc_desc() IAsc_descContext
	NULLS_() antlr.TerminalNode
	FIRST_() antlr.TerminalNode
	LAST_() antlr.TerminalNode

	// IsOrdering_termContext differentiates from other interfaces.
	IsOrdering_termContext()
}

type Ordering_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdering_termContext() *Ordering_termContext {
	var p = new(Ordering_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_ordering_term
	return p
}

func (*Ordering_termContext) IsOrdering_termContext() {}

func NewOrdering_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ordering_termContext {
	var p = new(Ordering_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_ordering_term

	return p
}

func (s *Ordering_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Ordering_termContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Ordering_termContext) COLLATE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserCOLLATE_, 0)
}

func (s *Ordering_termContext) Collation_name() ICollation_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollation_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollation_nameContext)
}

func (s *Ordering_termContext) Asc_desc() IAsc_descContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsc_descContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsc_descContext)
}

func (s *Ordering_termContext) NULLS_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserNULLS_, 0)
}

func (s *Ordering_termContext) FIRST_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserFIRST_, 0)
}

func (s *Ordering_termContext) LAST_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLAST_, 0)
}

func (s *Ordering_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ordering_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ordering_termContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitOrdering_term(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Ordering_term() (localctx IOrdering_termContext) {
	this := p
	_ = this

	localctx = NewOrdering_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SQLiteParserRULE_ordering_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(656)
		p.expr(0)
	}
	p.SetState(659)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserCOLLATE_ {
		{
			p.SetState(657)
			p.Match(SQLiteParserCOLLATE_)
		}
		{
			p.SetState(658)
			p.Collation_name()
		}

	}
	p.SetState(662)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserASC_ || _la == SQLiteParserDESC_ {
		{
			p.SetState(661)
			p.Asc_desc()
		}

	}
	p.SetState(666)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserNULLS_ {
		{
			p.SetState(664)
			p.Match(SQLiteParserNULLS_)
		}
		{
			p.SetState(665)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLiteParserFIRST_ || _la == SQLiteParserLAST_) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IAsc_descContext is an interface to support dynamic dispatch.
type IAsc_descContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASC_() antlr.TerminalNode
	DESC_() antlr.TerminalNode

	// IsAsc_descContext differentiates from other interfaces.
	IsAsc_descContext()
}

type Asc_descContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsc_descContext() *Asc_descContext {
	var p = new(Asc_descContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_asc_desc
	return p
}

func (*Asc_descContext) IsAsc_descContext() {}

func NewAsc_descContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asc_descContext {
	var p = new(Asc_descContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_asc_desc

	return p
}

func (s *Asc_descContext) GetParser() antlr.Parser { return s.parser }

func (s *Asc_descContext) ASC_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserASC_, 0)
}

func (s *Asc_descContext) DESC_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserDESC_, 0)
}

func (s *Asc_descContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asc_descContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Asc_descContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitAsc_desc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Asc_desc() (localctx IAsc_descContext) {
	this := p
	_ = this

	localctx = NewAsc_descContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SQLiteParserRULE_asc_desc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(668)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLiteParserASC_ || _la == SQLiteParserDESC_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_keywordContext is an interface to support dynamic dispatch.
type IFunction_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GLOB_() antlr.TerminalNode
	LIKE_() antlr.TerminalNode
	REPLACE_() antlr.TerminalNode

	// IsFunction_keywordContext differentiates from other interfaces.
	IsFunction_keywordContext()
}

type Function_keywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_keywordContext() *Function_keywordContext {
	var p = new(Function_keywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_function_keyword
	return p
}

func (*Function_keywordContext) IsFunction_keywordContext() {}

func NewFunction_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_keywordContext {
	var p = new(Function_keywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_function_keyword

	return p
}

func (s *Function_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_keywordContext) GLOB_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserGLOB_, 0)
}

func (s *Function_keywordContext) LIKE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserLIKE_, 0)
}

func (s *Function_keywordContext) REPLACE_() antlr.TerminalNode {
	return s.GetToken(SQLiteParserREPLACE_, 0)
}

func (s *Function_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitFunction_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Function_keyword() (localctx IFunction_keywordContext) {
	this := p
	_ = this

	localctx = NewFunction_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SQLiteParserRULE_function_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(670)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-56)) & ^0x3f) == 0 && ((int64(1)<<(_la-56))&8590000129) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Function_keyword() IFunction_keywordContext

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_function_name
	return p
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Function_nameContext) Function_keyword() IFunction_keywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_keywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_keywordContext)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitFunction_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Function_name() (localctx IFunction_nameContext) {
	this := p
	_ = this

	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SQLiteParserRULE_function_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(674)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(672)
			p.Match(SQLiteParserIDENTIFIER)
		}

	case SQLiteParserGLOB_, SQLiteParserLIKE_, SQLiteParserREPLACE_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(673)
			p.Function_keyword()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitTable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Table_name() (localctx ITable_nameContext) {
	this := p
	_ = this

	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SQLiteParserRULE_table_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(676)
		p.Match(SQLiteParserIDENTIFIER)
	}

	return localctx
}

// ITable_aliasContext is an interface to support dynamic dispatch.
type ITable_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTable_aliasContext differentiates from other interfaces.
	IsTable_aliasContext()
}

type Table_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_aliasContext() *Table_aliasContext {
	var p = new(Table_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_table_alias
	return p
}

func (*Table_aliasContext) IsTable_aliasContext() {}

func NewTable_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_aliasContext {
	var p = new(Table_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_table_alias

	return p
}

func (s *Table_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Table_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitTable_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Table_alias() (localctx ITable_aliasContext) {
	this := p
	_ = this

	localctx = NewTable_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SQLiteParserRULE_table_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(678)
		p.Match(SQLiteParserIDENTIFIER)
	}

	return localctx
}

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_column_name
	return p
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitColumn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Column_name() (localctx IColumn_nameContext) {
	this := p
	_ = this

	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SQLiteParserRULE_column_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(680)
		p.Match(SQLiteParserIDENTIFIER)
	}

	return localctx
}

// IColumn_aliasContext is an interface to support dynamic dispatch.
type IColumn_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsColumn_aliasContext differentiates from other interfaces.
	IsColumn_aliasContext()
}

type Column_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_aliasContext() *Column_aliasContext {
	var p = new(Column_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_column_alias
	return p
}

func (*Column_aliasContext) IsColumn_aliasContext() {}

func NewColumn_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_aliasContext {
	var p = new(Column_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_column_alias

	return p
}

func (s *Column_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Column_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitColumn_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Column_alias() (localctx IColumn_aliasContext) {
	this := p
	_ = this

	localctx = NewColumn_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SQLiteParserRULE_column_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(682)
		p.Match(SQLiteParserIDENTIFIER)
	}

	return localctx
}

// ICollation_nameContext is an interface to support dynamic dispatch.
type ICollation_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsCollation_nameContext differentiates from other interfaces.
	IsCollation_nameContext()
}

type Collation_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollation_nameContext() *Collation_nameContext {
	var p = new(Collation_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_collation_name
	return p
}

func (*Collation_nameContext) IsCollation_nameContext() {}

func NewCollation_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collation_nameContext {
	var p = new(Collation_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_collation_name

	return p
}

func (s *Collation_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Collation_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Collation_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collation_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collation_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitCollation_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Collation_name() (localctx ICollation_nameContext) {
	this := p
	_ = this

	localctx = NewCollation_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SQLiteParserRULE_collation_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(684)
		p.Match(SQLiteParserIDENTIFIER)
	}

	return localctx
}

// IIndex_nameContext is an interface to support dynamic dispatch.
type IIndex_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIndex_nameContext differentiates from other interfaces.
	IsIndex_nameContext()
}

type Index_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_nameContext() *Index_nameContext {
	var p = new(Index_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SQLiteParserRULE_index_name
	return p
}

func (*Index_nameContext) IsIndex_nameContext() {}

func NewIndex_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_nameContext {
	var p = new(Index_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLiteParserRULE_index_name

	return p
}

func (s *Index_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLiteParserIDENTIFIER, 0)
}

func (s *Index_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLiteParserVisitor:
		return t.VisitIndex_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLiteParser) Index_name() (localctx IIndex_nameContext) {
	this := p
	_ = this

	localctx = NewIndex_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SQLiteParserRULE_index_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(686)
		p.Match(SQLiteParserIDENTIFIER)
	}

	return localctx
}

func (p *SQLiteParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SQLiteParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
