package contracthighlow

import (
	"fmt"
	"time"

	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	"github.com/zerodha/gokiteconnect/v4/models"
)

// UserParam represent params to connect to kite connect APIs and fetch historical data
type UserParam struct {
	ApiKey      string
	AccessToken string
	Token       int
	Day         int
	Month       int
	Year        int
}

// HighLowAttr represents high and low data attributes response
type HighLowAttr struct {
	High     float64
	HighDate models.Time
	Low      float64
	LowDate  models.Time
}

// GetHighLow fetches high and low attributes response for the given time range
func GetHighLow(userInput UserParam) HighLowAttr {
	var (
		high     float64
		highDate models.Time
		low      float64
		lowDate  models.Time
	)
	// Create a new Kite connect instance
	kc := kiteconnect.New(userInput.ApiKey)
	// Set access token
	kc.SetAccessToken(userInput.AccessToken)
	// Calculate time range(start and end)
	endTime := time.Now()
	fromTime := endTime.AddDate(-userInput.Year, -userInput.Month, -userInput.Day)
	// Make historical data APIs request
	historicalData, err := kc.GetHistoricalData(userInput.Token, "day", fromTime, endTime, false, false)
	if err != nil {
		fmt.Printf("Error fetching historical data: %v", err)
	}
	// Set default value for all high and low attributes
	high = historicalData[0].High
	highDate = historicalData[0].Date
	low = historicalData[0].Low
	lowDate = historicalData[0].Date

	// Calculates high and low data for fetched historical data
	for _, candle := range historicalData {
		if candle.High > high {
			high = candle.High
			highDate = candle.Date
		}
		if candle.Low < low {
			low = candle.Low
			lowDate = candle.Date
		}
	}
	return HighLowAttr{
		High:     high,
		Low:      low,
		HighDate: highDate,
		LowDate:  lowDate,
	}
}
