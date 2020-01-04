package models

import (
  "fmt"
)

type ApiError struct {
  Err string
  Code int
  Metadata string
}

func (ae ApiError) Error() string{
  return fmt.Sprintf(ae.Err)
}
