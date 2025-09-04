package utils

import (
	"errors"
	"regexp"

	"github.com/cristian-yw/Minitask_W10D2/internal/models"
)

func ValidatePost(body models.Body) error {
	if body.Id <= 0 {
		return errors.New("id tidak boleh kosong")
	}
	if len(body.Massage) <= 8 {
		return errors.New("message tidak boleh kurang dari 8 karakter")
	}
	re, err := regexp.Compile(`^[lLpPmMfF]$`)
	if err != nil {
		return err
	}
	if !re.MatchString(body.Gender) {
		return errors.New("gender tidak valid")
	}
	ne, err := regexp.Compile(`^[a-zA-Z\s]+$`)
	if err != nil {
		return err
	}
	if !ne.MatchString(body.Name) {
		return errors.New("Nama tidak boleh ada angka")
	}
	return nil
}
