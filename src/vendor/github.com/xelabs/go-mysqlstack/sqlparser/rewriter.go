// Code generated by visitorgen/main/main.go. DO NOT EDIT.

package sqlparser

//go:generate go run ./visitorgen/main -input=ast.go -output=rewriter.go

import (
	"reflect"
)

type replacerFunc func(newNode, parent SQLNode)

// application carries all the shared data so we can pass it around cheaply.
type application struct {
	pre, post ApplyFunc
	cursor    Cursor
}

func replaceAliasedExprAs(newNode, parent SQLNode) {
	parent.(*AliasedExpr).As = newNode.(ColIdent)
}

func replaceAliasedExprExpr(newNode, parent SQLNode) {
	parent.(*AliasedExpr).Expr = newNode.(Expr)
}

func replaceAliasedTableExprAs(newNode, parent SQLNode) {
	parent.(*AliasedTableExpr).As = newNode.(TableIdent)
}

func replaceAliasedTableExprExpr(newNode, parent SQLNode) {
	parent.(*AliasedTableExpr).Expr = newNode.(SimpleTableExpr)
}

func replaceAliasedTableExprHints(newNode, parent SQLNode) {
	parent.(*AliasedTableExpr).Hints = newNode.(*IndexHints)
}

func replaceAndExprLeft(newNode, parent SQLNode) {
	parent.(*AndExpr).Left = newNode.(Expr)
}

func replaceAndExprRight(newNode, parent SQLNode) {
	parent.(*AndExpr).Right = newNode.(Expr)
}

func replaceBinaryExprLeft(newNode, parent SQLNode) {
	parent.(*BinaryExpr).Left = newNode.(Expr)
}

func replaceBinaryExprRight(newNode, parent SQLNode) {
	parent.(*BinaryExpr).Right = newNode.(Expr)
}

func replaceCaseExprElse(newNode, parent SQLNode) {
	parent.(*CaseExpr).Else = newNode.(Expr)
}

func replaceCaseExprExpr(newNode, parent SQLNode) {
	parent.(*CaseExpr).Expr = newNode.(Expr)
}

type replaceCaseExprWhens int

func (r *replaceCaseExprWhens) replace(newNode, container SQLNode) {
	container.(*CaseExpr).Whens[int(*r)] = newNode.(*When)
}

func (r *replaceCaseExprWhens) inc() {
	*r++
}

func replaceChecksumTable(newNode, parent SQLNode) {
	parent.(*Checksum).Table = newNode.(TableName)
}

func replaceColNameName(newNode, parent SQLNode) {
	parent.(*ColName).Name = newNode.(ColIdent)
}

func replaceColNameQualifier(newNode, parent SQLNode) {
	parent.(*ColName).Qualifier = newNode.(TableName)
}

func replaceCollateExprExpr(newNode, parent SQLNode) {
	parent.(*CollateExpr).Expr = newNode.(Expr)
}

func replaceColumnDefinitionName(newNode, parent SQLNode) {
	parent.(*ColumnDefinition).Name = newNode.(ColIdent)
}

func replaceColumnTypeAutoincrement(newNode, parent SQLNode) {
	parent.(*ColumnType).Autoincrement = newNode.(BoolVal)
}

func replaceColumnTypeCollate(newNode, parent SQLNode) {
	parent.(*ColumnType).Collate = newNode.(*SQLVal)
}

func replaceColumnTypeComment(newNode, parent SQLNode) {
	parent.(*ColumnType).Comment = newNode.(*SQLVal)
}

func replaceColumnTypeDefault(newNode, parent SQLNode) {
	parent.(*ColumnType).Default = newNode.(*SQLVal)
}

func replaceColumnTypeLength(newNode, parent SQLNode) {
	parent.(*ColumnType).Length = newNode.(*SQLVal)
}

func replaceColumnTypeNotNull(newNode, parent SQLNode) {
	parent.(*ColumnType).NotNull = newNode.(BoolVal)
}

func replaceColumnTypeScale(newNode, parent SQLNode) {
	parent.(*ColumnType).Scale = newNode.(*SQLVal)
}

func replaceColumnTypeUnsigned(newNode, parent SQLNode) {
	parent.(*ColumnType).Unsigned = newNode.(BoolVal)
}

func replaceColumnTypeZerofill(newNode, parent SQLNode) {
	parent.(*ColumnType).Zerofill = newNode.(BoolVal)
}

type replaceColumnsItems int

