package scenmodel

import (
	"github.com/subrahamanyam341/andes-core-16/core"
	txDataBuilder "github.com/subrahamanyam341/andes-vm-common-123/txDataBuilder"
)

// CreateMultiTransferData builds data for a multiTransferDCT
func CreateMultiTransferData(to []byte, dctData []*DCTTxData, endpointName string, arguments [][]byte) []byte {
	multiTransferData := make([]byte, 0)
	multiTransferData = append(multiTransferData, []byte(core.BuiltInFunctionMultiDCTNFTTransfer)...)
	tdb := txDataBuilder.NewBuilder()
	tdb.Bytes(to)
	tdb.Int(len(dctData))

	for _, dctDataTransfer := range dctData {
		tdb.Bytes(dctDataTransfer.TokenIdentifier.Value)
		tdb.Int64(int64(dctDataTransfer.Nonce.Value))
		tdb.BigInt(dctDataTransfer.Value.Value)
	}

	if len(endpointName) > 0 {
		tdb.Str(endpointName)

		for _, arg := range arguments {
			tdb.Bytes(arg)
		}
	}
	multiTransferData = append(multiTransferData, tdb.ToBytes()...)
	return multiTransferData
}
