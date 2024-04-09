package types

// PayloadField contains the names of the payload fields. Used in [ValidationError].
type PayloadField string

const (
	FieldTransactionID   PayloadField = "trans_id"
	FieldAmount          PayloadField = "amount"
	FieldCurrency        PayloadField = "currency"
	FieldClientIPAddress PayloadField = "client_ip_addr"
	FieldDescription     PayloadField = "description"
	FieldLanguage        PayloadField = "language"
	FieldBillerClientID  PayloadField = "biller_client_id"
	FieldPerspayeeExpiry PayloadField = "prespayee_expiry"
	FieldCommand         PayloadField = "command"
)
