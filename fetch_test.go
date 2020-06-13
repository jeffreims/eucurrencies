package eucurrencies

import (
	"testing"
)

func TestCurrency(t *testing.T) {
	var data EUCurrencies
	if err := data.Fetch(); err != nil {
		t.Log(err)
		t.FailNow()
	}
	if data.NOK == 0 && data.GBP == 0 {
		t.Log("NOK and GBP seem to be 0")
		t.FailNow()
	}
}
