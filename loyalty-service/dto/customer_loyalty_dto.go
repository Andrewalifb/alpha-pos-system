package dto

import "errors"

// CUSTOMER_LOYALTY Failed Messages
const (
	MESSAGE_FAILED_CREATE_CUSTOMER_LOYALTY = "failed to create customer loyalty"
	MESSAGE_FAILED_UPDATE_CUSTOMER_LOYALTY = "failed to update customer loyalty"
	MESSAGE_FAILED_DELETE_CUSTOMER_LOYALTY = "failed to delete customer loyalty"
	MESSAGE_FAILED_GET_CUSTOMER_LOYALTY    = "failed to get customer loyalty"
)

// CUSTOMER_LOYALTY Success Messages
const (
	MESSAGE_SUCCESS_CREATE_CUSTOMER_LOYALTY = "success create customer loyalty"
	MESSAGE_SUCCESS_UPDATE_CUSTOMER_LOYALTY = "success update customer loyalty"
	MESSAGE_SUCCESS_DELETE_CUSTOMER_LOYALTY = "success delete customer loyalty"
	MESSAGE_SUCCESS_GET_CUSTOMER_LOYALTY    = "success get customer loyalty"
)

// CUSTOMER_LOYALTY Custom Errors
var (
	ErrCreateCustomerLoyalty = errors.New(MESSAGE_FAILED_CREATE_CUSTOMER_LOYALTY)
	ErrUpdateCustomerLoyalty = errors.New(MESSAGE_FAILED_UPDATE_CUSTOMER_LOYALTY)
	ErrDeleteCustomerLoyalty = errors.New(MESSAGE_FAILED_DELETE_CUSTOMER_LOYALTY)
	ErrGetCustomerLoyalty    = errors.New(MESSAGE_FAILED_GET_CUSTOMER_LOYALTY)
)
