/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
//	"strconv"
//	"encoding/json"
//  "time"
//	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type login struct{
	Id string `json:"id"`					//User login JSON Schema
	Password string `json:"password"`
	Name string `json:"name"`
	Email string `json:"email"`
	User  string `json:"user"`
}


// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error

	if len(args) != 2 {
	   return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
    //Write the User Id "mail Id" arg[0] and password arg[1]
	userid := args[0]															//argument for UserID
	password := args[1]  	//argument for password
	str := `{"userid": "` + userid+ `", "password": "` + password + `"}`
	
	err = stub.PutState(userid, []byte(str))								//Put the userid and password in blockchain
	
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point for Invocations -  22/03/2016
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "create_user"{										//create a new User
		return t.create_user(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	return nil, nil
}

// ============================================================================================================================
// Creare User - create a new user, store into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) create_user(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	//   0       1          2     3       5
	// "id", "password", "name", "email","user"
	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 5")
	}

	//input sanitation
	fmt.Println("- start create user")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
		
	}	
	
	id := args[0]
	password := args[1]
	name := args[2]
	email := args[3]
	user := args[4]
	
	////check if user already exists
	//marbleAsBytes, err := stub.GetState(id)
	//if err != nil {
	//	return nil, errors.New("Failed to get user")
	//}
	//res := Marble{}
	//json.Unmarshal(marbleAsBytes, &res)
	//if res.Name == name{
	//	fmt.Println("This marble arleady exists: " + name)
	//	fmt.Println(res);
	//	return nil, errors.New("This marble arleady exists")				//all stop a marble by this name exists
	//}
	
	//build the login json string manually
	str := `{"id": "` + id+ `", "password": "` + password + `", "name": ` + name + `, "email": "` + email + `","user": "` + user + `"}`
	err = stub.PutState(id, []byte(str))									//store user with id as key
	if err != nil {
		return nil, err
	}
	    return nil,err
	}