// Code generated from SQLiteParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package sqlgrammar // SQLiteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseSQLiteParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSQLiteParserVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitSql_stmt_list(ctx *Sql_stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitSql_stmt(ctx *Sql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitIndexed_column(ctx *Indexed_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitCte_table_name(ctx *Cte_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitCommon_table_expression(ctx *Common_table_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitCommon_table_stmt(ctx *Common_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitDelete_stmt(ctx *Delete_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitLiteral_value(ctx *Literal_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitValue_row(ctx *Value_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitValues_clause(ctx *Values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitInsert_stmt(ctx *Insert_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitReturning_clause(ctx *Returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitUpsert_update(ctx *Upsert_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitUpsert_clause(ctx *Upsert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitSelect_stmt_core(ctx *Select_stmt_coreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitSelect_stmt(ctx *Select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitJoin_clause(ctx *Join_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitSelect_core(ctx *Select_coreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitTable_or_subquery(ctx *Table_or_subqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitResult_column(ctx *Result_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitReturning_clause_result_column(ctx *Returning_clause_result_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitJoin_operator(ctx *Join_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitJoin_constraint(ctx *Join_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitCompound_operator(ctx *Compound_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitUpdate_set_subclause(ctx *Update_set_subclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitUpdate_stmt(ctx *Update_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitColumn_name_list(ctx *Column_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitQualified_table_name(ctx *Qualified_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitOrder_by_stmt(ctx *Order_by_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitLimit_stmt(ctx *Limit_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitOrdering_term(ctx *Ordering_termContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitAsc_desc(ctx *Asc_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitColumn_alias(ctx *Column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitCollation_name(ctx *Collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitIndex_name(ctx *Index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLiteParserVisitor) VisitAny_name(ctx *Any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}
