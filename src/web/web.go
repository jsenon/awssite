// Package web CloudTab.
//
// the purpose of this package is to provide Web HTML Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>

package web

import (
	"apigw"
	"ec2"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
	// "encoding/json"
	// "lambda"
)

// TO DO

// Present Information on Dedicated WebPortal

type Ec2instance struct {
	Reservations []struct {
		Instances []struct {
			AmiLaunchIndex      int    `json:"AmiLaunchIndex"`
			Architecture        string `json:"Architecture"`
			BlockDeviceMappings []struct {
				DeviceName string `json:"DeviceName"`
				Ebs        struct {
					AttachTime          string `json:"AttachTime"`
					DeleteOnTermination bool   `json:"DeleteOnTermination"`
					Status              string `json:"Status"`
					VolumeID            string `json:"VolumeId"`
				} `json:"Ebs"`
			} `json:"BlockDeviceMappings"`
			ClientToken        string `json:"ClientToken"`
			EbsOptimized       bool   `json:"EbsOptimized"`
			Hypervisor         string `json:"Hypervisor"`
			IamInstanceProfile struct {
				Arn string `json:"Arn"`
				ID  string `json:"Id"`
			} `json:"IamInstanceProfile"`
			ImageID      string `json:"ImageId"`
			InstanceID   string `json:"InstanceId"`
			InstanceType string `json:"InstanceType"`
			KernelID     string `json:"KernelId"`
			KeyName      string `json:"KeyName"`
			LaunchTime   string `json:"LaunchTime"`
			Monitoring   struct {
				State string `json:"State"`
			} `json:"Monitoring"`
			NetworkInterfaces []struct {
				Attachment struct {
					AttachTime          string `json:"AttachTime"`
					AttachmentID        string `json:"AttachmentId"`
					DeleteOnTermination bool   `json:"DeleteOnTermination"`
					DeviceIndex         int    `json:"DeviceIndex"`
					Status              string `json:"Status"`
				} `json:"Attachment"`
				Description string `json:"Description"`
				Groups      []struct {
					GroupID   string `json:"GroupId"`
					GroupName string `json:"GroupName"`
				} `json:"Groups"`
				MacAddress         string `json:"MacAddress"`
				NetworkInterfaceID string `json:"NetworkInterfaceId"`
				OwnerID            string `json:"OwnerId"`
				PrivateDNSName     string `json:"PrivateDnsName"`
				PrivateIPAddress   string `json:"PrivateIpAddress"`
				PrivateIPAddresses []struct {
					Primary          bool   `json:"Primary"`
					PrivateDNSName   string `json:"PrivateDnsName"`
					PrivateIPAddress string `json:"PrivateIpAddress"`
				} `json:"PrivateIpAddresses"`
				SourceDestCheck bool   `json:"SourceDestCheck"`
				Status          string `json:"Status"`
				SubnetID        string `json:"SubnetId"`
				VpcID           string `json:"VpcId"`
				Association     struct {
					IPOwnerID     string `json:"IpOwnerId"`
					PublicDNSName string `json:"PublicDnsName"`
					PublicIP      string `json:"PublicIp"`
				} `json:"Association,omitempty"`
			} `json:"NetworkInterfaces"`
			Placement struct {
				AvailabilityZone string `json:"AvailabilityZone"`
				GroupName        string `json:"GroupName"`
				Tenancy          string `json:"Tenancy"`
			} `json:"Placement"`
			PrivateDNSName   string `json:"PrivateDnsName"`
			PrivateIPAddress string `json:"PrivateIpAddress"`
			PublicDNSName    string `json:"PublicDnsName"`
			PublicIPAddress  string `json:"PublicIpAddress"`
			RootDeviceName   string `json:"RootDeviceName"`
			RootDeviceType   string `json:"RootDeviceType"`
			SecurityGroups   []struct {
				GroupID   string `json:"GroupId"`
				GroupName string `json:"GroupName"`
			} `json:"SecurityGroups"`
			SourceDestCheck bool `json:"SourceDestCheck"`
			State           struct {
				Code int    `json:"Code"`
				Name string `json:"Name"`
			} `json:"State"`
			StateTransitionReason string `json:"StateTransitionReason"`
			SubnetID              string `json:"SubnetId"`
			Tags                  []struct {
				Key   string `json:"Key"`
				Value string `json:"Value"`
			} `json:"Tags"`
			VirtualizationType string `json:"VirtualizationType"`
			VpcID              string `json:"VpcId"`
		} `json:"Instances"`
		OwnerID       string `json:"OwnerId"`
		ReservationID string `json:"ReservationId"`
	} `json:"Reservations"`
}

