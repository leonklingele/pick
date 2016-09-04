package safe

import (
	"fmt"
)

func (s *Safe) Remove(name string) error {
	if _, exists := s.Accounts[name]; !exists {
		return fmt.Errorf("Account not found")
	}

	delete(s.Accounts, name)

	return s.save()
}
