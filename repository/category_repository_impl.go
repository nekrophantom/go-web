package repository

import (
	"context"
	"database/sql"
	"errors"
	"restapi/helper"
	"restapi/model/domain"
)

type CategoryRepositoryImpl struct {

}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category{
	SQL := "insert into customer(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category{
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

	return category
}


func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category){
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error){
	SQL := "select id, name from category where id = ?"
	result, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category{
	SQL := "select id, name from category"
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categories []domain.Category
	for result.Next(){
		category := domain.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}