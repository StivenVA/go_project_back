package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
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

func ParseToken(token string) string {
	return ""
}
