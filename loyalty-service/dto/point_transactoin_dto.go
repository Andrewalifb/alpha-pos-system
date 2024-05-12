package dto

import "errors"

// POINT_TRANSACTION Failed Messages
const (
	MESSAGE_FAILED_CREATE_POINT_TRANSACTION = "failed to create point transaction"
	MESSAGE_FAILED_UPDATE_POINT_TRANSACTION = "failed to update point transaction"
	MESSAGE_FAILED_DELETE_POINT_TRANSACTION = "failed to delete point transaction"
	MESSAGE_FAILED_GET_POINT_TRANSACTION    = "failed to get point transaction"
)

// POINT_TRANSACTION Success Messages
const (
	MESSAGE_SUCCESS_CREATE_POINT_TRANSACTION = "success create point transaction"
	MESSAGE_SUCCESS_UPDATE_POINT_TRANSACTION = "success update point transaction"
	MESSAGE_SUCCESS_DELETE_POINT_TRANSACTION = "success delete point transaction"
	MESSAGE_SUCCESS_GET_POINT_TRANSACTION    = "success get point transaction"
)

// POINT_TRANSACTION Custom Errors
var (
	ErrCreatePointTransaction = errors.New(MESSAGE_FAILED_CREATE_POINT_TRANSACTION)
	ErrUpdatePointTransaction = errors.New(MESSAGE_FAILED_UPDATE_POINT_TRANSACTION)
	ErrDeletePointTransaction = errors.New(MESSAGE_FAILED_DELETE_POINT_TRANSACTION)
	ErrGetPointTransaction    = errors.New(MESSAGE_FAILED_GET_POINT_TRANSACTION)
)
