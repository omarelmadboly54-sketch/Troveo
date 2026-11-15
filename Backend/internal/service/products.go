package service

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"gorm.io/gorm"
)

type productService struct {
	productRepository repository.ProductRepository
	cld               *cloudinary.Cloudinary
}

type ProductService interface {
	CreateProduct(product *models.Product, fileHeader *multipart.FileHeader, UserID int64) error
	GetAll(page,limit int)([]models.Product,error)
	GetProductByID(id int64)(*models.Product,error)
	GetMyProducts(UserID int64,page,limit int)([]models.Product,error)
	GetMyProduct(id,userID int64)(*models.Product,error)
	DeleteProductByID(id,UserID int64)error
	GetProductsByCategory(CategoryID int64)([]models.Product,error)
}

func NewProductService(productRepository repository.ProductRepository,cld *cloudinary.Cloudinary) ProductService {
	return &productService{
		productRepository: productRepository,
		cld: cld,
	}
}

func (s *productService) CreateProduct(product *models.Product, fileHeader *multipart.FileHeader, UserID int64) error {

	var maxfilesize int64 = 5 * 1024 * 1024

	if fileHeader == nil || fileHeader.Size == 0 {
		return apperrors.ErrInvalidSize
	}

	if fileHeader.Size > maxfilesize {
		return apperrors.ErrInvalidSize
	}

	file, err := fileHeader.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	buffer := make([]byte, 512)

	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}

	contentType := http.DetectContentType(buffer[:n])

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":true,
		"image/png":  true,
		"image/webp": true,
	}

	if !allowedTypes[contentType] {
		return apperrors.ErrInvalidImageType
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	ctx := context.Background()
	uploadResult, err := s.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "Trovero/products",
	})
	if err != nil {
		return err
	}

	product.Photo = uploadResult.SecureURL
	product.PhotoPublicID=uploadResult.PublicID
	product.SellerID = UserID

	err = s.productRepository.CreateProduct(product)
	if err != nil {
		_, deleteErr := s.cld.Upload.Destroy(
			ctx,
			uploader.DestroyParams{
				PublicID: uploadResult.PublicID,
			},
		)

		if deleteErr != nil {
			log.Printf("Failed to delete uploaded Cloudinary image %s: %v",
				uploadResult.PublicID, deleteErr)

		}
		return err
	}

	return nil
}


func (s *productService)GetAll(page,limit int)([]models.Product,error){
	return s.productRepository.GetAll(page,limit)
}


func (s *productService)GetProductByID(id int64)(*models.Product,error){
	product,err:=s.productRepository.GetProductByID(id)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,apperrors.ErrProductNotFound
		}
		return nil,err
	}

	return product,nil
}

func (s *productService)GetMyProducts(UserID int64,page,limit int)([]models.Product,error){
	return s.productRepository.GetMyProducts(UserID,page,limit)
}


func (s *productService)GetMyProduct(id,userID int64)(*models.Product,error){
	product,err:=s.productRepository.GetMyProduct(id,userID)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,apperrors.ErrProductNotFound
		}
		return nil,err
	}

	return product,nil
}

func (s *productService)DeleteProductByID(id,UserID int64)error{
	product,err:=s.productRepository.GetMyProduct(id,UserID)
	if err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return apperrors.ErrProductNotFound
		}
		return err
	}

	ctx:=context.Background()
	_,err=s.cld.Upload.Destroy(ctx,uploader.DestroyParams{
		PublicID: product.PhotoPublicID,
	})
	if err!=nil{
		return err
	}

	err=s.productRepository.DeleteProductByID(id,UserID)
	if err!=nil{
		return err
	}
	return nil

}

func (s *productService)GetProductsByCategory(CategoryID int64)([]models.Product,error){
	return s.productRepository.GetProductsByCategory(CategoryID)
}
