package models

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"math"
	"time"
)

type Runtime time.Duration

var RuntimeMarshaler = json.MarshalToFunc(func(enc *jsontext.Encoder, val Runtime) error {
	mins := math.Round(time.Duration(val).Minutes())
	token := jsontext.String(fmt.Sprintf("%dm", int(mins)))
	return enc.WriteToken(token)
})
