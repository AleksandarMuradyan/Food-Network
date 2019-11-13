package main

import (
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

)

type SmartContract struct {
}

type Intity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Stock struct {
	Id          string `json:"id"`
	Meat        string `json:"meat"`
	Importer    string `json:"importer"`
	Distributor string `json:"distributor"`
	Customer    string `json:"customer"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}
func (s *SmartContract) CreateUser(stub shim.ChaincodeStubInterface,args []string) pb.Response {
	Id := args[0]
	var user = Intity{ Id:Id,Name: args[1], Type: args[2]}
	UserBytes, _ := json.Marshal(user)
	stub.PutState(Id, UserBytes)
	return shim.Success(nil)
}
func (s *SmartContract) queryUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting UserID")
	}
	UserBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(UserBytes)
}
