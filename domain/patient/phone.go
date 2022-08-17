package patient

import (
	"strconv"
	"zmed_patient_manager/pkg/app_errors"
)

func (s *service) validatePhone(cellphone string) bool {
	_, err := strconv.Atoi(cellphone)
	if err != nil {
		return false
	}
	if len(cellphone) >= 11 {
		if cellphone[len(cellphone)-9:len(cellphone)-8] != "9" {
			return false
		}
		return true
	}
	return false
}

func (s *service) encryptPhone(cellphone string) (string, app_errors.AppError) {
	cellphoneCipher, appErr := s.crypto.Encrypt(cellphone)
	if appErr != nil {
		return "", appErr
	}
	return cellphoneCipher, nil
}

func (s *service) decryptPhone(cipher string) (string, app_errors.AppError) {
	cellphone, appErr := s.crypto.Decrypt(cipher)
	if appErr != nil {
		return "", appErr
	}
	return cellphone, nil
}
