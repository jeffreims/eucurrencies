package eucurrencies

type cubeCurrencies struct {
	Currency string  `xml:"currency,attr"`
	Rate     float64 `xml:"rate,attr"`
}

type cubeContainer struct {
	Currencies []cubeCurrencies `xml:"Cube>Cube"`
}

type euData struct {
	Cube cubeContainer `xml:"Cube"`
}

type EUCurrencies struct {
	USD float64
	JPY float64
	BGN float64
	CZK float64
	DKK float64
	GBP float64
	HUF float64
	PLN float64
	RON float64
	SEK float64
	CHF float64
	ISK float64
	NOK float64
	HRK float64
	RUB float64
	TRY float64
	AUD float64
	BRL float64
	CAD float64
	CNY float64
	HKD float64
	IDR float64
	ILS float64
	INR float64
	KRW float64
	MXN float64
	MYR float64
	NZD float64
	PHP float64
	SGD float64
	THB float64
	ZAR float64
}