func (r *replaceColumnsItems) replace(newNode, container SQLNode) {
	container.(Columns)[int(*r)] = newNode.(ColIdent)
}

func (r *replaceColumnsItems) inc() {
	*r++
}

func replaceComparisonExprEscape(newNode, parent SQLNode) {
	parent.(*ComparisonExpr).Escape = newNode.(Expr)
}

func replaceComparisonExprLeft(newNode, parent SQLNode) {
	parent.(*ComparisonExpr).Left = newNode.(Expr)
}

func replaceComparisonExprRight(newNode, parent SQLNode) {
	parent.(*ComparisonExpr).Right = newNode.(Expr)
}

func replaceConvertExprExpr(newNode, parent SQLNode) {
	parent.(*ConvertExpr).Expr = newNode.(Expr)
}

func replaceConvertExprType(newNode, parent SQLNode) {
	parent.(*ConvertExpr).Type = newNode.(*ConvertType)
}

func replaceConvertTypeLength(newNode, parent SQLNode) {
	parent.(*ConvertType).Length = newNode.(*SQLVal)
}

func replaceConvertTypeScale(newNode, parent SQLNode) {
	parent.(*ConvertType).Scale = newNode.(*SQLVal)
}

func replaceConvertUsingExprExpr(newNode, parent SQLNode) {
	parent.(*ConvertUsingExpr).Expr = newNode.(Expr)
}

func replaceDDLDatabase(newNode, parent SQLNode) {
	parent.(*DDL).Database = newNode.(TableIdent)
}

func replaceDDLDatabaseOptions(newNode, parent SQLNode) {
	parent.(*DDL).DatabaseOptions = newNode.(DatabaseOptionListOpt)
}

func replaceDDLIndexOpts(newNode, parent SQLNode) {
	parent.(*DDL).IndexOpts = newNode.(*IndexOptions)
}

func replaceDDLModifyColumnDef(newNode, parent SQLNode) {
	parent.(*DDL).ModifyColumnDef = newNode.(*ColumnDefinition)
}

func replaceDDLNewName(newNode, parent SQLNode) {
	parent.(*DDL).NewName = newNode.(TableName)
}

func replaceDDLTable(newNode, parent SQLNode) {
	parent.(*DDL).Table = newNode.(TableName)
}

func replaceDDLTableSpec(newNode, parent SQLNode) {
	parent.(*DDL).TableSpec = newNode.(*TableSpec)
}

func replaceDDLTables(newNode, parent SQLNode) {
	parent.(*DDL).Tables = newNode.(TableNames)
}

func replaceDDLindexLockAndAlgorithm(newNode, parent SQLNode) {
	parent.(*DDL).indexLockAndAlgorithm = newNode.(*IndexLockAndAlgorithm)
}

func replaceDeleteComments(newNode, parent SQLNode) {
	parent.(*Delete).Comments = newNode.(Comments)
}

func replaceDeleteLimit(newNode, parent SQLNode) {
	parent.(*Delete).Limit = newNode.(*Limit)
}

func replaceDeleteOrderBy(newNode, parent SQLNode) {
	parent.(*Delete).OrderBy = newNode.(OrderBy)
}

func replaceDeleteTable(newNode, parent SQLNode) {
	parent.(*Delete).Table = newNode.(TableName)
}

func replaceDeleteWhere(newNode, parent SQLNode) {
	parent.(*Delete).Where = newNode.(*Where)
}

func replaceExistsExprSubquery(newNode, parent SQLNode) {
	parent.(*ExistsExpr).Subquery = newNode.(*Subquery)
}

func replaceExplainStatement(newNode, parent SQLNode) {
	parent.(*Explain).Statement = newNode.(Statement)
}

type replaceExprsItems int

func (r *replaceExprsItems) replace(newNode, container SQLNode) {
	container.(Exprs)[int(*r)] = newNode.(Expr)
}

func (r *replaceExprsItems) inc() {
	*r++
}

func replaceFuncExprExprs(newNode, parent SQLNode) {
	parent.(*FuncExpr).Exprs = newNode.(SelectExprs)
}

func replaceFuncExprName(newNode, parent SQLNode) {
	parent.(*FuncExpr).Name = newNode.(ColIdent)
}

func replaceFuncExprQualifier(newNode, parent SQLNode) {
	parent.(*FuncExpr).Qualifier = newNode.(TableIdent)
}

type replaceGroupByItems int

