package canonical

import "fmt"

var (
	// G: Generic Errors
	GenericError         = Error{Code: "G101", Message: "Unexpected error occurred"}
	MethodNotImplemented = Error{Code: "G102", Message: "Method Not Implemented"}

	// I: Infrastrucuture Errors
	ConnectDatabaseError = Error{Code: "I201", Message: "Error Connecting into Database"}

	// D: Data Errors
	DBEntityNotFound   = Error{Code: "D301", Message: "Entity not found on Database"}
	DBConfigNotLoaded  = Error{Code: "D302", Message: "Database string connection is missing"}
	DBCreatingError    = Error{Code: "D303", Message: "Database error when creating a new entity"}
	DBNoEntityToUpdate = Error{Code: "D304", Message: "No entities found to update"}
	DBDeletingError    = Error{Code: "D305", Message: "Database error when deleting an entity"}

	// B: Business Errors
	MissingFieldsError = Error{Code: "B401", Message: "Missing required fields"}
)

type Error struct {
	Code    string
	Message string
}

func (r Error) Error() string {
	return fmt.Sprintf("%s - %s", r.Code, r.Message)
}
