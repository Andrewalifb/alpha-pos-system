package dto

import "errors"

// LOYALTY_LEVEL Failed Messages
const (
	MESSAGE_FAILED_CREATE_LOYALTY_LEVEL = "failed to create loyalty level"
	MESSAGE_FAILED_UPDATE_LOYALTY_LEVEL = "failed to update loyalty level"
	MESSAGE_FAILED_DELETE_LOYALTY_LEVEL = "failed to delete loyalty level"
	MESSAGE_FAILED_GET_LOYALTY_LEVEL    = "failed to get loyalty level"
)

// LOYALTY_LEVEL Success Messages
const (
	MESSAGE_SUCCESS_CREATE_LOYALTY_LEVEL = "success create loyalty level"
	MESSAGE_SUCCESS_UPDATE_LOYALTY_LEVEL = "success update loyalty level"
	MESSAGE_SUCCESS_DELETE_LOYALTY_LEVEL = "success delete loyalty level"
	MESSAGE_SUCCESS_GET_LOYALTY_LEVEL    = "success get loyalty level"
)

// LOYALTY_LEVEL Custom Errors
var (
	ErrCreateLoyaltyLevel = errors.New(MESSAGE_FAILED_CREATE_LOYALTY_LEVEL)
	ErrUpdateLoyaltyLevel = errors.New(MESSAGE_FAILED_UPDATE_LOYALTY_LEVEL)
	ErrDeleteLoyaltyLevel = errors.New(MESSAGE_FAILED_DELETE_LOYALTY_LEVEL)
	ErrGetLoyaltyLevel    = errors.New(MESSAGE_FAILED_GET_LOYALTY_LEVEL)
)
