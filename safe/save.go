package safe

import (
	"encoding/json"
)

func (s *Safe) save() error {
	plaintext, err := json.Marshal(s)
	if err != nil {
		return err
	}

	ciphertext, err := s.crypto.Encrypt(plaintext, s.password)
	if err != nil {
		return err
	}

	dto := safeDTO{
		Config:     &s.Config.Encryption,
		Ciphertext: ciphertext,
	}

	data, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	// Before saving the new safe, do an auto-backup if enabled
	if s.Config.Storage.Backup.AutoEnabled {
		Backup(s.backend)
	}

	return s.backend.Save(data)
}
