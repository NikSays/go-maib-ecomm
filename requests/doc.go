/*
Package requests contains the request structs for each of the commands available
in the MAIB ECommerce.

Each request struct implements Request interface from the base package, and is
accompanied by a result struct. E.g., [CloseDay] has [CloseDayResult]. Function
[DecodeResponse] should be used to parse the response map into the result
struct.

Additional fields in the payload are not supported natively, but you can create
custom requests like this:

	// A custom type that embeds an existing request.
	type RegisterTransactionWithFlightInfo struct {
		requests.RegisterTransaction
		Airline string
	}

	// Implement the Request interface by running the "super" function,
	// then attach the additional field.
	func (payload RegisterTransactionWithFlightInfo) Values() (url.Values, error) {
		vals, _ := req.RegisterTransaction.Values()
		vals.Add("airline", payload.Airline)
		return vals, nil
	}

	// Now RegisterTransactionWithFlightInfo can be used in Client.Send.
*/
package requests
