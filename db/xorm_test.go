package db

import "testing"

func TestXOrm_Delete(t *testing.T) {
	new(XOrm).Delete(3)
}
