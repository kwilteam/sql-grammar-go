// Code generated from SQLParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package sqlgrammar // SQLParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SQLParser struct {
	*antlr.BaseParser
}

var SQLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlparserParserInit() {
	staticData := &SQLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'.'", "'('", "')'", "','", "'='", "'*'", "'+'", "'-'", "'/'",
		"'%'", "'<'", "'<='", "'>'", "'>='", "'!='", "'<>'", "'::'", "'ADD'",
		"'ALL'", "'AND'", "'ASC'", "'AS'", "'BETWEEN'", "'BY'", "'CASE'", "'COLLATE'",
		"'COMMIT'", "'CONFLICT'", "'CREATE'", "'CROSS'", "'DEFAULT'", "'DELETE'",
		"'DESC'", "'DISTINCT'", "'DO'", "'ELSE'", "'END'", "'ESCAPE'", "'EXCEPT'",
		"'EXISTS'", "'FALSE'", "'FILTER'", "'FIRST'", "'FROM'", "'FULL'", "'GROUPS'",
		"'GROUP'", "'HAVING'", "'INNER'", "'INSERT'", "'INTERSECT'", "'INTO'",
		"'IN'", "'ISNULL'", "'IS'", "'JOIN'", "'LAST'", "'LEFT'", "'LIKE'",
		"'LIMIT'", "'NOTHING'", "'NOTNULL'", "'NOT'", "'NULLS'", "'NULL'", "'OFFSET'",
		"'OF'", "'ON'", "'ORDER'", "'OR'", "'OUTER'", "'RAISE'", "'REPLACE'",
		"'RETURNING'", "'RIGHT'", "'SELECT'", "'SET'", "'THEN'", "'TRUE'", "'UNION'",
		"'UPDATE'", "'USING'", "'VALUES'", "'WHEN'", "'WHERE'", "'WITH'",
	}
	staticData.SymbolicNames = []string{
		"", "SCOL", "DOT", "OPEN_PAR", "CLOSE_PAR", "COMMA", "ASSIGN", "STAR",
		"PLUS", "MINUS", "DIV", "MOD", "LT", "LT_EQ", "GT", "GT_EQ", "NOT_EQ1",
		"NOT_EQ2", "TYPE_CAST", "ADD_", "ALL_", "AND_", "ASC_", "AS_", "BETWEEN_",
		"BY_", "CASE_", "COLLATE_", "COMMIT_", "CONFLICT_", "CREATE_", "CROSS_",
		"DEFAULT_", "DELETE_", "DESC_", "DISTINCT_", "DO_", "ELSE_", "END_",
		"ESCAPE_", "EXCEPT_", "EXISTS_", "FALSE_", "FILTER_", "FIRST_", "FROM_",
		"FULL_", "GROUPS_", "GROUP_", "HAVING_", "INNER_", "INSERT_", "INTERSECT_",
		"INTO_", "IN_", "ISNULL_", "IS_", "JOIN_", "LAST_", "LEFT_", "LIKE_",
		"LIMIT_", "NOTHING_", "NOTNULL_", "NOT_", "NULLS_", "NULL_", "OFFSET_",
		"OF_", "ON_", "ORDER_", "OR_", "OUTER_", "RAISE_", "REPLACE_", "RETURNING_",
		"RIGHT_", "SELECT_", "SET_", "THEN_", "TRUE_", "UNION_", "UPDATE_",
		"USING_", "VALUES_", "WHEN_", "WHERE_", "WITH_", "IDENTIFIER", "NUMERIC_LITERAL",
		"BIND_PARAMETER", "STRING_LITERAL", "BLOB_LITERAL", "SINGLE_LINE_COMMENT",
		"MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	}
	staticData.RuleNames = []string{
		"statements", "sql_stmt_list", "sql_stmt", "indexed_column", "cte_table_name",
		"common_table_expression", "common_table_stmt", "delete_stmt", "variable",
		"function_call", "column_ref", "when_clause", "expr", "subquery", "expr_list",
		"comparisonOperator", "cast_type", "type_cast", "boolean_value", "string_value",
		"numeric_value", "literal", "value_row", "values_clause", "insert_stmt",
		"returning_clause", "upsert_update", "upsert_clause", "select_stmt_core",
		"select_stmt", "join_clause", "select_core", "table_or_subquery", "result_column",
		"returning_clause_result_column", "join_operator", "join_constraint",
		"compound_operator", "update_set_subclause", "update_stmt", "column_name_list",
		"qualified_table_name", "order_by_stmt", "limit_stmt", "ordering_term",
		"asc_desc", "function_keyword", "function_name", "table_name", "table_alias",
		"column_name", "column_alias", "collation_name", "index_name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 96, 704, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 1, 0, 5, 0, 110, 8, 0, 10, 0, 12, 0, 113, 9, 0, 1, 0,
		1, 0, 1, 1, 5, 1, 118, 8, 1, 10, 1, 12, 1, 121, 9, 1, 1, 1, 1, 1, 4, 1,
		125, 8, 1, 11, 1, 12, 1, 126, 1, 1, 5, 1, 130, 8, 1, 10, 1, 12, 1, 133,
		9, 1, 1, 1, 5, 1, 136, 8, 1, 10, 1, 12, 1, 139, 9, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 145, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		154, 8, 4, 10, 4, 12, 4, 157, 9, 4, 1, 4, 1, 4, 3, 4, 161, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 173, 8, 6,
		10, 6, 12, 6, 176, 9, 6, 1, 7, 3, 7, 179, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 186, 8, 7, 1, 7, 3, 7, 189, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 3, 9, 196, 8, 9, 1, 9, 1, 9, 1, 9, 5, 9, 201, 8, 9, 10, 9, 12, 9, 204,
		9, 9, 1, 9, 3, 9, 207, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 214,
		8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 3, 12, 226, 8, 12, 1, 12, 1, 12, 3, 12, 230, 8, 12, 1, 12, 1, 12, 3,
		12, 234, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 242, 8,
		12, 1, 12, 3, 12, 245, 8, 12, 1, 12, 3, 12, 248, 8, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 253, 8, 12, 1, 12, 4, 12, 256, 8, 12, 11, 12, 12, 12, 257, 1,
		12, 1, 12, 3, 12, 262, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 272, 8, 12, 1, 12, 1, 12, 3, 12, 276, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 286, 8, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 308,
		8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 314, 8, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 323, 8, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 329, 8, 12, 1, 12, 1, 12, 1, 12, 3, 12, 334, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 341, 8, 12, 1, 12, 1, 12, 5, 12, 345,
		8, 12, 10, 12, 12, 12, 348, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 5, 14, 357, 8, 14, 10, 14, 12, 14, 360, 9, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 379, 8, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 5, 22, 385, 8, 22, 10, 22, 12, 22, 388, 9, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 5, 23, 396, 8, 23, 10, 23, 12, 23, 399, 9, 23,
		1, 24, 3, 24, 402, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 409,
		8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 415, 8, 24, 10, 24, 12, 24, 418,
		9, 24, 1, 24, 1, 24, 3, 24, 422, 8, 24, 1, 24, 1, 24, 3, 24, 426, 8, 24,
		1, 24, 3, 24, 429, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 435, 8, 25,
		10, 25, 12, 25, 438, 9, 25, 1, 26, 1, 26, 3, 26, 442, 8, 26, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 453, 8, 27,
		10, 27, 12, 27, 456, 9, 27, 1, 27, 1, 27, 1, 27, 3, 27, 461, 8, 27, 3,
		27, 463, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27,
		472, 8, 27, 10, 27, 12, 27, 475, 9, 27, 1, 27, 1, 27, 3, 27, 479, 8, 27,
		3, 27, 481, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 487, 8, 28, 10, 28,
		12, 28, 490, 9, 28, 1, 28, 3, 28, 493, 8, 28, 1, 28, 3, 28, 496, 8, 28,
		1, 29, 3, 29, 499, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 5, 30, 508, 8, 30, 10, 30, 12, 30, 511, 9, 30, 1, 31, 1, 31, 3, 31,
		515, 8, 31, 1, 31, 1, 31, 1, 31, 5, 31, 520, 8, 31, 10, 31, 12, 31, 523,
		9, 31, 1, 31, 1, 31, 1, 31, 3, 31, 528, 8, 31, 3, 31, 530, 8, 31, 1, 31,
		1, 31, 3, 31, 534, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 541,
		8, 31, 10, 31, 12, 31, 544, 9, 31, 1, 31, 1, 31, 3, 31, 548, 8, 31, 3,
		31, 550, 8, 31, 1, 32, 1, 32, 1, 32, 3, 32, 555, 8, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 3, 32, 562, 8, 32, 3, 32, 564, 8, 32, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 574, 8, 33, 3, 33, 576, 8,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 582, 8, 34, 3, 34, 584, 8, 34, 1,
		35, 1, 35, 3, 35, 588, 8, 35, 1, 35, 3, 35, 591, 8, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 3, 37, 600, 8, 37, 1, 37, 1, 37, 3, 37,
		604, 8, 37, 1, 38, 1, 38, 3, 38, 608, 8, 38, 1, 38, 1, 38, 1, 38, 1, 39,
		3, 39, 614, 8, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 622,
		8, 39, 10, 39, 12, 39, 625, 9, 39, 1, 39, 1, 39, 1, 39, 3, 39, 630, 8,
		39, 3, 39, 632, 8, 39, 1, 39, 1, 39, 3, 39, 636, 8, 39, 1, 39, 3, 39, 639,
		8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 645, 8, 40, 10, 40, 12, 40, 648,
		9, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 3, 41, 655, 8, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 5, 42, 662, 8, 42, 10, 42, 12, 42, 665, 9, 42,
		1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 671, 8, 43, 1, 44, 1, 44, 3, 44, 675,
		8, 44, 1, 44, 1, 44, 3, 44, 679, 8, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 3, 46, 686, 8, 46, 1, 47, 1, 47, 3, 47, 690, 8, 47, 1, 48, 1, 48, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53,
		0, 1, 24, 54, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102,
		104, 106, 0, 8, 1, 0, 8, 9, 2, 0, 7, 7, 10, 11, 2, 0, 55, 55, 63, 63, 2,
		0, 6, 6, 12, 17, 2, 0, 42, 42, 80, 80, 3, 0, 46, 46, 59, 59, 76, 76, 2,
		0, 44, 44, 58, 58, 2, 0, 22, 22, 34, 34, 770, 0, 111, 1, 0, 0, 0, 2, 119,
		1, 0, 0, 0, 4, 144, 1, 0, 0, 0, 6, 146, 1, 0, 0, 0, 8, 148, 1, 0, 0, 0,
		10, 162, 1, 0, 0, 0, 12, 168, 1, 0, 0, 0, 14, 178, 1, 0, 0, 0, 16, 190,
		1, 0, 0, 0, 18, 192, 1, 0, 0, 0, 20, 213, 1, 0, 0, 0, 22, 217, 1, 0, 0,
		0, 24, 275, 1, 0, 0, 0, 26, 349, 1, 0, 0, 0, 28, 353, 1, 0, 0, 0, 30, 361,
		1, 0, 0, 0, 32, 363, 1, 0, 0, 0, 34, 365, 1, 0, 0, 0, 36, 368, 1, 0, 0,
		0, 38, 370, 1, 0, 0, 0, 40, 372, 1, 0, 0, 0, 42, 378, 1, 0, 0, 0, 44, 380,
		1, 0, 0, 0, 46, 391, 1, 0, 0, 0, 48, 401, 1, 0, 0, 0, 50, 430, 1, 0, 0,
		0, 52, 441, 1, 0, 0, 0, 54, 446, 1, 0, 0, 0, 56, 482, 1, 0, 0, 0, 58, 498,
		1, 0, 0, 0, 60, 502, 1, 0, 0, 0, 62, 512, 1, 0, 0, 0, 64, 563, 1, 0, 0,
		0, 66, 575, 1, 0, 0, 0, 68, 583, 1, 0, 0, 0, 70, 590, 1, 0, 0, 0, 72, 594,
		1, 0, 0, 0, 74, 603, 1, 0, 0, 0, 76, 607, 1, 0, 0, 0, 78, 613, 1, 0, 0,
		0, 80, 640, 1, 0, 0, 0, 82, 651, 1, 0, 0, 0, 84, 656, 1, 0, 0, 0, 86, 666,
		1, 0, 0, 0, 88, 672, 1, 0, 0, 0, 90, 680, 1, 0, 0, 0, 92, 685, 1, 0, 0,
		0, 94, 689, 1, 0, 0, 0, 96, 691, 1, 0, 0, 0, 98, 693, 1, 0, 0, 0, 100,
		695, 1, 0, 0, 0, 102, 697, 1, 0, 0, 0, 104, 699, 1, 0, 0, 0, 106, 701,
		1, 0, 0, 0, 108, 110, 3, 2, 1, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0,
		0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0,
		113, 111, 1, 0, 0, 0, 114, 115, 5, 0, 0, 1, 115, 1, 1, 0, 0, 0, 116, 118,
		5, 1, 0, 0, 117, 116, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0,
		0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		122, 131, 3, 4, 2, 0, 123, 125, 5, 1, 0, 0, 124, 123, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 130, 3, 4, 2, 0, 129, 124, 1, 0, 0, 0, 130, 133, 1, 0,
		0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 137, 1, 0, 0, 0,
		133, 131, 1, 0, 0, 0, 134, 136, 5, 1, 0, 0, 135, 134, 1, 0, 0, 0, 136,
		139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 3, 1,
		0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 145, 3, 14, 7, 0, 141, 145, 3, 48,
		24, 0, 142, 145, 3, 58, 29, 0, 143, 145, 3, 78, 39, 0, 144, 140, 1, 0,
		0, 0, 144, 141, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 143, 1, 0, 0, 0,
		145, 5, 1, 0, 0, 0, 146, 147, 3, 100, 50, 0, 147, 7, 1, 0, 0, 0, 148, 160,
		3, 96, 48, 0, 149, 150, 5, 3, 0, 0, 150, 155, 3, 100, 50, 0, 151, 152,
		5, 5, 0, 0, 152, 154, 3, 100, 50, 0, 153, 151, 1, 0, 0, 0, 154, 157, 1,
		0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0,
		0, 157, 155, 1, 0, 0, 0, 158, 159, 5, 4, 0, 0, 159, 161, 1, 0, 0, 0, 160,
		149, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 9, 1, 0, 0, 0, 162, 163, 3,
		8, 4, 0, 163, 164, 5, 23, 0, 0, 164, 165, 5, 3, 0, 0, 165, 166, 3, 56,
		28, 0, 166, 167, 5, 4, 0, 0, 167, 11, 1, 0, 0, 0, 168, 169, 5, 87, 0, 0,
		169, 174, 3, 10, 5, 0, 170, 171, 5, 5, 0, 0, 171, 173, 3, 10, 5, 0, 172,
		170, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175,
		1, 0, 0, 0, 175, 13, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 179, 3, 12,
		6, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0,
		180, 181, 5, 33, 0, 0, 181, 182, 5, 45, 0, 0, 182, 185, 3, 82, 41, 0, 183,
		184, 5, 86, 0, 0, 184, 186, 3, 24, 12, 0, 185, 183, 1, 0, 0, 0, 185, 186,
		1, 0, 0, 0, 186, 188, 1, 0, 0, 0, 187, 189, 3, 50, 25, 0, 188, 187, 1,
		0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 15, 1, 0, 0, 0, 190, 191, 5, 90, 0,
		0, 191, 17, 1, 0, 0, 0, 192, 193, 3, 94, 47, 0, 193, 206, 5, 3, 0, 0, 194,
		196, 5, 35, 0, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197,
		1, 0, 0, 0, 197, 202, 3, 24, 12, 0, 198, 199, 5, 5, 0, 0, 199, 201, 3,
		24, 12, 0, 200, 198, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0,
		0, 0, 202, 203, 1, 0, 0, 0, 203, 207, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0,
		205, 207, 5, 7, 0, 0, 206, 195, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 5, 4, 0, 0, 209, 19, 1,
		0, 0, 0, 210, 211, 3, 96, 48, 0, 211, 212, 5, 2, 0, 0, 212, 214, 1, 0,
		0, 0, 213, 210, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 216, 3, 100, 50, 0, 216, 21, 1, 0, 0, 0, 217, 218, 5, 85, 0, 0, 218,
		219, 3, 24, 12, 0, 219, 220, 5, 79, 0, 0, 220, 221, 3, 24, 12, 0, 221,
		23, 1, 0, 0, 0, 222, 223, 6, 12, -1, 0, 223, 225, 3, 42, 21, 0, 224, 226,
		3, 34, 17, 0, 225, 224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 276, 1,
		0, 0, 0, 227, 229, 3, 16, 8, 0, 228, 230, 3, 34, 17, 0, 229, 228, 1, 0,
		0, 0, 229, 230, 1, 0, 0, 0, 230, 276, 1, 0, 0, 0, 231, 233, 3, 20, 10,
		0, 232, 234, 3, 34, 17, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0,
		234, 276, 1, 0, 0, 0, 235, 236, 7, 0, 0, 0, 236, 276, 3, 24, 12, 19, 237,
		238, 5, 3, 0, 0, 238, 239, 3, 24, 12, 0, 239, 241, 5, 4, 0, 0, 240, 242,
		3, 34, 17, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 276, 1,
		0, 0, 0, 243, 245, 5, 64, 0, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1, 0, 0,
		0, 245, 246, 1, 0, 0, 0, 246, 248, 5, 41, 0, 0, 247, 244, 1, 0, 0, 0, 247,
		248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 276, 3, 26, 13, 0, 250, 252,
		5, 26, 0, 0, 251, 253, 3, 24, 12, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1,
		0, 0, 0, 253, 255, 1, 0, 0, 0, 254, 256, 3, 22, 11, 0, 255, 254, 1, 0,
		0, 0, 256, 257, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0,
		258, 261, 1, 0, 0, 0, 259, 260, 5, 37, 0, 0, 260, 262, 3, 24, 12, 0, 261,
		259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264,
		5, 38, 0, 0, 264, 276, 1, 0, 0, 0, 265, 266, 5, 3, 0, 0, 266, 267, 3, 28,
		14, 0, 267, 268, 5, 4, 0, 0, 268, 276, 1, 0, 0, 0, 269, 271, 3, 18, 9,
		0, 270, 272, 3, 34, 17, 0, 271, 270, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0,
		272, 276, 1, 0, 0, 0, 273, 274, 5, 64, 0, 0, 274, 276, 3, 24, 12, 3, 275,
		222, 1, 0, 0, 0, 275, 227, 1, 0, 0, 0, 275, 231, 1, 0, 0, 0, 275, 235,
		1, 0, 0, 0, 275, 237, 1, 0, 0, 0, 275, 247, 1, 0, 0, 0, 275, 250, 1, 0,
		0, 0, 275, 265, 1, 0, 0, 0, 275, 269, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0,
		276, 346, 1, 0, 0, 0, 277, 278, 10, 12, 0, 0, 278, 279, 7, 1, 0, 0, 279,
		345, 3, 24, 12, 13, 280, 281, 10, 11, 0, 0, 281, 282, 7, 0, 0, 0, 282,
		345, 3, 24, 12, 12, 283, 285, 10, 8, 0, 0, 284, 286, 5, 64, 0, 0, 285,
		284, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288,
		5, 24, 0, 0, 288, 289, 3, 24, 12, 0, 289, 290, 5, 21, 0, 0, 290, 291, 3,
		24, 12, 9, 291, 345, 1, 0, 0, 0, 292, 293, 10, 6, 0, 0, 293, 294, 3, 30,
		15, 0, 294, 295, 3, 24, 12, 7, 295, 345, 1, 0, 0, 0, 296, 297, 10, 2, 0,
		0, 297, 298, 5, 21, 0, 0, 298, 345, 3, 24, 12, 3, 299, 300, 10, 1, 0, 0,
		300, 301, 5, 71, 0, 0, 301, 345, 3, 24, 12, 2, 302, 303, 10, 18, 0, 0,
		303, 304, 5, 27, 0, 0, 304, 345, 3, 104, 52, 0, 305, 307, 10, 10, 0, 0,
		306, 308, 5, 64, 0, 0, 307, 306, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308,
		309, 1, 0, 0, 0, 309, 310, 5, 54, 0, 0, 310, 345, 3, 26, 13, 0, 311, 313,
		10, 9, 0, 0, 312, 314, 5, 64, 0, 0, 313, 312, 1, 0, 0, 0, 313, 314, 1,
		0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 5, 54, 0, 0, 316, 317, 5, 3, 0,
		0, 317, 318, 3, 28, 14, 0, 318, 319, 5, 4, 0, 0, 319, 345, 1, 0, 0, 0,
		320, 322, 10, 7, 0, 0, 321, 323, 5, 64, 0, 0, 322, 321, 1, 0, 0, 0, 322,
		323, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325, 5, 60, 0, 0, 325, 328,
		3, 24, 12, 0, 326, 327, 5, 39, 0, 0, 327, 329, 3, 24, 12, 0, 328, 326,
		1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 345, 1, 0, 0, 0, 330, 331, 10, 5,
		0, 0, 331, 333, 5, 56, 0, 0, 332, 334, 5, 64, 0, 0, 333, 332, 1, 0, 0,
		0, 333, 334, 1, 0, 0, 0, 334, 340, 1, 0, 0, 0, 335, 336, 5, 35, 0, 0, 336,
		337, 5, 45, 0, 0, 337, 341, 3, 24, 12, 0, 338, 341, 3, 36, 18, 0, 339,
		341, 5, 66, 0, 0, 340, 335, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 340, 339,
		1, 0, 0, 0, 341, 345, 1, 0, 0, 0, 342, 343, 10, 4, 0, 0, 343, 345, 7, 2,
		0, 0, 344, 277, 1, 0, 0, 0, 344, 280, 1, 0, 0, 0, 344, 283, 1, 0, 0, 0,
		344, 292, 1, 0, 0, 0, 344, 296, 1, 0, 0, 0, 344, 299, 1, 0, 0, 0, 344,
		302, 1, 0, 0, 0, 344, 305, 1, 0, 0, 0, 344, 311, 1, 0, 0, 0, 344, 320,
		1, 0, 0, 0, 344, 330, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 345, 348, 1, 0,
		0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 25, 1, 0, 0, 0,
		348, 346, 1, 0, 0, 0, 349, 350, 5, 3, 0, 0, 350, 351, 3, 56, 28, 0, 351,
		352, 5, 4, 0, 0, 352, 27, 1, 0, 0, 0, 353, 358, 3, 24, 12, 0, 354, 355,
		5, 5, 0, 0, 355, 357, 3, 24, 12, 0, 356, 354, 1, 0, 0, 0, 357, 360, 1,
		0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 29, 1, 0, 0,
		0, 360, 358, 1, 0, 0, 0, 361, 362, 7, 3, 0, 0, 362, 31, 1, 0, 0, 0, 363,
		364, 5, 88, 0, 0, 364, 33, 1, 0, 0, 0, 365, 366, 5, 18, 0, 0, 366, 367,
		3, 32, 16, 0, 367, 35, 1, 0, 0, 0, 368, 369, 7, 4, 0, 0, 369, 37, 1, 0,
		0, 0, 370, 371, 5, 91, 0, 0, 371, 39, 1, 0, 0, 0, 372, 373, 5, 89, 0, 0,
		373, 41, 1, 0, 0, 0, 374, 379, 5, 66, 0, 0, 375, 379, 3, 36, 18, 0, 376,
		379, 3, 38, 19, 0, 377, 379, 3, 40, 20, 0, 378, 374, 1, 0, 0, 0, 378, 375,
		1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 377, 1, 0, 0, 0, 379, 43, 1, 0,
		0, 0, 380, 381, 5, 3, 0, 0, 381, 386, 3, 24, 12, 0, 382, 383, 5, 5, 0,
		0, 383, 385, 3, 24, 12, 0, 384, 382, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0,
		386, 384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 389, 1, 0, 0, 0, 388,
		386, 1, 0, 0, 0, 389, 390, 5, 4, 0, 0, 390, 45, 1, 0, 0, 0, 391, 392, 5,
		84, 0, 0, 392, 397, 3, 44, 22, 0, 393, 394, 5, 5, 0, 0, 394, 396, 3, 44,
		22, 0, 395, 393, 1, 0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0,
		397, 398, 1, 0, 0, 0, 398, 47, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 400, 402,
		3, 12, 6, 0, 401, 400, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 403, 1, 0,
		0, 0, 403, 404, 5, 51, 0, 0, 404, 405, 5, 53, 0, 0, 405, 408, 3, 96, 48,
		0, 406, 407, 5, 23, 0, 0, 407, 409, 3, 98, 49, 0, 408, 406, 1, 0, 0, 0,
		408, 409, 1, 0, 0, 0, 409, 421, 1, 0, 0, 0, 410, 411, 5, 3, 0, 0, 411,
		416, 3, 100, 50, 0, 412, 413, 5, 5, 0, 0, 413, 415, 3, 100, 50, 0, 414,
		412, 1, 0, 0, 0, 415, 418, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 416, 417,
		1, 0, 0, 0, 417, 419, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 419, 420, 5, 4,
		0, 0, 420, 422, 1, 0, 0, 0, 421, 410, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0,
		422, 423, 1, 0, 0, 0, 423, 425, 3, 46, 23, 0, 424, 426, 3, 54, 27, 0, 425,
		424, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 428, 1, 0, 0, 0, 427, 429,
		3, 50, 25, 0, 428, 427, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 49, 1, 0,
		0, 0, 430, 431, 5, 75, 0, 0, 431, 436, 3, 68, 34, 0, 432, 433, 5, 5, 0,
		0, 433, 435, 3, 68, 34, 0, 434, 432, 1, 0, 0, 0, 435, 438, 1, 0, 0, 0,
		436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 51, 1, 0, 0, 0, 438, 436,
		1, 0, 0, 0, 439, 442, 3, 100, 50, 0, 440, 442, 3, 80, 40, 0, 441, 439,
		1, 0, 0, 0, 441, 440, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 444, 5, 6,
		0, 0, 444, 445, 3, 24, 12, 0, 445, 53, 1, 0, 0, 0, 446, 447, 5, 69, 0,
		0, 447, 462, 5, 29, 0, 0, 448, 449, 5, 3, 0, 0, 449, 454, 3, 6, 3, 0, 450,
		451, 5, 5, 0, 0, 451, 453, 3, 6, 3, 0, 452, 450, 1, 0, 0, 0, 453, 456,
		1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 457, 1, 0,
		0, 0, 456, 454, 1, 0, 0, 0, 457, 460, 5, 4, 0, 0, 458, 459, 5, 86, 0, 0,
		459, 461, 3, 24, 12, 0, 460, 458, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0, 461,
		463, 1, 0, 0, 0, 462, 448, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 464,
		1, 0, 0, 0, 464, 480, 5, 36, 0, 0, 465, 481, 5, 62, 0, 0, 466, 467, 5,
		82, 0, 0, 467, 468, 5, 78, 0, 0, 468, 473, 3, 52, 26, 0, 469, 470, 5, 5,
		0, 0, 470, 472, 3, 52, 26, 0, 471, 469, 1, 0, 0, 0, 472, 475, 1, 0, 0,
		0, 473, 471, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 478, 1, 0, 0, 0, 475,
		473, 1, 0, 0, 0, 476, 477, 5, 86, 0, 0, 477, 479, 3, 24, 12, 0, 478, 476,
		1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479, 481, 1, 0, 0, 0, 480, 465, 1, 0,
		0, 0, 480, 466, 1, 0, 0, 0, 481, 55, 1, 0, 0, 0, 482, 488, 3, 62, 31, 0,
		483, 484, 3, 74, 37, 0, 484, 485, 3, 62, 31, 0, 485, 487, 1, 0, 0, 0, 486,
		483, 1, 0, 0, 0, 487, 490, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 488, 489,
		1, 0, 0, 0, 489, 492, 1, 0, 0, 0, 490, 488, 1, 0, 0, 0, 491, 493, 3, 84,
		42, 0, 492, 491, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 495, 1, 0, 0, 0,
		494, 496, 3, 86, 43, 0, 495, 494, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496,
		57, 1, 0, 0, 0, 497, 499, 3, 12, 6, 0, 498, 497, 1, 0, 0, 0, 498, 499,
		1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501, 3, 56, 28, 0, 501, 59, 1, 0,
		0, 0, 502, 509, 3, 64, 32, 0, 503, 504, 3, 70, 35, 0, 504, 505, 3, 64,
		32, 0, 505, 506, 3, 72, 36, 0, 506, 508, 1, 0, 0, 0, 507, 503, 1, 0, 0,
		0, 508, 511, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510,
		61, 1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 512, 514, 5, 77, 0, 0, 513, 515,
		5, 35, 0, 0, 514, 513, 1, 0, 0, 0, 514, 515, 1, 0, 0, 0, 515, 516, 1, 0,
		0, 0, 516, 521, 3, 66, 33, 0, 517, 518, 5, 5, 0, 0, 518, 520, 3, 66, 33,
		0, 519, 517, 1, 0, 0, 0, 520, 523, 1, 0, 0, 0, 521, 519, 1, 0, 0, 0, 521,
		522, 1, 0, 0, 0, 522, 529, 1, 0, 0, 0, 523, 521, 1, 0, 0, 0, 524, 527,
		5, 45, 0, 0, 525, 528, 3, 64, 32, 0, 526, 528, 3, 60, 30, 0, 527, 525,
		1, 0, 0, 0, 527, 526, 1, 0, 0, 0, 528, 530, 1, 0, 0, 0, 529, 524, 1, 0,
		0, 0, 529, 530, 1, 0, 0, 0, 530, 533, 1, 0, 0, 0, 531, 532, 5, 86, 0, 0,
		532, 534, 3, 24, 12, 0, 533, 531, 1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534,
		549, 1, 0, 0, 0, 535, 536, 5, 48, 0, 0, 536, 537, 5, 25, 0, 0, 537, 542,
		3, 24, 12, 0, 538, 539, 5, 5, 0, 0, 539, 541, 3, 24, 12, 0, 540, 538, 1,
		0, 0, 0, 541, 544, 1, 0, 0, 0, 542, 540, 1, 0, 0, 0, 542, 543, 1, 0, 0,
		0, 543, 547, 1, 0, 0, 0, 544, 542, 1, 0, 0, 0, 545, 546, 5, 49, 0, 0, 546,
		548, 3, 24, 12, 0, 547, 545, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 550,
		1, 0, 0, 0, 549, 535, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 63, 1, 0,
		0, 0, 551, 554, 3, 96, 48, 0, 552, 553, 5, 23, 0, 0, 553, 555, 3, 98, 49,
		0, 554, 552, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555, 564, 1, 0, 0, 0, 556,
		557, 5, 3, 0, 0, 557, 558, 3, 56, 28, 0, 558, 561, 5, 4, 0, 0, 559, 560,
		5, 23, 0, 0, 560, 562, 3, 98, 49, 0, 561, 559, 1, 0, 0, 0, 561, 562, 1,
		0, 0, 0, 562, 564, 1, 0, 0, 0, 563, 551, 1, 0, 0, 0, 563, 556, 1, 0, 0,
		0, 564, 65, 1, 0, 0, 0, 565, 576, 5, 7, 0, 0, 566, 567, 3, 96, 48, 0, 567,
		568, 5, 2, 0, 0, 568, 569, 5, 7, 0, 0, 569, 576, 1, 0, 0, 0, 570, 573,
		3, 24, 12, 0, 571, 572, 5, 23, 0, 0, 572, 574, 3, 102, 51, 0, 573, 571,
		1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 576, 1, 0, 0, 0, 575, 565, 1, 0,
		0, 0, 575, 566, 1, 0, 0, 0, 575, 570, 1, 0, 0, 0, 576, 67, 1, 0, 0, 0,
		577, 584, 5, 7, 0, 0, 578, 581, 3, 24, 12, 0, 579, 580, 5, 23, 0, 0, 580,
		582, 3, 102, 51, 0, 581, 579, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582, 584,
		1, 0, 0, 0, 583, 577, 1, 0, 0, 0, 583, 578, 1, 0, 0, 0, 584, 69, 1, 0,
		0, 0, 585, 587, 7, 5, 0, 0, 586, 588, 5, 72, 0, 0, 587, 586, 1, 0, 0, 0,
		587, 588, 1, 0, 0, 0, 588, 591, 1, 0, 0, 0, 589, 591, 5, 50, 0, 0, 590,
		585, 1, 0, 0, 0, 590, 589, 1, 0, 0, 0, 590, 591, 1, 0, 0, 0, 591, 592,
		1, 0, 0, 0, 592, 593, 5, 57, 0, 0, 593, 71, 1, 0, 0, 0, 594, 595, 5, 69,
		0, 0, 595, 596, 3, 24, 12, 0, 596, 73, 1, 0, 0, 0, 597, 599, 5, 81, 0,
		0, 598, 600, 5, 20, 0, 0, 599, 598, 1, 0, 0, 0, 599, 600, 1, 0, 0, 0, 600,
		604, 1, 0, 0, 0, 601, 604, 5, 52, 0, 0, 602, 604, 5, 40, 0, 0, 603, 597,
		1, 0, 0, 0, 603, 601, 1, 0, 0, 0, 603, 602, 1, 0, 0, 0, 604, 75, 1, 0,
		0, 0, 605, 608, 3, 100, 50, 0, 606, 608, 3, 80, 40, 0, 607, 605, 1, 0,
		0, 0, 607, 606, 1, 0, 0, 0, 608, 609, 1, 0, 0, 0, 609, 610, 5, 6, 0, 0,
		610, 611, 3, 24, 12, 0, 611, 77, 1, 0, 0, 0, 612, 614, 3, 12, 6, 0, 613,
		612, 1, 0, 0, 0, 613, 614, 1, 0, 0, 0, 614, 615, 1, 0, 0, 0, 615, 616,
		5, 82, 0, 0, 616, 617, 3, 82, 41, 0, 617, 618, 5, 78, 0, 0, 618, 623, 3,
		76, 38, 0, 619, 620, 5, 5, 0, 0, 620, 622, 3, 76, 38, 0, 621, 619, 1, 0,
		0, 0, 622, 625, 1, 0, 0, 0, 623, 621, 1, 0, 0, 0, 623, 624, 1, 0, 0, 0,
		624, 631, 1, 0, 0, 0, 625, 623, 1, 0, 0, 0, 626, 629, 5, 45, 0, 0, 627,
		630, 3, 64, 32, 0, 628, 630, 3, 60, 30, 0, 629, 627, 1, 0, 0, 0, 629, 628,
		1, 0, 0, 0, 630, 632, 1, 0, 0, 0, 631, 626, 1, 0, 0, 0, 631, 632, 1, 0,
		0, 0, 632, 635, 1, 0, 0, 0, 633, 634, 5, 86, 0, 0, 634, 636, 3, 24, 12,
		0, 635, 633, 1, 0, 0, 0, 635, 636, 1, 0, 0, 0, 636, 638, 1, 0, 0, 0, 637,
		639, 3, 50, 25, 0, 638, 637, 1, 0, 0, 0, 638, 639, 1, 0, 0, 0, 639, 79,
		1, 0, 0, 0, 640, 641, 5, 3, 0, 0, 641, 646, 3, 100, 50, 0, 642, 643, 5,
		5, 0, 0, 643, 645, 3, 100, 50, 0, 644, 642, 1, 0, 0, 0, 645, 648, 1, 0,
		0, 0, 646, 644, 1, 0, 0, 0, 646, 647, 1, 0, 0, 0, 647, 649, 1, 0, 0, 0,
		648, 646, 1, 0, 0, 0, 649, 650, 5, 4, 0, 0, 650, 81, 1, 0, 0, 0, 651, 654,
		3, 96, 48, 0, 652, 653, 5, 23, 0, 0, 653, 655, 3, 98, 49, 0, 654, 652,
		1, 0, 0, 0, 654, 655, 1, 0, 0, 0, 655, 83, 1, 0, 0, 0, 656, 657, 5, 70,
		0, 0, 657, 658, 5, 25, 0, 0, 658, 663, 3, 88, 44, 0, 659, 660, 5, 5, 0,
		0, 660, 662, 3, 88, 44, 0, 661, 659, 1, 0, 0, 0, 662, 665, 1, 0, 0, 0,
		663, 661, 1, 0, 0, 0, 663, 664, 1, 0, 0, 0, 664, 85, 1, 0, 0, 0, 665, 663,
		1, 0, 0, 0, 666, 667, 5, 61, 0, 0, 667, 670, 3, 24, 12, 0, 668, 669, 5,
		67, 0, 0, 669, 671, 3, 24, 12, 0, 670, 668, 1, 0, 0, 0, 670, 671, 1, 0,
		0, 0, 671, 87, 1, 0, 0, 0, 672, 674, 3, 24, 12, 0, 673, 675, 3, 90, 45,
		0, 674, 673, 1, 0, 0, 0, 674, 675, 1, 0, 0, 0, 675, 678, 1, 0, 0, 0, 676,
		677, 5, 65, 0, 0, 677, 679, 7, 6, 0, 0, 678, 676, 1, 0, 0, 0, 678, 679,
		1, 0, 0, 0, 679, 89, 1, 0, 0, 0, 680, 681, 7, 7, 0, 0, 681, 91, 1, 0, 0,
		0, 682, 686, 1, 0, 0, 0, 683, 686, 5, 60, 0, 0, 684, 686, 5, 74, 0, 0,
		685, 682, 1, 0, 0, 0, 685, 683, 1, 0, 0, 0, 685, 684, 1, 0, 0, 0, 686,
		93, 1, 0, 0, 0, 687, 690, 5, 88, 0, 0, 688, 690, 3, 92, 46, 0, 689, 687,
		1, 0, 0, 0, 689, 688, 1, 0, 0, 0, 690, 95, 1, 0, 0, 0, 691, 692, 5, 88,
		0, 0, 692, 97, 1, 0, 0, 0, 693, 694, 5, 88, 0, 0, 694, 99, 1, 0, 0, 0,
		695, 696, 5, 88, 0, 0, 696, 101, 1, 0, 0, 0, 697, 698, 5, 88, 0, 0, 698,
		103, 1, 0, 0, 0, 699, 700, 5, 88, 0, 0, 700, 105, 1, 0, 0, 0, 701, 702,
		5, 88, 0, 0, 702, 107, 1, 0, 0, 0, 93, 111, 119, 126, 131, 137, 144, 155,
		160, 174, 178, 185, 188, 195, 202, 206, 213, 225, 229, 233, 241, 244, 247,
		252, 257, 261, 271, 275, 285, 307, 313, 322, 328, 333, 340, 344, 346, 358,
		378, 386, 397, 401, 408, 416, 421, 425, 428, 436, 441, 454, 460, 462, 473,
		478, 480, 488, 492, 495, 498, 509, 514, 521, 527, 529, 533, 542, 547, 549,
		554, 561, 563, 573, 575, 581, 583, 587, 590, 599, 603, 607, 613, 623, 629,
		631, 635, 638, 646, 654, 663, 670, 674, 678, 685, 689,
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

// SQLParserInit initializes any static state used to implement SQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SQLParserInit() {
	staticData := &SQLParserParserStaticData
	staticData.once.Do(sqlparserParserInit)
}

// NewSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSQLParser(input antlr.TokenStream) *SQLParser {
	SQLParserInit()
	this := new(SQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SQLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SQLParser.g4"

	return this
}

// SQLParser tokens.
const (
	SQLParserEOF                 = antlr.TokenEOF
	SQLParserSCOL                = 1
	SQLParserDOT                 = 2
	SQLParserOPEN_PAR            = 3
	SQLParserCLOSE_PAR           = 4
	SQLParserCOMMA               = 5
	SQLParserASSIGN              = 6
	SQLParserSTAR                = 7
	SQLParserPLUS                = 8
	SQLParserMINUS               = 9
	SQLParserDIV                 = 10
	SQLParserMOD                 = 11
	SQLParserLT                  = 12
	SQLParserLT_EQ               = 13
	SQLParserGT                  = 14
	SQLParserGT_EQ               = 15
	SQLParserNOT_EQ1             = 16
	SQLParserNOT_EQ2             = 17
	SQLParserTYPE_CAST           = 18
	SQLParserADD_                = 19
	SQLParserALL_                = 20
	SQLParserAND_                = 21
	SQLParserASC_                = 22
	SQLParserAS_                 = 23
	SQLParserBETWEEN_            = 24
	SQLParserBY_                 = 25
	SQLParserCASE_               = 26
	SQLParserCOLLATE_            = 27
	SQLParserCOMMIT_             = 28
	SQLParserCONFLICT_           = 29
	SQLParserCREATE_             = 30
	SQLParserCROSS_              = 31
	SQLParserDEFAULT_            = 32
	SQLParserDELETE_             = 33
	SQLParserDESC_               = 34
	SQLParserDISTINCT_           = 35
	SQLParserDO_                 = 36
	SQLParserELSE_               = 37
	SQLParserEND_                = 38
	SQLParserESCAPE_             = 39
	SQLParserEXCEPT_             = 40
	SQLParserEXISTS_             = 41
	SQLParserFALSE_              = 42
	SQLParserFILTER_             = 43
	SQLParserFIRST_              = 44
	SQLParserFROM_               = 45
	SQLParserFULL_               = 46
	SQLParserGROUPS_             = 47
	SQLParserGROUP_              = 48
	SQLParserHAVING_             = 49
	SQLParserINNER_              = 50
	SQLParserINSERT_             = 51
	SQLParserINTERSECT_          = 52
	SQLParserINTO_               = 53
	SQLParserIN_                 = 54
	SQLParserISNULL_             = 55
	SQLParserIS_                 = 56
	SQLParserJOIN_               = 57
	SQLParserLAST_               = 58
	SQLParserLEFT_               = 59
	SQLParserLIKE_               = 60
	SQLParserLIMIT_              = 61
	SQLParserNOTHING_            = 62
	SQLParserNOTNULL_            = 63
	SQLParserNOT_                = 64
	SQLParserNULLS_              = 65
	SQLParserNULL_               = 66
	SQLParserOFFSET_             = 67
	SQLParserOF_                 = 68
	SQLParserON_                 = 69
	SQLParserORDER_              = 70
	SQLParserOR_                 = 71
	SQLParserOUTER_              = 72
	SQLParserRAISE_              = 73
	SQLParserREPLACE_            = 74
	SQLParserRETURNING_          = 75
	SQLParserRIGHT_              = 76
	SQLParserSELECT_             = 77
	SQLParserSET_                = 78
	SQLParserTHEN_               = 79
	SQLParserTRUE_               = 80
	SQLParserUNION_              = 81
	SQLParserUPDATE_             = 82
	SQLParserUSING_              = 83
	SQLParserVALUES_             = 84
	SQLParserWHEN_               = 85
	SQLParserWHERE_              = 86
	SQLParserWITH_               = 87
	SQLParserIDENTIFIER          = 88
	SQLParserNUMERIC_LITERAL     = 89
	SQLParserBIND_PARAMETER      = 90
	SQLParserSTRING_LITERAL      = 91
	SQLParserBLOB_LITERAL        = 92
	SQLParserSINGLE_LINE_COMMENT = 93
	SQLParserMULTILINE_COMMENT   = 94
	SQLParserSPACES              = 95
	SQLParserUNEXPECTED_CHAR     = 96
)

// SQLParser rules.
const (
	SQLParserRULE_statements                     = 0
	SQLParserRULE_sql_stmt_list                  = 1
	SQLParserRULE_sql_stmt                       = 2
	SQLParserRULE_indexed_column                 = 3
	SQLParserRULE_cte_table_name                 = 4
	SQLParserRULE_common_table_expression        = 5
	SQLParserRULE_common_table_stmt              = 6
	SQLParserRULE_delete_stmt                    = 7
	SQLParserRULE_variable                       = 8
	SQLParserRULE_function_call                  = 9
	SQLParserRULE_column_ref                     = 10
	SQLParserRULE_when_clause                    = 11
	SQLParserRULE_expr                           = 12
	SQLParserRULE_subquery                       = 13
	SQLParserRULE_expr_list                      = 14
	SQLParserRULE_comparisonOperator             = 15
	SQLParserRULE_cast_type                      = 16
	SQLParserRULE_type_cast                      = 17
	SQLParserRULE_boolean_value                  = 18
	SQLParserRULE_string_value                   = 19
	SQLParserRULE_numeric_value                  = 20
	SQLParserRULE_literal                        = 21
	SQLParserRULE_value_row                      = 22
	SQLParserRULE_values_clause                  = 23
	SQLParserRULE_insert_stmt                    = 24
	SQLParserRULE_returning_clause               = 25
	SQLParserRULE_upsert_update                  = 26
	SQLParserRULE_upsert_clause                  = 27
	SQLParserRULE_select_stmt_core               = 28
	SQLParserRULE_select_stmt                    = 29
	SQLParserRULE_join_clause                    = 30
	SQLParserRULE_select_core                    = 31
	SQLParserRULE_table_or_subquery              = 32
	SQLParserRULE_result_column                  = 33
	SQLParserRULE_returning_clause_result_column = 34
	SQLParserRULE_join_operator                  = 35
	SQLParserRULE_join_constraint                = 36
	SQLParserRULE_compound_operator              = 37
	SQLParserRULE_update_set_subclause           = 38
	SQLParserRULE_update_stmt                    = 39
	SQLParserRULE_column_name_list               = 40
	SQLParserRULE_qualified_table_name           = 41
	SQLParserRULE_order_by_stmt                  = 42
	SQLParserRULE_limit_stmt                     = 43
	SQLParserRULE_ordering_term                  = 44
	SQLParserRULE_asc_desc                       = 45
	SQLParserRULE_function_keyword               = 46
	SQLParserRULE_function_name                  = 47
	SQLParserRULE_table_name                     = 48
	SQLParserRULE_table_alias                    = 49
	SQLParserRULE_column_name                    = 50
	SQLParserRULE_column_alias                   = 51
	SQLParserRULE_collation_name                 = 52
	SQLParserRULE_index_name                     = 53
)

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllSql_stmt_list() []ISql_stmt_listContext
	Sql_stmt_list(i int) ISql_stmt_listContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) EOF() antlr.TerminalNode {
	return s.GetToken(SQLParserEOF, 0)
}

