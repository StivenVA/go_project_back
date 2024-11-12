package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/api/idtoken"
	"os"
	"proyecto_go/DTO/request"
)

func CognitoIdentityProvider() *cognitoidentityprovider.CognitoIdentityProvider {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"), // Reemplaza con tu región
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), ""),
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

func GetAttributes(googleToken string) (map[string]interface{}, error) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, googleToken, "")
	if err != nil {
		return nil, fmt.Errorf("error verificando token de Google: %v", err)
	}

	return payload.Claims, nil
}

func RegisterWithGoogleToken(googleToken string) (*cognitoidentityprovider.AuthenticationResultType, error) {
	payload, err := GetAttributes(googleToken)
	// Extraer el correo y el ID único (sub) del payload de Google
	email, ok := payload["email"].(string)
	sub, ok := payload["sub"].(string)

	password := passwordGenerator(sub)

	if !ok || email == "" || sub == "" {
		return nil, fmt.Errorf("no se encontró el correo o el ID de usuario en el token de Google")
	}

	// Registrar en Cognito usando el correo como username y el sub como contraseña
	cognitoClient := CognitoIdentityProvider()
	_, err = CognitoSignUp(email, password)

	if err != nil && err.Error() == "UsernameExistsException: User already exists" {

		return CognitoLogin(request.AuthUser{
			Email:    email,
			Password: password,
		})
	}
	if err != nil {
		return nil, fmt.Errorf("error registrando en Cognito: %v", err)
	}

	// Confirmar el usuario (activar la cuenta) en Cognito
	_, err = cognitoClient.AdminConfirmSignUp(&cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String("us-east-1_j9e5Ladu8"), // Reemplaza con tu UserPoolId de Cognito
		Username:   aws.String(email),
	})
	if err != nil {
		return nil, fmt.Errorf("error confirmando el registro en Cognito: %v", err)
	}

	// Confirmar el correo como verificado
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
		Password: password,
	})
}

func passwordGenerator(sub string) string {
	return sub + "Aa!"
}

func ExtractSubClaim(tokenStr string) (string, error) {
	// Parsear el token sin validar la firma.
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return "", fmt.Errorf("error al parsear el token: %v", err)
	}

	// Obtener los claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
		return "", fmt.Errorf("claim 'sub' no encontrado en el token")
	}
	return "", fmt.Errorf("no se pudieron obtener los claims")
}
