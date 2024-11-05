package services

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"net/http"
	"proyecto_go/DTO/request"
)

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req request.AuthUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Inicia una sesión con AWS
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Reemplaza con tu región
	}))

	// Crea el cliente de Cognito
	cognitoClient := cognitoidentityprovider.New(sess)

	// Configura los parámetros de autenticación
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		ClientId: aws.String("6ko7nskvvfd8km2gmq2ij2tb81"), // Reemplaza con el ID de tu app client de Cognito
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(req.Email),
			"PASSWORD": aws.String(req.Password),
		},
	}

	// Llama a Cognito para autenticar al usuario
	authResp, err := cognitoClient.InitiateAuth(authInput)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Si la autenticación es exitosa, devuelve el token
	token := *authResp.AuthenticationResult.IdToken
	resp := LoginResponse{
		Message: "Login successful",
		Token:   token,
	}

	// Configura el encabezado Content-Type y escribe la respuesta en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

type RegisterResponse struct {
	Message string `json:"message"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Inicia una sesión con AWS
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Reemplaza con tu región
	}))

	// Crea el cliente de Cognito
	cognitoClient := cognitoidentityprovider.New(sess)

	// Configura los parámetros de registro
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String("6ko7nskvvfd8km2gmq2ij2tb81"), // Reemplaza con el ID de tu app client de Cognito
		Username: aws.String(req.Email),
		Password: aws.String(req.Password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(req.Email),
			},
		},
	}
	// Llama a Cognito para registrar al usuario
	_, err = cognitoClient.SignUp(signUpInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := RegisterResponse{
		Message: "User registered successfully. Please confirm your email.",
	}

	// Configura el encabezado Content-Type y escribe la respuesta en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func ConfirmSignUp(w http.ResponseWriter, r *http.Request) {
	// Decodifica el JSON del cuerpo de la solicitud
	var req request.ConfirmSignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Inicializa el cliente de Cognito
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Cambia esto a tu región
	})
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	cognitoSvc := cognitoidentityprovider.New(sess)

	// Llama a la API para confirmar el registro
	_, err = cognitoSvc.ConfirmSignUp(&cognitoidentityprovider.ConfirmSignUpInput{
		Username:         aws.String(req.Email),
		ConfirmationCode: aws.String(req.Code),
		ClientId:         aws.String("6ko7nskvvfd8km2gmq2ij2tb81"), // Cambia esto por tu ClientId
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to confirm sign up: %v", err), http.StatusInternalServerError)
		return
	}

	// Responde con un mensaje de éxito
	resp := RegisterResponse{
		Message: "User confirmed successfully.",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
