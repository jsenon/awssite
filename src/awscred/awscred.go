package awscred

import (
	// "fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Func to manage the Assume role needed when you switch your account in AWS
// Need in input :
// 		ARN need to switch
// 		Region
//		Session PTR
func CreateConfig(arn string, externalID string, region string, sess *session.Session) aws.Config {

	conf := aws.Config{Region: aws.String(region)}
	// fmt.Println("arn",arn)
	if arn != "" {
		// if ARN flag is passed in, we need to be able ot assume role here
		var creds *credentials.Credentials
		// fmt.Println("externalID",externalID)
		if externalID != "" {
			// If externalID flag is passed, we need to include it in credentials struct
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {
				p.ExternalID = &externalID
			})
		} else {
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {
			})
		}
		conf.Credentials = creds
	}
	return conf
}