func (r *replaceGroupByItems) replace(newNode, container SQLNode) {
	container.(GroupBy)[int(*r)] = newNode.(Expr)
}

func (r *replaceGroupByItems) inc() {
	*r++
}

func replaceGroupConcatExprExprs(newNode, parent SQLNode) {
	parent.(*GroupConcatExpr).Exprs = newNode.(SelectExprs)
}

func replaceGroupConcatExprOrderBy(newNode, parent SQLNode) {
	parent.(*GroupConcatExpr).OrderBy = newNode.(OrderBy)
}

func replaceIndexDefinitionName(newNode, parent SQLNode) {
	parent.(*IndexDefinition).Name = newNode.(ColIdent)
}

func replaceIndexDefinitionOpts(newNode, parent SQLNode) {
	parent.(*IndexDefinition).Opts = newNode.(*IndexOptions)
}

type replaceIndexHintsIndexes int

func (r *replaceIndexHintsIndexes) replace(newNode, container SQLNode) {
	container.(*IndexHints).Indexes[int(*r)] = newNode.(ColIdent)
}

func (r *replaceIndexHintsIndexes) inc() {
	*r++
}

func replaceIndexOptionsBlockSize(newNode, parent SQLNode) {
	parent.(*IndexOptions).BlockSize = newNode.(*SQLVal)
}

func replaceInsertColumns(newNode, parent SQLNode) {
	parent.(*Insert).Columns = newNode.(Columns)
}

func replaceInsertComments(newNode, parent SQLNode) {
	parent.(*Insert).Comments = newNode.(Comments)
}

func replaceInsertOnDup(newNode, parent SQLNode) {
	parent.(*Insert).OnDup = newNode.(OnDup)
}

func replaceInsertRows(newNode, parent SQLNode) {
	parent.(*Insert).Rows = newNode.(InsertRows)
}

func replaceInsertTable(newNode, parent SQLNode) {
	parent.(*Insert).Table = newNode.(TableName)
}

func replaceIntervalExprExpr(newNode, parent SQLNode) {
	parent.(*IntervalExpr).Expr = newNode.(Expr)
}

func replaceIsExprExpr(newNode, parent SQLNode) {
	parent.(*IsExpr).Expr = newNode.(Expr)
}

func replaceJoinTableExprLeftExpr(newNode, parent SQLNode) {
	parent.(*JoinTableExpr).LeftExpr = newNode.(TableExpr)
}

func replaceJoinTableExprOn(newNode, parent SQLNode) {
	parent.(*JoinTableExpr).On = newNode.(Expr)
}

func replaceJoinTableExprRightExpr(newNode, parent SQLNode) {
	parent.(*JoinTableExpr).RightExpr = newNode.(TableExpr)
}

func replaceLimitOffset(newNode, parent SQLNode) {
	parent.(*Limit).Offset = newNode.(Expr)
}

func replaceLimitRowcount(newNode, parent SQLNode) {
	parent.(*Limit).Rowcount = newNode.(Expr)
}

func replaceMatchExprColumns(newNode, parent SQLNode) {
	parent.(*MatchExpr).Columns = newNode.(SelectExprs)
}

func replaceMatchExprExpr(newNode, parent SQLNode) {
	parent.(*MatchExpr).Expr = newNode.(Expr)
}

func replaceNextvalExpr(newNode, parent SQLNode) {
	tmp := parent.(Nextval)
	tmp.Expr = newNode.(Expr)
}

func replaceNotExprExpr(newNode, parent SQLNode) {
	parent.(*NotExpr).Expr = newNode.(Expr)
}

type replaceOnDupItems int

func (r *replaceOnDupItems) replace(newNode, container SQLNode) {
	container.(OnDup)[int(*r)] = newNode.(*UpdateExpr)
}

func (r *replaceOnDupItems) inc() {
	*r++
}

func replaceOptValValue(newNode, parent SQLNode) {
	parent.(*OptVal).Value = newNode.(Expr)
}

func replaceOrExprLeft(newNode, parent SQLNode) {
	parent.(*OrExpr).Left = newNode.(Expr)
}

func replaceOrExprRight(newNode, parent SQLNode) {
	parent.(*OrExpr).Right = newNode.(Expr)
}

func replaceOrderExpr(newNode, parent SQLNode) {
	parent.(*Order).Expr = newNode.(Expr)
}

type replaceOrderByItems int

func (r *replaceOrderByItems) replace(newNode, container SQLNode) {
	container.(OrderBy)[int(*r)] = newNode.(*Order)
}

