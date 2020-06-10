package service

import (
	"fmt"
	"testing"
)

func TestSelectNameReviewer(t *testing.T) {

			gotQu, _ := SelectNameReviewer(6299)
			fmt.Println(gotQu)

}