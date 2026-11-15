package apperrors

import "errors"

var (
	//User
	ErrUserAlreadyExists = errors.New("User already exists")
	ErrInvalidCredentials=errors.New("Invalid credentials")
	ErrUserNotFound=errors.New("User not found")
	ErrUserNotAllowed=errors.New("User not allowed")
	//Address
	ErrAddressNotFound=errors.New("Address not found")
	//Category
	ErrCategoryNotFound=errors.New("Category not found")
	//Product
	ErrInvalidSize=errors.New("file size exceeds the 5MB limit")
	ErrInvalidImageType=errors.New("invalid file format. Only JPEG, PNG, and WebP are allowed")
	ErrProductNotFound=errors.New("Product not found")
	//Coupon:
	ErrCouponCodeExists        = errors.New("coupon code already exists")
    ErrInvalidCouponValue      = errors.New("percentage discount cannot exceed 100")
    ErrInvalidCouponExpiration = errors.New("coupon expiration must be in the future")
)