func (r *replaceOrderByItems) inc() {
	*r++
}

func replaceParenExprExpr(newNode, parent SQLNode) {
	parent.(*ParenExpr).Expr = newNode.(Expr)
}

func replaceParenSelectSelect(newNode, parent SQLNode) {
	parent.(*ParenSelect).Select = newNode.(SelectStatement)
}

func replaceParenTableExprExprs(newNode, parent SQLNode) {
	parent.(*ParenTableExpr).Exprs = newNode.(TableExprs)
}

func replaceRadonNewName(newNode, parent SQLNode) {
	parent.(*Radon).NewName = newNode.(TableName)
}

func replaceRadonRow(newNode, parent SQLNode) {
	parent.(*Radon).Row = newNode.(ValTuple)
}

func replaceRadonTable(newNode, parent SQLNode) {
	parent.(*Radon).Table = newNode.(TableName)
}

func replaceRangeCondFrom(newNode, parent SQLNode) {
	parent.(*RangeCond).From = newNode.(Expr)
}

func replaceRangeCondLeft(newNode, parent SQLNode) {
	parent.(*RangeCond).Left = newNode.(Expr)
}

func replaceRangeCondTo(newNode, parent SQLNode) {
	parent.(*RangeCond).To = newNode.(Expr)
}

func replaceSelectComments(newNode, parent SQLNode) {
	parent.(*Select).Comments = newNode.(Comments)
}

func replaceSelectFrom(newNode, parent SQLNode) {
	parent.(*Select).From = newNode.(TableExprs)
}

func replaceSelectGroupBy(newNode, parent SQLNode) {
	parent.(*Select).GroupBy = newNode.(GroupBy)
}

func replaceSelectHaving(newNode, parent SQLNode) {
	parent.(*Select).Having = newNode.(*Where)
}

func replaceSelectLimit(newNode, parent SQLNode) {
	parent.(*Select).Limit = newNode.(*Limit)
}

func replaceSelectOrderBy(newNode, parent SQLNode) {
	parent.(*Select).OrderBy = newNode.(OrderBy)
}

func replaceSelectSelectExprs(newNode, parent SQLNode) {
	parent.(*Select).SelectExprs = newNode.(SelectExprs)
}

func replaceSelectWhere(newNode, parent SQLNode) {
	parent.(*Select).Where = newNode.(*Where)
}

type replaceSelectExprsItems int

func (r *replaceSelectExprsItems) replace(newNode, container SQLNode) {
	container.(SelectExprs)[int(*r)] = newNode.(SelectExpr)
}

func (r *replaceSelectExprsItems) inc() {
	*r++
}

func replaceSetComments(newNode, parent SQLNode) {
	parent.(*Set).Comments = newNode.(Comments)
}

func replaceSetExprs(newNode, parent SQLNode) {
	parent.(*Set).Exprs = newNode.(SetExprs)
}

func replaceSetExprType(newNode, parent SQLNode) {
	parent.(*SetExpr).Type = newNode.(ColIdent)
}

func replaceSetExprVal(newNode, parent SQLNode) {
	parent.(*SetExpr).Val = newNode.(SetVal)
}

type replaceSetExprsItems int

func (r *replaceSetExprsItems) replace(newNode, container SQLNode) {
	container.(SetExprs)[int(*r)] = newNode.(*SetExpr)
}

func (r *replaceSetExprsItems) inc() {
	*r++
}

func replaceShowFilter(newNode, parent SQLNode) {
	parent.(*Show).Filter = newNode.(*ShowFilter)
}

func replaceShowLimit(newNode, parent SQLNode) {
	parent.(*Show).Limit = newNode.(*Limit)
}

func replaceShowTable(newNode, parent SQLNode) {
	parent.(*Show).Table = newNode.(TableName)
}

func replaceShowFilterFilter(newNode, parent SQLNode) {
	parent.(*ShowFilter).Filter = newNode.(Expr)
}

func replaceStarExprTableName(newNode, parent SQLNode) {
	parent.(*StarExpr).TableName = newNode.(TableName)
}

func replaceSubquerySelect(newNode, parent SQLNode) {
	parent.(*Subquery).Select = newNode.(SelectStatement)
}

type replaceTableExprsItems int

func (r *replaceTableExprsItems) replace(newNode, container SQLNode) {
	container.(TableExprs)[int(*r)] = newNode.(TableExpr)
}

