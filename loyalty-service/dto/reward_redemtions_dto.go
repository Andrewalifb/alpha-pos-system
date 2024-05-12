package dto

import "errors"

// REWARD_REDEMTIONS Failed Messages
const (
	MESSAGE_FAILED_CREATE_REWARD_REDEMTIONS = "failed to create reward redemptions"
	MESSAGE_FAILED_UPDATE_REWARD_REDEMTIONS = "failed to update reward redemptions"
	MESSAGE_FAILED_DELETE_REWARD_REDEMTIONS = "failed to delete reward redemptions"
	MESSAGE_FAILED_GET_REWARD_REDEMTIONS    = "failed to get reward redemptions"
)

// REWARD_REDEMTIONS Success Messages
const (
	MESSAGE_SUCCESS_CREATE_REWARD_REDEMTIONS = "success create reward redemptions"
	MESSAGE_SUCCESS_UPDATE_REWARD_REDEMTIONS = "success update reward redemptions"
	MESSAGE_SUCCESS_DELETE_REWARD_REDEMTIONS = "success delete reward redemptions"
	MESSAGE_SUCCESS_GET_REWARD_REDEMTIONS    = "success get reward redemptions"
)

// REWARD_REDEMTIONS Custom Errors
var (
	ErrCreateRewardRedemtions = errors.New(MESSAGE_FAILED_CREATE_REWARD_REDEMTIONS)
	ErrUpdateRewardRedemtions = errors.New(MESSAGE_FAILED_UPDATE_REWARD_REDEMTIONS)
	ErrDeleteRewardRedemtions = errors.New(MESSAGE_FAILED_DELETE_REWARD_REDEMTIONS)
	ErrGetRewardRedemtions    = errors.New(MESSAGE_FAILED_GET_REWARD_REDEMTIONS)
)