type myEC2struct struct {
	Instanceid   string   `json:"Instanceid"`
	InstanceType string   `json:"InstanceType"`
	Tags         []Mytags `json:"Mytags"`
}

type Mytags struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

// Func to display all server
func Index(res http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, req)

}

func Help(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/help.html")
	t.Execute(res, req)
}

func LaunchApigw(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	myurl := req.FormValue("URL")
	myunauthourl := req.FormValue("UnURL")

	numrequest := req.FormValue("NumRequest")
	mynumrequest, err := strconv.ParseInt(numrequest, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("%v %s", mynumrequest, reflect.TypeOf(mynumrequest))

	status403 := req.FormValue("403")
	status401 := req.FormValue("401")

	// t, _ := template.ParseFiles("templates/Inprogress.html")
	// t.Execute(res, Server)
	fmt.Println("------> Launch BURST ApiGW")

	// Call api without Credential
	// Need URL, number of Send
	apigw.Getburst(myurl, mynumrequest)

	// Loop on wrong URL to generate 403
	if status403 == "on" {
		mywrongurl := myurl + "/fake"
		apigw.Getburst(mywrongurl, mynumrequest)
	}

	// Loop on Unauthorized Ressource 401 Based on POST
	if status401 == "on" {
		apigw.Postburst(myunauthourl, mynumrequest)
	}

	// Need URL redirect
	http.Redirect(res, req, "/index", http.StatusSeeOther)

	// Call apigw.go with Credential, URL, Number of burst

}

func Ec2mgmt(res http.ResponseWriter, req *http.Request) {
	// Function Not Used in HTML for the moment
	// Function will print ec2 instances in stdout

	var myEC2 []myEC2struct
	var tabtags []Mytags

	arn := "YOURARN"
	req.ParseForm()

	fmt.Println("------> Launch BURST EC2 Create")
	// Create Sesion to EC2 and use switch role
	sess := ec2.Ec2AWSRole(arn)
	// fmt.Println ("Session:", sess)
	// Retrieve instance info
	data := ec2.Ec2Getinstance(sess)
	// fmt.Println ("MyData:", data)
	for idx, _ := range data.Reservations {
		for _, inst := range data.Reservations[idx].Instances {
			fmt.Println("instance", *inst.InstanceId)
			result := *inst.InstanceId
			tabtags = make([]Mytags, 0, 20)

			for _, t := range inst.Tags {

				// fmt.Println("Key:\n", *t.Key)
				// fmt.Println("Value:\n", *t.Value)
				tabtags = append(tabtags, Mytags{Key: *t.Key, Value: *t.Value})

				// fmt.Println ("tgs:", tabtags)
			}

			myEC2 = append(myEC2, myEC2struct{

				Instanceid: result,
				Tags:       tabtags,
			})

			fmt.Println("myEC2:\n", myEC2)

		}

	}

	t, _ := template.ParseFiles("templates/ec2mgmt.html")
	t.Execute(res, myEC2)
	// t.Execute(res, data)

}

func LaunchLambda(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	// t, _ := template.ParseFiles("templates/Inprogress.html")
	// t.Execute(res, Server)
	fmt.Println("------> Launch BURST Lambda")
}
