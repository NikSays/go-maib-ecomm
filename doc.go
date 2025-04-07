/*
Package maib provides tools to interact with the MAIB ECommerce system in a type
safe way.

# Requirements

To use this module you should:
  - Understand how MAIB ECommerce works
  - Have a .pfx certificate
  - Register as a merchant in MAIB

# Usage

 1. Use [NewClient] to set up a [Client] that communicates with the MAIB
    ECommerce system.
 2. Send a [Request] with [Client.Send] (The requests described in the
    ECommerce documentation are implemented in the `requests` package).
 3. Decode the returned map into a result struct with requests.DecodeResult.

# Error Handling

Use [errors.As] to check the type and the contents of the errors returned by
the [Client]:
  - [ValidationError] is returned before sending the request if it
    has failed validation.
  - [ECommError] is returned if the response has a non-200 code, or its
    body starts with "error:".
  - [ParseError] is returned if the response has an invalid structure, or
    a response field has an unexpected datatype.

See the example to get an understanding of the full flow.
*/
package maib
