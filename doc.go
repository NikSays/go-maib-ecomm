/*
Package maib provides tools to interact with MAIB ECommerce in a type safe way.

# Requirements

To use this module you should:
  - Understand how MAIB EComm works
  - Have a .pfx certificate
  - Register as a merchant in MAIB

# Usage

The main part of the module is the [ECommClient] struct. It contains an inner [http.Client] with Transport
set up to support mTLS, which is required for communication with EComm.

The requests described in EComm documentation are implemented in "requests" directory. Running [ECommClient.Send] on a
[Request] does the following:
 1. Validates the request
 2. Encodes it into a querystring
 3. Sends it to the MAIB EComm server
 4. Decodes the response into map[string]any

The response map can be decoded into a struct.

See the example to get an understanding of the full flow.
*/
package maib
