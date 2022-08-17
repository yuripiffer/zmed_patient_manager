package patient

import (
	"strings"
	"zmed_patient_manager/pkg/app_errors"
	"zmed_patient_manager/pkg/email/blocklist"
)

func (s *service) validateEmail(email string) bool {
	emailSplit := strings.Split(email, "@")
	if emailSplit[0] == "" || emailSplit[1] == "" {
		return false
	}
	if blocklist.IsDomainInBlockList(emailSplit[1]) {
		return false
	}
	return true
}

func (s *service) encryptEmail(email string) (string, app_errors.AppError) {
	emailCipher, appErr := s.crypto.Encrypt(email)
	if appErr != nil {
		return "", appErr
	}
	return emailCipher, nil
}

func (s *service) decryptEmail(cipher string) (string, app_errors.AppError) {
	email, appErr := s.crypto.Decrypt(cipher)
	if appErr != nil {
		return "", appErr
	}
	return email, nil
}
