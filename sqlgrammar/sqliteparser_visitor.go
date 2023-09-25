// Code generated from SQLiteParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package sqlgrammar // SQLiteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by SQLiteParser.
type SQLiteParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SQLiteParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by SQLiteParser#sql_stmt_list.
	VisitSql_stmt_list(ctx *Sql_stmt_listContext) interface{}

	// Visit a parse tree produced by SQLiteParser#sql_stmt.
	VisitSql_stmt(ctx *Sql_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#indexed_column.
	VisitIndexed_column(ctx *Indexed_columnContext) interface{}

	// Visit a parse tree produced by SQLiteParser#cte_table_name.
	VisitCte_table_name(ctx *Cte_table_nameContext) interface{}

	// Visit a parse tree produced by SQLiteParser#common_table_expression.
	VisitCommon_table_expression(ctx *Common_table_expressionContext) interface{}

	// Visit a parse tree produced by SQLiteParser#common_table_stmt.
	VisitCommon_table_stmt(ctx *Common_table_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#delete_stmt.
	VisitDelete_stmt(ctx *Delete_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SQLiteParser#literal_value.
	VisitLiteral_value(ctx *Literal_valueContext) interface{}

	// Visit a parse tree produced by SQLiteParser#value_row.
	VisitValue_row(ctx *Value_rowContext) interface{}

	// Visit a parse tree produced by SQLiteParser#values_clause.
	VisitValues_clause(ctx *Values_clauseContext) interface{}

	// Visit a parse tree produced by SQLiteParser#insert_stmt.
	VisitInsert_stmt(ctx *Insert_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#returning_clause.
	VisitReturning_clause(ctx *Returning_clauseContext) interface{}

	// Visit a parse tree produced by SQLiteParser#upsert_update.
	VisitUpsert_update(ctx *Upsert_updateContext) interface{}

	// Visit a parse tree produced by SQLiteParser#upsert_clause.
	VisitUpsert_clause(ctx *Upsert_clauseContext) interface{}

	// Visit a parse tree produced by SQLiteParser#select_stmt_core.
	VisitSelect_stmt_core(ctx *Select_stmt_coreContext) interface{}

	// Visit a parse tree produced by SQLiteParser#select_stmt.
	VisitSelect_stmt(ctx *Select_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#join_clause.
	VisitJoin_clause(ctx *Join_clauseContext) interface{}

	// Visit a parse tree produced by SQLiteParser#select_core.
	VisitSelect_core(ctx *Select_coreContext) interface{}

	// Visit a parse tree produced by SQLiteParser#table_or_subquery.
	VisitTable_or_subquery(ctx *Table_or_subqueryContext) interface{}

	// Visit a parse tree produced by SQLiteParser#result_column.
	VisitResult_column(ctx *Result_columnContext) interface{}

	// Visit a parse tree produced by SQLiteParser#returning_clause_result_column.
	VisitReturning_clause_result_column(ctx *Returning_clause_result_columnContext) interface{}

	// Visit a parse tree produced by SQLiteParser#join_operator.
	VisitJoin_operator(ctx *Join_operatorContext) interface{}

	// Visit a parse tree produced by SQLiteParser#join_constraint.
	VisitJoin_constraint(ctx *Join_constraintContext) interface{}

	// Visit a parse tree produced by SQLiteParser#compound_operator.
	VisitCompound_operator(ctx *Compound_operatorContext) interface{}

	// Visit a parse tree produced by SQLiteParser#update_set_subclause.
	VisitUpdate_set_subclause(ctx *Update_set_subclauseContext) interface{}

	// Visit a parse tree produced by SQLiteParser#update_stmt.
	VisitUpdate_stmt(ctx *Update_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#column_name_list.
	VisitColumn_name_list(ctx *Column_name_listContext) interface{}

	// Visit a parse tree produced by SQLiteParser#qualified_table_name.
	VisitQualified_table_name(ctx *Qualified_table_nameContext) interface{}

	// Visit a parse tree produced by SQLiteParser#order_by_stmt.
	VisitOrder_by_stmt(ctx *Order_by_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#limit_stmt.
	VisitLimit_stmt(ctx *Limit_stmtContext) interface{}

	// Visit a parse tree produced by SQLiteParser#ordering_term.
	VisitOrdering_term(ctx *Ordering_termContext) interface{}

	// Visit a parse tree produced by SQLiteParser#asc_desc.
	VisitAsc_desc(ctx *Asc_descContext) interface{}

	// Visit a parse tree produced by SQLiteParser#function_keyword.
	VisitFunction_keyword(ctx *Function_keywordContext) interface{}

	// Visit a parse tree produced by SQLiteParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by SQLiteParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by SQLiteParser#table_alias.
	VisitTable_alias(ctx *Table_aliasContext) interface{}

	// Visit a parse tree produced by SQLiteParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by SQLiteParser#column_alias.
	VisitColumn_alias(ctx *Column_aliasContext) interface{}

	// Visit a parse tree produced by SQLiteParser#collation_name.
	VisitCollation_name(ctx *Collation_nameContext) interface{}

	// Visit a parse tree produced by SQLiteParser#index_name.
	VisitIndex_name(ctx *Index_nameContext) interface{}
}
