package request

//envio do usuario
type UserRequest struct {
	//Sem gingonic -> Email    string `json:"email" validate:""` (gingonic tem um validator por debaixo dos panos)
	//com gingonic -> Email    string `json:"email" binding:""` (usando a lib validator)
	//Como vamos utilizar o gingonic que esta fazendo o 'ShouldByJson' nós vamos utilizar 'binding'
	//se nao tivesse gingonic usariamos 'validator' no lugar de 'binding'
	//ou seja para fazer a validação com o 'ShouldByJson' do gingonic eu preciso utilizar o binding
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"` //containsAny=!@#$%* -> é obrigatório a senha ter algum desses caracteres
	Name     string `json:"name" binding:"required,min=4,max=50"`                 //required -> obrigatório //min-max -> entre 4 e 50 caracteres
	Age      int8   `json:"age" binding:"required,min=1,max=140"`
}
