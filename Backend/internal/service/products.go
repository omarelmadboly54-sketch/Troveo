package service

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/repository"
	"context"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type productService struct {
	productRepository repository.ProductRepository
	cld               *cloudinary.Cloudinary
}

type ProductService interface {
	CreateProduct(product *models.Product, fileHeader *multipart.FileHeader, UserID int64) error
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
