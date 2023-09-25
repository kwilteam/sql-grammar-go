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
		"'>='", "'=='", "'!='", "'<>'", "'ABORT'", "'ADD'", "'ALL'", "'AND'",
		"'ASC'", "'AS'", "'BETWEEN'", "'BY'", "'CASE'", "'COLLATE'", "'COMMIT'",
		"'CONFLICT'", "'CREATE'", "'CROSS'", "'DEFAULT'", "'DELETE'", "'DESC'",
		"'DISTINCT'", "'DO'", "'ELSE'", "'END'", "'ESCAPE'", "'EXCEPT'", "'EXISTS'",
		"'FAIL'", "'FALSE'", "'FILTER'", "'FIRST'", "'FROM'", "'FULL'", "'GLOB'",
		"'GROUPS'", "'GROUP'", "'HAVING'", "'IGNORE'", "'INDEXED'", "'INNER'",
		"'INSERT'", "'INTERSECT'", "'INTO'", "'IN'", "'ISNULL'", "'IS'", "'JOIN'",
		"'LAST'", "'LEFT'", "'LIKE'", "'LIMIT'", "'MATCH'", "'NATURAL'", "'NOTHING'",
		"'NOTNULL'", "'NOT'", "'NULLS'", "'NULL'", "'OFFSET'", "'OF'", "'ON'",
		"'ORDER'", "'OR'", "'OUTER'", "'RAISE'", "'REGEXP'", "'REPLACE'", "'RETURNING'",
		"'RIGHT'", "'ROLLBACK'", "'SELECT'", "'SET'", "'THEN'", "'TRUE'", "'UNION'",
		"'UPDATE'", "'USING'", "'VALUES'", "'WHEN'", "'WHERE'", "'WITH'",
	}
	staticData.symbolicNames = []string{
		"", "SCOL", "DOT", "OPEN_PAR", "CLOSE_PAR", "COMMA", "ASSIGN", "STAR",
		"PLUS", "MINUS", "TILDE", "PIPE2", "DIV", "MOD", "LT2", "GT2", "AMP",
		"PIPE", "LT", "LT_EQ", "GT", "GT_EQ", "EQ", "NOT_EQ1", "NOT_EQ2", "ABORT_",
		"ADD_", "ALL_", "AND_", "ASC_", "AS_", "BETWEEN_", "BY_", "CASE_", "COLLATE_",
		"COMMIT_", "CONFLICT_", "CREATE_", "CROSS_", "DEFAULT_", "DELETE_",
		"DESC_", "DISTINCT_", "DO_", "ELSE_", "END_", "ESCAPE_", "EXCEPT_",
		"EXISTS_", "FAIL_", "FALSE_", "FILTER_", "FIRST_", "FROM_", "FULL_",
		"GLOB_", "GROUPS_", "GROUP_", "HAVING_", "IGNORE_", "INDEXED_", "INNER_",
		"INSERT_", "INTERSECT_", "INTO_", "IN_", "ISNULL_", "IS_", "JOIN_",
		"LAST_", "LEFT_", "LIKE_", "LIMIT_", "MATCH_", "NATURAL_", "NOTHING_",
		"NOTNULL_", "NOT_", "NULLS_", "NULL_", "OFFSET_", "OF_", "ON_", "ORDER_",
		"OR_", "OUTER_", "RAISE_", "REGEXP_", "REPLACE_", "RETURNING_", "RIGHT_",
		"ROLLBACK_", "SELECT_", "SET_", "THEN_", "TRUE_", "UNION_", "UPDATE_",
		"USING_", "VALUES_", "WHEN_", "WHERE_", "WITH_", "IDENTIFIER", "NUMERIC_LITERAL",
		"BIND_PARAMETER", "STRING_LITERAL", "BLOB_LITERAL", "SINGLE_LINE_COMMENT",
		"MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	}
	staticData.ruleNames = []string{
		"parse", "sql_stmt_list", "sql_stmt", "indexed_column", "cte_table_name",
		"common_table_expression", "common_table_stmt", "delete_stmt", "expr",
		"literal_value", "value_row", "values_clause", "insert_stmt", "returning_clause",
		"upsert_update", "upsert_clause", "select_stmt_core", "select_stmt",
		"join_clause", "select_core", "table_or_subquery", "result_column",
		"returning_clause_result_column", "join_operator", "join_constraint",
		"compound_operator", "update_set_subclause", "update_stmt", "column_name_list",
		"qualified_table_name", "order_by_stmt", "limit_stmt", "ordering_term",
		"asc_desc", "function_keyword", "function_name", "table_name", "table_alias",
		"column_name", "column_alias", "collation_name", "index_name",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 111, 667, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 1,
		0, 5, 0, 86, 8, 0, 10, 0, 12, 0, 89, 9, 0, 1, 0, 1, 0, 1, 1, 5, 1, 94,
		8, 1, 10, 1, 12, 1, 97, 9, 1, 1, 1, 1, 1, 4, 1, 101, 8, 1, 11, 1, 12, 1,
		102, 1, 1, 5, 1, 106, 8, 1, 10, 1, 12, 1, 109, 9, 1, 1, 1, 5, 1, 112, 8,
		1, 10, 1, 12, 1, 115, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 121, 8, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 130, 8, 4, 10, 4, 12, 4, 133,
		9, 4, 1, 4, 1, 4, 3, 4, 137, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 149, 8, 6, 10, 6, 12, 6, 152, 9, 6, 1, 7,
		3, 7, 155, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 162, 8, 7, 1, 7, 3,
		7, 165, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 173, 8, 8, 1, 8,
		1, 8, 3, 8, 177, 8, 8, 1, 8, 3, 8, 180, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 198, 8, 8, 10, 8, 12, 8, 201, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 208, 8, 8, 1, 8, 1, 8, 1, 8, 5, 8, 213, 8, 8, 10, 8, 12, 8, 216, 9,
		8, 1, 8, 3, 8, 219, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 225, 8, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 232, 8, 8, 11, 8, 12, 8, 233, 1, 8, 1, 8,
		3, 8, 238, 8, 8, 1, 8, 1, 8, 3, 8, 242, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 266, 8, 8, 1, 8, 1, 8, 3, 8, 270,
		8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 275, 8, 8, 1, 8, 3, 8, 278, 8, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 283, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 301, 8, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 307, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 314, 8, 8, 5, 8, 316, 8, 8, 10, 8, 12, 8, 319, 9, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 5, 10, 327, 8, 10, 10, 10, 12, 10, 330, 9, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 338, 8, 11, 10, 11, 12,
		11, 341, 9, 11, 1, 12, 3, 12, 344, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 351, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 357, 8, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 5, 12, 363, 8, 12, 10, 12, 12, 12, 366, 9, 12,
		1, 12, 1, 12, 3, 12, 370, 8, 12, 1, 12, 1, 12, 3, 12, 374, 8, 12, 1, 12,
		3, 12, 377, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 383, 8, 13, 10, 13,
		12, 13, 386, 9, 13, 1, 14, 1, 14, 3, 14, 390, 8, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 401, 8, 15, 10, 15, 12,
		15, 404, 9, 15, 1, 15, 1, 15, 1, 15, 3, 15, 409, 8, 15, 3, 15, 411, 8,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 420, 8, 15,
		10, 15, 12, 15, 423, 9, 15, 1, 15, 1, 15, 3, 15, 427, 8, 15, 3, 15, 429,
		8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 435, 8, 16, 10, 16, 12, 16, 438,
		9, 16, 1, 16, 3, 16, 441, 8, 16, 1, 16, 3, 16, 444, 8, 16, 1, 17, 3, 17,
		447, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 456,
		8, 18, 10, 18, 12, 18, 459, 9, 18, 1, 19, 1, 19, 3, 19, 463, 8, 19, 1,
		19, 1, 19, 1, 19, 5, 19, 468, 8, 19, 10, 19, 12, 19, 471, 9, 19, 1, 19,
		1, 19, 1, 19, 3, 19, 476, 8, 19, 3, 19, 478, 8, 19, 1, 19, 1, 19, 3, 19,
		482, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 489, 8, 19, 10, 19,
		12, 19, 492, 9, 19, 1, 19, 1, 19, 3, 19, 496, 8, 19, 3, 19, 498, 8, 19,
		1, 20, 1, 20, 1, 20, 3, 20, 503, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 510, 8, 20, 3, 20, 512, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 3, 21, 522, 8, 21, 3, 21, 524, 8, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 530, 8, 22, 3, 22, 532, 8, 22, 1, 23, 3, 23, 535,
		8, 23, 1, 23, 1, 23, 3, 23, 539, 8, 23, 1, 23, 3, 23, 542, 8, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 551, 8, 25, 1, 25, 1,
		25, 3, 25, 555, 8, 25, 1, 26, 1, 26, 3, 26, 559, 8, 26, 1, 26, 1, 26, 1,
		26, 1, 27, 3, 27, 565, 8, 27, 1, 27, 1, 27, 1, 27, 3, 27, 570, 8, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 577, 8, 27, 10, 27, 12, 27, 580,
		9, 27, 1, 27, 1, 27, 1, 27, 3, 27, 585, 8, 27, 3, 27, 587, 8, 27, 1, 27,
		1, 27, 3, 27, 591, 8, 27, 1, 27, 3, 27, 594, 8, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 5, 28, 600, 8, 28, 10, 28, 12, 28, 603, 9, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 3, 29, 610, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		3, 29, 617, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 624, 8, 30,
		10, 30, 12, 30, 627, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 633, 8,
		31, 1, 32, 1, 32, 1, 32, 3, 32, 638, 8, 32, 1, 32, 3, 32, 641, 8, 32, 1,
		32, 1, 32, 3, 32, 645, 8, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35,
		3, 35, 653, 8, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 0, 1, 16, 42, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
		0, 13, 1, 0, 8, 10, 2, 0, 7, 7, 12, 13, 1, 0, 8, 9, 1, 0, 14, 17, 1, 0,
		18, 21, 4, 0, 55, 55, 65, 65, 73, 73, 87, 87, 5, 0, 50, 50, 79, 79, 95,
		95, 104, 104, 106, 106, 3, 0, 54, 54, 70, 70, 90, 90, 5, 0, 25, 25, 49,
		49, 59, 59, 88, 88, 91, 91, 2, 0, 5, 5, 80, 80, 2, 0, 52, 52, 69, 69, 2,
		0, 29, 29, 41, 41, 3, 0, 55, 55, 71, 71, 88, 88, 748, 0, 87, 1, 0, 0, 0,
		2, 95, 1, 0, 0, 0, 4, 120, 1, 0, 0, 0, 6, 122, 1, 0, 0, 0, 8, 124, 1, 0,
		0, 0, 10, 138, 1, 0, 0, 0, 12, 144, 1, 0, 0, 0, 14, 154, 1, 0, 0, 0, 16,
		241, 1, 0, 0, 0, 18, 320, 1, 0, 0, 0, 20, 322, 1, 0, 0, 0, 22, 333, 1,
		0, 0, 0, 24, 343, 1, 0, 0, 0, 26, 378, 1, 0, 0, 0, 28, 389, 1, 0, 0, 0,
		30, 394, 1, 0, 0, 0, 32, 430, 1, 0, 0, 0, 34, 446, 1, 0, 0, 0, 36, 450,
		1, 0, 0, 0, 38, 460, 1, 0, 0, 0, 40, 511, 1, 0, 0, 0, 42, 523, 1, 0, 0,
		0, 44, 531, 1, 0, 0, 0, 46, 534, 1, 0, 0, 0, 48, 545, 1, 0, 0, 0, 50, 554,
		1, 0, 0, 0, 52, 558, 1, 0, 0, 0, 54, 564, 1, 0, 0, 0, 56, 595, 1, 0, 0,
		0, 58, 606, 1, 0, 0, 0, 60, 618, 1, 0, 0, 0, 62, 628, 1, 0, 0, 0, 64, 634,
		1, 0, 0, 0, 66, 646, 1, 0, 0, 0, 68, 648, 1, 0, 0, 0, 70, 652, 1, 0, 0,
		0, 72, 654, 1, 0, 0, 0, 74, 656, 1, 0, 0, 0, 76, 658, 1, 0, 0, 0, 78, 660,
		1, 0, 0, 0, 80, 662, 1, 0, 0, 0, 82, 664, 1, 0, 0, 0, 84, 86, 3, 2, 1,
		0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88,
		1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 5, 0, 0, 1,
		91, 1, 1, 0, 0, 0, 92, 94, 5, 1, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0,
		0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 98, 107, 3, 4, 2, 0, 99, 101, 5, 1, 0, 0, 100, 99, 1, 0, 0,
		0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103,
		104, 1, 0, 0, 0, 104, 106, 3, 4, 2, 0, 105, 100, 1, 0, 0, 0, 106, 109,
		1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 113, 1, 0,
		0, 0, 109, 107, 1, 0, 0, 0, 110, 112, 5, 1, 0, 0, 111, 110, 1, 0, 0, 0,
		112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		3, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 121, 3, 14, 7, 0, 117, 121, 3,
		24, 12, 0, 118, 121, 3, 34, 17, 0, 119, 121, 3, 54, 27, 0, 120, 116, 1,
		0, 0, 0, 120, 117, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 119, 1, 0, 0,
		0, 121, 5, 1, 0, 0, 0, 122, 123, 3, 76, 38, 0, 123, 7, 1, 0, 0, 0, 124,
		136, 3, 72, 36, 0, 125, 126, 5, 3, 0, 0, 126, 131, 3, 76, 38, 0, 127, 128,
		5, 5, 0, 0, 128, 130, 3, 76, 38, 0, 129, 127, 1, 0, 0, 0, 130, 133, 1,
		0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0, 0,
		0, 133, 131, 1, 0, 0, 0, 134, 135, 5, 4, 0, 0, 135, 137, 1, 0, 0, 0, 136,
		125, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 9, 1, 0, 0, 0, 138, 139, 3,
		8, 4, 0, 139, 140, 5, 30, 0, 0, 140, 141, 5, 3, 0, 0, 141, 142, 3, 32,
		16, 0, 142, 143, 5, 4, 0, 0, 143, 11, 1, 0, 0, 0, 144, 145, 5, 102, 0,
		0, 145, 150, 3, 10, 5, 0, 146, 147, 5, 5, 0, 0, 147, 149, 3, 10, 5, 0,
		148, 146, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150,
		151, 1, 0, 0, 0, 151, 13, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 155, 3,
		12, 6, 0, 154, 153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 1, 0, 0,
		0, 156, 157, 5, 40, 0, 0, 157, 158, 5, 53, 0, 0, 158, 161, 3, 58, 29, 0,
		159, 160, 5, 101, 0, 0, 160, 162, 3, 16, 8, 0, 161, 159, 1, 0, 0, 0, 161,
		162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 165, 3, 26, 13, 0, 164, 163,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 15, 1, 0, 0, 0, 166, 167, 6, 8,
		-1, 0, 167, 242, 3, 18, 9, 0, 168, 242, 5, 105, 0, 0, 169, 170, 3, 72,
		36, 0, 170, 171, 5, 2, 0, 0, 171, 173, 1, 0, 0, 0, 172, 169, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 242, 3, 76, 38, 0, 175,
		177, 5, 77, 0, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 180, 5, 48, 0, 0, 179, 176, 1, 0, 0, 0, 179, 180, 1, 0,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 5, 3, 0, 0, 182, 183, 3, 32, 16,
		0, 183, 184, 5, 4, 0, 0, 184, 242, 1, 0, 0, 0, 185, 186, 5, 3, 0, 0, 186,
		187, 3, 16, 8, 0, 187, 188, 5, 4, 0, 0, 188, 242, 1, 0, 0, 0, 189, 190,
		7, 0, 0, 0, 190, 242, 3, 16, 8, 17, 191, 192, 5, 77, 0, 0, 192, 242, 3,
		16, 8, 6, 193, 194, 5, 3, 0, 0, 194, 199, 3, 16, 8, 0, 195, 196, 5, 5,
		0, 0, 196, 198, 3, 16, 8, 0, 197, 195, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0,
		199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201,
		199, 1, 0, 0, 0, 202, 203, 5, 4, 0, 0, 203, 242, 1, 0, 0, 0, 204, 205,
		3, 70, 35, 0, 205, 218, 5, 3, 0, 0, 206, 208, 5, 42, 0, 0, 207, 206, 1,
		0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 214, 3, 16, 8,
		0, 210, 211, 5, 5, 0, 0, 211, 213, 3, 16, 8, 0, 212, 210, 1, 0, 0, 0, 213,
		216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 219,
		1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 219, 5, 7, 0, 0, 218, 207, 1, 0,
		0, 0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0,
		220, 221, 5, 4, 0, 0, 221, 242, 1, 0, 0, 0, 222, 224, 5, 33, 0, 0, 223,
		225, 3, 16, 8, 0, 224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 231,
		1, 0, 0, 0, 226, 227, 5, 100, 0, 0, 227, 228, 3, 16, 8, 0, 228, 229, 5,
		94, 0, 0, 229, 230, 3, 16, 8, 0, 230, 232, 1, 0, 0, 0, 231, 226, 1, 0,
		0, 0, 232, 233, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0,
		234, 237, 1, 0, 0, 0, 235, 236, 5, 44, 0, 0, 236, 238, 3, 16, 8, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240,
		5, 45, 0, 0, 240, 242, 1, 0, 0, 0, 241, 166, 1, 0, 0, 0, 241, 168, 1, 0,
		0, 0, 241, 172, 1, 0, 0, 0, 241, 179, 1, 0, 0, 0, 241, 185, 1, 0, 0, 0,
		241, 189, 1, 0, 0, 0, 241, 191, 1, 0, 0, 0, 241, 193, 1, 0, 0, 0, 241,
		204, 1, 0, 0, 0, 241, 222, 1, 0, 0, 0, 242, 317, 1, 0, 0, 0, 243, 244,
		10, 15, 0, 0, 244, 245, 5, 11, 0, 0, 245, 316, 3, 16, 8, 16, 246, 247,
		10, 14, 0, 0, 247, 248, 7, 1, 0, 0, 248, 316, 3, 16, 8, 15, 249, 250, 10,
		13, 0, 0, 250, 251, 7, 2, 0, 0, 251, 316, 3, 16, 8, 14, 252, 253, 10, 12,
		0, 0, 253, 254, 7, 3, 0, 0, 254, 316, 3, 16, 8, 13, 255, 256, 10, 11, 0,
		0, 256, 257, 7, 4, 0, 0, 257, 316, 3, 16, 8, 12, 258, 277, 10, 10, 0, 0,
		259, 278, 5, 6, 0, 0, 260, 278, 5, 22, 0, 0, 261, 278, 5, 23, 0, 0, 262,
		278, 5, 24, 0, 0, 263, 265, 5, 67, 0, 0, 264, 266, 5, 77, 0, 0, 265, 264,
		1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 278, 1, 0, 0, 0, 267, 269, 5, 67,
		0, 0, 268, 270, 5, 77, 0, 0, 269, 268, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0,
		270, 271, 1, 0, 0, 0, 271, 272, 5, 42, 0, 0, 272, 278, 5, 53, 0, 0, 273,
		275, 5, 77, 0, 0, 274, 273, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 278, 7, 5, 0, 0, 277, 259, 1, 0, 0, 0, 277, 260, 1, 0,
		0, 0, 277, 261, 1, 0, 0, 0, 277, 262, 1, 0, 0, 0, 277, 263, 1, 0, 0, 0,
		277, 267, 1, 0, 0, 0, 277, 274, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279,
		316, 3, 16, 8, 11, 280, 282, 10, 8, 0, 0, 281, 283, 5, 77, 0, 0, 282, 281,
		1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 5, 31,
		0, 0, 285, 286, 3, 16, 8, 0, 286, 287, 5, 28, 0, 0, 287, 288, 3, 16, 8,
		9, 288, 316, 1, 0, 0, 0, 289, 290, 10, 5, 0, 0, 290, 291, 5, 28, 0, 0,
		291, 316, 3, 16, 8, 6, 292, 293, 10, 4, 0, 0, 293, 294, 5, 84, 0, 0, 294,
		316, 3, 16, 8, 5, 295, 296, 10, 16, 0, 0, 296, 297, 5, 34, 0, 0, 297, 316,
		3, 80, 40, 0, 298, 300, 10, 9, 0, 0, 299, 301, 5, 77, 0, 0, 300, 299, 1,
		0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 5, 71, 0,
		0, 303, 306, 3, 16, 8, 0, 304, 305, 5, 46, 0, 0, 305, 307, 3, 16, 8, 0,
		306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 316, 1, 0, 0, 0, 308,
		313, 10, 7, 0, 0, 309, 314, 5, 66, 0, 0, 310, 314, 5, 76, 0, 0, 311, 312,
		5, 77, 0, 0, 312, 314, 5, 79, 0, 0, 313, 309, 1, 0, 0, 0, 313, 310, 1,
		0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 243, 1, 0, 0,
		0, 315, 246, 1, 0, 0, 0, 315, 249, 1, 0, 0, 0, 315, 252, 1, 0, 0, 0, 315,
		255, 1, 0, 0, 0, 315, 258, 1, 0, 0, 0, 315, 280, 1, 0, 0, 0, 315, 289,
		1, 0, 0, 0, 315, 292, 1, 0, 0, 0, 315, 295, 1, 0, 0, 0, 315, 298, 1, 0,
		0, 0, 315, 308, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0,
		317, 318, 1, 0, 0, 0, 318, 17, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 321,
		7, 6, 0, 0, 321, 19, 1, 0, 0, 0, 322, 323, 5, 3, 0, 0, 323, 328, 3, 16,
		8, 0, 324, 325, 5, 5, 0, 0, 325, 327, 3, 16, 8, 0, 326, 324, 1, 0, 0, 0,
		327, 330, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329,
		331, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 332, 5, 4, 0, 0, 332, 21, 1,
		0, 0, 0, 333, 334, 5, 99, 0, 0, 334, 339, 3, 20, 10, 0, 335, 336, 5, 5,
		0, 0, 336, 338, 3, 20, 10, 0, 337, 335, 1, 0, 0, 0, 338, 341, 1, 0, 0,
		0, 339, 337, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 23, 1, 0, 0, 0, 341,
		339, 1, 0, 0, 0, 342, 344, 3, 12, 6, 0, 343, 342, 1, 0, 0, 0, 343, 344,
		1, 0, 0, 0, 344, 350, 1, 0, 0, 0, 345, 351, 5, 88, 0, 0, 346, 351, 5, 62,
		0, 0, 347, 348, 5, 62, 0, 0, 348, 349, 5, 84, 0, 0, 349, 351, 5, 88, 0,
		0, 350, 345, 1, 0, 0, 0, 350, 346, 1, 0, 0, 0, 350, 347, 1, 0, 0, 0, 351,
		352, 1, 0, 0, 0, 352, 353, 5, 64, 0, 0, 353, 356, 3, 72, 36, 0, 354, 355,
		5, 30, 0, 0, 355, 357, 3, 74, 37, 0, 356, 354, 1, 0, 0, 0, 356, 357, 1,
		0, 0, 0, 357, 369, 1, 0, 0, 0, 358, 359, 5, 3, 0, 0, 359, 364, 3, 76, 38,
		0, 360, 361, 5, 5, 0, 0, 361, 363, 3, 76, 38, 0, 362, 360, 1, 0, 0, 0,
		363, 366, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365,
		367, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 368, 5, 4, 0, 0, 368, 370,
		1, 0, 0, 0, 369, 358, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0,
		0, 0, 371, 373, 3, 22, 11, 0, 372, 374, 3, 30, 15, 0, 373, 372, 1, 0, 0,
		0, 373, 374, 1, 0, 0, 0, 374, 376, 1, 0, 0, 0, 375, 377, 3, 26, 13, 0,
		376, 375, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 25, 1, 0, 0, 0, 378, 379,
		5, 89, 0, 0, 379, 384, 3, 44, 22, 0, 380, 381, 5, 5, 0, 0, 381, 383, 3,
		44, 22, 0, 382, 380, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0,
		0, 0, 384, 385, 1, 0, 0, 0, 385, 27, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0,
		387, 390, 3, 76, 38, 0, 388, 390, 3, 56, 28, 0, 389, 387, 1, 0, 0, 0, 389,
		388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 392, 5, 6, 0, 0, 392, 393,
		3, 16, 8, 0, 393, 29, 1, 0, 0, 0, 394, 395, 5, 82, 0, 0, 395, 410, 5, 36,
		0, 0, 396, 397, 5, 3, 0, 0, 397, 402, 3, 6, 3, 0, 398, 399, 5, 5, 0, 0,
		399, 401, 3, 6, 3, 0, 400, 398, 1, 0, 0, 0, 401, 404, 1, 0, 0, 0, 402,
		400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 405, 1, 0, 0, 0, 404, 402,
		1, 0, 0, 0, 405, 408, 5, 4, 0, 0, 406, 407, 5, 101, 0, 0, 407, 409, 3,
		16, 8, 0, 408, 406, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 411, 1, 0, 0,
		0, 410, 396, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412,
		428, 5, 43, 0, 0, 413, 429, 5, 75, 0, 0, 414, 415, 5, 97, 0, 0, 415, 416,
		5, 93, 0, 0, 416, 421, 3, 28, 14, 0, 417, 418, 5, 5, 0, 0, 418, 420, 3,
		28, 14, 0, 419, 417, 1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421, 419, 1, 0,
		0, 0, 421, 422, 1, 0, 0, 0, 422, 426, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0,
		424, 425, 5, 101, 0, 0, 425, 427, 3, 16, 8, 0, 426, 424, 1, 0, 0, 0, 426,
		427, 1, 0, 0, 0, 427, 429, 1, 0, 0, 0, 428, 413, 1, 0, 0, 0, 428, 414,
		1, 0, 0, 0, 429, 31, 1, 0, 0, 0, 430, 436, 3, 38, 19, 0, 431, 432, 3, 50,
		25, 0, 432, 433, 3, 38, 19, 0, 433, 435, 1, 0, 0, 0, 434, 431, 1, 0, 0,
		0, 435, 438, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437,
		440, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 439, 441, 3, 60, 30, 0, 440, 439,
		1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 443, 1, 0, 0, 0, 442, 444, 3, 62,
		31, 0, 443, 442, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 33, 1, 0, 0, 0,
		445, 447, 3, 12, 6, 0, 446, 445, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447,
		448, 1, 0, 0, 0, 448, 449, 3, 32, 16, 0, 449, 35, 1, 0, 0, 0, 450, 457,
		3, 40, 20, 0, 451, 452, 3, 46, 23, 0, 452, 453, 3, 40, 20, 0, 453, 454,
		3, 48, 24, 0, 454, 456, 1, 0, 0, 0, 455, 451, 1, 0, 0, 0, 456, 459, 1,
		0, 0, 0, 457, 455, 1, 0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 37, 1, 0, 0,
		0, 459, 457, 1, 0, 0, 0, 460, 462, 5, 92, 0, 0, 461, 463, 5, 42, 0, 0,
		462, 461, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464,
		469, 3, 42, 21, 0, 465, 466, 5, 5, 0, 0, 466, 468, 3, 42, 21, 0, 467, 465,
		1, 0, 0, 0, 468, 471, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469, 470, 1, 0,
		0, 0, 470, 477, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 472, 475, 5, 53, 0, 0,
		473, 476, 3, 40, 20, 0, 474, 476, 3, 36, 18, 0, 475, 473, 1, 0, 0, 0, 475,
		474, 1, 0, 0, 0, 476, 478, 1, 0, 0, 0, 477, 472, 1, 0, 0, 0, 477, 478,
		1, 0, 0, 0, 478, 481, 1, 0, 0, 0, 479, 480, 5, 101, 0, 0, 480, 482, 3,
		16, 8, 0, 481, 479, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 497, 1, 0, 0,
		0, 483, 484, 5, 57, 0, 0, 484, 485, 5, 32, 0, 0, 485, 490, 3, 16, 8, 0,
		486, 487, 5, 5, 0, 0, 487, 489, 3, 16, 8, 0, 488, 486, 1, 0, 0, 0, 489,
		492, 1, 0, 0, 0, 490, 488, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 495,
		1, 0, 0, 0, 492, 490, 1, 0, 0, 0, 493, 494, 5, 58, 0, 0, 494, 496, 3, 16,
		8, 0, 495, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 498, 1, 0, 0, 0,
		497, 483, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 39, 1, 0, 0, 0, 499, 502,
		3, 72, 36, 0, 500, 501, 5, 30, 0, 0, 501, 503, 3, 74, 37, 0, 502, 500,
		1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 512, 1, 0, 0, 0, 504, 505, 5, 3,
		0, 0, 505, 506, 3, 32, 16, 0, 506, 509, 5, 4, 0, 0, 507, 508, 5, 30, 0,
		0, 508, 510, 3, 74, 37, 0, 509, 507, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0,
		510, 512, 1, 0, 0, 0, 511, 499, 1, 0, 0, 0, 511, 504, 1, 0, 0, 0, 512,
		41, 1, 0, 0, 0, 513, 524, 5, 7, 0, 0, 514, 515, 3, 72, 36, 0, 515, 516,
		5, 2, 0, 0, 516, 517, 5, 7, 0, 0, 517, 524, 1, 0, 0, 0, 518, 521, 3, 16,
		8, 0, 519, 520, 5, 30, 0, 0, 520, 522, 3, 78, 39, 0, 521, 519, 1, 0, 0,
		0, 521, 522, 1, 0, 0, 0, 522, 524, 1, 0, 0, 0, 523, 513, 1, 0, 0, 0, 523,
		514, 1, 0, 0, 0, 523, 518, 1, 0, 0, 0, 524, 43, 1, 0, 0, 0, 525, 532, 5,
		7, 0, 0, 526, 529, 3, 16, 8, 0, 527, 528, 5, 30, 0, 0, 528, 530, 3, 78,
		39, 0, 529, 527, 1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530, 532, 1, 0, 0, 0,
		531, 525, 1, 0, 0, 0, 531, 526, 1, 0, 0, 0, 532, 45, 1, 0, 0, 0, 533, 535,
		5, 74, 0, 0, 534, 533, 1, 0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 541, 1, 0,
		0, 0, 536, 538, 7, 7, 0, 0, 537, 539, 5, 85, 0, 0, 538, 537, 1, 0, 0, 0,
		538, 539, 1, 0, 0, 0, 539, 542, 1, 0, 0, 0, 540, 542, 5, 61, 0, 0, 541,
		536, 1, 0, 0, 0, 541, 540, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 543,
		1, 0, 0, 0, 543, 544, 5, 68, 0, 0, 544, 47, 1, 0, 0, 0, 545, 546, 5, 82,
		0, 0, 546, 547, 3, 16, 8, 0, 547, 49, 1, 0, 0, 0, 548, 550, 5, 96, 0, 0,
		549, 551, 5, 27, 0, 0, 550, 549, 1, 0, 0, 0, 550, 551, 1, 0, 0, 0, 551,
		555, 1, 0, 0, 0, 552, 555, 5, 63, 0, 0, 553, 555, 5, 47, 0, 0, 554, 548,
		1, 0, 0, 0, 554, 552, 1, 0, 0, 0, 554, 553, 1, 0, 0, 0, 555, 51, 1, 0,
		0, 0, 556, 559, 3, 76, 38, 0, 557, 559, 3, 56, 28, 0, 558, 556, 1, 0, 0,
		0, 558, 557, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0, 560, 561, 5, 6, 0, 0, 561,
		562, 3, 16, 8, 0, 562, 53, 1, 0, 0, 0, 563, 565, 3, 12, 6, 0, 564, 563,
		1, 0, 0, 0, 564, 565, 1, 0, 0, 0, 565, 566, 1, 0, 0, 0, 566, 569, 5, 97,
		0, 0, 567, 568, 5, 84, 0, 0, 568, 570, 7, 8, 0, 0, 569, 567, 1, 0, 0, 0,
		569, 570, 1, 0, 0, 0, 570, 571, 1, 0, 0, 0, 571, 572, 3, 58, 29, 0, 572,
		573, 5, 93, 0, 0, 573, 578, 3, 52, 26, 0, 574, 575, 5, 5, 0, 0, 575, 577,
		3, 52, 26, 0, 576, 574, 1, 0, 0, 0, 577, 580, 1, 0, 0, 0, 578, 576, 1,
		0, 0, 0, 578, 579, 1, 0, 0, 0, 579, 586, 1, 0, 0, 0, 580, 578, 1, 0, 0,
		0, 581, 584, 5, 53, 0, 0, 582, 585, 3, 40, 20, 0, 583, 585, 3, 36, 18,
		0, 584, 582, 1, 0, 0, 0, 584, 583, 1, 0, 0, 0, 585, 587, 1, 0, 0, 0, 586,
		581, 1, 0, 0, 0, 586, 587, 1, 0, 0, 0, 587, 590, 1, 0, 0, 0, 588, 589,
		5, 101, 0, 0, 589, 591, 3, 16, 8, 0, 590, 588, 1, 0, 0, 0, 590, 591, 1,
		0, 0, 0, 591, 593, 1, 0, 0, 0, 592, 594, 3, 26, 13, 0, 593, 592, 1, 0,
		0, 0, 593, 594, 1, 0, 0, 0, 594, 55, 1, 0, 0, 0, 595, 596, 5, 3, 0, 0,
		596, 601, 3, 76, 38, 0, 597, 598, 5, 5, 0, 0, 598, 600, 3, 76, 38, 0, 599,
		597, 1, 0, 0, 0, 600, 603, 1, 0, 0, 0, 601, 599, 1, 0, 0, 0, 601, 602,
		1, 0, 0, 0, 602, 604, 1, 0, 0, 0, 603, 601, 1, 0, 0, 0, 604, 605, 5, 4,
		0, 0, 605, 57, 1, 0, 0, 0, 606, 609, 3, 72, 36, 0, 607, 608, 5, 30, 0,
		0, 608, 610, 3, 74, 37, 0, 609, 607, 1, 0, 0, 0, 609, 610, 1, 0, 0, 0,
		610, 616, 1, 0, 0, 0, 611, 612, 5, 60, 0, 0, 612, 613, 5, 32, 0, 0, 613,
		617, 3, 82, 41, 0, 614, 615, 5, 77, 0, 0, 615, 617, 5, 60, 0, 0, 616, 611,
		1, 0, 0, 0, 616, 614, 1, 0, 0, 0, 616, 617, 1, 0, 0, 0, 617, 59, 1, 0,
		0, 0, 618, 619, 5, 83, 0, 0, 619, 620, 5, 32, 0, 0, 620, 625, 3, 64, 32,
		0, 621, 622, 5, 5, 0, 0, 622, 624, 3, 64, 32, 0, 623, 621, 1, 0, 0, 0,
		624, 627, 1, 0, 0, 0, 625, 623, 1, 0, 0, 0, 625, 626, 1, 0, 0, 0, 626,
		61, 1, 0, 0, 0, 627, 625, 1, 0, 0, 0, 628, 629, 5, 72, 0, 0, 629, 632,
		3, 16, 8, 0, 630, 631, 7, 9, 0, 0, 631, 633, 3, 16, 8, 0, 632, 630, 1,
		0, 0, 0, 632, 633, 1, 0, 0, 0, 633, 63, 1, 0, 0, 0, 634, 637, 3, 16, 8,
		0, 635, 636, 5, 34, 0, 0, 636, 638, 3, 80, 40, 0, 637, 635, 1, 0, 0, 0,
		637, 638, 1, 0, 0, 0, 638, 640, 1, 0, 0, 0, 639, 641, 3, 66, 33, 0, 640,
		639, 1, 0, 0, 0, 640, 641, 1, 0, 0, 0, 641, 644, 1, 0, 0, 0, 642, 643,
		5, 78, 0, 0, 643, 645, 7, 10, 0, 0, 644, 642, 1, 0, 0, 0, 644, 645, 1,
		0, 0, 0, 645, 65, 1, 0, 0, 0, 646, 647, 7, 11, 0, 0, 647, 67, 1, 0, 0,
		0, 648, 649, 7, 12, 0, 0, 649, 69, 1, 0, 0, 0, 650, 653, 5, 103, 0, 0,
		651, 653, 3, 68, 34, 0, 652, 650, 1, 0, 0, 0, 652, 651, 1, 0, 0, 0, 653,
		71, 1, 0, 0, 0, 654, 655, 5, 103, 0, 0, 655, 73, 1, 0, 0, 0, 656, 657,
		5, 103, 0, 0, 657, 75, 1, 0, 0, 0, 658, 659, 5, 103, 0, 0, 659, 77, 1,
		0, 0, 0, 660, 661, 5, 103, 0, 0, 661, 79, 1, 0, 0, 0, 662, 663, 5, 103,
		0, 0, 663, 81, 1, 0, 0, 0, 664, 665, 5, 103, 0, 0, 665, 83, 1, 0, 0, 0,
		92, 87, 95, 102, 107, 113, 120, 131, 136, 150, 154, 161, 164, 172, 176,
		179, 199, 207, 214, 218, 224, 233, 237, 241, 265, 269, 274, 277, 282, 300,
		306, 313, 315, 317, 328, 339, 343, 350, 356, 364, 369, 373, 376, 384, 389,
		402, 408, 410, 421, 426, 428, 436, 440, 443, 446, 457, 462, 469, 475, 477,
		481, 490, 495, 497, 502, 509, 511, 521, 523, 529, 531, 534, 538, 541, 550,
		554, 558, 564, 569, 578, 584, 586, 590, 593, 601, 609, 616, 625, 632, 637,
		640, 644, 652,
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
	SQLiteParserABORT_              = 25
	SQLiteParserADD_                = 26
	SQLiteParserALL_                = 27
	SQLiteParserAND_                = 28
	SQLiteParserASC_                = 29
	SQLiteParserAS_                 = 30
	SQLiteParserBETWEEN_            = 31
	SQLiteParserBY_                 = 32
	SQLiteParserCASE_               = 33
	SQLiteParserCOLLATE_            = 34
	SQLiteParserCOMMIT_             = 35
	SQLiteParserCONFLICT_           = 36
	SQLiteParserCREATE_             = 37
	SQLiteParserCROSS_              = 38
	SQLiteParserDEFAULT_            = 39
	SQLiteParserDELETE_             = 40
	SQLiteParserDESC_               = 41
	SQLiteParserDISTINCT_           = 42
	SQLiteParserDO_                 = 43
	SQLiteParserELSE_               = 44
	SQLiteParserEND_                = 45
	SQLiteParserESCAPE_             = 46
	SQLiteParserEXCEPT_             = 47
	SQLiteParserEXISTS_             = 48
	SQLiteParserFAIL_               = 49
	SQLiteParserFALSE_              = 50
	SQLiteParserFILTER_             = 51
	SQLiteParserFIRST_              = 52
	SQLiteParserFROM_               = 53
	SQLiteParserFULL_               = 54
	SQLiteParserGLOB_               = 55
	SQLiteParserGROUPS_             = 56
	SQLiteParserGROUP_              = 57
	SQLiteParserHAVING_             = 58
	SQLiteParserIGNORE_             = 59
	SQLiteParserINDEXED_            = 60
	SQLiteParserINNER_              = 61
	SQLiteParserINSERT_             = 62
	SQLiteParserINTERSECT_          = 63
	SQLiteParserINTO_               = 64
	SQLiteParserIN_                 = 65
	SQLiteParserISNULL_             = 66
	SQLiteParserIS_                 = 67
	SQLiteParserJOIN_               = 68
	SQLiteParserLAST_               = 69
	SQLiteParserLEFT_               = 70
	SQLiteParserLIKE_               = 71
	SQLiteParserLIMIT_              = 72
	SQLiteParserMATCH_              = 73
	SQLiteParserNATURAL_            = 74
	SQLiteParserNOTHING_            = 75
	SQLiteParserNOTNULL_            = 76
	SQLiteParserNOT_                = 77
	SQLiteParserNULLS_              = 78
	SQLiteParserNULL_               = 79
	SQLiteParserOFFSET_             = 80
	SQLiteParserOF_                 = 81
	SQLiteParserON_                 = 82
	SQLiteParserORDER_              = 83
	SQLiteParserOR_                 = 84
	SQLiteParserOUTER_              = 85
	SQLiteParserRAISE_              = 86
	SQLiteParserREGEXP_             = 87
	SQLiteParserREPLACE_            = 88
	SQLiteParserRETURNING_          = 89
	SQLiteParserRIGHT_              = 90
	SQLiteParserROLLBACK_           = 91
	SQLiteParserSELECT_             = 92
	SQLiteParserSET_                = 93
	SQLiteParserTHEN_               = 94
	SQLiteParserTRUE_               = 95
	SQLiteParserUNION_              = 96
	SQLiteParserUPDATE_             = 97
	SQLiteParserUSING_              = 98
	SQLiteParserVALUES_             = 99
	SQLiteParserWHEN_               = 100
	SQLiteParserWHERE_              = 101
	SQLiteParserWITH_               = 102
	SQLiteParserIDENTIFIER          = 103
	SQLiteParserNUMERIC_LITERAL     = 104
	SQLiteParserBIND_PARAMETER      = 105
	SQLiteParserSTRING_LITERAL      = 106
	SQLiteParserBLOB_LITERAL        = 107
	SQLiteParserSINGLE_LINE_COMMENT = 108
	SQLiteParserMULTILINE_COMMENT   = 109
	SQLiteParserSPACES              = 110
	SQLiteParserUNEXPECTED_CHAR     = 111
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
	SQLiteParserRULE_literal_value                  = 9
	SQLiteParserRULE_value_row                      = 10
	SQLiteParserRULE_values_clause                  = 11
	SQLiteParserRULE_insert_stmt                    = 12
	SQLiteParserRULE_returning_clause               = 13
	SQLiteParserRULE_upsert_update                  = 14
	SQLiteParserRULE_upsert_clause                  = 15
	SQLiteParserRULE_select_stmt_core               = 16
	SQLiteParserRULE_select_stmt                    = 17
	SQLiteParserRULE_join_clause                    = 18
	SQLiteParserRULE_select_core                    = 19
	SQLiteParserRULE_table_or_subquery              = 20
	SQLiteParserRULE_result_column                  = 21
	SQLiteParserRULE_returning_clause_result_column = 22
	SQLiteParserRULE_join_operator                  = 23
	SQLiteParserRULE_join_constraint                = 24
	SQLiteParserRULE_compound_operator              = 25
	SQLiteParserRULE_update_set_subclause           = 26
	SQLiteParserRULE_update_stmt                    = 27
	SQLiteParserRULE_column_name_list               = 28
	SQLiteParserRULE_qualified_table_name           = 29
	SQLiteParserRULE_order_by_stmt                  = 30
	SQLiteParserRULE_limit_stmt                     = 31
	SQLiteParserRULE_ordering_term                  = 32
	SQLiteParserRULE_asc_desc                       = 33
	SQLiteParserRULE_function_keyword               = 34
	SQLiteParserRULE_function_name                  = 35
	SQLiteParserRULE_table_name                     = 36
	SQLiteParserRULE_table_alias                    = 37
	SQLiteParserRULE_column_name                    = 38
	SQLiteParserRULE_column_alias                   = 39
	SQLiteParserRULE_collation_name                 = 40
	SQLiteParserRULE_index_name                     = 41
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4611687117939015682) != 0) || ((int64((_la-88)) & ^0x3f) == 0 && ((int64(1)<<(_la-88))&16913) != 0) {
		{
			p.SetState(84)
			p.Sql_stmt_list()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserSCOL {
		{
			p.SetState(92)
			p.Match(SQLiteParserSCOL)
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Sql_stmt()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == SQLiteParserSCOL {
				{
					p.SetState(99)
					p.Match(SQLiteParserSCOL)
				}

				p.SetState(102)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(104)
				p.Sql_stmt()
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(110)
				p.Match(SQLiteParserSCOL)
			}

		}
		p.SetState(115)
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
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(116)
			p.Delete_stmt()
		}

	case 2:
		{
			p.SetState(117)
			p.Insert_stmt()
		}

	case 3:
		{
			p.SetState(118)
			p.Select_stmt()
		}

	case 4:
		{
			p.SetState(119)
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
		p.SetState(122)
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
		p.SetState(124)
		p.Table_name()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOPEN_PAR {
		{
			p.SetState(125)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(126)
			p.Column_name()
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(127)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(128)
				p.Column_name()
			}

			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(134)
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
		p.SetState(138)
		p.Cte_table_name()
	}
	{
		p.SetState(139)
		p.Match(SQLiteParserAS_)
	}
	{
		p.SetState(140)
		p.Match(SQLiteParserOPEN_PAR)
	}
	{
		p.SetState(141)
		p.Select_stmt_core()
	}
	{
		p.SetState(142)
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
		p.SetState(144)
		p.Match(SQLiteParserWITH_)
	}
	{
		p.SetState(145)
		p.Common_table_expression()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(146)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(147)
			p.Common_table_expression()
		}

		p.SetState(152)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(153)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(156)
		p.Match(SQLiteParserDELETE_)
	}
	{
		p.SetState(157)
		p.Match(SQLiteParserFROM_)
	}
	{
		p.SetState(158)
		p.Qualified_table_name()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWHERE_ {
		{
			p.SetState(159)
			p.Match(SQLiteParserWHERE_)
		}
		{
			p.SetState(160)
			p.expr(0)
		}

	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserRETURNING_ {
		{
			p.SetState(163)
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
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(167)
			p.Literal_value()
		}

	case 2:
		{
			p.SetState(168)
			p.Match(SQLiteParserBIND_PARAMETER)
		}

	case 3:
		p.SetState(172)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(169)
				p.Table_name()
			}
			{
				p.SetState(170)
				p.Match(SQLiteParserDOT)
			}

		}
		{
			p.SetState(174)
			p.Column_name()
		}

	case 4:
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserEXISTS_ || _la == SQLiteParserNOT_ {
			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == SQLiteParserNOT_ {
				{
					p.SetState(175)
					p.Match(SQLiteParserNOT_)
				}

			}
			{
				p.SetState(178)
				p.Match(SQLiteParserEXISTS_)
			}

		}
		{
			p.SetState(181)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(182)
			p.Select_stmt_core()
		}
		{
			p.SetState(183)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	case 5:
		{
			p.SetState(185)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(186)

			var _x = p.expr(0)

			localctx.(*ExprContext).elevate_expr = _x
		}
		{
			p.SetState(187)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	case 6:
		{
			p.SetState(189)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1792) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(190)

			var _x = p.expr(17)

			localctx.(*ExprContext).unary_expr = _x
		}

	case 7:
		{
			p.SetState(191)
			p.Match(SQLiteParserNOT_)
		}
		{
			p.SetState(192)

			var _x = p.expr(6)

			localctx.(*ExprContext).unary_expr = _x
		}

	case 8:
		{
			p.SetState(193)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(194)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).expr_list = append(localctx.(*ExprContext).expr_list, localctx.(*ExprContext)._expr)
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(195)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(196)

				var _x = p.expr(0)

				localctx.(*ExprContext)._expr = _x
			}
			localctx.(*ExprContext).expr_list = append(localctx.(*ExprContext).expr_list, localctx.(*ExprContext)._expr)

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(202)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	case 9:
		{
			p.SetState(204)
			p.Function_name()
		}
		{
			p.SetState(205)
			p.Match(SQLiteParserOPEN_PAR)
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SQLiteParserOPEN_PAR, SQLiteParserPLUS, SQLiteParserMINUS, SQLiteParserTILDE, SQLiteParserCASE_, SQLiteParserDISTINCT_, SQLiteParserEXISTS_, SQLiteParserFALSE_, SQLiteParserGLOB_, SQLiteParserLIKE_, SQLiteParserNOT_, SQLiteParserNULL_, SQLiteParserREPLACE_, SQLiteParserTRUE_, SQLiteParserIDENTIFIER, SQLiteParserNUMERIC_LITERAL, SQLiteParserBIND_PARAMETER, SQLiteParserSTRING_LITERAL:
			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == SQLiteParserDISTINCT_ {
				{
					p.SetState(206)
					p.Match(SQLiteParserDISTINCT_)
				}

			}
			{
				p.SetState(209)
				p.expr(0)
			}
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
					p.expr(0)
				}

				p.SetState(216)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		case SQLiteParserSTAR:
			{
				p.SetState(217)
				p.Match(SQLiteParserSTAR)
			}

		case SQLiteParserCLOSE_PAR:

		default:
		}
		{
			p.SetState(220)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	case 10:
		{
			p.SetState(222)
			p.Match(SQLiteParserCASE_)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&37436180492453640) != 0) || ((int64((_la-71)) & ^0x3f) == 0 && ((int64(1)<<(_la-71))&64441418049) != 0) {
			{
				p.SetState(223)

				var _x = p.expr(0)

				localctx.(*ExprContext).case_expr = _x
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SQLiteParserWHEN_ {
			{
				p.SetState(226)
				p.Match(SQLiteParserWHEN_)
			}
			{
				p.SetState(227)

				var _x = p.expr(0)

				localctx.(*ExprContext)._expr = _x
			}
			localctx.(*ExprContext).when_expr = append(localctx.(*ExprContext).when_expr, localctx.(*ExprContext)._expr)
			{
				p.SetState(228)
				p.Match(SQLiteParserTHEN_)
			}
			{
				p.SetState(229)

				var _x = p.expr(0)

				localctx.(*ExprContext)._expr = _x
			}
			localctx.(*ExprContext).then_expr = append(localctx.(*ExprContext).then_expr, localctx.(*ExprContext)._expr)

			p.SetState(233)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserELSE_ {
			{
				p.SetState(235)
				p.Match(SQLiteParserELSE_)
			}
			{
				p.SetState(236)

				var _x = p.expr(0)

				localctx.(*ExprContext).else_expr = _x
			}

		}
		{
			p.SetState(239)
			p.Match(SQLiteParserEND_)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(244)
					p.Match(SQLiteParserPIPE2)
				}
				{
					p.SetState(245)
					p.expr(16)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(247)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12416) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(248)
					p.expr(15)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(249)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(250)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SQLiteParserPLUS || _la == SQLiteParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(251)
					p.expr(14)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(253)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(254)
					p.expr(13)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(256)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3932160) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(257)
					p.expr(12)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(277)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(259)
						p.Match(SQLiteParserASSIGN)
					}

				case 2:
					{
						p.SetState(260)
						p.Match(SQLiteParserEQ)
					}

				case 3:
					{
						p.SetState(261)
						p.Match(SQLiteParserNOT_EQ1)
					}

				case 4:
					{
						p.SetState(262)
						p.Match(SQLiteParserNOT_EQ2)
					}

				case 5:
					{
						p.SetState(263)
						p.Match(SQLiteParserIS_)
					}
					p.SetState(265)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
						{
							p.SetState(264)
							p.Match(SQLiteParserNOT_)
						}

					}

				case 6:
					{
						p.SetState(267)
						p.Match(SQLiteParserIS_)
					}
					p.SetState(269)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == SQLiteParserNOT_ {
						{
							p.SetState(268)
							p.Match(SQLiteParserNOT_)
						}

					}
					{
						p.SetState(271)
						p.Match(SQLiteParserDISTINCT_)
					}
					{
						p.SetState(272)
						p.Match(SQLiteParserFROM_)
					}

				case 7:
					p.SetState(274)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == SQLiteParserNOT_ {
						{
							p.SetState(273)
							p.Match(SQLiteParserNOT_)
						}

					}
					{
						p.SetState(276)
						_la = p.GetTokenStream().LA(1)

						if !((int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&4295230465) != 0) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				}
				{
					p.SetState(279)
					p.expr(11)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				p.SetState(282)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SQLiteParserNOT_ {
					{
						p.SetState(281)
						p.Match(SQLiteParserNOT_)
					}

				}
				{
					p.SetState(284)
					p.Match(SQLiteParserBETWEEN_)
				}
				{
					p.SetState(285)
					p.expr(0)
				}
				{
					p.SetState(286)
					p.Match(SQLiteParserAND_)
				}
				{
					p.SetState(287)
					p.expr(9)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(290)
					p.Match(SQLiteParserAND_)
				}
				{
					p.SetState(291)
					p.expr(6)
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(293)
					p.Match(SQLiteParserOR_)
				}
				{
					p.SetState(294)
					p.expr(5)
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(296)
					p.Match(SQLiteParserCOLLATE_)
				}
				{
					p.SetState(297)
					p.Collation_name()
				}

			case 11:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(300)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SQLiteParserNOT_ {
					{
						p.SetState(299)
						p.Match(SQLiteParserNOT_)
					}

				}
				{
					p.SetState(302)
					p.Match(SQLiteParserLIKE_)
				}
				{
					p.SetState(303)
					p.expr(0)
				}
				p.SetState(306)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(304)
						p.Match(SQLiteParserESCAPE_)
					}
					{
						p.SetState(305)
						p.expr(0)
					}

				}

			case 12:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SQLiteParserRULE_expr)
				p.SetState(308)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				p.SetState(313)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case SQLiteParserISNULL_:
					{
						p.SetState(309)
						p.Match(SQLiteParserISNULL_)
					}

				case SQLiteParserNOTNULL_:
					{
						p.SetState(310)
						p.Match(SQLiteParserNOTNULL_)
					}

				case SQLiteParserNOT_:
					{
						p.SetState(311)
						p.Match(SQLiteParserNOT_)
					}
					{
						p.SetState(312)
						p.Match(SQLiteParserNULL_)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}

		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, SQLiteParserRULE_literal_value)
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
		p.SetState(320)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-50)) & ^0x3f) == 0 && ((int64(1)<<(_la-50))&90107177456369665) != 0) {
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
	p.EnterRule(localctx, 20, SQLiteParserRULE_value_row)
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
		p.SetState(322)
		p.Match(SQLiteParserOPEN_PAR)
	}
	{
		p.SetState(323)
		p.expr(0)
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(324)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(325)
			p.expr(0)
		}

		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(331)
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
	p.EnterRule(localctx, 22, SQLiteParserRULE_values_clause)
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
		p.SetState(333)
		p.Match(SQLiteParserVALUES_)
	}
	{
		p.SetState(334)
		p.Value_row()
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(335)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(336)
			p.Value_row()
		}

		p.SetState(341)
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
	p.EnterRule(localctx, 24, SQLiteParserRULE_insert_stmt)
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
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(342)
			p.Common_table_stmt()
		}

	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(345)
			p.Match(SQLiteParserREPLACE_)
		}

	case 2:
		{
			p.SetState(346)
			p.Match(SQLiteParserINSERT_)
		}

	case 3:
		{
			p.SetState(347)
			p.Match(SQLiteParserINSERT_)
		}
		{
			p.SetState(348)
			p.Match(SQLiteParserOR_)
		}
		{
			p.SetState(349)
			p.Match(SQLiteParserREPLACE_)
		}

	}
	{
		p.SetState(352)
		p.Match(SQLiteParserINTO_)
	}
	{
		p.SetState(353)
		p.Table_name()
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserAS_ {
		{
			p.SetState(354)
			p.Match(SQLiteParserAS_)
		}
		{
			p.SetState(355)
			p.Table_alias()
		}

	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOPEN_PAR {
		{
			p.SetState(358)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(359)
			p.Column_name()
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(360)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(361)
				p.Column_name()
			}

			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(367)
			p.Match(SQLiteParserCLOSE_PAR)
		}

	}
	{
		p.SetState(371)
		p.Values_clause()
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserON_ {
		{
			p.SetState(372)
			p.Upsert_clause()
		}

	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserRETURNING_ {
		{
			p.SetState(375)
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
	p.EnterRule(localctx, 26, SQLiteParserRULE_returning_clause)
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
		p.SetState(378)
		p.Match(SQLiteParserRETURNING_)
	}
	{
		p.SetState(379)
		p.Returning_clause_result_column()
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(380)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(381)
			p.Returning_clause_result_column()
		}

		p.SetState(386)
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
	p.EnterRule(localctx, 28, SQLiteParserRULE_upsert_update)

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
	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		{
			p.SetState(387)
			p.Column_name()
		}

	case SQLiteParserOPEN_PAR:
		{
			p.SetState(388)
			p.Column_name_list()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(391)
		p.Match(SQLiteParserASSIGN)
	}
	{
		p.SetState(392)
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
	p.EnterRule(localctx, 30, SQLiteParserRULE_upsert_clause)
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
		p.SetState(394)
		p.Match(SQLiteParserON_)
	}
	{
		p.SetState(395)
		p.Match(SQLiteParserCONFLICT_)
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOPEN_PAR {
		{
			p.SetState(396)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(397)
			p.Indexed_column()
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(398)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(399)
				p.Indexed_column()
			}

			p.SetState(404)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(405)
			p.Match(SQLiteParserCLOSE_PAR)
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserWHERE_ {
			{
				p.SetState(406)
				p.Match(SQLiteParserWHERE_)
			}
			{
				p.SetState(407)

				var _x = p.expr(0)

				localctx.(*Upsert_clauseContext).target_expr = _x
			}

		}

	}
	{
		p.SetState(412)
		p.Match(SQLiteParserDO_)
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserNOTHING_:
		{
			p.SetState(413)
			p.Match(SQLiteParserNOTHING_)
		}

	case SQLiteParserUPDATE_:
		{
			p.SetState(414)
			p.Match(SQLiteParserUPDATE_)
		}
		{
			p.SetState(415)
			p.Match(SQLiteParserSET_)
		}

		{
			p.SetState(416)
			p.Upsert_update()
		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(417)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(418)
				p.Upsert_update()
			}

			p.SetState(423)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserWHERE_ {
			{
				p.SetState(424)
				p.Match(SQLiteParserWHERE_)
			}
			{
				p.SetState(425)

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
	p.EnterRule(localctx, 32, SQLiteParserRULE_select_stmt_core)
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
		p.SetState(430)
		p.Select_core()
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-47)) & ^0x3f) == 0 && ((int64(1)<<(_la-47))&562949953486849) != 0 {
		{
			p.SetState(431)
			p.Compound_operator()
		}
		{
			p.SetState(432)
			p.Select_core()
		}

		p.SetState(438)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserORDER_ {
		{
			p.SetState(439)
			p.Order_by_stmt()
		}

	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserLIMIT_ {
		{
			p.SetState(442)
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
	p.EnterRule(localctx, 34, SQLiteParserRULE_select_stmt)
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
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(445)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(448)
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
	p.EnterRule(localctx, 36, SQLiteParserRULE_join_clause)
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
		p.SetState(450)
		p.Table_or_subquery()
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-54)) & ^0x3f) == 0 && ((int64(1)<<(_la-54))&68720607361) != 0 {
		{
			p.SetState(451)
			p.Join_operator()
		}
		{
			p.SetState(452)
			p.Table_or_subquery()
		}
		{
			p.SetState(453)
			p.Join_constraint()
		}

		p.SetState(459)
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
	p.EnterRule(localctx, 38, SQLiteParserRULE_select_core)
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
		p.SetState(460)
		p.Match(SQLiteParserSELECT_)
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserDISTINCT_ {
		{
			p.SetState(461)
			p.Match(SQLiteParserDISTINCT_)
		}

	}
	{
		p.SetState(464)
		p.Result_column()
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(465)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(466)
			p.Result_column()
		}

		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserFROM_ {
		{
			p.SetState(472)
			p.Match(SQLiteParserFROM_)
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(473)
				p.Table_or_subquery()
			}

		case 2:
			{
				p.SetState(474)
				p.Join_clause()
			}

		}

	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWHERE_ {
		{
			p.SetState(479)
			p.Match(SQLiteParserWHERE_)
		}
		{
			p.SetState(480)

			var _x = p.expr(0)

			localctx.(*Select_coreContext).whereExpr = _x
		}

	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserGROUP_ {
		{
			p.SetState(483)
			p.Match(SQLiteParserGROUP_)
		}
		{
			p.SetState(484)
			p.Match(SQLiteParserBY_)
		}
		{
			p.SetState(485)

			var _x = p.expr(0)

			localctx.(*Select_coreContext)._expr = _x
		}
		localctx.(*Select_coreContext).groupByExpr = append(localctx.(*Select_coreContext).groupByExpr, localctx.(*Select_coreContext)._expr)
		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SQLiteParserCOMMA {
			{
				p.SetState(486)
				p.Match(SQLiteParserCOMMA)
			}
			{
				p.SetState(487)

				var _x = p.expr(0)

				localctx.(*Select_coreContext)._expr = _x
			}
			localctx.(*Select_coreContext).groupByExpr = append(localctx.(*Select_coreContext).groupByExpr, localctx.(*Select_coreContext)._expr)

			p.SetState(492)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(495)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserHAVING_ {
			{
				p.SetState(493)
				p.Match(SQLiteParserHAVING_)
			}
			{
				p.SetState(494)

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
	p.EnterRule(localctx, 40, SQLiteParserRULE_table_or_subquery)
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

	p.SetState(511)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Table_name()
		}
		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(500)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(501)
				p.Table_alias()
			}

		}

	case SQLiteParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(504)
			p.Match(SQLiteParserOPEN_PAR)
		}
		{
			p.SetState(505)
			p.Select_stmt_core()
		}
		{
			p.SetState(506)
			p.Match(SQLiteParserCLOSE_PAR)
		}
		p.SetState(509)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(507)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(508)
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
	p.EnterRule(localctx, 42, SQLiteParserRULE_result_column)
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

	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(513)
			p.Match(SQLiteParserSTAR)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(514)
			p.Table_name()
		}
		{
			p.SetState(515)
			p.Match(SQLiteParserDOT)
		}
		{
			p.SetState(516)
			p.Match(SQLiteParserSTAR)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(518)
			p.expr(0)
		}
		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(519)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(520)
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
	p.EnterRule(localctx, 44, SQLiteParserRULE_returning_clause_result_column)
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

	p.SetState(531)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(525)
			p.Match(SQLiteParserSTAR)
		}

	case SQLiteParserOPEN_PAR, SQLiteParserPLUS, SQLiteParserMINUS, SQLiteParserTILDE, SQLiteParserCASE_, SQLiteParserEXISTS_, SQLiteParserFALSE_, SQLiteParserGLOB_, SQLiteParserLIKE_, SQLiteParserNOT_, SQLiteParserNULL_, SQLiteParserREPLACE_, SQLiteParserTRUE_, SQLiteParserIDENTIFIER, SQLiteParserNUMERIC_LITERAL, SQLiteParserBIND_PARAMETER, SQLiteParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(526)
			p.expr(0)
		}
		p.SetState(529)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserAS_ {
			{
				p.SetState(527)
				p.Match(SQLiteParserAS_)
			}
			{
				p.SetState(528)
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
	p.EnterRule(localctx, 46, SQLiteParserRULE_join_operator)
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
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserNATURAL_ {
		{
			p.SetState(533)
			p.Match(SQLiteParserNATURAL_)
		}

	}
	p.SetState(541)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserFULL_, SQLiteParserLEFT_, SQLiteParserRIGHT_:
		{
			p.SetState(536)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-54)) & ^0x3f) == 0 && ((int64(1)<<(_la-54))&68719542273) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(538)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserOUTER_ {
			{
				p.SetState(537)
				p.Match(SQLiteParserOUTER_)
			}

		}

	case SQLiteParserINNER_:
		{
			p.SetState(540)
			p.Match(SQLiteParserINNER_)
		}

	case SQLiteParserJOIN_:

	default:
	}
	{
		p.SetState(543)
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
	p.EnterRule(localctx, 48, SQLiteParserRULE_join_constraint)

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
		p.SetState(545)
		p.Match(SQLiteParserON_)
	}
	{
		p.SetState(546)
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
	p.EnterRule(localctx, 50, SQLiteParserRULE_compound_operator)
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

	p.SetState(554)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserUNION_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(548)
			p.Match(SQLiteParserUNION_)
		}
		p.SetState(550)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SQLiteParserALL_ {
			{
				p.SetState(549)
				p.Match(SQLiteParserALL_)
			}

		}

	case SQLiteParserINTERSECT_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.Match(SQLiteParserINTERSECT_)
		}

	case SQLiteParserEXCEPT_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(553)
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
	p.EnterRule(localctx, 52, SQLiteParserRULE_update_set_subclause)

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
	p.SetState(558)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		{
			p.SetState(556)
			p.Column_name()
		}

	case SQLiteParserOPEN_PAR:
		{
			p.SetState(557)
			p.Column_name_list()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(560)
		p.Match(SQLiteParserASSIGN)
	}
	{
		p.SetState(561)
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
	p.EnterRule(localctx, 54, SQLiteParserRULE_update_stmt)
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
	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWITH_ {
		{
			p.SetState(563)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(566)
		p.Match(SQLiteParserUPDATE_)
	}
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserOR_ {
		{
			p.SetState(567)
			p.Match(SQLiteParserOR_)
		}
		{
			p.SetState(568)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577023702290399232) != 0) || _la == SQLiteParserREPLACE_ || _la == SQLiteParserROLLBACK_) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(571)
		p.Qualified_table_name()
	}
	{
		p.SetState(572)
		p.Match(SQLiteParserSET_)
	}
	{
		p.SetState(573)
		p.Update_set_subclause()
	}
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(574)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(575)
			p.Update_set_subclause()
		}

		p.SetState(580)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserFROM_ {
		{
			p.SetState(581)
			p.Match(SQLiteParserFROM_)
		}
		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(582)
				p.Table_or_subquery()
			}

		case 2:
			{
				p.SetState(583)
				p.Join_clause()
			}

		}

	}
	p.SetState(590)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserWHERE_ {
		{
			p.SetState(588)
			p.Match(SQLiteParserWHERE_)
		}
		{
			p.SetState(589)
			p.expr(0)
		}

	}
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserRETURNING_ {
		{
			p.SetState(592)
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
	p.EnterRule(localctx, 56, SQLiteParserRULE_column_name_list)
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
		p.SetState(595)
		p.Match(SQLiteParserOPEN_PAR)
	}
	{
		p.SetState(596)
		p.Column_name()
	}
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(597)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(598)
			p.Column_name()
		}

		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(604)
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
	p.EnterRule(localctx, 58, SQLiteParserRULE_qualified_table_name)
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
		p.SetState(606)
		p.Table_name()
	}
	p.SetState(609)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserAS_ {
		{
			p.SetState(607)
			p.Match(SQLiteParserAS_)
		}
		{
			p.SetState(608)
			p.Table_alias()
		}

	}
	p.SetState(616)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserINDEXED_:
		{
			p.SetState(611)
			p.Match(SQLiteParserINDEXED_)
		}
		{
			p.SetState(612)
			p.Match(SQLiteParserBY_)
		}
		{
			p.SetState(613)
			p.Index_name()
		}

	case SQLiteParserNOT_:
		{
			p.SetState(614)
			p.Match(SQLiteParserNOT_)
		}
		{
			p.SetState(615)
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
	p.EnterRule(localctx, 60, SQLiteParserRULE_order_by_stmt)
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
		p.SetState(618)
		p.Match(SQLiteParserORDER_)
	}
	{
		p.SetState(619)
		p.Match(SQLiteParserBY_)
	}
	{
		p.SetState(620)
		p.Ordering_term()
	}
	p.SetState(625)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SQLiteParserCOMMA {
		{
			p.SetState(621)
			p.Match(SQLiteParserCOMMA)
		}
		{
			p.SetState(622)
			p.Ordering_term()
		}

		p.SetState(627)
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
	p.EnterRule(localctx, 62, SQLiteParserRULE_limit_stmt)
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
		p.Match(SQLiteParserLIMIT_)
	}
	{
		p.SetState(629)
		p.expr(0)
	}
	p.SetState(632)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserCOMMA || _la == SQLiteParserOFFSET_ {
		{
			p.SetState(630)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLiteParserCOMMA || _la == SQLiteParserOFFSET_) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(631)
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
	p.EnterRule(localctx, 64, SQLiteParserRULE_ordering_term)
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
		p.SetState(634)
		p.expr(0)
	}
	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserCOLLATE_ {
		{
			p.SetState(635)
			p.Match(SQLiteParserCOLLATE_)
		}
		{
			p.SetState(636)
			p.Collation_name()
		}

	}
	p.SetState(640)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserASC_ || _la == SQLiteParserDESC_ {
		{
			p.SetState(639)
			p.Asc_desc()
		}

	}
	p.SetState(644)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SQLiteParserNULLS_ {
		{
			p.SetState(642)
			p.Match(SQLiteParserNULLS_)
		}
		{
			p.SetState(643)
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
	p.EnterRule(localctx, 66, SQLiteParserRULE_asc_desc)
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
		p.SetState(646)
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
	p.EnterRule(localctx, 68, SQLiteParserRULE_function_keyword)
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
		p.SetState(648)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&8590000129) != 0) {
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
	p.EnterRule(localctx, 70, SQLiteParserRULE_function_name)

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

	p.SetState(652)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SQLiteParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(650)
			p.Match(SQLiteParserIDENTIFIER)
		}

	case SQLiteParserGLOB_, SQLiteParserLIKE_, SQLiteParserREPLACE_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(651)
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
	p.EnterRule(localctx, 72, SQLiteParserRULE_table_name)

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
		p.SetState(654)
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
	p.EnterRule(localctx, 74, SQLiteParserRULE_table_alias)

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
	p.EnterRule(localctx, 76, SQLiteParserRULE_column_name)

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
		p.SetState(658)
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
	p.EnterRule(localctx, 78, SQLiteParserRULE_column_alias)

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
		p.SetState(660)
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
	p.EnterRule(localctx, 80, SQLiteParserRULE_collation_name)

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
		p.SetState(662)
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
	p.EnterRule(localctx, 82, SQLiteParserRULE_index_name)

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
		p.SetState(664)
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
