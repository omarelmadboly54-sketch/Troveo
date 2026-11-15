package postgres

import (
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(product *models.Product) error{
	return r.db.Create(product).Error
}


func (r *productRepository)GetAll(page,limit int)([]models.Product,error){
	var products []models.Product
	offset:=(page-1)*limit
	result:=r.db.Limit(limit).Offset(offset).Find(&products)
	if result.Error!=nil{
		return []models.Product{},nil
	}

	return products,nil
}


func (r *productRepository)GetProductByID(id int64)(*models.Product,error){
	var product models.Product

	result:=r.db.Where("id=?",id).First(&product)
	if result.Error!=nil{
		return nil,result.Error
	}

	return &product,nil
}

func (r *productRepository)GetMyProducts(UserID int64,page,limit int)([]models.Product,error){
	var products []models.Product

	offset:=(page-1)*limit

	result:=r.db.Where("seller_id=?",UserID).Limit(limit).Offset(offset).Find(&products)
	if result.Error!=nil{
		return []models.Product{},result.Error
	}

	return products,nil
}


func (r *productRepository)GetMyProduct(id,userID int64)(*models.Product,error){
	var product models.Product

	result:=r.db.Where("id=? AND seller_id=?",id,userID).First(&product)
	if result.Error!=nil{
		return nil,result.Error
	}

	return &product,nil
}


func (r *productRepository)DeleteProductByID(id,UserID int64)error{
	
	result:=r.db.Where("id=? AND seller_id=?",id,UserID).Delete(&models.Product{})
	if result.Error!=nil{
		return result.Error
	}

	if result.RowsAffected==0{
		return gorm.ErrRecordNotFound
	}

	return  nil
}

func (r *productRepository)GetProductsByCategory(CategoryID int64)([]models.Product,error){

	var products []models.Product

	err:=r.db.Raw(`
	WITH RECURSIVE category_tree AS(
	SELECT id FROM categories WHERE id=?

	UNION ALL 

	SELECT c.id FROM categories c
	JOIN category_tree ct ON c.parent_id=ct.id 
	)
	SELECT * FROM products 
	WHERE category_id IN(SELECT id FROM category_tree)
	`,CategoryID).Scan(&products).Error

	return products,err
}

