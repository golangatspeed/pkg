package notenough

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T){
	ne := MyApp{}
	res:= fmt.Sprint(ne.SimpleMethod())

	if res != "Golang At Speed the API of notenough only comprises a really simple exported method" {
		t.Error("expected a different string, got", res )
	}
}
