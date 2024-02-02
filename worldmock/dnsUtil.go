package worldmock

import (
	"fmt"

	mei "github.com/subrahamanyam341/andes-scenario-1234/scenario/expression/interpreter"
)

var numDNSAddresses = uint8(0xFF)
var dnsAddressVMType = []byte{5, 0}

func makeDNSAddresses(numAddresses uint8) map[string]struct{} {
	ei := mei.ExprInterpreter{
		VMType: dnsAddressVMType,
	}

	dnsMap := make(map[string]struct{}, numAddresses)
	for i := uint8(0); i < numAddresses; i++ {
		// using the value interpreter to generate the addresses
		// consistently to how they appear in the DNS scenario tests
		dnsAddress, _ := ei.InterpretString(fmt.Sprintf("sc:dns#%02x", i))
		dnsMap[string(dnsAddress)] = struct{}{}
	}

	return dnsMap
}
