//Chain code Commercial Paper
// Function: This is an agnostic chaincode paper server that can take arguments and manage transactions
//Overview: A commercial paper is a promissory note that allows for a short term loan for small assets 
// and quickly appreciating loans.

package main

import (
"github.com/hyperledger/fabric/core/chaincode/shim"
//"errors"
"fmt"
"time"
)


type SimpleChaincode struct {

}



type CommercialPaper struct {

        issuer string
        recipient string
        amount int
	creation_date time.Time
        target string
        interest float64  // Decimals are percentages

}



//INIT
func (t *SimpleChaincode) Init (shim *shim.ChaincodeStub, function string, args []string) (bool){

    sample := CommercialPaper{issuer : "John Doe", 
                           recipient : "Jane Doe",
                           amount : 100000,
                           creation_date : time.Now(),
                           ameleoration_date : time.Now(),
                           }


    fmt.Printf(sample.issuer)

return false
}


//QUERY
func (t *SimpleChaincode) Query(shim *shim.ChaincodeStub, function string, args []string)(bool){
println(t)
return false
}

//INVOKE
func (t *SimpleChaincode) Invoke (shim *shim.ChaincodeStub, function string, args []string)(bool){


return false
}



//QUERY HELPERS

//INVOKE HELPERS

//MAIN

func main(){
err := shim.Start(new(SimpleChaincode))

    if err != nil{
        fmt.Println("BOOM BITCH",err)
    }
}

