package main

import (
	"fmt"
    "bytes"
	"encoding/json"
    "strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	 pb "github.com/hyperledger/fabric/protos/peer"
	)

// Define the chaincode structure
type Chaincode1 struct {
}

//Define DWPB structure 
type DWPB struct {
	UniqueID     			string    `json:"UniqueID"`
	Latest_approved_value	float64   `json:"Latest_approved_value"`
	Latest_utilised_value	float64   `json:"Latest_utilised_value"`
	Latest_remaining_value	float64   `json:"Latest_remaining_value"`
	ModifiedDate 			int64	  `json:"Modified_Date"` // timestamp
	Version					string    `json:"Version"`
	Expired      			bool	  `json:"Expired"`
	Budget_sub_status		string    `json:"Budget_sub_status"`
}

func (s *Chaincode1) Init (APIstub shim.ChaincodeStubInterface, args[] string) pb.Response {
	fmt.Println ("Chaincode1 init")
}

func (s *Chaincode1) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
    if function == "createDWPB" {
	return s.createDWPB(APIstub, args)
  	} else if function == "queryDWPB" {
	return s.queryDWPB(APIstub, args)
    }
    return shim.Error("Invalid Smart Contract function name.")
}    


func (s *Chaincode1) CreateDWPB (APIstub shim.ChaincodeStubInterface) pb.Response {
    if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

// populate DWPB
	DWPB := DWPB {}
	UniqueID  = args[0]
	Latest_approved_value = args[1]
	Latest_utilised_value = args[2]
	Latest_remaining_value = args[3]
	ModifiedDate = args[4]
	Version = args[5]
	Expired  = args[6]
	Budget_sub_status = args[7]

//store state in ledger
	DWPBAsBytes,_ := json.Marshal(DWPB)	
	err := APIstub.PutState(UniqueID,DWPBAsBytes)
	if err != nil {
	return shim.Error(fmt.Sprintf("Failed to create DWPB"))
	}

//function to update the existing DWPB

//func (s *Chaincode1) UpdateDWPB (APIstub shim.ChaincodeStubInterface) pb.Response {
//}
// to be written
//update for one unique ID 	

//function to query a DWPB
func (s *Chaincode1) queryDWPB (APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting UniqueID")
  }
 DWPBAsBytes, err := APIstub.GetState(args[0])
  if err != nil {
    return shim.Error(err.Error())
  }
  return shim.Success(DWPBAsBytes)
}
