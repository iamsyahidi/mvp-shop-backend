package utils

import (
	"fmt"
	"time"
)

func GenerateInvoice() string {
	times := time.Now()

	epoch := times.UnixMilli()
	id := "INV/" + fmt.Sprintf("%v", epoch)

	return id
}
