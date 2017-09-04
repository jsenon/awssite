//go:generate swagger generate spec
// Package main burstaws.
//
// the purpose of this application is to provide an burst interface to generate log and traffic for demo purpose
// of monitoring and log tools
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes:
//     Host:
//     BasePath:
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
package main

import (
	"web"
	// "apigw"
	// "ec2"
	// "lambda"
	// "fmt"
	"api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// TO FIX

func main() {
	r := mux.NewRouter()

	// Remove CORS Header check to allow swagger and application on same host and port
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// To be changed
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH"})

	// Web Part
	r.HandleFunc("/index", web.Index)
	r.HandleFunc("/help", web.Help)
	r.HandleFunc("/apigwsend", web.LaunchApigw)
	r.HandleFunc("/ec2mgmt", web.Ec2mgmt)
	r.HandleFunc("/lambdasend", web.LaunchLambda)

	// API Part
	// r.HandleFunc("/apigwsend", apigw.Getburst).Methods("PUT")
	// r.HandleFunc("/ec2create", ec2.Create).Methods("POST")
	// r.HandleFunc("/lambdasend", lambda.Send).Methods("PUT")


	// Static dir
	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/static/"))))

	// Health Check-
	r.HandleFunc("/healthy/am-i-up", api.Statusamiup).Methods("GET")
	r.HandleFunc("/healthy/about", api.Statusabout).Methods("GET")

	http.ListenAndServe(":9040", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}