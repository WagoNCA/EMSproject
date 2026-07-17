package models

type MeterReading struct {
	MeterID string  `json:"meter_id"`
	Value   float64 `json:"value"`
}
