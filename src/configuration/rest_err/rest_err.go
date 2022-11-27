//arquivo para tratar erros em toda nossa aplicação
package rest_err

import "net/http"

//json:"message"` -> nome que vai aparecer no JSON
type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
	//caso eu nao queria que o causes apareça no json caso for vazio eu coloco o 'omitempty'
}

//lista de quais campos estão incorretos
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

//construtor para nosso objeto
//caso eu queira criar um objeto que não tenha aqui dentro usaremos o contructor
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

//para satisfazer a interface de error (porque a gente pode usa alguma lib que elas recebem erro por parametro, ou retornam erro por parametro)
//e para nós não precisar criar uma variavel error colocar o valor de RestError e retornar
//criamos essa função que é o restError
func (r *RestErr) Error() string {
	return r.Message
}

//ao invés da gente fica passando o erro se é 200, se é 400 etc..
//vamos criar um metodo para cada tipo de erro

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
