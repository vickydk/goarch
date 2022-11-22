package constants

import "net/http"

var (
	ErrorGeneralMsg           = "Internal Server Error"
	ErrorBadRequestMsg        = "Bad Request"
	ErrorValidationMsg        = "Validation Error"
	ErrorInvalidRequestMsg    = "Invalid request"
	ErrorNotFoundMsg          = "Data not found"
	ErrorInvalidEmailMsg      = "Invalid email"
	ErrorNotAuthorizedMsg     = "Not authorized"
	ErrorInRequestValidParams = "Invalid Request Params"
)

var (
	ErrorGeneral           = NewError(http.StatusInternalServerError, ErrorGeneralMsg)
	ErrorValidation        = NewError(http.StatusBadRequest, ErrorValidationMsg)
	ErrorInvalidRequest    = NewError(http.StatusBadRequest, ErrorInvalidRequestMsg)
	ErrorDataNotFound      = NewError(http.StatusNotFound, ErrorNotFoundMsg)
	ErrorInvalidEmail      = NewError(http.StatusBadRequest, ErrorInvalidEmailMsg)
	ErrorUserCannotAccess  = NewError(http.StatusUnauthorized, ErrorNotAuthorizedMsg)
	ErrorDatabase          = NewError(http.StatusInternalServerError, "Database Error")
	ErrorUserNotFound      = NewError(http.StatusNotFound, "User Not Found")
	ErrorPasswordNotMatch  = NewError(http.StatusUnauthorized, "Password not match")
	ErrorInvoiceExisted    = NewError(http.StatusBadRequest, "Invoice already exist")
	ErrorSuratJalanExisted = NewError(http.StatusBadRequest, "No surat jalan already exist")
	ErrorInvoiceNotExisted = NewError(http.StatusBadRequest, "Invoice not exist")
	ErrorOpCostExisted     = NewError(http.StatusBadRequest, "operational cost already exist")
	ErrorOpCostNotExisted  = NewError(http.StatusBadRequest, "operational cost not exist")
	ErrorBulanValidation   = NewError(http.StatusBadRequest, "Invalid bulan")
	ErrorTanggalMasuk      = NewError(http.StatusBadRequest, "Tanggal masuk tidak boleh kosong")
	ErrorTanggalKeluar     = NewError(http.StatusBadRequest, "Tanggal keluar tidak boleh kosong")
)

func NewError(errorCode int, message string) error {
	return &ApplicationError{
		Code:    errorCode,
		Message: message,
	}
}

type ApplicationError struct {
	Code    int
	Message string
}

func (e *ApplicationError) Error() string {
	return e.Message
}
