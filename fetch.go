package eucurrencies

import "net/http"
import "time"
import "encoding/xml"

var client = http.Client{
	Timeout: 5 * time.Second,
}

//Fetch data from europa.eu
func (r *EUCurrencies) Fetch() error {
	var data euData

	res, err := client.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = xml.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	for _, v := range data.Cube.Currencies {
		switch v.Currency {
		case "USD":
			r.USD = v.Rate
		case "JPY":
			r.JPY = v.Rate
		case "BGN":
			r.BGN = v.Rate
		case "CZK":
			r.CZK = v.Rate
		case "DKK":
			r.DKK = v.Rate
		case "GBP":
			r.GBP = v.Rate
		case "HUF":
			r.HUF = v.Rate
		case "PLN":
			r.PLN = v.Rate
		case "RON":
			r.RON = v.Rate
		case "SEK":
			r.SEK = v.Rate
		case "CHF":
			r.CHF = v.Rate
		case "ISK":
			r.ISK = v.Rate
		case "NOK":
			r.NOK = v.Rate
		case "HRK":
			r.HRK = v.Rate
		case "RUB":
			r.RUB = v.Rate
		case "TRY":
			r.TRY = v.Rate
		case "AUD":
			r.AUD = v.Rate
		case "BRL":
			r.BRL = v.Rate
		case "CAD":
			r.CAD = v.Rate
		case "CNY":
			r.CNY = v.Rate
		case "HKD":
			r.HKD = v.Rate
		case "IDR":
			r.IDR = v.Rate
		case "ILS":
			r.ILS = v.Rate
		case "INR":
			r.INR = v.Rate
		case "KRW":
			r.KRW = v.Rate
		case "MXN":
			r.MXN = v.Rate
		case "MYR":
			r.MYR = v.Rate
		case "NZD":
			r.NZD = v.Rate
		case "PHP":
			r.PHP = v.Rate
		case "SGD":
			r.SGD = v.Rate
		case "THB":
			r.THB = v.Rate
		case "ZAR":
			r.ZAR = v.Rate
		}
	}
	return nil
}
