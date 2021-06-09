package db

import (
	"xorm.io/builder"
)

const maxInsertRows = 1000

var batchCountCalcFunc = func(total, maxCount int) int {
	if total <= maxCount {
		return 1
	}
	if total%maxCount > 0 {
		return total/maxCount + 1
	}
	return total / maxCount
}

var model = &Model{}

type Model struct{}

// -------------------------------------------------------------------------
// ------------------------------WriteDatabase------------------------------
// -------------------------------------------------------------------------

// InsertOne .
func InsertOne(modelPtr Interface, cols ...string) (int64, error) {
	return model.InsertOne(modelPtr, cols...)
}

// InsertOne .
func (*Model) InsertOne(modelPtr Interface, cols ...string) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).InsertOne(modelPtr)
}

// Insert .
func Insert(modelPtr Interface, modelsPtr interface{}, cols ...string) (int64, error) {
	return model.Insert(modelPtr, modelsPtr, cols...)
}

// Insert .
func (*Model) Insert(modelPtr Interface, modelsPtr interface{}, cols ...string) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Insert(modelsPtr)
}

// InsertBatch .
func InsertBatch(modelPtr Interface, modelsPtr ...interface{}) (int64, error) {
	return model.InsertBatch(modelPtr, modelsPtr...)
}

// InsertBatch .
func (*Model) InsertBatch(modelPtr Interface, models ...interface{}) (int64, error) {
	count := len(models)
	if count == 0 {
		return 0, nil
	}

	batchCount := batchCountCalcFunc(count, maxInsertRows)

	var rows int64
	var records []interface{}
	for i := 0; i < batchCount; i++ {
		if i == batchCount-1 {
			records = models[i*maxInsertRows:]
		} else {
			records = models[i*maxInsertRows : (i+1)*maxInsertRows]
		}

		affected, err := modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Insert(records)
		if err != nil {
			return 0, err
		}

		rows += affected
	}

	return rows, nil
}

// Update .
func Update(modelPtr Interface, id interface{}, cols ...string) (int64, error) {
	return model.Update(modelPtr, id, cols...)
}

// Update .
func (*Model) Update(modelPtr Interface, id interface{}, cols ...string) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).ID(id).Update(modelPtr)
}

// UpdateCond .
func UpdateCond(modelPtr Interface, cond interface{}, cols ...string) (int64, error) {
	return model.UpdateCond(modelPtr, cond, cols...)
}

// UpdateCond .
func (*Model) UpdateCond(modelPtr Interface, cond interface{}, cols ...string) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Where(cond).NoAutoCondition(true).Update(modelPtr)
}

// Delete .
func Delete(modelPtr Interface) (int64, error) {
	return model.Delete(modelPtr)
}

// Delete .
func (*Model) Delete(modelPtr Interface) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Delete(modelPtr)
}

// DeleteByID .
func DeleteByID(modelPtr Interface, id interface{}) (int64, error) {
	return model.DeleteByID(modelPtr, id)
}

// DeleteByID .
func (*Model) DeleteByID(modelPtr Interface, id interface{}) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).ID(id).Delete(modelPtr)
}

// DeleteByCond .
func DeleteByCond(modelPtr Interface, cond interface{}) (int64, error) {
	return model.DeleteByCond(modelPtr, cond)
}

// DeleteByCond .
func (*Model) DeleteByCond(modelPtr Interface, cond interface{}) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Where(cond).NoAutoCondition(true).Delete(modelPtr)
}

// -------------------------------------------------------------------------
// ------------------------------ReadDatabase-------------------------------
// -------------------------------------------------------------------------

// SelectOne .
func SelectOne(modelPtr Interface, cols ...string) (bool, error) {
	return model.SelectOne(modelPtr, cols...)
}

// SelectOne .
func (*Model) SelectOne(modelPtr Interface, cols ...string) (bool, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Cols(cols...).Get(modelPtr)
}

// SelectOne2 .
func SelectOne2(modelPtr Interface, cols ...string) (bool, error) {
	return model.SelectOne2(modelPtr, cols...)
}

// SelectOne2 .
func (*Model) SelectOne2(modelPtr Interface, cols ...string) (bool, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Get(modelPtr)
}

// SelectOneCond .
func SelectOneCond(modelPtr Interface, cond builder.Cond, cols ...string) (bool, error) {
	return model.SelectOneCond(modelPtr, cond, cols...)
}

// SelectOneCond .
func (*Model) SelectOneCond(modelPtr Interface, cond builder.Cond, cols ...string) (bool, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Cols(cols...).Where(cond).NoAutoCondition(true).Get(modelPtr)
}

