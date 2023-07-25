package types

type PayloadField string

// PayloadField contains the names of the parameters and payload fields
const (
	FieldTransactionID   PayloadField = "trans_id"
	FieldAmount          PayloadField = "amount"
	FieldCurrency        PayloadField = "currency"
	FieldClientIPAddress PayloadField = "client_ip_addr"
	FieldDescription     PayloadField = "description"
	FieldLanguage        PayloadField = "language"
	FieldBillerClientID  PayloadField = "biller_client_id"
	FieldPerspayeeExpiry PayloadField = "prespayee_expiry"
	FieldTransactionType PayloadField = "transaction_type"
)