func (s *StatementsContext) AllSql_stmt_list() []ISql_stmt_listContext {
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

func (s *StatementsContext) Sql_stmt_list(i int) ISql_stmt_listContext {
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

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251808403619842) != 0) || ((int64((_la-77)) & ^0x3f) == 0 && ((int64(1)<<(_la-77))&1057) != 0) {
		{
			p.SetState(108)
			p.Sql_stmt_list()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
		p.Match(SQLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmt_listContext() *Sql_stmt_listContext {
	var p = new(Sql_stmt_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_sql_stmt_list
	return p
}

func InitEmptySql_stmt_listContext(p *Sql_stmt_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_sql_stmt_list
}

func (*Sql_stmt_listContext) IsSql_stmt_listContext() {}

func NewSql_stmt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmt_listContext {
	var p = new(Sql_stmt_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_sql_stmt_list

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
	return s.GetTokens(SQLParserSCOL)
}

func (s *Sql_stmt_listContext) SCOL(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserSCOL, i)
}

func (s *Sql_stmt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_stmt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_stmt_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitSql_stmt_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Sql_stmt_list() (localctx ISql_stmt_listContext) {
	localctx = NewSql_stmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLParserRULE_sql_stmt_list)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserSCOL {
		{
			p.SetState(116)
			p.Match(SQLParserSCOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)
		p.Sql_stmt()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == SQLParserSCOL {
				{
					p.SetState(123)
					p.Match(SQLParserSCOL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(126)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(128)
				p.Sql_stmt()
			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(134)
				p.Match(SQLParserSCOL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmtContext() *Sql_stmtContext {
	var p = new(Sql_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_sql_stmt
	return p
}

func InitEmptySql_stmtContext(p *Sql_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_sql_stmt
}

func (*Sql_stmtContext) IsSql_stmtContext() {}

func NewSql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmtContext {
	var p = new(Sql_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_sql_stmt

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
	case SQLParserVisitor:
		return t.VisitSql_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Sql_stmt() (localctx ISql_stmtContext) {
	localctx = NewSql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLParserRULE_sql_stmt)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(140)
			p.Delete_stmt()
		}

	case 2:
		{
			p.SetState(141)
			p.Insert_stmt()
		}

	case 3:
		{
			p.SetState(142)
			p.Select_stmt()
		}

	case 4:
		{
			p.SetState(143)
			p.Update_stmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexed_columnContext() *Indexed_columnContext {
	var p = new(Indexed_columnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_indexed_column
	return p
}

func InitEmptyIndexed_columnContext(p *Indexed_columnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_indexed_column
}

func (*Indexed_columnContext) IsIndexed_columnContext() {}

func NewIndexed_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Indexed_columnContext {
	var p = new(Indexed_columnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_indexed_column

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
	case SQLParserVisitor:
		return t.VisitIndexed_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Indexed_column() (localctx IIndexed_columnContext) {
	localctx = NewIndexed_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLParserRULE_indexed_column)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Column_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCte_table_nameContext() *Cte_table_nameContext {
	var p = new(Cte_table_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_cte_table_name
	return p
}

func InitEmptyCte_table_nameContext(p *Cte_table_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_cte_table_name
}

func (*Cte_table_nameContext) IsCte_table_nameContext() {}

func NewCte_table_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cte_table_nameContext {
	var p = new(Cte_table_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_cte_table_name

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
	return s.GetToken(SQLParserOPEN_PAR, 0)
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
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Cte_table_nameContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Cte_table_nameContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Cte_table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cte_table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cte_table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCte_table_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Cte_table_name() (localctx ICte_table_nameContext) {
	localctx = NewCte_table_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLParserRULE_cte_table_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Table_name()
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOPEN_PAR {
		{
			p.SetState(149)
			p.Match(SQLParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Column_name()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(151)
				p.Match(SQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(152)
				p.Column_name()
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(158)
			p.Match(SQLParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommon_table_expressionContext() *Common_table_expressionContext {
	var p = new(Common_table_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_common_table_expression
	return p
}

func InitEmptyCommon_table_expressionContext(p *Common_table_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_common_table_expression
}

func (*Common_table_expressionContext) IsCommon_table_expressionContext() {}

func NewCommon_table_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Common_table_expressionContext {
	var p = new(Common_table_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_common_table_expression

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
	return s.GetToken(SQLParserAS_, 0)
}

func (s *Common_table_expressionContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
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
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Common_table_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Common_table_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Common_table_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCommon_table_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Common_table_expression() (localctx ICommon_table_expressionContext) {
	localctx = NewCommon_table_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLParserRULE_common_table_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Cte_table_name()
	}
	{
		p.SetState(163)
		p.Match(SQLParserAS_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(SQLParserOPEN_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Select_stmt_core()
	}
	{
		p.SetState(166)
		p.Match(SQLParserCLOSE_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommon_table_stmtContext() *Common_table_stmtContext {
	var p = new(Common_table_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_common_table_stmt
	return p
}

func InitEmptyCommon_table_stmtContext(p *Common_table_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_common_table_stmt
}

func (*Common_table_stmtContext) IsCommon_table_stmtContext() {}

func NewCommon_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Common_table_stmtContext {
	var p = new(Common_table_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_common_table_stmt

	return p
}

func (s *Common_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Common_table_stmtContext) WITH_() antlr.TerminalNode {
	return s.GetToken(SQLParserWITH_, 0)
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
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Common_table_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Common_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Common_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Common_table_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCommon_table_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Common_table_stmt() (localctx ICommon_table_stmtContext) {
	localctx = NewCommon_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLParserRULE_common_table_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(SQLParserWITH_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Common_table_expression()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(170)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Common_table_expression()
		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_stmtContext() *Delete_stmtContext {
	var p = new(Delete_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_delete_stmt
	return p
}

func InitEmptyDelete_stmtContext(p *Delete_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_delete_stmt
}

func (*Delete_stmtContext) IsDelete_stmtContext() {}

func NewDelete_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_stmtContext {
	var p = new(Delete_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_delete_stmt

	return p
}

func (s *Delete_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_stmtContext) DELETE_() antlr.TerminalNode {
	return s.GetToken(SQLParserDELETE_, 0)
}

func (s *Delete_stmtContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLParserFROM_, 0)
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
	return s.GetToken(SQLParserWHERE_, 0)
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
	case SQLParserVisitor:
		return t.VisitDelete_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Delete_stmt() (localctx IDelete_stmtContext) {
	localctx = NewDelete_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLParserRULE_delete_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWITH_ {
		{
			p.SetState(177)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(180)
		p.Match(SQLParserDELETE_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(SQLParserFROM_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Qualified_table_name()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE_ {
		{
			p.SetState(183)
			p.Match(SQLParserWHERE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.expr(0)
		}

	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserRETURNING_ {
		{
			p.SetState(187)
			p.Returning_clause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BIND_PARAMETER() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) BIND_PARAMETER() antlr.TerminalNode {
	return s.GetToken(SQLParserBIND_PARAMETER, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(SQLParserBIND_PARAMETER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_name() IFunction_nameContext
	OPEN_PAR() antlr.TerminalNode
	CLOSE_PAR() antlr.TerminalNode
	STAR() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	DISTINCT_() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Function_name() IFunction_nameContext {
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

func (s *Function_callContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
}

func (s *Function_callContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Function_callContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLParserSTAR, 0)
}

func (s *Function_callContext) AllExpr() []IExprContext {
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

func (s *Function_callContext) Expr(i int) IExprContext {
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

func (s *Function_callContext) DISTINCT_() antlr.TerminalNode {
	return s.GetToken(SQLParserDISTINCT_, 0)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLParserRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Function_name()
	}
	{
		p.SetState(193)
		p.Match(SQLParserOPEN_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case SQLParserOPEN_PAR, SQLParserPLUS, SQLParserMINUS, SQLParserCASE_, SQLParserDISTINCT_, SQLParserEXISTS_, SQLParserFALSE_, SQLParserLIKE_, SQLParserNOT_, SQLParserNULL_, SQLParserREPLACE_, SQLParserTRUE_, SQLParserIDENTIFIER, SQLParserNUMERIC_LITERAL, SQLParserBIND_PARAMETER, SQLParserSTRING_LITERAL:
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserDISTINCT_ {
			{
				p.SetState(194)
				p.Match(SQLParserDISTINCT_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(197)
			p.expr(0)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(198)
				p.Match(SQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(199)
				p.expr(0)
			}

			p.SetState(204)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SQLParserSTAR:
		{
			p.SetState(205)
			p.Match(SQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserCLOSE_PAR:

	default:
	}
	{
		p.SetState(208)
		p.Match(SQLParserCLOSE_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumn_refContext is an interface to support dynamic dispatch.
type IColumn_refContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Column_name() IColumn_nameContext
	Table_name() ITable_nameContext
	DOT() antlr.TerminalNode

	// IsColumn_refContext differentiates from other interfaces.
	IsColumn_refContext()
}

type Column_refContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_refContext() *Column_refContext {
	var p = new(Column_refContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_ref
	return p
}

func InitEmptyColumn_refContext(p *Column_refContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_ref
}

func (*Column_refContext) IsColumn_refContext() {}

func NewColumn_refContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_refContext {
	var p = new(Column_refContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_column_ref

	return p
}

func (s *Column_refContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_refContext) Column_name() IColumn_nameContext {
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

func (s *Column_refContext) Table_name() ITable_nameContext {
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

func (s *Column_refContext) DOT() antlr.TerminalNode {
	return s.GetToken(SQLParserDOT, 0)
}

func (s *Column_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_refContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_refContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitColumn_ref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Column_ref() (localctx IColumn_refContext) {
	localctx = NewColumn_refContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SQLParserRULE_column_ref)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(210)
			p.Table_name()
		}
		{
			p.SetState(211)
			p.Match(SQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(215)
		p.Column_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhen_clauseContext is an interface to support dynamic dispatch.
type IWhen_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCondition returns the condition rule contexts.
	GetCondition() IExprContext

	// GetResult returns the result rule contexts.
	GetResult() IExprContext

	// SetCondition sets the condition rule contexts.
	SetCondition(IExprContext)

	// SetResult sets the result rule contexts.
	SetResult(IExprContext)

	// Getter signatures
	WHEN_() antlr.TerminalNode
	THEN_() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsWhen_clauseContext differentiates from other interfaces.
	IsWhen_clauseContext()
}

type When_clauseContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	condition IExprContext
	result    IExprContext
}

func NewEmptyWhen_clauseContext() *When_clauseContext {
	var p = new(When_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_when_clause
	return p
}

func InitEmptyWhen_clauseContext(p *When_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_when_clause
}

func (*When_clauseContext) IsWhen_clauseContext() {}

func NewWhen_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *When_clauseContext {
	var p = new(When_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_when_clause

	return p
}

func (s *When_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *When_clauseContext) GetCondition() IExprContext { return s.condition }

func (s *When_clauseContext) GetResult() IExprContext { return s.result }

func (s *When_clauseContext) SetCondition(v IExprContext) { s.condition = v }

func (s *When_clauseContext) SetResult(v IExprContext) { s.result = v }

func (s *When_clauseContext) WHEN_() antlr.TerminalNode {
	return s.GetToken(SQLParserWHEN_, 0)
}

func (s *When_clauseContext) THEN_() antlr.TerminalNode {
	return s.GetToken(SQLParserTHEN_, 0)
}

func (s *When_clauseContext) AllExpr() []IExprContext {
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

func (s *When_clauseContext) Expr(i int) IExprContext {
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

func (s *When_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *When_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *When_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitWhen_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) When_clause() (localctx IWhen_clauseContext) {
	localctx = NewWhen_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SQLParserRULE_when_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(SQLParserWHEN_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)

		var _x = p.expr(0)

		localctx.(*When_clauseContext).condition = _x
	}
	{
		p.SetState(219)
		p.Match(SQLParserTHEN_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)

		var _x = p.expr(0)

		localctx.(*When_clauseContext).result = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Subquery_exprContext struct {
	ExprContext
}

func NewSubquery_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Subquery_exprContext {
	var p = new(Subquery_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Subquery_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subquery_exprContext) Subquery() ISubqueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubqueryContext)
}

func (s *Subquery_exprContext) EXISTS_() antlr.TerminalNode {
	return s.GetToken(SQLParserEXISTS_, 0)
}

func (s *Subquery_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_, 0)
}

func (s *Subquery_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitSubquery_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Logical_not_exprContext struct {
	ExprContext
}

func NewLogical_not_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Logical_not_exprContext {
	var p = new(Logical_not_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Logical_not_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_not_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_, 0)
}

func (s *Logical_not_exprContext) Expr() IExprContext {
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

func (s *Logical_not_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitLogical_not_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Comparison_exprContext struct {
	ExprContext
	left  IExprContext
	right IExprContext
}

func NewComparison_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Comparison_exprContext {
	var p = new(Comparison_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Comparison_exprContext) GetLeft() IExprContext { return s.left }

func (s *Comparison_exprContext) GetRight() IExprContext { return s.right }

func (s *Comparison_exprContext) SetLeft(v IExprContext) { s.left = v }

func (s *Comparison_exprContext) SetRight(v IExprContext) { s.right = v }

func (s *Comparison_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_exprContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *Comparison_exprContext) AllExpr() []IExprContext {
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

func (s *Comparison_exprContext) Expr(i int) IExprContext {
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

func (s *Comparison_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitComparison_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Like_exprContext struct {
	ExprContext
	elem     IExprContext
	operator antlr.Token
	pattern  IExprContext
	escape   IExprContext
}

func NewLike_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Like_exprContext {
	var p = new(Like_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Like_exprContext) GetOperator() antlr.Token { return s.operator }

func (s *Like_exprContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *Like_exprContext) GetElem() IExprContext { return s.elem }

func (s *Like_exprContext) GetPattern() IExprContext { return s.pattern }

func (s *Like_exprContext) GetEscape() IExprContext { return s.escape }

func (s *Like_exprContext) SetElem(v IExprContext) { s.elem = v }

func (s *Like_exprContext) SetPattern(v IExprContext) { s.pattern = v }

func (s *Like_exprContext) SetEscape(v IExprContext) { s.escape = v }

func (s *Like_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Like_exprContext) AllExpr() []IExprContext {
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

func (s *Like_exprContext) Expr(i int) IExprContext {
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

func (s *Like_exprContext) LIKE_() antlr.TerminalNode {
	return s.GetToken(SQLParserLIKE_, 0)
}

func (s *Like_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_, 0)
}

func (s *Like_exprContext) ESCAPE_() antlr.TerminalNode {
	return s.GetToken(SQLParserESCAPE_, 0)
}

func (s *Like_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitLike_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Null_exprContext struct {
	ExprContext
}

func NewNull_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Null_exprContext {
	var p = new(Null_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Null_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Null_exprContext) Expr() IExprContext {
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

func (s *Null_exprContext) ISNULL_() antlr.TerminalNode {
	return s.GetToken(SQLParserISNULL_, 0)
}

func (s *Null_exprContext) NOTNULL_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOTNULL_, 0)
}

func (s *Null_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitNull_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Column_exprContext struct {
	ExprContext
}

func NewColumn_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Column_exprContext {
	var p = new(Column_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Column_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_exprContext) Column_ref() IColumn_refContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_refContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_refContext)
}

func (s *Column_exprContext) Type_cast() IType_castContext {
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

func (s *Column_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitColumn_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type In_subquery_exprContext struct {
	ExprContext
	elem     IExprContext
	operator antlr.Token
}

func NewIn_subquery_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *In_subquery_exprContext {
	var p = new(In_subquery_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *In_subquery_exprContext) GetOperator() antlr.Token { return s.operator }

func (s *In_subquery_exprContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *In_subquery_exprContext) GetElem() IExprContext { return s.elem }

func (s *In_subquery_exprContext) SetElem(v IExprContext) { s.elem = v }

func (s *In_subquery_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *In_subquery_exprContext) Subquery() ISubqueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubqueryContext)
}

func (s *In_subquery_exprContext) Expr() IExprContext {
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

func (s *In_subquery_exprContext) IN_() antlr.TerminalNode {
	return s.GetToken(SQLParserIN_, 0)
}

func (s *In_subquery_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_, 0)
}

func (s *In_subquery_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitIn_subquery_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Arithmetic_exprContext struct {
	ExprContext
	left     IExprContext
	operator antlr.Token
	right    IExprContext
}

func NewArithmetic_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Arithmetic_exprContext {
	var p = new(Arithmetic_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Arithmetic_exprContext) GetOperator() antlr.Token { return s.operator }

func (s *Arithmetic_exprContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *Arithmetic_exprContext) GetLeft() IExprContext { return s.left }

func (s *Arithmetic_exprContext) GetRight() IExprContext { return s.right }

func (s *Arithmetic_exprContext) SetLeft(v IExprContext) { s.left = v }

func (s *Arithmetic_exprContext) SetRight(v IExprContext) { s.right = v }

func (s *Arithmetic_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_exprContext) AllExpr() []IExprContext {
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

func (s *Arithmetic_exprContext) Expr(i int) IExprContext {
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

func (s *Arithmetic_exprContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLParserSTAR, 0)
}

func (s *Arithmetic_exprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SQLParserDIV, 0)
}

func (s *Arithmetic_exprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SQLParserMOD, 0)
}

func (s *Arithmetic_exprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SQLParserPLUS, 0)
}

func (s *Arithmetic_exprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SQLParserMINUS, 0)
}

func (s *Arithmetic_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitArithmetic_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Logical_binary_exprContext struct {
	ExprContext
	left     IExprContext
	operator antlr.Token
	right    IExprContext
}

func NewLogical_binary_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Logical_binary_exprContext {
	var p = new(Logical_binary_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Logical_binary_exprContext) GetOperator() antlr.Token { return s.operator }

func (s *Logical_binary_exprContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *Logical_binary_exprContext) GetLeft() IExprContext { return s.left }

func (s *Logical_binary_exprContext) GetRight() IExprContext { return s.right }

func (s *Logical_binary_exprContext) SetLeft(v IExprContext) { s.left = v }

func (s *Logical_binary_exprContext) SetRight(v IExprContext) { s.right = v }

func (s *Logical_binary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_binary_exprContext) AllExpr() []IExprContext {
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

func (s *Logical_binary_exprContext) Expr(i int) IExprContext {
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

func (s *Logical_binary_exprContext) AND_() antlr.TerminalNode {
	return s.GetToken(SQLParserAND_, 0)
}

func (s *Logical_binary_exprContext) OR_() antlr.TerminalNode {
	return s.GetToken(SQLParserOR_, 0)
}

func (s *Logical_binary_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitLogical_binary_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Variable_exprContext struct {
	ExprContext
}

func NewVariable_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Variable_exprContext {
	var p = new(Variable_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Variable_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_exprContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Variable_exprContext) Type_cast() IType_castContext {
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

func (s *Variable_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitVariable_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Unary_exprContext struct {
	ExprContext
	operator antlr.Token
}

func NewUnary_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Unary_exprContext {
	var p = new(Unary_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Unary_exprContext) GetOperator() antlr.Token { return s.operator }

func (s *Unary_exprContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *Unary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_exprContext) Expr() IExprContext {
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

func (s *Unary_exprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SQLParserMINUS, 0)
}

func (s *Unary_exprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SQLParserPLUS, 0)
}

func (s *Unary_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitUnary_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Collate_exprContext struct {
	ExprContext
}

func NewCollate_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Collate_exprContext {
	var p = new(Collate_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Collate_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collate_exprContext) Expr() IExprContext {
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

func (s *Collate_exprContext) COLLATE_() antlr.TerminalNode {
	return s.GetToken(SQLParserCOLLATE_, 0)
}

func (s *Collate_exprContext) Collation_name() ICollation_nameContext {
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

func (s *Collate_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCollate_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Parenthesized_exprContext struct {
	ExprContext
}

func NewParenthesized_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Parenthesized_exprContext {
	var p = new(Parenthesized_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Parenthesized_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parenthesized_exprContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
}

func (s *Parenthesized_exprContext) Expr() IExprContext {
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

func (s *Parenthesized_exprContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Parenthesized_exprContext) Type_cast() IType_castContext {
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

func (s *Parenthesized_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitParenthesized_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Between_exprContext struct {
	ExprContext
	elem     IExprContext
	operator antlr.Token
	low      IExprContext
	high     IExprContext
}

func NewBetween_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Between_exprContext {
	var p = new(Between_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Between_exprContext) GetOperator() antlr.Token { return s.operator }

func (s *Between_exprContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *Between_exprContext) GetElem() IExprContext { return s.elem }

func (s *Between_exprContext) GetLow() IExprContext { return s.low }

func (s *Between_exprContext) GetHigh() IExprContext { return s.high }

func (s *Between_exprContext) SetElem(v IExprContext) { s.elem = v }

func (s *Between_exprContext) SetLow(v IExprContext) { s.low = v }

func (s *Between_exprContext) SetHigh(v IExprContext) { s.high = v }

func (s *Between_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Between_exprContext) AND_() antlr.TerminalNode {
	return s.GetToken(SQLParserAND_, 0)
}

func (s *Between_exprContext) AllExpr() []IExprContext {
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

func (s *Between_exprContext) Expr(i int) IExprContext {
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

func (s *Between_exprContext) BETWEEN_() antlr.TerminalNode {
	return s.GetToken(SQLParserBETWEEN_, 0)
}

func (s *Between_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_, 0)
}

func (s *Between_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitBetween_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_list_exprContext struct {
	ExprContext
}

func NewExpr_list_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_list_exprContext {
	var p = new(Expr_list_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Expr_list_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_list_exprContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
}

func (s *Expr_list_exprContext) Expr_list() IExpr_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_listContext)
}

func (s *Expr_list_exprContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Expr_list_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitExpr_list_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type In_list_exprContext struct {
	ExprContext
	elem     IExprContext
	operator antlr.Token
}

func NewIn_list_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *In_list_exprContext {
	var p = new(In_list_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *In_list_exprContext) GetOperator() antlr.Token { return s.operator }

func (s *In_list_exprContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *In_list_exprContext) GetElem() IExprContext { return s.elem }

func (s *In_list_exprContext) SetElem(v IExprContext) { s.elem = v }

func (s *In_list_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *In_list_exprContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
}

func (s *In_list_exprContext) Expr_list() IExpr_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_listContext)
}

func (s *In_list_exprContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *In_list_exprContext) Expr() IExprContext {
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

func (s *In_list_exprContext) IN_() antlr.TerminalNode {
	return s.GetToken(SQLParserIN_, 0)
}

func (s *In_list_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_, 0)
}

func (s *In_list_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitIn_list_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Literal_exprContext struct {
	ExprContext
}

func NewLiteral_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Literal_exprContext {
	var p = new(Literal_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Literal_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_exprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Literal_exprContext) Type_cast() IType_castContext {
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

func (s *Literal_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitLiteral_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Is_exprContext struct {
	ExprContext
}

func NewIs_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Is_exprContext {
	var p = new(Is_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Is_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Is_exprContext) AllExpr() []IExprContext {
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

func (s *Is_exprContext) Expr(i int) IExprContext {
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

func (s *Is_exprContext) IS_() antlr.TerminalNode {
	return s.GetToken(SQLParserIS_, 0)
}

func (s *Is_exprContext) DISTINCT_() antlr.TerminalNode {
	return s.GetToken(SQLParserDISTINCT_, 0)
}

func (s *Is_exprContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLParserFROM_, 0)
}

func (s *Is_exprContext) Boolean_value() IBoolean_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_valueContext)
}

func (s *Is_exprContext) NULL_() antlr.TerminalNode {
	return s.GetToken(SQLParserNULL_, 0)
}

func (s *Is_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_, 0)
}

func (s *Is_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitIs_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Case_exprContext struct {
	ExprContext
	case_clause IExprContext
	else_clause IExprContext
}

func NewCase_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Case_exprContext {
	var p = new(Case_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Case_exprContext) GetCase_clause() IExprContext { return s.case_clause }

func (s *Case_exprContext) GetElse_clause() IExprContext { return s.else_clause }

func (s *Case_exprContext) SetCase_clause(v IExprContext) { s.case_clause = v }

func (s *Case_exprContext) SetElse_clause(v IExprContext) { s.else_clause = v }

func (s *Case_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_exprContext) CASE_() antlr.TerminalNode {
	return s.GetToken(SQLParserCASE_, 0)
}

func (s *Case_exprContext) END_() antlr.TerminalNode {
	return s.GetToken(SQLParserEND_, 0)
}

func (s *Case_exprContext) AllWhen_clause() []IWhen_clauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhen_clauseContext); ok {
			len++
		}
	}

	tst := make([]IWhen_clauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhen_clauseContext); ok {
			tst[i] = t.(IWhen_clauseContext)
			i++
		}
	}

	return tst
}

func (s *Case_exprContext) When_clause(i int) IWhen_clauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhen_clauseContext); ok {
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

	return t.(IWhen_clauseContext)
}

func (s *Case_exprContext) ELSE_() antlr.TerminalNode {
	return s.GetToken(SQLParserELSE_, 0)
}

func (s *Case_exprContext) AllExpr() []IExprContext {
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

func (s *Case_exprContext) Expr(i int) IExprContext {
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

func (s *Case_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCase_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

type Function_exprContext struct {
	ExprContext
}

func NewFunction_exprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Function_exprContext {
	var p = new(Function_exprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *Function_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_exprContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Function_exprContext) Type_cast() IType_castContext {
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

func (s *Function_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitFunction_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SQLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, SQLParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteral_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(223)
			p.Literal()
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(224)
				p.Type_cast()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		localctx = NewVariable_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(227)
			p.Variable()
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(228)
				p.Type_cast()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 3:
		localctx = NewColumn_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(231)
			p.Column_ref()
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(232)
				p.Type_cast()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 4:
		localctx = NewUnary_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Unary_exprContext).operator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLParserPLUS || _la == SQLParserMINUS) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Unary_exprContext).operator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(236)
			p.expr(19)
		}

	case 5:
		localctx = NewParenthesized_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(237)
			p.Match(SQLParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.expr(0)
		}
		{
			p.SetState(239)
			p.Match(SQLParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(240)
				p.Type_cast()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 6:
		localctx = NewSubquery_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserEXISTS_ || _la == SQLParserNOT_ {
			p.SetState(244)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SQLParserNOT_ {
				{
					p.SetState(243)
					p.Match(SQLParserNOT_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(246)
				p.Match(SQLParserEXISTS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(249)
			p.Subquery()
		}

	case 7:
		localctx = NewCase_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(250)
			p.Match(SQLParserCASE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1152928101743723272) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&251724805) != 0) {
			{
				p.SetState(251)

				var _x = p.expr(0)

				localctx.(*Case_exprContext).case_clause = _x
			}

		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SQLParserWHEN_ {
			{
				p.SetState(254)
				p.When_clause()
			}

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserELSE_ {
			{
				p.SetState(259)
				p.Match(SQLParserELSE_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(260)

				var _x = p.expr(0)

				localctx.(*Case_exprContext).else_clause = _x
			}

		}
		{
			p.SetState(263)
			p.Match(SQLParserEND_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExpr_list_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(265)
			p.Match(SQLParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Expr_list()
		}
		{
			p.SetState(267)
			p.Match(SQLParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewFunction_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(269)
			p.Function_call()
		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(270)
				p.Type_cast()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 10:
		localctx = NewLogical_not_exprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(273)
			p.Match(SQLParserNOT_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.expr(3)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(344)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmetic_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*Arithmetic_exprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(278)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Arithmetic_exprContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3200) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Arithmetic_exprContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(279)

					var _x = p.expr(13)

					localctx.(*Arithmetic_exprContext).right = _x
				}

			case 2:
				localctx = NewArithmetic_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*Arithmetic_exprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(281)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Arithmetic_exprContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SQLParserPLUS || _la == SQLParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Arithmetic_exprContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(282)

					var _x = p.expr(12)

					localctx.(*Arithmetic_exprContext).right = _x
				}

			case 3:
				localctx = NewBetween_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*Between_exprContext).elem = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				p.SetState(285)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SQLParserNOT_ {
					{
						p.SetState(284)
						p.Match(SQLParserNOT_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(287)

					var _m = p.Match(SQLParserBETWEEN_)

					localctx.(*Between_exprContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(288)

					var _x = p.expr(0)

					localctx.(*Between_exprContext).low = _x
				}
				{
					p.SetState(289)
					p.Match(SQLParserAND_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(290)

					var _x = p.expr(9)

					localctx.(*Between_exprContext).high = _x
				}

			case 4:
				localctx = NewComparison_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*Comparison_exprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(293)
					p.ComparisonOperator()
				}
				{
					p.SetState(294)

					var _x = p.expr(7)

					localctx.(*Comparison_exprContext).right = _x
				}

			case 5:
				localctx = NewLogical_binary_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*Logical_binary_exprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(297)

					var _m = p.Match(SQLParserAND_)

					localctx.(*Logical_binary_exprContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(298)

					var _x = p.expr(3)

					localctx.(*Logical_binary_exprContext).right = _x
				}

			case 6:
				localctx = NewLogical_binary_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*Logical_binary_exprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(300)

					var _m = p.Match(SQLParserOR_)

					localctx.(*Logical_binary_exprContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(301)

					var _x = p.expr(2)

					localctx.(*Logical_binary_exprContext).right = _x
				}

			case 7:
				localctx = NewCollate_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(303)
					p.Match(SQLParserCOLLATE_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(304)
					p.Collation_name()
				}

			case 8:
				localctx = NewIn_subquery_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*In_subquery_exprContext).elem = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				p.SetState(307)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SQLParserNOT_ {
					{
						p.SetState(306)
						p.Match(SQLParserNOT_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(309)

					var _m = p.Match(SQLParserIN_)

					localctx.(*In_subquery_exprContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(310)
					p.Subquery()
				}

			case 9:
				localctx = NewIn_list_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*In_list_exprContext).elem = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				p.SetState(313)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SQLParserNOT_ {
					{
						p.SetState(312)
						p.Match(SQLParserNOT_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(315)

					var _m = p.Match(SQLParserIN_)

					localctx.(*In_list_exprContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(316)
					p.Match(SQLParserOPEN_PAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(317)
					p.Expr_list()
				}
				{
					p.SetState(318)
					p.Match(SQLParserCLOSE_PAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewLike_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*Like_exprContext).elem = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(320)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				p.SetState(322)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SQLParserNOT_ {
					{
						p.SetState(321)
						p.Match(SQLParserNOT_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(324)

					var _m = p.Match(SQLParserLIKE_)

					localctx.(*Like_exprContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(325)

					var _x = p.expr(0)

					localctx.(*Like_exprContext).pattern = _x
				}
				p.SetState(328)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(326)
						p.Match(SQLParserESCAPE_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(327)

						var _x = p.expr(0)

						localctx.(*Like_exprContext).escape = _x
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case 11:
				localctx = NewIs_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(330)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(331)
					p.Match(SQLParserIS_)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(333)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SQLParserNOT_ {
					{
						p.SetState(332)
						p.Match(SQLParserNOT_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				p.SetState(340)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case SQLParserDISTINCT_:
					{
						p.SetState(335)
						p.Match(SQLParserDISTINCT_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(336)
						p.Match(SQLParserFROM_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(337)
						p.expr(0)
					}

				case SQLParserFALSE_, SQLParserTRUE_:
					{
						p.SetState(338)
						p.Boolean_value()
					}

				case SQLParserNULL_:
					{
						p.SetState(339)
						p.Match(SQLParserNULL_)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

			case 12:
				localctx = NewNull_exprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SQLParserRULE_expr)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(343)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SQLParserISNULL_ || _la == SQLParserNOTNULL_) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubqueryContext is an interface to support dynamic dispatch.
type ISubqueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAR() antlr.TerminalNode
	Select_stmt_core() ISelect_stmt_coreContext
	CLOSE_PAR() antlr.TerminalNode

	// IsSubqueryContext differentiates from other interfaces.
	IsSubqueryContext()
}

type SubqueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubqueryContext() *SubqueryContext {
	var p = new(SubqueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_subquery
	return p
}

func InitEmptySubqueryContext(p *SubqueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_subquery
}

func (*SubqueryContext) IsSubqueryContext() {}

func NewSubqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubqueryContext {
	var p = new(SubqueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_subquery

	return p
}

func (s *SubqueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SubqueryContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
}

func (s *SubqueryContext) Select_stmt_core() ISelect_stmt_coreContext {
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

func (s *SubqueryContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *SubqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubqueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitSubquery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Subquery() (localctx ISubqueryContext) {
	localctx = NewSubqueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SQLParserRULE_subquery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(SQLParserOPEN_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
		p.Select_stmt_core()
	}
	{
		p.SetState(351)
		p.Match(SQLParserCLOSE_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpr_listContext is an interface to support dynamic dispatch.
type IExpr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpr_listContext differentiates from other interfaces.
	IsExpr_listContext()
}

type Expr_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_listContext() *Expr_listContext {
	var p = new(Expr_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_expr_list
	return p
}

func InitEmptyExpr_listContext(p *Expr_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_expr_list
}

func (*Expr_listContext) IsExpr_listContext() {}

func NewExpr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_listContext {
	var p = new(Expr_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_expr_list

	return p
}

func (s *Expr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_listContext) AllExpr() []IExprContext {
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

func (s *Expr_listContext) Expr(i int) IExprContext {
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

func (s *Expr_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Expr_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Expr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitExpr_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Expr_list() (localctx IExpr_listContext) {
	localctx = NewExpr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SQLParserRULE_expr_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.expr(0)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(354)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.expr(0)
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	LT_EQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	GT_EQ() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	NOT_EQ1() antlr.TerminalNode
	NOT_EQ2() antlr.TerminalNode

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_comparisonOperator
	return p
}

func InitEmptyComparisonOperatorContext(p *ComparisonOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_comparisonOperator
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(SQLParserLT, 0)
}

func (s *ComparisonOperatorContext) LT_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserLT_EQ, 0)
}

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SQLParserGT, 0)
}

func (s *ComparisonOperatorContext) GT_EQ() antlr.TerminalNode {
	return s.GetToken(SQLParserGT_EQ, 0)
}

func (s *ComparisonOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SQLParserASSIGN, 0)
}

func (s *ComparisonOperatorContext) NOT_EQ1() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_EQ1, 0)
}

func (s *ComparisonOperatorContext) NOT_EQ2() antlr.TerminalNode {
	return s.GetToken(SQLParserNOT_EQ2, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SQLParserRULE_comparisonOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&258112) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCast_typeContext() *Cast_typeContext {
	var p = new(Cast_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_cast_type
	return p
}

func InitEmptyCast_typeContext(p *Cast_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_cast_type
}

func (*Cast_typeContext) IsCast_typeContext() {}

func NewCast_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cast_typeContext {
	var p = new(Cast_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_cast_type

	return p
}

func (s *Cast_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Cast_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Cast_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cast_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cast_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCast_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Cast_type() (localctx ICast_typeContext) {
	localctx = NewCast_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SQLParserRULE_cast_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(SQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_castContext() *Type_castContext {
	var p = new(Type_castContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_type_cast
	return p
}

func InitEmptyType_castContext(p *Type_castContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_type_cast
}

func (*Type_castContext) IsType_castContext() {}

func NewType_castContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_castContext {
	var p = new(Type_castContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_type_cast

	return p
}

func (s *Type_castContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_castContext) TYPE_CAST() antlr.TerminalNode {
	return s.GetToken(SQLParserTYPE_CAST, 0)
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
	case SQLParserVisitor:
		return t.VisitType_cast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Type_cast() (localctx IType_castContext) {
	localctx = NewType_castContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SQLParserRULE_type_cast)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(SQLParserTYPE_CAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Cast_type()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_valueContext is an interface to support dynamic dispatch.
type IBoolean_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE_() antlr.TerminalNode
	FALSE_() antlr.TerminalNode

	// IsBoolean_valueContext differentiates from other interfaces.
	IsBoolean_valueContext()
}

type Boolean_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_valueContext() *Boolean_valueContext {
	var p = new(Boolean_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_boolean_value
	return p
}

func InitEmptyBoolean_valueContext(p *Boolean_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_boolean_value
}

func (*Boolean_valueContext) IsBoolean_valueContext() {}

func NewBoolean_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_valueContext {
	var p = new(Boolean_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_boolean_value

	return p
}

func (s *Boolean_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_valueContext) TRUE_() antlr.TerminalNode {
	return s.GetToken(SQLParserTRUE_, 0)
}

func (s *Boolean_valueContext) FALSE_() antlr.TerminalNode {
	return s.GetToken(SQLParserFALSE_, 0)
}

func (s *Boolean_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitBoolean_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Boolean_value() (localctx IBoolean_valueContext) {
	localctx = NewBoolean_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SQLParserRULE_boolean_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLParserFALSE_ || _la == SQLParserTRUE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_valueContext is an interface to support dynamic dispatch.
type IString_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsString_valueContext differentiates from other interfaces.
	IsString_valueContext()
}

type String_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_valueContext() *String_valueContext {
	var p = new(String_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_string_value
	return p
}

func InitEmptyString_valueContext(p *String_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_string_value
}

func (*String_valueContext) IsString_valueContext() {}

func NewString_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_valueContext {
	var p = new(String_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_string_value

	return p
}

func (s *String_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *String_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLParserSTRING_LITERAL, 0)
}

func (s *String_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitString_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) String_value() (localctx IString_valueContext) {
	localctx = NewString_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SQLParserRULE_string_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(SQLParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumeric_valueContext is an interface to support dynamic dispatch.
type INumeric_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode

	// IsNumeric_valueContext differentiates from other interfaces.
	IsNumeric_valueContext()
}

type Numeric_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_valueContext() *Numeric_valueContext {
	var p = new(Numeric_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_numeric_value
	return p
}

func InitEmptyNumeric_valueContext(p *Numeric_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_numeric_value
}

func (*Numeric_valueContext) IsNumeric_valueContext() {}

func NewNumeric_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_valueContext {
	var p = new(Numeric_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_numeric_value

	return p
}

func (s *Numeric_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_valueContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLParserNUMERIC_LITERAL, 0)
}

func (s *Numeric_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitNumeric_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Numeric_value() (localctx INumeric_valueContext) {
	localctx = NewNumeric_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SQLParserRULE_numeric_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(SQLParserNUMERIC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NULL_() antlr.TerminalNode
	Boolean_value() IBoolean_valueContext
	String_value() IString_valueContext
	Numeric_value() INumeric_valueContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NULL_() antlr.TerminalNode {
	return s.GetToken(SQLParserNULL_, 0)
}

func (s *LiteralContext) Boolean_value() IBoolean_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_valueContext)
}

func (s *LiteralContext) String_value() IString_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_valueContext)
}

func (s *LiteralContext) Numeric_value() INumeric_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumeric_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumeric_valueContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SQLParserRULE_literal)
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserNULL_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(SQLParserNULL_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserFALSE_, SQLParserTRUE_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Boolean_value()
		}

	case SQLParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(376)
			p.String_value()
		}

	case SQLParserNUMERIC_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(377)
			p.Numeric_value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_rowContext() *Value_rowContext {
	var p = new(Value_rowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_value_row
	return p
}

func InitEmptyValue_rowContext(p *Value_rowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_value_row
}

func (*Value_rowContext) IsValue_rowContext() {}

func NewValue_rowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_rowContext {
	var p = new(Value_rowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_value_row

	return p
}

func (s *Value_rowContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_rowContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
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
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Value_rowContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Value_rowContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Value_rowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_rowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_rowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitValue_row(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Value_row() (localctx IValue_rowContext) {
	localctx = NewValue_rowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SQLParserRULE_value_row)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(SQLParserOPEN_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.expr(0)
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(382)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.expr(0)
		}

		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(389)
		p.Match(SQLParserCLOSE_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValues_clauseContext() *Values_clauseContext {
	var p = new(Values_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_values_clause
	return p
}

func InitEmptyValues_clauseContext(p *Values_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_values_clause
}

func (*Values_clauseContext) IsValues_clauseContext() {}

func NewValues_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Values_clauseContext {
	var p = new(Values_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_values_clause

	return p
}

func (s *Values_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Values_clauseContext) VALUES_() antlr.TerminalNode {
	return s.GetToken(SQLParserVALUES_, 0)
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
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Values_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Values_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Values_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Values_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitValues_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Values_clause() (localctx IValues_clauseContext) {
	localctx = NewValues_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SQLParserRULE_values_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(SQLParserVALUES_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Value_row()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(393)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.Value_row()
		}

		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInsert_stmtContext is an interface to support dynamic dispatch.
type IInsert_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSERT_() antlr.TerminalNode
	INTO_() antlr.TerminalNode
	Table_name() ITable_nameContext
	Values_clause() IValues_clauseContext
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_stmtContext() *Insert_stmtContext {
	var p = new(Insert_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_insert_stmt
	return p
}

func InitEmptyInsert_stmtContext(p *Insert_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_insert_stmt
}

func (*Insert_stmtContext) IsInsert_stmtContext() {}

func NewInsert_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_stmtContext {
	var p = new(Insert_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_insert_stmt

	return p
}

func (s *Insert_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_stmtContext) INSERT_() antlr.TerminalNode {
	return s.GetToken(SQLParserINSERT_, 0)
}

func (s *Insert_stmtContext) INTO_() antlr.TerminalNode {
	return s.GetToken(SQLParserINTO_, 0)
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
	return s.GetToken(SQLParserAS_, 0)
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
	return s.GetToken(SQLParserOPEN_PAR, 0)
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
	return s.GetToken(SQLParserCLOSE_PAR, 0)
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
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Insert_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Insert_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitInsert_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Insert_stmt() (localctx IInsert_stmtContext) {
	localctx = NewInsert_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SQLParserRULE_insert_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWITH_ {
		{
			p.SetState(400)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(403)
		p.Match(SQLParserINSERT_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.Match(SQLParserINTO_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.Table_name()
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserAS_ {
		{
			p.SetState(406)
			p.Match(SQLParserAS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Table_alias()
		}

	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOPEN_PAR {
		{
			p.SetState(410)
			p.Match(SQLParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Column_name()
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(412)
				p.Match(SQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(413)
				p.Column_name()
			}

			p.SetState(418)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(419)
			p.Match(SQLParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(423)
		p.Values_clause()
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserON_ {
		{
			p.SetState(424)
			p.Upsert_clause()
		}

	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserRETURNING_ {
		{
			p.SetState(427)
			p.Returning_clause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturning_clauseContext() *Returning_clauseContext {
	var p = new(Returning_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_returning_clause
	return p
}

func InitEmptyReturning_clauseContext(p *Returning_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_returning_clause
}

func (*Returning_clauseContext) IsReturning_clauseContext() {}

func NewReturning_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Returning_clauseContext {
	var p = new(Returning_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_returning_clause

	return p
}

func (s *Returning_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Returning_clauseContext) RETURNING_() antlr.TerminalNode {
	return s.GetToken(SQLParserRETURNING_, 0)
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
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Returning_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Returning_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Returning_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Returning_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitReturning_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Returning_clause() (localctx IReturning_clauseContext) {
	localctx = NewReturning_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SQLParserRULE_returning_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(SQLParserRETURNING_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Returning_clause_result_column()
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(432)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.Returning_clause_result_column()
		}

		p.SetState(438)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpsert_updateContext() *Upsert_updateContext {
	var p = new(Upsert_updateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_upsert_update
	return p
}

func InitEmptyUpsert_updateContext(p *Upsert_updateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_upsert_update
}

func (*Upsert_updateContext) IsUpsert_updateContext() {}

func NewUpsert_updateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upsert_updateContext {
	var p = new(Upsert_updateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_upsert_update

	return p
}

func (s *Upsert_updateContext) GetParser() antlr.Parser { return s.parser }

func (s *Upsert_updateContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SQLParserASSIGN, 0)
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
	case SQLParserVisitor:
		return t.VisitUpsert_update(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Upsert_update() (localctx IUpsert_updateContext) {
	localctx = NewUpsert_updateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SQLParserRULE_upsert_update)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		{
			p.SetState(439)
			p.Column_name()
		}

	case SQLParserOPEN_PAR:
		{
			p.SetState(440)
			p.Column_name_list()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(443)
		p.Match(SQLParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	target_expr IExprContext
	update_expr IExprContext
}

func NewEmptyUpsert_clauseContext() *Upsert_clauseContext {
	var p = new(Upsert_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_upsert_clause
	return p
}

func InitEmptyUpsert_clauseContext(p *Upsert_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_upsert_clause
}

func (*Upsert_clauseContext) IsUpsert_clauseContext() {}

func NewUpsert_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upsert_clauseContext {
	var p = new(Upsert_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_upsert_clause

	return p
}

func (s *Upsert_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Upsert_clauseContext) GetTarget_expr() IExprContext { return s.target_expr }

func (s *Upsert_clauseContext) GetUpdate_expr() IExprContext { return s.update_expr }

func (s *Upsert_clauseContext) SetTarget_expr(v IExprContext) { s.target_expr = v }

func (s *Upsert_clauseContext) SetUpdate_expr(v IExprContext) { s.update_expr = v }

func (s *Upsert_clauseContext) ON_() antlr.TerminalNode {
	return s.GetToken(SQLParserON_, 0)
}

func (s *Upsert_clauseContext) CONFLICT_() antlr.TerminalNode {
	return s.GetToken(SQLParserCONFLICT_, 0)
}

func (s *Upsert_clauseContext) DO_() antlr.TerminalNode {
	return s.GetToken(SQLParserDO_, 0)
}

func (s *Upsert_clauseContext) NOTHING_() antlr.TerminalNode {
	return s.GetToken(SQLParserNOTHING_, 0)
}

func (s *Upsert_clauseContext) UPDATE_() antlr.TerminalNode {
	return s.GetToken(SQLParserUPDATE_, 0)
}

func (s *Upsert_clauseContext) SET_() antlr.TerminalNode {
	return s.GetToken(SQLParserSET_, 0)
}

func (s *Upsert_clauseContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
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
	return s.GetToken(SQLParserCLOSE_PAR, 0)
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
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Upsert_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Upsert_clauseContext) AllWHERE_() []antlr.TerminalNode {
	return s.GetTokens(SQLParserWHERE_)
}

func (s *Upsert_clauseContext) WHERE_(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserWHERE_, i)
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
	case SQLParserVisitor:
		return t.VisitUpsert_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Upsert_clause() (localctx IUpsert_clauseContext) {
	localctx = NewUpsert_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SQLParserRULE_upsert_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(SQLParserON_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(447)
		p.Match(SQLParserCONFLICT_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOPEN_PAR {
		{
			p.SetState(448)
			p.Match(SQLParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.Indexed_column()
		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(450)
				p.Match(SQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(451)
				p.Indexed_column()
			}

			p.SetState(456)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(457)
			p.Match(SQLParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserWHERE_ {
			{
				p.SetState(458)
				p.Match(SQLParserWHERE_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(459)

				var _x = p.expr(0)

				localctx.(*Upsert_clauseContext).target_expr = _x
			}

		}

	}
	{
		p.SetState(464)
		p.Match(SQLParserDO_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserNOTHING_:
		{
			p.SetState(465)
			p.Match(SQLParserNOTHING_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserUPDATE_:
		{
			p.SetState(466)
			p.Match(SQLParserUPDATE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.Match(SQLParserSET_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(468)
			p.Upsert_update()
		}
		p.SetState(473)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(469)
				p.Match(SQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(470)
				p.Upsert_update()
			}

			p.SetState(475)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserWHERE_ {
			{
				p.SetState(476)
				p.Match(SQLParserWHERE_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(477)

				var _x = p.expr(0)

				localctx.(*Upsert_clauseContext).update_expr = _x
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_stmt_coreContext() *Select_stmt_coreContext {
	var p = new(Select_stmt_coreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_stmt_core
	return p
}

func InitEmptySelect_stmt_coreContext(p *Select_stmt_coreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_stmt_core
}

func (*Select_stmt_coreContext) IsSelect_stmt_coreContext() {}

func NewSelect_stmt_coreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_stmt_coreContext {
	var p = new(Select_stmt_coreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_select_stmt_core

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
	case SQLParserVisitor:
		return t.VisitSelect_stmt_core(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Select_stmt_core() (localctx ISelect_stmt_coreContext) {
	localctx = NewSelect_stmt_coreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SQLParserRULE_select_stmt_core)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(482)
		p.Select_core()
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-40)) & ^0x3f) == 0 && ((int64(1)<<(_la-40))&2199023259649) != 0 {
		{
			p.SetState(483)
			p.Compound_operator()
		}
		{
			p.SetState(484)
			p.Select_core()
		}

		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserORDER_ {
		{
			p.SetState(491)
			p.Order_by_stmt()
		}

	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserLIMIT_ {
		{
			p.SetState(494)
			p.Limit_stmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_stmtContext() *Select_stmtContext {
	var p = new(Select_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_stmt
	return p
}

func InitEmptySelect_stmtContext(p *Select_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_stmt
}

func (*Select_stmtContext) IsSelect_stmtContext() {}

func NewSelect_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_stmtContext {
	var p = new(Select_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_select_stmt

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
	case SQLParserVisitor:
		return t.VisitSelect_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Select_stmt() (localctx ISelect_stmtContext) {
	localctx = NewSelect_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SQLParserRULE_select_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWITH_ {
		{
			p.SetState(497)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(500)
		p.Select_stmt_core()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_clauseContext() *Join_clauseContext {
	var p = new(Join_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_join_clause
	return p
}

func InitEmptyJoin_clauseContext(p *Join_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_join_clause
}

func (*Join_clauseContext) IsJoin_clauseContext() {}

func NewJoin_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_clauseContext {
	var p = new(Join_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_join_clause

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
	case SQLParserVisitor:
		return t.VisitJoin_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Join_clause() (localctx IJoin_clauseContext) {
	localctx = NewJoin_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SQLParserRULE_join_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		p.Table_or_subquery()
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-46)) & ^0x3f) == 0 && ((int64(1)<<(_la-46))&1073752081) != 0 {
		{
			p.SetState(503)
			p.Join_operator()
		}
		{
			p.SetState(504)
			p.Table_or_subquery()
		}
		{
			p.SetState(505)
			p.Join_constraint()
		}

		p.SetState(511)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	whereExpr   IExprContext
	_expr       IExprContext
	groupByExpr []IExprContext
	havingExpr  IExprContext
}

func NewEmptySelect_coreContext() *Select_coreContext {
	var p = new(Select_coreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_core
	return p
}

func InitEmptySelect_coreContext(p *Select_coreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_select_core
}

func (*Select_coreContext) IsSelect_coreContext() {}

func NewSelect_coreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_coreContext {
	var p = new(Select_coreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_select_core

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
	return s.GetToken(SQLParserSELECT_, 0)
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
	return s.GetToken(SQLParserDISTINCT_, 0)
}

func (s *Select_coreContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Select_coreContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Select_coreContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLParserFROM_, 0)
}

func (s *Select_coreContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SQLParserWHERE_, 0)
}

func (s *Select_coreContext) GROUP_() antlr.TerminalNode {
	return s.GetToken(SQLParserGROUP_, 0)
}

func (s *Select_coreContext) BY_() antlr.TerminalNode {
	return s.GetToken(SQLParserBY_, 0)
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
	return s.GetToken(SQLParserHAVING_, 0)
}

func (s *Select_coreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_coreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_coreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitSelect_core(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Select_core() (localctx ISelect_coreContext) {
	localctx = NewSelect_coreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SQLParserRULE_select_core)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Match(SQLParserSELECT_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserDISTINCT_ {
		{
			p.SetState(513)
			p.Match(SQLParserDISTINCT_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(516)
		p.Result_column()
	}
	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(517)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)
			p.Result_column()
		}

		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserFROM_ {
		{
			p.SetState(524)
			p.Match(SQLParserFROM_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(525)
				p.Table_or_subquery()
			}

		case 2:
			{
				p.SetState(526)
				p.Join_clause()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	}
	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE_ {
		{
			p.SetState(531)
			p.Match(SQLParserWHERE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)

			var _x = p.expr(0)

			localctx.(*Select_coreContext).whereExpr = _x
		}

	}
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserGROUP_ {
		{
			p.SetState(535)
			p.Match(SQLParserGROUP_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(536)
			p.Match(SQLParserBY_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(537)

			var _x = p.expr(0)

			localctx.(*Select_coreContext)._expr = _x
		}
		localctx.(*Select_coreContext).groupByExpr = append(localctx.(*Select_coreContext).groupByExpr, localctx.(*Select_coreContext)._expr)
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SQLParserCOMMA {
			{
				p.SetState(538)
				p.Match(SQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(539)

				var _x = p.expr(0)

				localctx.(*Select_coreContext)._expr = _x
			}
			localctx.(*Select_coreContext).groupByExpr = append(localctx.(*Select_coreContext).groupByExpr, localctx.(*Select_coreContext)._expr)

			p.SetState(544)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserHAVING_ {
			{
				p.SetState(545)
				p.Match(SQLParserHAVING_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(546)

				var _x = p.expr(0)

				localctx.(*Select_coreContext).havingExpr = _x
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_or_subqueryContext() *Table_or_subqueryContext {
	var p = new(Table_or_subqueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_or_subquery
	return p
}

func InitEmptyTable_or_subqueryContext(p *Table_or_subqueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_or_subquery
}

func (*Table_or_subqueryContext) IsTable_or_subqueryContext() {}

func NewTable_or_subqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_or_subqueryContext {
	var p = new(Table_or_subqueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_table_or_subquery

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
	return s.GetToken(SQLParserAS_, 0)
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
	return s.GetToken(SQLParserOPEN_PAR, 0)
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
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Table_or_subqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_or_subqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_or_subqueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitTable_or_subquery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Table_or_subquery() (localctx ITable_or_subqueryContext) {
	localctx = NewTable_or_subqueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SQLParserRULE_table_or_subquery)
	var _la int

	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Table_name()
		}
		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserAS_ {
			{
				p.SetState(552)
				p.Match(SQLParserAS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(553)
				p.Table_alias()
			}

		}

	case SQLParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(556)
			p.Match(SQLParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.Select_stmt_core()
		}
		{
			p.SetState(558)
			p.Match(SQLParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(561)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserAS_ {
			{
				p.SetState(559)
				p.Match(SQLParserAS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(560)
				p.Table_alias()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResult_columnContext() *Result_columnContext {
	var p = new(Result_columnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_result_column
	return p
}

func InitEmptyResult_columnContext(p *Result_columnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_result_column
}

func (*Result_columnContext) IsResult_columnContext() {}

func NewResult_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Result_columnContext {
	var p = new(Result_columnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_result_column

	return p
}

func (s *Result_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Result_columnContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLParserSTAR, 0)
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
	return s.GetToken(SQLParserDOT, 0)
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
	return s.GetToken(SQLParserAS_, 0)
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
	case SQLParserVisitor:
		return t.VisitResult_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Result_column() (localctx IResult_columnContext) {
	localctx = NewResult_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SQLParserRULE_result_column)
	var _la int

	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 71, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(565)
			p.Match(SQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(566)
			p.Table_name()
		}
		{
			p.SetState(567)
			p.Match(SQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)
			p.Match(SQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(570)
			p.expr(0)
		}
		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserAS_ {
			{
				p.SetState(571)
				p.Match(SQLParserAS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(572)
				p.Column_alias()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturning_clause_result_columnContext() *Returning_clause_result_columnContext {
	var p = new(Returning_clause_result_columnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_returning_clause_result_column
	return p
}

func InitEmptyReturning_clause_result_columnContext(p *Returning_clause_result_columnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_returning_clause_result_column
}

func (*Returning_clause_result_columnContext) IsReturning_clause_result_columnContext() {}

func NewReturning_clause_result_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Returning_clause_result_columnContext {
	var p = new(Returning_clause_result_columnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_returning_clause_result_column

	return p
}

func (s *Returning_clause_result_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Returning_clause_result_columnContext) STAR() antlr.TerminalNode {
	return s.GetToken(SQLParserSTAR, 0)
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
	return s.GetToken(SQLParserAS_, 0)
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
	case SQLParserVisitor:
		return t.VisitReturning_clause_result_column(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Returning_clause_result_column() (localctx IReturning_clause_result_columnContext) {
	localctx = NewReturning_clause_result_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SQLParserRULE_returning_clause_result_column)
	var _la int

	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserSTAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(577)
			p.Match(SQLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserOPEN_PAR, SQLParserPLUS, SQLParserMINUS, SQLParserCASE_, SQLParserEXISTS_, SQLParserFALSE_, SQLParserLIKE_, SQLParserNOT_, SQLParserNULL_, SQLParserREPLACE_, SQLParserTRUE_, SQLParserIDENTIFIER, SQLParserNUMERIC_LITERAL, SQLParserBIND_PARAMETER, SQLParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(578)
			p.expr(0)
		}
		p.SetState(581)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserAS_ {
			{
				p.SetState(579)
				p.Match(SQLParserAS_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(580)
				p.Column_alias()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJoin_operatorContext is an interface to support dynamic dispatch.
type IJoin_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JOIN_() antlr.TerminalNode
	INNER_() antlr.TerminalNode
	LEFT_() antlr.TerminalNode
	RIGHT_() antlr.TerminalNode
	FULL_() antlr.TerminalNode
	OUTER_() antlr.TerminalNode

	// IsJoin_operatorContext differentiates from other interfaces.
	IsJoin_operatorContext()
}

type Join_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_operatorContext() *Join_operatorContext {
	var p = new(Join_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_join_operator
	return p
}

func InitEmptyJoin_operatorContext(p *Join_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_join_operator
}

func (*Join_operatorContext) IsJoin_operatorContext() {}

func NewJoin_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_operatorContext {
	var p = new(Join_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_join_operator

	return p
}

func (s *Join_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Join_operatorContext) JOIN_() antlr.TerminalNode {
	return s.GetToken(SQLParserJOIN_, 0)
}

func (s *Join_operatorContext) INNER_() antlr.TerminalNode {
	return s.GetToken(SQLParserINNER_, 0)
}

func (s *Join_operatorContext) LEFT_() antlr.TerminalNode {
	return s.GetToken(SQLParserLEFT_, 0)
}

func (s *Join_operatorContext) RIGHT_() antlr.TerminalNode {
	return s.GetToken(SQLParserRIGHT_, 0)
}

func (s *Join_operatorContext) FULL_() antlr.TerminalNode {
	return s.GetToken(SQLParserFULL_, 0)
}

func (s *Join_operatorContext) OUTER_() antlr.TerminalNode {
	return s.GetToken(SQLParserOUTER_, 0)
}

func (s *Join_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Join_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Join_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitJoin_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Join_operator() (localctx IJoin_operatorContext) {
	localctx = NewJoin_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SQLParserRULE_join_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(590)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case SQLParserFULL_, SQLParserLEFT_, SQLParserRIGHT_:
		{
			p.SetState(585)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-46)) & ^0x3f) == 0 && ((int64(1)<<(_la-46))&1073750017) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(587)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserOUTER_ {
			{
				p.SetState(586)
				p.Match(SQLParserOUTER_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SQLParserINNER_:
		{
			p.SetState(589)
			p.Match(SQLParserINNER_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserJOIN_:

	default:
	}
	{
		p.SetState(592)
		p.Match(SQLParserJOIN_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_constraintContext() *Join_constraintContext {
	var p = new(Join_constraintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_join_constraint
	return p
}

func InitEmptyJoin_constraintContext(p *Join_constraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_join_constraint
}

func (*Join_constraintContext) IsJoin_constraintContext() {}

func NewJoin_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_constraintContext {
	var p = new(Join_constraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_join_constraint

	return p
}

func (s *Join_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *Join_constraintContext) ON_() antlr.TerminalNode {
	return s.GetToken(SQLParserON_, 0)
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
	case SQLParserVisitor:
		return t.VisitJoin_constraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Join_constraint() (localctx IJoin_constraintContext) {
	localctx = NewJoin_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SQLParserRULE_join_constraint)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)
		p.Match(SQLParserON_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(595)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_operatorContext() *Compound_operatorContext {
	var p = new(Compound_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_compound_operator
	return p
}

func InitEmptyCompound_operatorContext(p *Compound_operatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_compound_operator
}

func (*Compound_operatorContext) IsCompound_operatorContext() {}

func NewCompound_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_operatorContext {
	var p = new(Compound_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_compound_operator

	return p
}

func (s *Compound_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_operatorContext) UNION_() antlr.TerminalNode {
	return s.GetToken(SQLParserUNION_, 0)
}

func (s *Compound_operatorContext) ALL_() antlr.TerminalNode {
	return s.GetToken(SQLParserALL_, 0)
}

func (s *Compound_operatorContext) INTERSECT_() antlr.TerminalNode {
	return s.GetToken(SQLParserINTERSECT_, 0)
}

func (s *Compound_operatorContext) EXCEPT_() antlr.TerminalNode {
	return s.GetToken(SQLParserEXCEPT_, 0)
}

func (s *Compound_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCompound_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Compound_operator() (localctx ICompound_operatorContext) {
	localctx = NewCompound_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SQLParserRULE_compound_operator)
	var _la int

	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserUNION_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(597)
			p.Match(SQLParserUNION_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(599)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLParserALL_ {
			{
				p.SetState(598)
				p.Match(SQLParserALL_)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case SQLParserINTERSECT_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(601)
			p.Match(SQLParserINTERSECT_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserEXCEPT_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(602)
			p.Match(SQLParserEXCEPT_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_set_subclauseContext() *Update_set_subclauseContext {
	var p = new(Update_set_subclauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_update_set_subclause
	return p
}

func InitEmptyUpdate_set_subclauseContext(p *Update_set_subclauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_update_set_subclause
}

func (*Update_set_subclauseContext) IsUpdate_set_subclauseContext() {}

func NewUpdate_set_subclauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_set_subclauseContext {
	var p = new(Update_set_subclauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_update_set_subclause

	return p
}

func (s *Update_set_subclauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_set_subclauseContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SQLParserASSIGN, 0)
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
	case SQLParserVisitor:
		return t.VisitUpdate_set_subclause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Update_set_subclause() (localctx IUpdate_set_subclauseContext) {
	localctx = NewUpdate_set_subclauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SQLParserRULE_update_set_subclause)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(607)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		{
			p.SetState(605)
			p.Column_name()
		}

	case SQLParserOPEN_PAR:
		{
			p.SetState(606)
			p.Column_name_list()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(609)
		p.Match(SQLParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(610)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	FROM_() antlr.TerminalNode
	WHERE_() antlr.TerminalNode
	Expr() IExprContext
	Returning_clause() IReturning_clauseContext
	Table_or_subquery() ITable_or_subqueryContext
	Join_clause() IJoin_clauseContext

	// IsUpdate_stmtContext differentiates from other interfaces.
	IsUpdate_stmtContext()
}

type Update_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_stmtContext() *Update_stmtContext {
	var p = new(Update_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_update_stmt
	return p
}

func InitEmptyUpdate_stmtContext(p *Update_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_update_stmt
}

func (*Update_stmtContext) IsUpdate_stmtContext() {}

func NewUpdate_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_stmtContext {
	var p = new(Update_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_update_stmt

	return p
}

func (s *Update_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_stmtContext) UPDATE_() antlr.TerminalNode {
	return s.GetToken(SQLParserUPDATE_, 0)
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
	return s.GetToken(SQLParserSET_, 0)
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

func (s *Update_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Update_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Update_stmtContext) FROM_() antlr.TerminalNode {
	return s.GetToken(SQLParserFROM_, 0)
}

func (s *Update_stmtContext) WHERE_() antlr.TerminalNode {
	return s.GetToken(SQLParserWHERE_, 0)
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
	case SQLParserVisitor:
		return t.VisitUpdate_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Update_stmt() (localctx IUpdate_stmtContext) {
	localctx = NewUpdate_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SQLParserRULE_update_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWITH_ {
		{
			p.SetState(612)
			p.Common_table_stmt()
		}

	}
	{
		p.SetState(615)
		p.Match(SQLParserUPDATE_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(616)
		p.Qualified_table_name()
	}
	{
		p.SetState(617)
		p.Match(SQLParserSET_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(618)
		p.Update_set_subclause()
	}
	p.SetState(623)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(619)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)
			p.Update_set_subclause()
		}

		p.SetState(625)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(631)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserFROM_ {
		{
			p.SetState(626)
			p.Match(SQLParserFROM_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(629)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 81, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(627)
				p.Table_or_subquery()
			}

		case 2:
			{
				p.SetState(628)
				p.Join_clause()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	}
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserWHERE_ {
		{
			p.SetState(633)
			p.Match(SQLParserWHERE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(634)
			p.expr(0)
		}

	}
	p.SetState(638)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserRETURNING_ {
		{
			p.SetState(637)
			p.Returning_clause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_name_listContext() *Column_name_listContext {
	var p = new(Column_name_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_name_list
	return p
}

func InitEmptyColumn_name_listContext(p *Column_name_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_name_list
}

func (*Column_name_listContext) IsColumn_name_listContext() {}

func NewColumn_name_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_name_listContext {
	var p = new(Column_name_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_column_name_list

	return p
}

func (s *Column_name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_name_listContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(SQLParserOPEN_PAR, 0)
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
	return s.GetToken(SQLParserCLOSE_PAR, 0)
}

func (s *Column_name_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Column_name_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Column_name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_name_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitColumn_name_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Column_name_list() (localctx IColumn_name_listContext) {
	localctx = NewColumn_name_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SQLParserRULE_column_name_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(640)
		p.Match(SQLParserOPEN_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(641)
		p.Column_name()
	}
	p.SetState(646)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(642)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(643)
			p.Column_name()
		}

		p.SetState(648)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(649)
		p.Match(SQLParserCLOSE_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// IsQualified_table_nameContext differentiates from other interfaces.
	IsQualified_table_nameContext()
}

type Qualified_table_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualified_table_nameContext() *Qualified_table_nameContext {
	var p = new(Qualified_table_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_qualified_table_name
	return p
}

func InitEmptyQualified_table_nameContext(p *Qualified_table_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_qualified_table_name
}

func (*Qualified_table_nameContext) IsQualified_table_nameContext() {}

func NewQualified_table_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Qualified_table_nameContext {
	var p = new(Qualified_table_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_qualified_table_name

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
	return s.GetToken(SQLParserAS_, 0)
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

func (s *Qualified_table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qualified_table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Qualified_table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitQualified_table_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Qualified_table_name() (localctx IQualified_table_nameContext) {
	localctx = NewQualified_table_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SQLParserRULE_qualified_table_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(651)
		p.Table_name()
	}
	p.SetState(654)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserAS_ {
		{
			p.SetState(652)
			p.Match(SQLParserAS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(653)
			p.Table_alias()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_stmtContext() *Order_by_stmtContext {
	var p = new(Order_by_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_order_by_stmt
	return p
}

func InitEmptyOrder_by_stmtContext(p *Order_by_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_order_by_stmt
}

func (*Order_by_stmtContext) IsOrder_by_stmtContext() {}

func NewOrder_by_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_stmtContext {
	var p = new(Order_by_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_order_by_stmt

	return p
}

func (s *Order_by_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_stmtContext) ORDER_() antlr.TerminalNode {
	return s.GetToken(SQLParserORDER_, 0)
}

func (s *Order_by_stmtContext) BY_() antlr.TerminalNode {
	return s.GetToken(SQLParserBY_, 0)
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
	return s.GetTokens(SQLParserCOMMA)
}

func (s *Order_by_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLParserCOMMA, i)
}

func (s *Order_by_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitOrder_by_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Order_by_stmt() (localctx IOrder_by_stmtContext) {
	localctx = NewOrder_by_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SQLParserRULE_order_by_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(656)
		p.Match(SQLParserORDER_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(657)
		p.Match(SQLParserBY_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(658)
		p.Ordering_term()
	}
	p.SetState(663)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLParserCOMMA {
		{
			p.SetState(659)
			p.Match(SQLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(660)
			p.Ordering_term()
		}

		p.SetState(665)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// IsLimit_stmtContext differentiates from other interfaces.
	IsLimit_stmtContext()
}

type Limit_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimit_stmtContext() *Limit_stmtContext {
	var p = new(Limit_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_limit_stmt
	return p
}

func InitEmptyLimit_stmtContext(p *Limit_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_limit_stmt
}

func (*Limit_stmtContext) IsLimit_stmtContext() {}

func NewLimit_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Limit_stmtContext {
	var p = new(Limit_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_limit_stmt

	return p
}

func (s *Limit_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Limit_stmtContext) LIMIT_() antlr.TerminalNode {
	return s.GetToken(SQLParserLIMIT_, 0)
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
	return s.GetToken(SQLParserOFFSET_, 0)
}

func (s *Limit_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Limit_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Limit_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitLimit_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Limit_stmt() (localctx ILimit_stmtContext) {
	localctx = NewLimit_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SQLParserRULE_limit_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(666)
		p.Match(SQLParserLIMIT_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(667)
		p.expr(0)
	}
	p.SetState(670)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserOFFSET_ {
		{
			p.SetState(668)
			p.Match(SQLParserOFFSET_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(669)
			p.expr(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrdering_termContext is an interface to support dynamic dispatch.
type IOrdering_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Asc_desc() IAsc_descContext
	NULLS_() antlr.TerminalNode
	FIRST_() antlr.TerminalNode
	LAST_() antlr.TerminalNode

	// IsOrdering_termContext differentiates from other interfaces.
	IsOrdering_termContext()
}

type Ordering_termContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdering_termContext() *Ordering_termContext {
	var p = new(Ordering_termContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_ordering_term
	return p
}

func InitEmptyOrdering_termContext(p *Ordering_termContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_ordering_term
}

func (*Ordering_termContext) IsOrdering_termContext() {}

func NewOrdering_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ordering_termContext {
	var p = new(Ordering_termContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_ordering_term

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
	return s.GetToken(SQLParserNULLS_, 0)
}

func (s *Ordering_termContext) FIRST_() antlr.TerminalNode {
	return s.GetToken(SQLParserFIRST_, 0)
}

func (s *Ordering_termContext) LAST_() antlr.TerminalNode {
	return s.GetToken(SQLParserLAST_, 0)
}

func (s *Ordering_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ordering_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ordering_termContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitOrdering_term(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Ordering_term() (localctx IOrdering_termContext) {
	localctx = NewOrdering_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SQLParserRULE_ordering_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(672)
		p.expr(0)
	}
	p.SetState(674)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserASC_ || _la == SQLParserDESC_ {
		{
			p.SetState(673)
			p.Asc_desc()
		}

	}
	p.SetState(678)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLParserNULLS_ {
		{
			p.SetState(676)
			p.Match(SQLParserNULLS_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(677)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SQLParserFIRST_ || _la == SQLParserLAST_) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsc_descContext() *Asc_descContext {
	var p = new(Asc_descContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_asc_desc
	return p
}

func InitEmptyAsc_descContext(p *Asc_descContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_asc_desc
}

func (*Asc_descContext) IsAsc_descContext() {}

func NewAsc_descContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asc_descContext {
	var p = new(Asc_descContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_asc_desc

	return p
}

func (s *Asc_descContext) GetParser() antlr.Parser { return s.parser }

func (s *Asc_descContext) ASC_() antlr.TerminalNode {
	return s.GetToken(SQLParserASC_, 0)
}

func (s *Asc_descContext) DESC_() antlr.TerminalNode {
	return s.GetToken(SQLParserDESC_, 0)
}

func (s *Asc_descContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asc_descContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Asc_descContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitAsc_desc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Asc_desc() (localctx IAsc_descContext) {
	localctx = NewAsc_descContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SQLParserRULE_asc_desc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(680)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLParserASC_ || _la == SQLParserDESC_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_keywordContext is an interface to support dynamic dispatch.
type IFunction_keywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIKE_() antlr.TerminalNode
	REPLACE_() antlr.TerminalNode

	// IsFunction_keywordContext differentiates from other interfaces.
	IsFunction_keywordContext()
}

type Function_keywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_keywordContext() *Function_keywordContext {
	var p = new(Function_keywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_keyword
	return p
}

func InitEmptyFunction_keywordContext(p *Function_keywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_keyword
}

func (*Function_keywordContext) IsFunction_keywordContext() {}

func NewFunction_keywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_keywordContext {
	var p = new(Function_keywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_keyword

	return p
}

func (s *Function_keywordContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_keywordContext) LIKE_() antlr.TerminalNode {
	return s.GetToken(SQLParserLIKE_, 0)
}

func (s *Function_keywordContext) REPLACE_() antlr.TerminalNode {
	return s.GetToken(SQLParserREPLACE_, 0)
}

func (s *Function_keywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_keywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_keywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitFunction_keyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Function_keyword() (localctx IFunction_keywordContext) {
	localctx = NewFunction_keywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SQLParserRULE_function_keyword)
	p.SetState(685)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 1)

	case SQLParserLIKE_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(683)
			p.Match(SQLParserLIKE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserREPLACE_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(684)
			p.Match(SQLParserREPLACE_)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_name
	return p
}

func InitEmptyFunction_nameContext(p *Function_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_function_name
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
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
	case SQLParserVisitor:
		return t.VisitFunction_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Function_name() (localctx IFunction_nameContext) {
	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SQLParserRULE_function_name)
	p.SetState(689)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(687)
			p.Match(SQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLParserOPEN_PAR, SQLParserLIKE_, SQLParserREPLACE_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(688)
			p.Function_keyword()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_name
	return p
}

func InitEmptyTable_nameContext(p *Table_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_name
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitTable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Table_name() (localctx ITable_nameContext) {
	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SQLParserRULE_table_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(691)
		p.Match(SQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_aliasContext() *Table_aliasContext {
	var p = new(Table_aliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_alias
	return p
}

func InitEmptyTable_aliasContext(p *Table_aliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_table_alias
}

func (*Table_aliasContext) IsTable_aliasContext() {}

func NewTable_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_aliasContext {
	var p = new(Table_aliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_table_alias

	return p
}

func (s *Table_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Table_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitTable_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Table_alias() (localctx ITable_aliasContext) {
	localctx = NewTable_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SQLParserRULE_table_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(693)
		p.Match(SQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_name
	return p
}

func InitEmptyColumn_nameContext(p *Column_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_name
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitColumn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Column_name() (localctx IColumn_nameContext) {
	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SQLParserRULE_column_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(695)
		p.Match(SQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_aliasContext() *Column_aliasContext {
	var p = new(Column_aliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_alias
	return p
}

func InitEmptyColumn_aliasContext(p *Column_aliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_column_alias
}

func (*Column_aliasContext) IsColumn_aliasContext() {}

func NewColumn_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_aliasContext {
	var p = new(Column_aliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_column_alias

	return p
}

func (s *Column_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_aliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Column_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_aliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitColumn_alias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Column_alias() (localctx IColumn_aliasContext) {
	localctx = NewColumn_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SQLParserRULE_column_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(697)
		p.Match(SQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollation_nameContext() *Collation_nameContext {
	var p = new(Collation_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_collation_name
	return p
}

func InitEmptyCollation_nameContext(p *Collation_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_collation_name
}

func (*Collation_nameContext) IsCollation_nameContext() {}

func NewCollation_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collation_nameContext {
	var p = new(Collation_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_collation_name

	return p
}

func (s *Collation_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Collation_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Collation_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collation_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collation_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitCollation_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Collation_name() (localctx ICollation_nameContext) {
	localctx = NewCollation_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SQLParserRULE_collation_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(699)
		p.Match(SQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_nameContext() *Index_nameContext {
	var p = new(Index_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_index_name
	return p
}

func InitEmptyIndex_nameContext(p *Index_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLParserRULE_index_name
}

func (*Index_nameContext) IsIndex_nameContext() {}

func NewIndex_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_nameContext {
	var p = new(Index_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLParserRULE_index_name

	return p
}

func (s *Index_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLParserIDENTIFIER, 0)
}

func (s *Index_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLParserVisitor:
		return t.VisitIndex_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLParser) Index_name() (localctx IIndex_nameContext) {
	localctx = NewIndex_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SQLParserRULE_index_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(701)
		p.Match(SQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SQLParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
