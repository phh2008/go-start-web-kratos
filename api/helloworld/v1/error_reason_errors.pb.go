// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsNoLogin(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NoLogin.String() && e.Code == 401
}

func ErrorNoLogin(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_NoLogin.String(), fmt.Sprintf(format, args...))
}

func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Unauthorized.String() && e.Code == 402
}

func ErrorUnauthorized(format string, args ...interface{}) *errors.Error {
	return errors.New(402, ErrorReason_Unauthorized.String(), fmt.Sprintf(format, args...))
}

func IsSignVerifyError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SignVerifyError.String() && e.Code == 403
}

func ErrorSignVerifyError(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_SignVerifyError.String(), fmt.Sprintf(format, args...))
}

func IsSysError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SysError.String() && e.Code == 500
}

func ErrorSysError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_SysError.String(), fmt.Sprintf(format, args...))
}

func IsParamError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ParamError.String() && e.Code == 501
}

func ErrorParamError(format string, args ...interface{}) *errors.Error {
	return errors.New(501, ErrorReason_ParamError.String(), fmt.Sprintf(format, args...))
}

func IsDBError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DBError.String() && e.Code == 502
}

func ErrorDBError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, ErrorReason_DBError.String(), fmt.Sprintf(format, args...))
}