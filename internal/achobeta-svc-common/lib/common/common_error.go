package common

import (
	errorsv1 "achobeta-svc/internal/achobeta-svc-proto/gen/go/common/errors/v1"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CommonError struct {
	errCode errorsv1.Code
	errMsg  string
}

func (e *CommonError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewCommonError(errCode errorsv1.Code, errMsg string) *CommonError {
	return &CommonError{errCode: errCode, errMsg: errMsg}
}
func NewCommonErrorWithCodeDefaultMsg(errCode errorsv1.Code) *CommonError {
	return &CommonError{errCode: errCode, errMsg: extractMessage(errCode)}
}
func extractMessage(errCode errorsv1.Code) string {
	m := errCode.String()
	if strings.HasPrefix(m, "CODE_") {
		m = strings.TrimLeft(m, "CODE_")
	}
	return cases.Title(language.AmericanEnglish).String(strings.ReplaceAll(strings.ToLower(m), "_", " "))
}
func NewCommonErrorWithStack(errCode errorsv1.Code, errMsg string) error {
	return errors.WithStack(&CommonError{errCode: errCode, errMsg: errMsg})
}
