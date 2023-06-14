package uuid

import (
	"fmt"
	"github.com/google/uuid"
)

func New() string {
	src := uuid.New()
	return fmt.Sprintf("%x%x%x%x%x", src[0:4], src[4:6], src[6:8], src[8:10], src[10:16])

}
