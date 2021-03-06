package pttbbs

import (
	"strings"
	"testing"
)

func TestReadUserFavoriteRecordsFileNoFile(t *testing.T) {
	c := Connector{}
	_, err := c.ReadUserFavoriteRecordsFile("test")

	if !strings.Contains(err.Error(), "OpenFavFile") {
		t.Errorf("err not return OpenFavFile")
	}
}