func (r *replaceTableExprsItems) inc() {
	*r++
}

func replaceTableNameName(newNode, parent SQLNode) {
	tmp := parent.(TableName)
	tmp.Name = newNode.(TableIdent)
}

func replaceTableNameQualifier(newNode, parent SQLNode) {
	tmp := parent.(TableName)
	tmp.Qualifier = newNode.(TableIdent)
}

type replaceTableNamesItems int

func (r *replaceTableNamesItems) replace(newNode, container SQLNode) {
	container.(TableNames)[int(*r)] = newNode.(TableName)
}

func (r *replaceTableNamesItems) inc() {
	*r++
}

type replaceTableSpecColumns int

func (r *replaceTableSpecColumns) replace(newNode, container SQLNode) {
	container.(*TableSpec).Columns[int(*r)] = newNode.(*ColumnDefinition)
}

func (r *replaceTableSpecColumns) inc() {
	*r++
}

type replaceTableSpecIndexes int

func (r *replaceTableSpecIndexes) replace(newNode, container SQLNode) {
	container.(*TableSpec).Indexes[int(*r)] = newNode.(*IndexDefinition)
}

func (r *replaceTableSpecIndexes) inc() {
	*r++
}

func replaceTableSpecOptions(newNode, parent SQLNode) {
	parent.(*TableSpec).Options = newNode.(TableOptions)
}

func replaceUnaryExprExpr(newNode, parent SQLNode) {
	parent.(*UnaryExpr).Expr = newNode.(Expr)
}

func replaceUnionLeft(newNode, parent SQLNode) {
	parent.(*Union).Left = newNode.(SelectStatement)
}

func replaceUnionLimit(newNode, parent SQLNode) {
	parent.(*Union).Limit = newNode.(*Limit)
}

func replaceUnionOrderBy(newNode, parent SQLNode) {
	parent.(*Union).OrderBy = newNode.(OrderBy)
}

func replaceUnionRight(newNode, parent SQLNode) {
	parent.(*Union).Right = newNode.(SelectStatement)
}

func replaceUpdateComments(newNode, parent SQLNode) {
	parent.(*Update).Comments = newNode.(Comments)
}

func replaceUpdateExprs(newNode, parent SQLNode) {
	parent.(*Update).Exprs = newNode.(UpdateExprs)
}

func replaceUpdateLimit(newNode, parent SQLNode) {
	parent.(*Update).Limit = newNode.(*Limit)
}

func replaceUpdateOrderBy(newNode, parent SQLNode) {
	parent.(*Update).OrderBy = newNode.(OrderBy)
}

func replaceUpdateTable(newNode, parent SQLNode) {
	parent.(*Update).Table = newNode.(TableName)
}

func replaceUpdateWhere(newNode, parent SQLNode) {
	parent.(*Update).Where = newNode.(*Where)
}

func replaceUpdateExprExpr(newNode, parent SQLNode) {
	parent.(*UpdateExpr).Expr = newNode.(Expr)
}

func replaceUpdateExprName(newNode, parent SQLNode) {
	parent.(*UpdateExpr).Name = newNode.(*ColName)
}

type replaceUpdateExprsItems int

func (r *replaceUpdateExprsItems) replace(newNode, container SQLNode) {
	container.(UpdateExprs)[int(*r)] = newNode.(*UpdateExpr)
}

func (r *replaceUpdateExprsItems) inc() {
	*r++
}

func replaceUseDBName(newNode, parent SQLNode) {
	parent.(*Use).DBName = newNode.(TableIdent)
}

type replaceValTupleItems int

func (r *replaceValTupleItems) replace(newNode, container SQLNode) {
	container.(ValTuple)[int(*r)] = newNode.(Expr)
}

func (r *replaceValTupleItems) inc() {
	*r++
}

type replaceValuesItems int

func (r *replaceValuesItems) replace(newNode, container SQLNode) {
	container.(Values)[int(*r)] = newNode.(ValTuple)
}

func (r *replaceValuesItems) inc() {
	*r++
}

func replaceValuesFuncExprName(newNode, parent SQLNode) {
	parent.(*ValuesFuncExpr).Name = newNode.(ColIdent)
}

func replaceValuesFuncExprResolved(newNode, parent SQLNode) {
	parent.(*ValuesFuncExpr).Resolved = newNode.(Expr)
}

func replaceWhenCond(newNode, parent SQLNode) {
	parent.(*When).Cond = newNode.(Expr)
}

