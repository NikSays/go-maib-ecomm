package requests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/NikSays/go-maib-ecomm/types"
)

func TestRegisterRecurring(t *testing.T) {
	cases := []struct {
		name               string
		payload            RegisterRecurring
		expectedErrorField types.PayloadField
		expectedEncoded    string
	}{
		{
			name: "OK SMS",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedEncoded: "amount=1234&biller_client_id=biller&client_ip_addr=127.0.0.1&command=z&currency=498&description=Description&language=en&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "OK DMS",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringDMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedEncoded: "amount=1234&biller_client_id=biller&client_ip_addr=127.0.0.1&command=d&currency=498&description=Description&language=en&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "OK no BillerClientID",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedEncoded: "amount=1234&biller_client_id=&client_ip_addr=127.0.0.1&command=z&currency=498&description=Description&language=en&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "OK overwrite",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				PerspayeeExpiry:   "1224",
				BillerClientID:    "biller",
				OverwriteExisting: true,
			},
			expectedEncoded: "amount=1234&biller_client_id=biller&client_ip_addr=127.0.0.1&command=z&currency=498&description=Description&language=en&perspayee_expiry=1224&perspayee_overwrite=1",
		},
		{
			name: "OK without payment",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringWithoutPayment,
				Amount:            0,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				PerspayeeExpiry:   "1224",
				BillerClientID:    "biller",
				OverwriteExisting: false,
			},
			expectedEncoded: "biller_client_id=biller&client_ip_addr=127.0.0.1&command=p&currency=498&description=Description&language=en&perspayee_expiry=1224&perspayee_gen=1",
		},
		{
			name: "TransactionType invalid",
			payload: RegisterRecurring{
				TransactionType:   -9,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: types.FieldCommand,
		},
		{
			name: "Amount invalid",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            0,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: types.FieldAmount,
		},
		{
			name: "Currency invalid",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          1000,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: types.FieldCurrency,
		},
		{
			name: "ClientIPAddress invalid",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "927.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: types.FieldClientIPAddress,
		},
		{
			name: "Description invalid",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       strings.Repeat("-", 130),
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: types.FieldDescription,
		},
		{
			name: "Language invalid",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          "",
				BillerClientID:    "biller",
				PerspayeeExpiry:   "1224",
				OverwriteExisting: false,
			},
			expectedErrorField: types.FieldLanguage,
		},
		{
			name: "PerspayeeExpiry invalid",
			payload: RegisterRecurring{
				TransactionType:   RegisterRecurringSMS,
				Amount:            1234,
				Currency:          types.CurrencyMDL,
				ClientIPAddress:   "127.0.0.1",
				Description:       "Description",
				Language:          types.LanguageEnglish,
				BillerClientID:    "biller",
				PerspayeeExpiry:   "wrong",
				OverwriteExisting: false,
			},
			expectedErrorField: types.FieldPerspayeeExpiry,
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
				assert.Equal(t, c.expectedErrorField, err.(types.ValidationError).Field)
			}
		})
	}
}
