package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"google.golang.org/api/idtoken"
	"proyecto_go/DTO/request"
)

func CognitoIdentityProvider() *cognitoidentityprovider.CognitoIdentityProvider {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Reemplaza con tu región
	}))
	return cognitoidentityprovider.New(sess)

}

func CognitoLogin(req request.AuthUser) (*cognitoidentityprovider.AuthenticationResultType, error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		ClientId: aws.String("6ko7nskvvfd8km2gmq2ij2tb81"), // Reemplaza con el ID de tu app client de Cognito
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(req.Email),
			"PASSWORD": aws.String(req.Password),
		},
	}

	cognitoClient := CognitoIdentityProvider()

	authResp, err := cognitoClient.InitiateAuth(authInput)
	if err != nil {
		return nil, err
	}

	return authResp.AuthenticationResult, nil
}

func CognitoSignUp(email string, password string) (*cognitoidentityprovider.SignUpOutput, error) {
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String("6ko7nskvvfd8km2gmq2ij2tb81"), // Reemplaza con el ID de tu app client de Cognito
		Username: aws.String(email),
		Password: aws.String(password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	cognitoClient := CognitoIdentityProvider()

	success, err := cognitoClient.SignUp(signUpInput)

	if err != nil {
		return nil, err
	}

	return success, nil

}

func CognitoConfirmSignUp(confirmSignUp request.ConfirmSignUpRequest) error {

	cognitoCLient := CognitoIdentityProvider()

	_, err := cognitoCLient.ConfirmSignUp(&cognitoidentityprovider.ConfirmSignUpInput{
		Username:         aws.String(confirmSignUp.Email),
		ConfirmationCode: aws.String(confirmSignUp.Code),
		ClientId:         aws.String("6ko7nskvvfd8km2gmq2ij2tb81"), // Cambia esto por tu ClientId
	})

	if err != nil {
		return err
	}

	return nil
}

func RegisterWithGoogleToken(googleToken string) (*cognitoidentityprovider.AuthenticationResultType, error) {
	// Crear contexto
	ctx := context.Background()

	// Verificar token de Google y obtener payload
	payload, err := idtoken.Validate(ctx, googleToken, "")
	if err != nil {
		return nil, fmt.Errorf("error verificando token de Google: %v", err)
	}

	// Extraer el correo y el ID único (sub) del payload de Google
	email, ok := payload.Claims["email"].(string)
	sub, ok := payload.Claims["sub"].(string)

	if !ok || email == "" || sub == "" {
		return nil, fmt.Errorf("no se encontró el correo o el ID de usuario en el token de Google")
	}

	// Registrar en Cognito usando el correo como username y el sub como contraseña
	cognitoClient := CognitoIdentityProvider()
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String("6ko7nskvvfd8km2gmq2ij2tb81"), // Reemplaza con tu ClientId de Cognito
		Username: aws.String(email),
		Password: aws.String(sub),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	// Intentar registro en Cognito
	_, err = cognitoClient.SignUp(signUpInput)
	if err != nil {
		return nil, fmt.Errorf("error registrando en Cognito: %v", err)
	}

	// Confirmar automáticamente el correo
	_, err = cognitoClient.AdminUpdateUserAttributes(&cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserPoolId: aws.String("us-east-1_j9e5Ladu8"), // Reemplaza con tu UserPoolId de Cognito
		Username:   aws.String(email),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error confirmando el correo en Cognito: %v", err)
	}

	return CognitoLogin(request.AuthUser{
		Email:    email,
		Password: sub,
	})
}

func ParseToken(token string) string {
	return ""
}
