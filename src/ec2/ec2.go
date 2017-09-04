package ec2

import (
	"awscred"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"


	// "os"
)

type Credential struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	TokenID string `json:TokenID`
}

// TO DO


func Ec2AWSRole(arn string)  *ec2.EC2 {
	// externalID Not Used
	externalID := ""

	
	fmt.Println("Arn:", arn)

	// Credential need $HOME/.aws/credential and $HOME/.aws/config
	sess := session.Must(session.NewSession())
	region := "eu-central-1"
	fmt.Println("Region:", region)
	conf := awscred.CreateConfig(arn, externalID, region, sess)

	// Create Session to EC2 with aws-go SDK
	ec2Svc := ec2.New(sess,&conf)
	return ec2Svc
}

func Ec2Getinstance(ec2Svc *ec2.EC2)  *ec2.DescribeInstancesOutput {

	// Create myFilter
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
		},
	}
	resp, _ := ec2Svc.DescribeInstances(params)
	

	// for idx, _ := range resp.Reservations {
	// 	for _, inst := range resp.Reservations[idx].Instances {
	// 		fmt.Println("instance", inst.InstanceId)
	// 	}
	// }
	return resp
}

