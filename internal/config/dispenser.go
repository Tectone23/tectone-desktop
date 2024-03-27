package config

import (
	"context"
	"encoding/json"
	"tectone/tectone-desktop/internal/model"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const ExampleDispenserConfig = `{
  "bank.testnet.algorand.network": {
    "recaptcha_sitekey": "YOUR SITEKEY HERE"
    "recaptcha_secret": "YOUR SECRET HERE",
    "amount": 1000,
    "fee": 1,
    "wallet": "M5XGQQMKZJQI2GIBSMUYRYB32QT23MBPAEWDQXQE223IQM7PY22JPLZOHY",
    "data_dir": "http://127.0.0.1:8161",
  },

  "bank.devnet.algodev.network": {
    "recaptcha_sitekey": "YOUR SITEKEY HERE",
    "recaptcha_secret": "YOUR SECRET HERE",
    "amount": 1000,
    "fee": 1,
    "wallet": "A4NEDM7HE23PYDDMGY4RKSNGIYFKO5SWDLEFAUIB3Y7XUSQ7754E4YR364",
    "data_dir": "http://127.0.0.1:8160",
  }
}`

//  EXAMPLE DISPENSER CONFIG
//
// {
//   "bank.testnet.algorand.network": {
//     "recaptcha_sitekey": "YOUR SITEKEY HERE"
//     "recaptcha_secret": "YOUR SECRET HERE",
//     "amount": 1000,
//     "fee": 1,
//     "wallet": "M5XGQQMKZJQI2GIBSMUYRYB32QT23MBPAEWDQXQE223IQM7PY22JPLZOHY",
//     "data_dir": "http://127.0.0.1:8161",
//   },

//	  "bank.devnet.algodev.network": {
//	    "recaptcha_sitekey": "YOUR SITEKEY HERE",
//	    "recaptcha_secret": "YOUR SECRET HERE",
//	    "amount": 1000,
//	    "fee": 1,
//	    "wallet": "A4NEDM7HE23PYDDMGY4RKSNGIYFKO5SWDLEFAUIB3Y7XUSQ7754E4YR364",
//	    "data_dir": "http://127.0.0.1:8160",
//	  }
//	}
type DispenserConfig struct {
	AlgorandTestnetBankNetwork model.BankNetwork `json:"bank.testnet.algorand.network"`
	AlgodevDevnetBankNetwork   model.BankNetwork `json:"bank.devnet.algodev.network"`
}

func GetExampleDispenserConfig(ctx context.Context) *DispenserConfig {
	var c *DispenserConfig
	if err := json.Unmarshal([]byte(ExampleDispenserConfig), c); err != nil {
		runtime.LogErrorf(ctx, "error could not unmarshal example dispenser config json: %s\n", err)
		return c
	}
	return c
}