func replaceWhenVal(newNode, parent SQLNode) {
	parent.(*When).Val = newNode.(Expr)
}

func replaceWhereExpr(newNode, parent SQLNode) {
	parent.(*Where).Expr = newNode.(Expr)
}

// apply is where the visiting happens. Here is where we keep the big switch-case that will be used
// to do the actual visiting of SQLNodes
func (a *application) apply(parent, node SQLNode, replacer replacerFunc) {
	if node == nil || isNilValue(node) {
		return
	}

	// avoid heap-allocating a new cursor for each apply call; reuse a.cursor instead
	saved := a.cursor
	a.cursor.replacer = replacer
	a.cursor.node = node
	a.cursor.parent = parent

	if a.pre != nil && !a.pre(&a.cursor) {
		a.cursor = saved
		return
	}

	// walk children
	// (the order of the cases is alphabetical)
	switch n := node.(type) {
	case nil:
	case *AliasedExpr:
		a.apply(node, n.As, replaceAliasedExprAs)
		a.apply(node, n.Expr, replaceAliasedExprExpr)

	case *AliasedTableExpr:
		a.apply(node, n.As, replaceAliasedTableExprAs)
		a.apply(node, n.Expr, replaceAliasedTableExprExpr)
		a.apply(node, n.Hints, replaceAliasedTableExprHints)

	case *AndExpr:
		a.apply(node, n.Left, replaceAndExprLeft)
		a.apply(node, n.Right, replaceAndExprRight)

	case *BinaryExpr:
		a.apply(node, n.Left, replaceBinaryExprLeft)
		a.apply(node, n.Right, replaceBinaryExprRight)

	case BoolVal:

	case *CaseExpr:
		a.apply(node, n.Else, replaceCaseExprElse)
		a.apply(node, n.Expr, replaceCaseExprExpr)
		replacerWhens := replaceCaseExprWhens(0)
		replacerWhensB := &replacerWhens
		for _, item := range n.Whens {
			a.apply(node, item, replacerWhensB.replace)
			replacerWhensB.inc()
		}

	case *Checksum:
		a.apply(node, n.Table, replaceChecksumTable)

	case ColIdent:

	case *ColName:
		a.apply(node, n.Name, replaceColNameName)
		a.apply(node, n.Qualifier, replaceColNameQualifier)

	case *CollateExpr:
		a.apply(node, n.Expr, replaceCollateExprExpr)

	case *ColumnDefinition:
		a.apply(node, n.Name, replaceColumnDefinitionName)

	case *ColumnType:
		a.apply(node, n.Autoincrement, replaceColumnTypeAutoincrement)
		a.apply(node, n.Collate, replaceColumnTypeCollate)
		a.apply(node, n.Comment, replaceColumnTypeComment)
		a.apply(node, n.Default, replaceColumnTypeDefault)
		a.apply(node, n.Length, replaceColumnTypeLength)
		a.apply(node, n.NotNull, replaceColumnTypeNotNull)
		a.apply(node, n.Scale, replaceColumnTypeScale)
		a.apply(node, n.Unsigned, replaceColumnTypeUnsigned)
		a.apply(node, n.Zerofill, replaceColumnTypeZerofill)

	case Columns:
		replacer := replaceColumnsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case Comments:

	case *ComparisonExpr:
		a.apply(node, n.Escape, replaceComparisonExprEscape)
		a.apply(node, n.Left, replaceComparisonExprLeft)
		a.apply(node, n.Right, replaceComparisonExprRight)

	case *ConvertExpr:
		a.apply(node, n.Expr, replaceConvertExprExpr)
		a.apply(node, n.Type, replaceConvertExprType)

	case *ConvertType:
		a.apply(node, n.Length, replaceConvertTypeLength)
		a.apply(node, n.Scale, replaceConvertTypeScale)

	case *ConvertUsingExpr:
		a.apply(node, n.Expr, replaceConvertUsingExprExpr)

	case *DDL:
		a.apply(node, n.Database, replaceDDLDatabase)
		a.apply(node, n.DatabaseOptions, replaceDDLDatabaseOptions)
		a.apply(node, n.IndexOpts, replaceDDLIndexOpts)
		a.apply(node, n.ModifyColumnDef, replaceDDLModifyColumnDef)
		a.apply(node, n.NewName, replaceDDLNewName)
		a.apply(node, n.Table, replaceDDLTable)
		a.apply(node, n.TableSpec, replaceDDLTableSpec)
		a.apply(node, n.Tables, replaceDDLTables)
		a.apply(node, n.indexLockAndAlgorithm, replaceDDLindexLockAndAlgorithm)

	case DatabaseOptionListOpt:

	case *Default:

	case *Delete:
		a.apply(node, n.Comments, replaceDeleteComments)
		a.apply(node, n.Limit, replaceDeleteLimit)
		a.apply(node, n.OrderBy, replaceDeleteOrderBy)
		a.apply(node, n.Table, replaceDeleteTable)
		a.apply(node, n.Where, replaceDeleteWhere)

	case *ExistsExpr:
		a.apply(node, n.Subquery, replaceExistsExprSubquery)

	case *Explain:
		a.apply(node, n.Statement, replaceExplainStatement)

	case Exprs:
		replacer := replaceExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *FuncExpr:
		a.apply(node, n.Exprs, replaceFuncExprExprs)
		a.apply(node, n.Name, replaceFuncExprName)
		a.apply(node, n.Qualifier, replaceFuncExprQualifier)

	case GroupBy:
		replacer := replaceGroupByItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *GroupConcatExpr:
		a.apply(node, n.Exprs, replaceGroupConcatExprExprs)
		a.apply(node, n.OrderBy, replaceGroupConcatExprOrderBy)

	case *IndexDefinition:
		a.apply(node, n.Name, replaceIndexDefinitionName)
		a.apply(node, n.Opts, replaceIndexDefinitionOpts)

	case *IndexHints:
		replacerIndexes := replaceIndexHintsIndexes(0)
		replacerIndexesB := &replacerIndexes
		for _, item := range n.Indexes {
			a.apply(node, item, replacerIndexesB.replace)
			replacerIndexesB.inc()
		}

	case *IndexLockAndAlgorithm:

	case *IndexOptions:
		a.apply(node, n.BlockSize, replaceIndexOptionsBlockSize)

	case *Insert:
		a.apply(node, n.Columns, replaceInsertColumns)
		a.apply(node, n.Comments, replaceInsertComments)
		a.apply(node, n.OnDup, replaceInsertOnDup)
		a.apply(node, n.Rows, replaceInsertRows)
		a.apply(node, n.Table, replaceInsertTable)

	case *IntervalExpr:
		a.apply(node, n.Expr, replaceIntervalExprExpr)

	case *IsExpr:
		a.apply(node, n.Expr, replaceIsExprExpr)

	case *JoinTableExpr:
		a.apply(node, n.LeftExpr, replaceJoinTableExprLeftExpr)
		a.apply(node, n.On, replaceJoinTableExprOn)
		a.apply(node, n.RightExpr, replaceJoinTableExprRightExpr)

	case *Kill:

	case *Limit:
		a.apply(node, n.Offset, replaceLimitOffset)
		a.apply(node, n.Rowcount, replaceLimitRowcount)

	case ListArg:

	case *MatchExpr:
		a.apply(node, n.Columns, replaceMatchExprColumns)
		a.apply(node, n.Expr, replaceMatchExprExpr)

	case Nextval:
		a.apply(node, n.Expr, replaceNextvalExpr)

	case *NotExpr:
		a.apply(node, n.Expr, replaceNotExprExpr)

	case *NullVal:

	case OnDup:
		replacer := replaceOnDupItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *OptVal:
		a.apply(node, n.Value, replaceOptValValue)

	case *OrExpr:
		a.apply(node, n.Left, replaceOrExprLeft)
		a.apply(node, n.Right, replaceOrExprRight)

	case *Order:
		a.apply(node, n.Expr, replaceOrderExpr)

	case OrderBy:
		replacer := replaceOrderByItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *OtherAdmin:

	case *OtherRead:

	case *ParenExpr:
		a.apply(node, n.Expr, replaceParenExprExpr)

	case *ParenSelect:
		a.apply(node, n.Select, replaceParenSelectSelect)

	case *ParenTableExpr:
		a.apply(node, n.Exprs, replaceParenTableExprExprs)

	case *Radon:
		a.apply(node, n.NewName, replaceRadonNewName)
		a.apply(node, n.Row, replaceRadonRow)
		a.apply(node, n.Table, replaceRadonTable)

	case *RangeCond:
		a.apply(node, n.From, replaceRangeCondFrom)
		a.apply(node, n.Left, replaceRangeCondLeft)
		a.apply(node, n.To, replaceRangeCondTo)

	case *SQLVal:

	case *Select:
		a.apply(node, n.Comments, replaceSelectComments)
		a.apply(node, n.From, replaceSelectFrom)
		a.apply(node, n.GroupBy, replaceSelectGroupBy)
		a.apply(node, n.Having, replaceSelectHaving)
		a.apply(node, n.Limit, replaceSelectLimit)
		a.apply(node, n.OrderBy, replaceSelectOrderBy)
		a.apply(node, n.SelectExprs, replaceSelectSelectExprs)
		a.apply(node, n.Where, replaceSelectWhere)

	case SelectExprs:
		replacer := replaceSelectExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *Set:
		a.apply(node, n.Comments, replaceSetComments)
		a.apply(node, n.Exprs, replaceSetExprs)

	case *SetExpr:
		a.apply(node, n.Type, replaceSetExprType)
		a.apply(node, n.Val, replaceSetExprVal)

	case SetExprs:
		replacer := replaceSetExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *Show:
		a.apply(node, n.Filter, replaceShowFilter)
		a.apply(node, n.Limit, replaceShowLimit)
		a.apply(node, n.Table, replaceShowTable)

	case *ShowFilter:
		a.apply(node, n.Filter, replaceShowFilterFilter)

	case *StarExpr:
		a.apply(node, n.TableName, replaceStarExprTableName)

	case *Subquery:
		a.apply(node, n.Select, replaceSubquerySelect)

	case TableExprs:
		replacer := replaceTableExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case TableIdent:

	case TableName:
		a.apply(node, n.Name, replaceTableNameName)
		a.apply(node, n.Qualifier, replaceTableNameQualifier)

	case TableNames:
		replacer := replaceTableNamesItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case TableOptions:

	case *TableSpec:
		replacerColumns := replaceTableSpecColumns(0)
		replacerColumnsB := &replacerColumns
		for _, item := range n.Columns {
			a.apply(node, item, replacerColumnsB.replace)
			replacerColumnsB.inc()
		}
		replacerIndexes := replaceTableSpecIndexes(0)
		replacerIndexesB := &replacerIndexes
		for _, item := range n.Indexes {
			a.apply(node, item, replacerIndexesB.replace)
			replacerIndexesB.inc()
		}
		a.apply(node, n.Options, replaceTableSpecOptions)

	case *Transaction:

	case *TxnVal:

	case *UnaryExpr:
		a.apply(node, n.Expr, replaceUnaryExprExpr)

	case *Union:
		a.apply(node, n.Left, replaceUnionLeft)
		a.apply(node, n.Limit, replaceUnionLimit)
		a.apply(node, n.OrderBy, replaceUnionOrderBy)
		a.apply(node, n.Right, replaceUnionRight)

	case *Update:
		a.apply(node, n.Comments, replaceUpdateComments)
		a.apply(node, n.Exprs, replaceUpdateExprs)
		a.apply(node, n.Limit, replaceUpdateLimit)
		a.apply(node, n.OrderBy, replaceUpdateOrderBy)
		a.apply(node, n.Table, replaceUpdateTable)
		a.apply(node, n.Where, replaceUpdateWhere)

	case *UpdateExpr:
		a.apply(node, n.Expr, replaceUpdateExprExpr)
		a.apply(node, n.Name, replaceUpdateExprName)

	case UpdateExprs:
		replacer := replaceUpdateExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *Use:
		a.apply(node, n.DBName, replaceUseDBName)

	case ValTuple:
		replacer := replaceValTupleItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case Values:
		replacer := replaceValuesItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *ValuesFuncExpr:
		a.apply(node, n.Name, replaceValuesFuncExprName)
		a.apply(node, n.Resolved, replaceValuesFuncExprResolved)

	case *When:
		a.apply(node, n.Cond, replaceWhenCond)
		a.apply(node, n.Val, replaceWhenVal)

	case *Where:
		a.apply(node, n.Expr, replaceWhereExpr)

	case *Xa:

	default:
		panic("unknown ast type " + reflect.TypeOf(node).String())
	}

	if a.post != nil && !a.post(&a.cursor) {
		panic(abort)
	}

	a.cursor = saved
}

func isNilValue(i interface{}) bool {
	valueOf := reflect.ValueOf(i)
	kind := valueOf.Kind()
	isNullable := kind == reflect.Ptr || kind == reflect.Array || kind == reflect.Slice
	return isNullable && valueOf.IsNil()
}
