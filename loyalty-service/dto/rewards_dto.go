package dto

import "errors"

// REWARDs Failed Messages
const (
	MESSAGE_FAILED_CREATE_REWARD = "failed to create reward"
	MESSAGE_FAILED_UPDATE_REWARD = "failed to update reward"
	MESSAGE_FAILED_DELETE_REWARD = "failed to delete reward"
	MESSAGE_FAILED_GET_REWARD    = "failed to get reward"
)

// REWARDs Success Messages
const (
	MESSAGE_SUCCESS_CREATE_REWARD = "success create reward"
	MESSAGE_SUCCESS_UPDATE_REWARD = "success update reward"
	MESSAGE_SUCCESS_DELETE_REWARD = "success delete reward"
	MESSAGE_SUCCESS_GET_REWARD    = "success get reward"
)

// REWARDs Custom Errors
var (
	ErrCreateReward = errors.New(MESSAGE_FAILED_CREATE_REWARD)
	ErrUpdateReward = errors.New(MESSAGE_FAILED_UPDATE_REWARD)
	ErrDeleteReward = errors.New(MESSAGE_FAILED_DELETE_REWARD)
	ErrGetReward    = errors.New(MESSAGE_FAILED_GET_REWARD)
)