// SelectOneCond2 .
func SelectOneCond2(modelPtr Interface, cond builder.Cond, cols ...string) (bool, error) {
	return model.SelectOneCond2(modelPtr, cond, cols...)
}

// SelectOneCond2 .
func (*Model) SelectOneCond2(modelPtr Interface, cond builder.Cond, cols ...string) (bool, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Where(cond).NoAutoCondition(true).Get(modelPtr)
}

// Select .
func Select(modelPtr Interface, modelsPtr interface{}, paging Paging, cols ...string) error {
	return model.Select(modelPtr, modelsPtr, paging, cols...)
}

// Select .
func (*Model) Select(modelPtr Interface, modelsPtr interface{}, paging Paging, cols ...string) error {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Cols(cols...).Limit(paging.LimitOffset()).Find(modelsPtr, modelPtr)
}

// Select2 .
func Select2(modelPtr Interface, modelsPtr interface{}, paging Paging, cols ...string) error {
	return model.Select2(modelPtr, modelsPtr, paging, cols...)
}

// Select2 .
func (*Model) Select2(modelPtr Interface, modelsPtr interface{}, paging Paging, cols ...string) error {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Limit(paging.LimitOffset()).Find(modelsPtr, modelPtr)
}

// SelectAll .
func SelectAll(modelPtr Interface, modelsPtr interface{}, cols ...string) error {
	return model.SelectAll(modelPtr, modelsPtr, cols...)
}

// SelectAll .
func (*Model) SelectAll(modelPtr Interface, modelsPtr interface{}, cols ...string) error {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Cols(cols...).Find(modelsPtr, modelPtr)
}

// SelectAll2 .
func SelectAll2(modelPtr Interface, modelsPtr interface{}, cols ...string) error {
	return model.SelectAll2(modelPtr, modelsPtr, cols...)
}

// SelectAll2 .
func (*Model) SelectAll2(modelPtr Interface, modelsPtr interface{}, cols ...string) error {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Find(modelsPtr, modelPtr)
}

// SelectAllCond .
func SelectAllCond(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, cols ...string) error {
	return model.SelectAllCond(modelPtr, modelsPtr, cond, cols...)
}

// SelectAllCond .
func (*Model) SelectAllCond(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, cols ...string) error {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Cols(cols...).Where(cond).Find(modelsPtr)
}

// SelectAllCond2 .
func SelectAllCond2(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, cols ...string) error {
	return model.SelectAllCond2(modelPtr, modelsPtr, cond, cols...)
}

// SelectAllCond2 .
func (*Model) SelectAllCond2(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, cols ...string) error {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Where(cond).Find(modelsPtr)
}

// SelectSql .
func SelectSql(modelPtr Interface, modelsPtr []interface{}, sql string, args ...interface{}) error {
	return model.SelectSql(modelPtr, modelsPtr, sql, args...)
}

// SelectSql .
func (*Model) SelectSql(modelPtr Interface, modelsPtr interface{}, sql string, args ...interface{}) error {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).SQL(sql, args...).Find(modelsPtr)
}

// SelectSql2 .
func SelectSql2(modelPtr Interface, modelsPtr []interface{}, sql string, args ...interface{}) error {
	return model.SelectSql2(modelPtr, modelsPtr, sql, args...)
}

// SelectSql2 .
func (*Model) SelectSql2(modelPtr Interface, modelsPtr interface{}, sql string, args ...interface{}) error {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).SQL(sql, args...).Find(modelsPtr)
}

// SelectCond .
func SelectCond(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, orderBy string, paging Paging, cols ...string) error {
	return model.SelectCond(modelPtr, modelsPtr, cond, orderBy, paging, cols...)
}

// SelectCond .
func (*Model) SelectCond(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, orderBy string, paging Paging, cols ...string) error {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Cols(cols...).Where(cond).OrderBy(orderBy).Limit(paging.LimitOffset()).Find(modelsPtr)
}

// SelectCond2 .
func SelectCond2(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, orderBy string, paging Paging, cols ...string) error {
	return model.SelectCond(modelPtr, modelsPtr, cond, orderBy, paging, cols...)
}

// SelectCond2 .
func (*Model) SelectCond2(modelPtr Interface, modelsPtr interface{}, cond builder.Cond, orderBy string, paging Paging, cols ...string) error {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(cols...).Where(cond).OrderBy(orderBy).Limit(paging.LimitOffset()).Find(modelsPtr)
}

// Count .
func Count(modelPtr Interface) (int64, error) {
	return model.Count(modelPtr)
}

// Count .
func (*Model) Count(modelPtr Interface) (int64, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Count(modelPtr)
}

