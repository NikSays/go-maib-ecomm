package requests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm"
)

func TestOneClick(t *testing.T) {
	cases := []struct {
		name               string
		payload            RegisterOneClick
		expectedErrorField maib.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK SMS",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedEncoded: "amount=1234&biller_client_id=biller&client_ip_addr=127.0.0.1&command=z&currency=498&description=Description&language=en&oneclick=Y&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "OK no BillerClientID",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedEncoded: "amount=1234&biller_client_id=&client_ip_addr=127.0.0.1&command=z&currency=498&description=Description&language=en&oneclick=Y&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "OK overwrite",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				PerspayeeExpiry:   "1224",
				BillerClientID:    "biller",
				OverwriteExisting: true,
			},
			expectedEncoded: "amount=1234&biller_client_id=biller&client_ip_addr=127.0.0.1&command=z&currency=498&description=Description&language=en&oneclick=Y&perspayee_expiry=1224&perspayee_overwrite=1",
		},
		{
			name: "OK ask save",
			payload: RegisterOneClick{
				TransactionType: RegisterOneClickSMS,
				Amount:          1234,
				Currency:        maib.CurrencyMDL,
				ClientIPAddress: "127.0.0.1",
				Description:     "Description",
				Language:        maib.LanguageEnglish,
				PerspayeeExpiry: "1224",
				BillerClientID:  "biller",
				AskSaveCardData: true,
			},
			expectedEncoded: "amount=1234&ask_save_card_data=True&biller_client_id=biller&client_ip_addr=127.0.0.1&command=z&currency=498&description=Description&language=en&oneclick=Y&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "OK without payment",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickWithoutPayment,
				Amount:            0,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				PerspayeeExpiry:   "1224",
				BillerClientID:    "biller",
				OverwriteExisting: false,
			},
			expectedEncoded: "biller_client_id=biller&client_ip_addr=127.0.0.1&command=p&currency=498&description=Description&language=en&oneclick=Y&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "TransactionType invalid",
			payload: RegisterOneClick{
				TransactionType:   -9,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: maib.FieldCommand,
		},
		{
			name: "Amount invalid",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            0,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: maib.FieldAmount,
		},
		{
			name: "Currency invalid",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          1000,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: maib.FieldCurrency,
		},
		{
			name: "ClientIPAddress invalid",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "927.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: maib.FieldClientIPAddress,
		},
		{
			name: "Description invalid",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       strings.Repeat("-", 130),
				Language:          maib.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: maib.FieldDescription,
		},
		{
			name: "Language invalid",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          "",
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: maib.FieldLanguage,
		},
		{
			name: "PerspayeeExpiry invalid",
			payload: RegisterOneClick{
				TransactionType:   RegisterOneClickSMS,
				Amount:            1234,
				Currency:          maib.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          maib.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "wrong",
				OverwriteExisting: false,
			},
			expectedErrorField: maib.FieldPerspayeeExpiry,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.payload.Validate()
			if c.expectedErrorField == "" {
				assert.Nil(t, err)
				val, err := c.payload.Values()
				assert.Nil(t, err)
				assert.Equal(t, c.expectedEncoded, val.Encode())
			} else {
				assert.Equal(t, c.expectedErrorField, err.(*maib.ValidationError).Field)
			}
		})
	}
}
