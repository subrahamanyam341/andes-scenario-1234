package scenjsonwrite

import (
	oj "github.com/subrahamanyam341/andes-scenario-1234/orderedjson"
	scenmodel "github.com/subrahamanyam341/andes-scenario-1234/scenario/model"
)

func dctTxDataToOJ(dctItems []*scenmodel.DCTTxData) oj.OJsonObject {
	dctItemList := oj.OJsonList{}
	for _, dctItemRaw := range dctItems {
		dctItemOJ := dctTxRawEntryToOJ(dctItemRaw)
		dctItemList = append(dctItemList, dctItemOJ)
	}

	return &dctItemList

}

func dctTxRawEntryToOJ(dctItemRaw *scenmodel.DCTTxData) *oj.OJsonMap {
	dctItemOJ := oj.NewMap()

	if len(dctItemRaw.TokenIdentifier.Original) > 0 {
		dctItemOJ.Put("tokenIdentifier", bytesFromStringToOJ(dctItemRaw.TokenIdentifier))
	}
	if len(dctItemRaw.Nonce.Original) > 0 {
		dctItemOJ.Put("nonce", uint64ToOJ(dctItemRaw.Nonce))
	}
	if len(dctItemRaw.Value.Original) > 0 {
		dctItemOJ.Put("value", bigIntToOJ(dctItemRaw.Value))
	}

	return dctItemOJ
}

func dctDataToOJ(dctItems []*scenmodel.DCTData) *oj.OJsonMap {
	dctItemsOJ := oj.NewMap()
	for _, dctItem := range dctItems {
		dctItemsOJ.Put(dctItem.TokenIdentifier.Original, dctItemToOJ(dctItem))
	}
	return dctItemsOJ
}

func dctItemToOJ(dctItem *scenmodel.DCTData) oj.OJsonObject {
	if isCompactDCT(dctItem) {
		return bigIntToOJ(dctItem.Instances[0].Balance)
	}

	dctItemOJ := oj.NewMap()

	// instances
	if len(dctItem.Instances) > 0 {
		var convertedList []oj.OJsonObject
		for _, dctInstance := range dctItem.Instances {
			dctInstanceOJ := oj.NewMap()
			appendDCTInstanceToOJ(dctInstance, dctInstanceOJ)
			convertedList = append(convertedList, dctInstanceOJ)
		}
		instancesOJList := oj.OJsonList(convertedList)
		dctItemOJ.Put("instances", &instancesOJList)
	}

	if len(dctItem.LastNonce.Original) > 0 {
		dctItemOJ.Put("lastNonce", uint64ToOJ(dctItem.LastNonce))
	}

	// roles
	if len(dctItem.Roles) > 0 {
		var convertedList []oj.OJsonObject
		for _, roleStr := range dctItem.Roles {
			convertedList = append(convertedList, &oj.OJsonString{Value: roleStr})
		}
		rolesOJList := oj.OJsonList(convertedList)
		dctItemOJ.Put("roles", &rolesOJList)
	}
	if len(dctItem.Frozen.Original) > 0 {
		dctItemOJ.Put("frozen", uint64ToOJ(dctItem.Frozen))
	}

	return dctItemOJ
}

func appendDCTInstanceToOJ(dctInstance *scenmodel.DCTInstance, targetOj *oj.OJsonMap) {
	targetOj.Put("nonce", uint64ToOJ(dctInstance.Nonce))

	if len(dctInstance.Balance.Original) > 0 {
		targetOj.Put("balance", bigIntToOJ(dctInstance.Balance))
	}
	if len(dctInstance.Creator.Original) > 0 {
		targetOj.Put("creator", bytesFromStringToOJ(dctInstance.Creator))
	}
	if len(dctInstance.Royalties.Original) > 0 {
		targetOj.Put("royalties", uint64ToOJ(dctInstance.Royalties))
	}
	if len(dctInstance.Hash.Original) > 0 {
		targetOj.Put("hash", bytesFromStringToOJ(dctInstance.Hash))
	}
	if !dctInstance.Uris.IsUnspecified() {
		targetOj.Put("uri", valueListToOJ(dctInstance.Uris))
	}
	if len(dctInstance.Attributes.Value) > 0 {
		targetOj.Put("attributes", bytesFromTreeToOJ(dctInstance.Attributes))
	}
}

func isCompactDCT(dctItem *scenmodel.DCTData) bool {
	if len(dctItem.Instances) != 1 {
		return false
	}
	if len(dctItem.Instances[0].Nonce.Original) > 0 {
		return false
	}
	if len(dctItem.Roles) > 0 {
		return false
	}
	if len(dctItem.Frozen.Original) > 0 {
		return false
	}
	return true
}