// Count2 .
func Count2(modelPtr Interface) (int64, error) {
	return model.Count2(modelPtr)
}

// Count2 .
func (*Model) Count2(modelPtr Interface) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Count(modelPtr)
}

// CountCond .
func CountCond(modelPtr Interface, cond builder.Cond) (int64, error) {
	return model.CountCond(modelPtr, cond)
}

// CountCond .
func (*Model) CountCond(modelPtr Interface, cond builder.Cond) (int64, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Where(cond).Count()
}

// CountCond2 .
func CountCond2(modelPtr Interface, cond builder.Cond) (int64, error) {
	return model.CountCond2(modelPtr, cond)
}

// CountCond2 .
func (*Model) CountCond2(modelPtr Interface, cond builder.Cond) (int64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Where(cond).Count()
}

// Sum .
func Sum(modelPtr Interface, col string) (float64, error) {
	return model.Sum(modelPtr, col)
}

// Sum .
func (*Model) Sum(modelPtr Interface, col string) (float64, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Sum(modelPtr, col)
}

// Sum2 .
func Sum2(modelPtr Interface, col string) (float64, error) {
	return model.Sum2(modelPtr, col)
}

// Sum2 .
func (*Model) Sum2(modelPtr Interface, col string) (float64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Sum(modelPtr, col)
}

// SumCond .
func SumCond(modelPtr Interface, col string, cond builder.Cond) (float64, error) {
	return model.SumCond(modelPtr, col, cond)
}

// SumCond .
func (*Model) SumCond(modelPtr Interface, col string, cond builder.Cond) (float64, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Where(cond).NoAutoCondition(true).Sum(modelPtr, col)
}

// SumCond2 .
func SumCond2(modelPtr Interface, col string, cond builder.Cond) (float64, error) {
	return model.SumCond2(modelPtr, col, cond)
}

// SumCond2 .
func (*Model) SumCond2(modelPtr Interface, col string, cond builder.Cond) (float64, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Where(cond).NoAutoCondition(true).Sum(modelPtr, col)
}

// IsExists .
func IsExists(modelPtr Interface) (bool, error) {
	return model.IsExists(modelPtr)
}

// IsExists .
func (*Model) IsExists(modelPtr Interface) (bool, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Exist(modelPtr)
}

// IsExists2 .
func IsExists2(modelPtr Interface) (bool, error) {
	return model.IsExists2(modelPtr)
}

// IsExists2 .
func (*Model) IsExists2(modelPtr Interface) (bool, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Exist(modelPtr)
}

// IsExistsCond .
func IsExistsCond(modelPtr Interface, cond interface{}, col ...string) (bool, error) {
	return model.IsExistsCond(modelPtr, cond, col...)
}

// IsExistsCond .
func (*Model) IsExistsCond(modelPtr Interface, cond interface{}, col ...string) (bool, error) {
	return modelPtr.DatabaseReadEngine().Table(modelPtr.TableName()).Cols(col...).Where(cond).Exist()
}

// IsExistsCond2 .
func IsExistsCond2(modelPtr Interface, cond interface{}, col ...string) (bool, error) {
	return model.IsExistsCond2(modelPtr, cond, col...)
}

// IsExistsCond2 .
func (*Model) IsExistsCond2(modelPtr Interface, cond interface{}, col ...string) (bool, error) {
	return modelPtr.DatabaseWriteEngine().Table(modelPtr.TableName()).Cols(col...).Where(cond).Exist()
}

// IsColExist .
func IsColExist(modelPtr Interface, cols ...string) (bool, error) {
	return model.IsColExist(modelPtr, cols...)
}

// IsColExist .
func (*Model) IsColExist(modelPtr Interface, cols ...string) (bool, error) {
	if len(cols) == 0 {
		return true, nil
	}
	info, err := modelPtr.DatabaseReadEngine().TableInfo(modelPtr)
	if err != nil {
		return false, err
	}
	for _, col := range cols {
		if info.GetColumn(col) == nil {
			return false, nil
		}
	}
	return true, nil
}

// IsColExist2 .
func IsColExist2(modelPtr Interface, cols ...string) (bool, error) {
	return model.IsColExist2(modelPtr, cols...)
}

// IsColExist2 .
func (*Model) IsColExist2(modelPtr Interface, cols ...string) (bool, error) {
	if len(cols) == 0 {
		return true, nil
	}
	info, err := modelPtr.DatabaseWriteEngine().TableInfo(modelPtr)
	if err != nil {
		return false, err
	}
	for _, col := range cols {
		if info.GetColumn(col) == nil {
			return false, nil
		}
	}
	return true, nil
}
