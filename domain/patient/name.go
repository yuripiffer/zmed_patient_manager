package patient

import (
	"strings"
	"zmed_patient_manager/pkg/app_errors"
)

func (s *service) validateName(name string) bool {
	splitName := strings.Split(name, " ")
	if !(len(splitName) >= 2) {
		return false
	}
	if strings.Contains(name, "!") ||
		strings.Contains(name, "@") ||
		strings.Contains(name, "$") ||
		strings.Contains(name, "%") ||
		strings.Contains(name, "*") {
		return false
	}
	return true
}

func (s *service) encryptName(name string) (string, app_errors.AppError) {
	nameCipher, appErr := s.crypto.Encrypt(name)
	if appErr != nil {
		return "", appErr
	}
	return nameCipher, nil
}

func (s *service) decryptName(cipher string) (string, app_errors.AppError) {
	name, appErr := s.crypto.Decrypt(cipher)
	if appErr != nil {
		return "", appErr
	}
	return name, nil
}
