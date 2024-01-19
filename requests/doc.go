/*
Package requests contains the request structs for each of the commands available in the MAIB ECommerce.

Each request struct implements Request interface from the base package, and is accompanied by a result struct. E.g.,
[CloseDay] has [CloseDayResult]. Function [DecodeResponse] should be used to parse the response map into the result struct.

Additional fields in the payload are not supported.
*/
package requests
