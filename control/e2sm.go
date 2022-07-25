package control

/*
#include <e2sm/wrapper.h>
#cgo LDFLAGS: -le2smwrapper -lm
#cgo CFLAGS: -I/usr/local/include/e2sm
*/
import "C"

import (
	"errors"
	"strconv"
	"unsafe"
)

//type E2sm {}

//type E2Handler interface{}

func ConvertStr2Byte(str string) (val []byte) {
	length := len(str)
	for i := 0; i < length/2; i++ {
		SubStr := str[2*i : 2*i+2]
		v, _ := strconv.ParseUint(SubStr, 16, 8)
		val = append(val, byte(v))
	}
	return
}

func E2smRanFunctionDefinitionDecode(str string) (RanFuncDef *E2SM_KPM_RANfunction_Description, err error) {
	Buffer := ConvertStr2Byte(str)

	cptr := unsafe.Pointer(&Buffer[0])
	RanFuncDef = &E2SM_KPM_RANfunction_Description{}

	// Call E2SM Wrapper to decode
	DecodedRanFuncDef := C.Decode_RAN_Function_Description(cptr, C.size_t(len(Buffer)))
	if DecodedRanFuncDef == nil {
		return RanFuncDef, errors.New("e2sm wrapper is unable to decode RANFunctionDescription due to wrong or invalid input")
	}
	defer C.Free_RAN_Function_Dscription(DecodedRanFuncDef)

	//Parse decoded Ranfunction Definition C structure to Golang Structure
	RanFuncDef.ranFunction_Name = RANfunction_Name{}
	RanFuncDef.ranFunction_Name.ranFunction_ShortName.Buf = C.GoBytes(unsafe.Pointer(DecodedRanFuncDef.ranFunction_Name.ranFunction_ShortName.buf), C.int(DecodedRanFuncDef.ranFunction_Name.ranFunction_ShortName.size))
	RanFuncDef.ranFunction_Name.ranFunction_ShortName.Size = int(DecodedRanFuncDef.ranFunction_Name.ranFunction_ShortName.size)

	RanFuncDef.ranFunction_Name.ranFunction_E2SM_OID.Buf = C.GoBytes(unsafe.Pointer(DecodedRanFuncDef.ranFunction_Name.ranFunction_E2SM_OID.buf), C.int(DecodedRanFuncDef.ranFunction_Name.ranFunction_E2SM_OID.size))
	RanFuncDef.ranFunction_Name.ranFunction_E2SM_OID.Size = int(DecodedRanFuncDef.ranFunction_Name.ranFunction_E2SM_OID.size)

	RanFuncDef.ranFunction_Name.ranFunction_Description.Buf = C.GoBytes(unsafe.Pointer(DecodedRanFuncDef.ranFunction_Name.ranFunction_Description.buf), C.int(DecodedRanFuncDef.ranFunction_Name.ranFunction_Description.size))
	RanFuncDef.ranFunction_Name.ranFunction_Description.Size = int(DecodedRanFuncDef.ranFunction_Name.ranFunction_Description.size)

	if unsafe.Pointer(DecodedRanFuncDef.ranFunction_Name.ranFunction_Instance) != nil {
		RanFuncDef.ranFunction_Name.ranFunction_Instance = (*int64)(DecodedRanFuncDef.ranFunction_Name.ranFunction_Instance)
	}

	if DecodedRanFuncDef.ric_EventTriggerStyle_List != nil {
		RanFuncDef.ric_EventTriggerStyle_List = []RIC_EventTriggerStyle_Item{}
		ric_EventTriggerStyle_Item := RIC_EventTriggerStyle_Item{}

		// Iteratively parse each item to list with i
		for i := 0; i < int(DecodedRanFuncDef.ric_EventTriggerStyle_List.list.count); i++ {
			var sizeof_RIC_EventTriggerStyle_Item_t *C.RIC_EventTriggerStyle_Item_t
			RIC_EventTriggerStyle_Item_C := *(**C.RIC_EventTriggerStyle_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(DecodedRanFuncDef.ric_EventTriggerStyle_List.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_RIC_EventTriggerStyle_Item_t)))

			ric_EventTriggerStyle_Item.ric_EventTriggerStyle_Type = int64(RIC_EventTriggerStyle_Item_C.ric_EventTriggerStyle_Type)

			ric_EventTriggerStyle_Item.ric_EventTriggerStyle_Name.Buf = C.GoBytes(unsafe.Pointer(RIC_EventTriggerStyle_Item_C.ric_EventTriggerStyle_Name.buf), C.int(RIC_EventTriggerStyle_Item_C.ric_EventTriggerStyle_Name.size))
			ric_EventTriggerStyle_Item.ric_EventTriggerStyle_Name.Size = int(RIC_EventTriggerStyle_Item_C.ric_EventTriggerStyle_Name.size)

			ric_EventTriggerStyle_Item.ric_EventTriggerFormat_Type = int64(RIC_EventTriggerStyle_Item_C.ric_EventTriggerFormat_Type)

			//Append
			RanFuncDef.ric_EventTriggerStyle_List = append(RanFuncDef.ric_EventTriggerStyle_List, ric_EventTriggerStyle_Item)
		}
	}

	if DecodedRanFuncDef.ric_ReportStyle_List != nil {
		RanFuncDef.ric_ReportStyle_List = []RIC_ReportStyle_Item{}
		ric_ReportStyle_Item := RIC_ReportStyle_Item{}

		// Iteratively parse each item to list with i
		for i := 0; i < int(DecodedRanFuncDef.ric_ReportStyle_List.list.count); i++ {
			var sizeof_RIC_ReportStyle_Item_t *C.RIC_ReportStyle_Item_t
			RIC_ReportStyle_Item_C := *(**C.RIC_ReportStyle_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(DecodedRanFuncDef.ric_ReportStyle_List.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_RIC_ReportStyle_Item_t)))

			ric_ReportStyle_Item.ric_ReportStyle_Type = int64(RIC_ReportStyle_Item_C.ric_ReportStyle_Type)

			ric_ReportStyle_Item.ric_ReportStyle_Name.Buf = C.GoBytes(unsafe.Pointer(RIC_ReportStyle_Item_C.ric_ReportStyle_Name.buf), C.int(RIC_ReportStyle_Item_C.ric_ReportStyle_Name.size))
			ric_ReportStyle_Item.ric_ReportStyle_Name.Size = int(RIC_ReportStyle_Item_C.ric_ReportStyle_Name.size)

			ric_ReportStyle_Item.ric_ActionFormat_Type = int64(RIC_ReportStyle_Item_C.ric_ActionFormat_Type)

			ric_ReportStyle_Item.ric_IndicationHeaderFormat_Type = int64(RIC_ReportStyle_Item_C.ric_IndicationHeaderFormat_Type)

			ric_ReportStyle_Item.ric_IndicationMessageFormat_Type = int64(RIC_ReportStyle_Item_C.ric_IndicationMessageFormat_Type)

			//Handle measInfo_Action_List
			ric_ReportStyle_Item.measInfo_Action_List = []MeasurementInfo_Action_Item{}
			MeasurementInfo_Action_Item := MeasurementInfo_Action_Item{}

			// Iteratively parse each item to list with j
			for j := 0; j < int(RIC_ReportStyle_Item_C.measInfo_Action_List.list.count); j++ {
				var sizeof_MeasurementInfo_Action_Item_t *C.MeasurementInfo_Action_Item_t
				MeasurementInfo_Action_Item_C := *(**C.MeasurementInfo_Action_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(RIC_ReportStyle_Item_C.measInfo_Action_List.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_MeasurementInfo_Action_Item_t)))

				MeasurementInfo_Action_Item.measName.Buf = C.GoBytes(unsafe.Pointer(MeasurementInfo_Action_Item_C.measName.buf), C.int(MeasurementInfo_Action_Item_C.measName.size))
				MeasurementInfo_Action_Item.measName.Size = int(MeasurementInfo_Action_Item_C.measName.size)

				if MeasurementInfo_Action_Item_C.measID != nil {
					measID := int64(*MeasurementInfo_Action_Item_C.measID)
					MeasurementInfo_Action_Item.measID = &measID
				}

				//Append
				ric_ReportStyle_Item.measInfo_Action_List = append(ric_ReportStyle_Item.measInfo_Action_List, MeasurementInfo_Action_Item)
			}

			//Append
			RanFuncDef.ric_ReportStyle_List = append(RanFuncDef.ric_ReportStyle_List, ric_ReportStyle_Item)
		}
	}
	return
}

func E2smEventTriggerDefinitionEncode(Buffer []byte, Report_Period int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&Buffer[0])

	Size := C.Encode_Event_Trigger_Definition(cptr, C.size_t(len(Buffer)), C.long(Report_Period))
	if Size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to encode EventTriggerDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(Size)+7)/8)
	return
}

func E2smActionDefinitionEncode(Buffer []byte, Report_Period int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&Buffer[0])

	Size := C.Encode_Event_Trigger_Definition(cptr, C.size_t(len(Buffer)), C.long(Report_Period))
	if Size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to encode EventTriggerDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(Size)+7)/8)
	return
}

func E2smIndicationHeaderDecode(Buffer []byte) (IndiHdr *E2SM_KPM_IndicationHeader, err error) {
	cptr := unsafe.Pointer(&Buffer[0])
	IndiHdr = &E2SM_KPM_IndicationHeader{}

	DecodedIndiHdr := C.Decode_Indication_Header(cptr, C.size_t(len(Buffer)))
	if DecodedIndiHdr == nil {
		return IndiHdr, errors.New("e2sm wrapper is unable to decode IndicationHeader due to wrong or invalid input")
	}
	defer C.Free_Indication_Header(DecodedIndiHdr)

	IndiHdr.indicationHeader_FormatType = int32(DecodedIndiHdr.indicationHeader_formats.present)
	if IndiHdr.indicationHeader_FormatType == 1 {
		IndiHdr1 := &E2SM_KPM_IndicationHeader_Format1{}

		E2SM_KPM_IndicationHeader_Format1_C := *(**C.E2SM_KPM_IndicationHeader_Format1_t)(unsafe.Pointer(&DecodedIndiHdr.indicationHeader_formats.choice[0]))

		IndiHdr1.colletStartTime.Buf = C.GoBytes(unsafe.Pointer(E2SM_KPM_IndicationHeader_Format1_C.colletStartTime.buf), C.int(E2SM_KPM_IndicationHeader_Format1_C.colletStartTime.size))

		IndiHdr1.colletStartTime.Size = int(E2SM_KPM_IndicationHeader_Format1_C.colletStartTime.size)

		if E2SM_KPM_IndicationHeader_Format1_C.fileFormatversion != nil {
			fileFormatversion := PrintableString{}

			fileFormatversion.Buf = C.GoBytes(unsafe.Pointer(E2SM_KPM_IndicationHeader_Format1_C.fileFormatversion.buf), C.int(E2SM_KPM_IndicationHeader_Format1_C.fileFormatversion.size))
			fileFormatversion.Size = int(E2SM_KPM_IndicationHeader_Format1_C.fileFormatversion.size)

			IndiHdr1.fileFormatversion = &fileFormatversion
		}

		if E2SM_KPM_IndicationHeader_Format1_C.senderName != nil {
			senderName := PrintableString{}

			senderName.Buf = C.GoBytes(unsafe.Pointer(E2SM_KPM_IndicationHeader_Format1_C.senderName.buf), C.int(E2SM_KPM_IndicationHeader_Format1_C.senderName.size))
			senderName.Size = int(E2SM_KPM_IndicationHeader_Format1_C.senderName.size)

			IndiHdr1.senderName = &senderName
		}

		if E2SM_KPM_IndicationHeader_Format1_C.senderType != nil {
			senderType := PrintableString{}

			senderType.Buf = C.GoBytes(unsafe.Pointer(E2SM_KPM_IndicationHeader_Format1_C.senderType.buf), C.int(E2SM_KPM_IndicationHeader_Format1_C.senderType.size))
			senderType.Size = int(E2SM_KPM_IndicationHeader_Format1_C.senderType.size)

			IndiHdr1.senderType = &senderType
		}

		if E2SM_KPM_IndicationHeader_Format1_C.vendorName != nil {
			vendorName := PrintableString{}

			vendorName.Buf = C.GoBytes(unsafe.Pointer(E2SM_KPM_IndicationHeader_Format1_C.vendorName.buf), C.int(E2SM_KPM_IndicationHeader_Format1_C.vendorName.size))
			vendorName.Size = int(E2SM_KPM_IndicationHeader_Format1_C.vendorName.size)

			IndiHdr1.vendorName = &vendorName
		}

		IndiHdr.indicationHeader_Format = IndiHdr1

	} else {
		return IndiHdr, errors.New("Unknown RIC Indication Header type")
	}
	return
}

func E2smIndicationMessageDecode(Buffer []byte) (IndiMsg *E2SM_KPM_IndicationMessage, err error) {
	cptr := unsafe.Pointer(&Buffer[0])
	IndiMsg = &E2SM_KPM_IndicationMessage{}

	DecodedIndiMsg := C.Decode_Indication_Message(cptr, C.size_t(len(Buffer)))
	if DecodedIndiMsg == nil {
		return IndiMsg, errors.New("e2sm wrapper is unable to decode IndicationMessage due to wrong or invalid input")
	}
	defer C.Free_Indication_Message(DecodedIndiMsg)

	IndiMsg.indicationMessage_FormatsType = int32(DecodedIndiMsg.indicationMessage_formats.present)

	switch IndiMsg.indicationMessage_FormatsType {
	case 1: //Indication Message Format 1
		IndiMsgFmt1 := &E2SM_KPM_IndicationMessage_Format1{}
		E2SM_KPM_IndicationMessage_Format1_C := *(**C.E2SM_KPM_IndicationMessage_Format1_t)(unsafe.Pointer(&DecodedIndiMsg.indicationMessage_formats.choice[0]))

		// Handle MeasurementData
		IndiMsgFmt1.measData = []MeasurementDataItem{}
		measDataItem := MeasurementDataItem{}

		// Iteratively parse each item to list with i
		for i := 0; i < int(E2SM_KPM_IndicationMessage_Format1_C.measData.list.count); i++ {
			var sizeof_MeasurementDataItem_t *C.MeasurementDataItem_t
			MeasurementDataItem_C := *(**C.MeasurementDataItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(E2SM_KPM_IndicationMessage_Format1_C.measData.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_MeasurementDataItem_t)))

			if MeasurementDataItem_C.incompleteFlag != nil {
				incompleteFlag := int64(*MeasurementDataItem_C.incompleteFlag)

				measDataItem.incompleteFlag = &incompleteFlag
			}

			measDataItem.measRecord = []MeasurementRecordItem{}
			var measRecordItem interface{} // MeasurementRecordItem

			// Iteratively parse each item to list with j
			for j := 0; j < int(MeasurementDataItem_C.measRecord.list.count); j++ {
				var sizeof_MeasurementRecordItem_t *C.MeasurementRecordItem_t
				MeasurementRecordItem_C := *(**C.MeasurementRecordItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(MeasurementDataItem_C.measRecord.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_MeasurementRecordItem_t)))

				MeasurementRecordType := int32(MeasurementRecordItem_C.present)
				switch MeasurementRecordType {
				case 1:
					integer := uint64(MeasurementRecordItem_C.choice[0])
					measRecordItem = integer
				case 2:
					real := float64(MeasurementRecordItem_C.choice[1])
					measRecordItem = real
				case 3:
					noValue := int32(MeasurementRecordItem_C.choice[2])
					measRecordItem = noValue
				default:
				}

				measDataItem.measRecord = append(measDataItem.measRecord, measRecordItem)
			}

			IndiMsgFmt1.measData = append(IndiMsgFmt1.measData, measDataItem)
		}

		// Handle MeasurementInfoList *Optional*
		measInfoList := []MeasurementInfoItem{}
		measInfoItem := MeasurementInfoItem{}

		for i := 0; i < int(E2SM_KPM_IndicationMessage_Format1_C.measInfoList.list.count); i++ {
			var sizeof_MeasurementInfoItem_t *C.MeasurementInfoItem_t
			MeasurementInfoItem_C := *(**C.MeasurementInfoItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(E2SM_KPM_IndicationMessage_Format1_C.measInfoList.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_MeasurementInfoItem_t)))

			measType := int32(MeasurementInfoItem_C.measType.present)
			switch measType {
			case 1:
				measName := PrintableString{}

				measName_C := (*C.MeasurementTypeName_t)(unsafe.Pointer(&MeasurementInfoItem_C.measType.choice[0]))

				measName.Buf = C.GoBytes(unsafe.Pointer(measName_C.buf), C.int(measName_C.size))
				measName.Size = int(measName_C.size)

				measInfoItem.measType = measName
			case 2:
				measID := int64(MeasurementInfoItem_C.measType.choice[0])

				measInfoItem.measType = measID
			default:
			}

			measInfoItem.labelInfoList = []LabelInfoItem{}
			LabelInfoItem := LabelInfoItem{}

			for j := 0; j < int(MeasurementInfoItem_C.labelInfoList.list.count); j++ {
				var sizeof_LabelInfoItem_t *C.LabelInfoItem_t
				LabelInfoItem_C := *(**C.LabelInfoItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(MeasurementInfoItem_C.labelInfoList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_LabelInfoItem_t)))

				if LabelInfoItem_C.measLabel.noLabel != nil {
					noLabel := int64(*LabelInfoItem_C.measLabel.noLabel)
					LabelInfoItem.noLabel = &noLabel
				}

				if LabelInfoItem_C.measLabel.plmnID != nil {
					plmnID := OctetString{}

					plmnID.Buf = C.GoBytes(unsafe.Pointer(LabelInfoItem_C.measLabel.plmnID.buf), C.int(LabelInfoItem_C.measLabel.plmnID.size))
					plmnID.Size = int(LabelInfoItem_C.measLabel.plmnID.size)

					LabelInfoItem.plmnID = &plmnID
				}

				if LabelInfoItem_C.measLabel.sliceID != nil {
					LabelInfoItem.sliceID.sST.Buf = C.GoBytes(unsafe.Pointer(LabelInfoItem_C.measLabel.sliceID.sST.buf), C.int(LabelInfoItem_C.measLabel.sliceID.sST.size))
					LabelInfoItem.sliceID.sST.Size = int(LabelInfoItem_C.measLabel.sliceID.sST.size)

					if LabelInfoItem_C.measLabel.sliceID.sD != nil {
						sD := OctetString{}

						sD.Buf = C.GoBytes(unsafe.Pointer(LabelInfoItem_C.measLabel.sliceID.sD.buf), C.int(LabelInfoItem_C.measLabel.sliceID.sD.size))
						sD.Size = int(LabelInfoItem_C.measLabel.sliceID.sD.size)

						LabelInfoItem.sliceID.sD = &sD
					}
				}

				if LabelInfoItem_C.measLabel.fiveQI != nil {
					fiveQI := int64(*LabelInfoItem_C.measLabel.fiveQI)
					LabelInfoItem.fiveQI = &fiveQI
				}

				if LabelInfoItem_C.measLabel.qFI != nil {
					qFI := int64(*LabelInfoItem_C.measLabel.qFI)
					LabelInfoItem.qFI = &qFI
				}

				if LabelInfoItem_C.measLabel.qCI != nil {
					qCI := int64(*LabelInfoItem_C.measLabel.qCI)
					LabelInfoItem.qCI = &qCI
				}

				if LabelInfoItem_C.measLabel.qCImax != nil {
					qCImax := int64(*LabelInfoItem_C.measLabel.qCImax)
					LabelInfoItem.qCImax = &qCImax
				}

				if LabelInfoItem_C.measLabel.qCImin != nil {
					qCImin := int64(*LabelInfoItem_C.measLabel.qCImin)
					LabelInfoItem.noLabel = &qCImin
				}

				if LabelInfoItem_C.measLabel.aRPmax != nil {
					aRPmax := int64(*LabelInfoItem_C.measLabel.aRPmax)
					LabelInfoItem.noLabel = &aRPmax
				}

				if LabelInfoItem_C.measLabel.aRPmin != nil {
					aRPmin := int64(*LabelInfoItem_C.measLabel.aRPmin)
					LabelInfoItem.noLabel = &aRPmin
				}

				if LabelInfoItem_C.measLabel.bitrateRange != nil {
					bitrateRange := int64(*LabelInfoItem_C.measLabel.bitrateRange)
					LabelInfoItem.noLabel = &bitrateRange
				}

				if LabelInfoItem_C.measLabel.layerMU_MIMO != nil {
					layerMU_MIMO := int64(*LabelInfoItem_C.measLabel.layerMU_MIMO)
					LabelInfoItem.noLabel = &layerMU_MIMO
				}

				if LabelInfoItem_C.measLabel.sUM != nil {
					sUM := int64(*LabelInfoItem_C.measLabel.sUM)
					LabelInfoItem.noLabel = &sUM
				}

				if LabelInfoItem_C.measLabel.distBinX != nil {
					distBinX := int64(*LabelInfoItem_C.measLabel.distBinX)
					LabelInfoItem.noLabel = &distBinX
				}

				if LabelInfoItem_C.measLabel.distBinY != nil {
					distBinY := int64(*LabelInfoItem_C.measLabel.distBinY)
					LabelInfoItem.noLabel = &distBinY
				}

				if LabelInfoItem_C.measLabel.distBinZ != nil {
					distBinZ := int64(*LabelInfoItem_C.measLabel.distBinZ)
					LabelInfoItem.noLabel = &distBinZ
				}

				if LabelInfoItem_C.measLabel.preLabelOverride != nil {
					preLabelOverride := int64(*LabelInfoItem_C.measLabel.preLabelOverride)
					LabelInfoItem.noLabel = &preLabelOverride
				}

				if LabelInfoItem_C.measLabel.startEndInd != nil {
					startEndInd := int64(*LabelInfoItem_C.measLabel.startEndInd)
					LabelInfoItem.noLabel = &startEndInd
				}

				if LabelInfoItem_C.measLabel.min != nil {
					min := int64(*LabelInfoItem_C.measLabel.min)
					LabelInfoItem.noLabel = &min
				}

				if LabelInfoItem_C.measLabel.max != nil {
					max := int64(*LabelInfoItem_C.measLabel.max)
					LabelInfoItem.noLabel = &max
				}

				if LabelInfoItem_C.measLabel.avg != nil {
					avg := int64(*LabelInfoItem_C.measLabel.avg)
					LabelInfoItem.noLabel = &avg
				}
				measInfoItem.labelInfoList = append(measInfoItem.labelInfoList, LabelInfoItem)
			}
			measInfoList = append(measInfoList, measInfoItem)
		}
		IndiMsgFmt1.measInfoList = &measInfoList

		// Handle GranularityPeriod *Optional*
		if E2SM_KPM_IndicationMessage_Format1_C.granulPeriod != nil {
			granulPeriod := uint64(*E2SM_KPM_IndicationMessage_Format1_C.granulPeriod)
			IndiMsgFmt1.granulPeriod = &granulPeriod
		}

		IndiMsg.indicationMessage_Formats = IndiMsgFmt1

	case 2: //Indication Message Format 2
		IndiMsgFmt2 := &E2SM_KPM_IndicationMessage_Format2{}
		E2SM_KPM_IndicationMessage_Format2_C := *(**C.E2SM_KPM_IndicationMessage_Format2_t)(unsafe.Pointer(&DecodedIndiMsg.indicationMessage_formats.choice[0]))

		// Handle MeasurementData
		IndiMsgFmt2.measData = []MeasurementDataItem{}
		measDataItem := MeasurementDataItem{}

		// Iteratively parse each item to list with i
		for i := 0; i < int(E2SM_KPM_IndicationMessage_Format2_C.measData.list.count); i++ {
			var sizeof_MeasurementDataItem_t *C.MeasurementDataItem_t
			MeasurementDataItem_C := *(**C.MeasurementDataItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(E2SM_KPM_IndicationMessage_Format2_C.measData.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_MeasurementDataItem_t)))

			if MeasurementDataItem_C.incompleteFlag != nil {
				incompleteFlag := int64(*MeasurementDataItem_C.incompleteFlag)

				measDataItem.incompleteFlag = &incompleteFlag
			}

			measDataItem.measRecord = []MeasurementRecordItem{}
			var measRecordItem interface{}

			// Iteratively parse each item to list with j
			for j := 0; j < int(MeasurementDataItem_C.measRecord.list.count); j++ {
				var sizeof_MeasurementRecordItem_t *C.MeasurementRecordItem_t
				MeasurementRecordItem_C := *(**C.MeasurementRecordItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(MeasurementDataItem_C.measRecord.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_MeasurementRecordItem_t)))

				MeasurementRecordType := int32(MeasurementRecordItem_C.present)
				switch MeasurementRecordType {
				case 1:
					integer := uint64(MeasurementRecordItem_C.choice[0])
					measRecordItem = integer
				case 2:
					real := float64(MeasurementRecordItem_C.choice[1])
					measRecordItem = real
				case 3:
					noValue := int32(MeasurementRecordItem_C.choice[2])
					measRecordItem = noValue
				default:
				}

				measDataItem.measRecord = append(measDataItem.measRecord, measRecordItem)
			}

			IndiMsgFmt2.measData = append(IndiMsgFmt2.measData, measDataItem)
		}

		// Handle measCondUEidList
		IndiMsgFmt2.measCondUEidList = []MeasurementCondUEidItem{}
		measCondUEidItem := MeasurementCondUEidItem{}

		// Iteratively parse each item to list with i
		for i := 0; i < int(E2SM_KPM_IndicationMessage_Format2_C.measCondUEidList.list.count); i++ {
			var sizeof_MeasurementCondUEidItem_t *C.MeasurementCondUEidItem_t
			MeasurementCondUEidItem_C := *(**C.MeasurementCondUEidItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(E2SM_KPM_IndicationMessage_Format2_C.measCondUEidList.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_MeasurementCondUEidItem_t)))

			measType := int32(MeasurementCondUEidItem_C.measType.present)
			switch measType {
			case 1:
				measName := PrintableString{}

				measName_C := (*C.MeasurementTypeName_t)(unsafe.Pointer(&MeasurementCondUEidItem_C.measType.choice[0]))

				measName.Buf = C.GoBytes(unsafe.Pointer(measName_C.buf), C.int(measName_C.size))
				measName.Size = int(measName_C.size)

				measCondUEidItem.measType = measName
			case 2:
				measID := int64(MeasurementCondUEidItem_C.measType.choice[1])

				measCondUEidItem.measType = measID
			default:
			}

			measCondUEidItem.matchingCond = []MatchingCondItem{}
			var matchingCondItem interface{}

			// Iteratively parse each item to list with j
			for j := 0; j < int(MeasurementCondUEidItem_C.matchingCond.list.count); j++ {
				var sizeof_MatchingCondItem_t *C.MatchingCondItem_t
				MatchingCondItem_C := *(**C.MatchingCondItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(MeasurementCondUEidItem_C.matchingCond.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_MatchingCondItem_t)))

				matchingCondItemType := int32(MatchingCondItem_C.present)
				switch matchingCondItemType {
				case 1:
					measLabel := MeasurementLabel{}

					measLabel_C := (*C.MeasurementLabel_t)(unsafe.Pointer(&MatchingCondItem_C.choice[0]))

					if measLabel_C.noLabel != nil {
						noLabel := int64(*measLabel_C.noLabel)
						measLabel.noLabel = &noLabel
					}

					if measLabel_C.plmnID != nil {
						plmnID := OctetString{}

						plmnID_C := (*C.PLMNIdentity_t)(unsafe.Pointer(&measLabel_C.plmnID))

						plmnID.Buf = C.GoBytes(unsafe.Pointer(plmnID_C.buf), C.int(plmnID_C.size))
						plmnID.Size = int(plmnID_C.size)

						measLabel.plmnID = &plmnID
					}

					if measLabel_C.sliceID != nil {

						sliceID_C := (*C.S_NSSAI_t)(unsafe.Pointer(measLabel_C.sliceID))
						measLabel.sliceID.sST.Buf = C.GoBytes(unsafe.Pointer(sliceID_C.sST.buf), C.int(sliceID_C.sST.size))
						measLabel.sliceID.sST.Size = int(sliceID_C.sST.size)

						if sliceID_C.sD != nil {
							sD := OctetString{}

							sD.Buf = C.GoBytes(unsafe.Pointer(sliceID_C.sD.buf), C.int(sliceID_C.sD.size))
							sD.Size = int(sliceID_C.sD.size)

							measLabel.sliceID.sD = &sD
						}
					}

					if measLabel_C.fiveQI != nil {
						fiveQI := int64(*measLabel_C.fiveQI)
						measLabel.fiveQI = &fiveQI
					}

					if measLabel_C.qFI != nil {
						qFI := int64(*measLabel_C.qFI)
						measLabel.qFI = &qFI
					}

					if measLabel_C.qCI != nil {
						qCI := int64(*measLabel_C.qCI)
						measLabel.qCI = &qCI
					}

					if measLabel_C.qCImax != nil {
						qCImax := int64(*measLabel_C.qCImax)
						measLabel.qCImax = &qCImax
					}

					if measLabel_C.qCImin != nil {
						qCImin := int64(*measLabel_C.qCImin)
						measLabel.noLabel = &qCImin
					}

					if measLabel_C.aRPmax != nil {
						aRPmax := int64(*measLabel_C.aRPmax)
						measLabel.noLabel = &aRPmax
					}

					if measLabel_C.aRPmin != nil {
						aRPmin := int64(*measLabel_C.aRPmin)
						measLabel.noLabel = &aRPmin
					}

					if measLabel_C.bitrateRange != nil {
						bitrateRange := int64(*measLabel_C.bitrateRange)
						measLabel.noLabel = &bitrateRange
					}

					if measLabel_C.layerMU_MIMO != nil {
						layerMU_MIMO := int64(*measLabel_C.layerMU_MIMO)
						measLabel.noLabel = &layerMU_MIMO
					}

					if measLabel_C.sUM != nil {
						sUM := int64(*measLabel_C.sUM)
						measLabel.noLabel = &sUM
					}

					if measLabel_C.distBinX != nil {
						distBinX := int64(*measLabel_C.distBinX)
						measLabel.noLabel = &distBinX
					}

					if measLabel_C.distBinY != nil {
						distBinY := int64(*measLabel_C.distBinY)
						measLabel.noLabel = &distBinY
					}

					if measLabel_C.distBinZ != nil {
						distBinZ := int64(*measLabel_C.distBinZ)
						measLabel.noLabel = &distBinZ
					}

					if measLabel_C.preLabelOverride != nil {
						preLabelOverride := int64(*measLabel_C.preLabelOverride)
						measLabel.noLabel = &preLabelOverride
					}

					if measLabel_C.startEndInd != nil {
						startEndInd := int64(*measLabel_C.startEndInd)
						measLabel.noLabel = &startEndInd
					}

					if measLabel_C.min != nil {
						min := int64(*measLabel_C.min)
						measLabel.noLabel = &min
					}

					if measLabel_C.max != nil {
						max := int64(*measLabel_C.max)
						measLabel.noLabel = &max
					}

					if measLabel_C.avg != nil {
						avg := int64(*measLabel_C.avg)
						measLabel.noLabel = &avg
					}

					matchingCondItem = measLabel
				case 2:
					testCondInfo := TestCondInfo{}

					testCondInfo_C := (*C.TestCondInfo_t)(unsafe.Pointer(&MatchingCondItem_C.choice[0]))

					testType := int32(testCondInfo_C.testType.present)
					switch testType {
					case 1:
						testCondInfo.testType = 1 //gBR
					case 2:
						testCondInfo.testType = 2 //aMBR
					case 3:
						testCondInfo.testType = 3 //isStat
					case 4:
						testCondInfo.testType = 4 //isCatM
					case 5:
						testCondInfo.testType = 5 //rSRP
					case 6:
						testCondInfo.testType = 6 //rSRQ
					default:
					}

					testCondInfo.testExpr = int64(testCondInfo_C.testExpr)

					testValueType := int32(testCondInfo_C.testValue.present)
					switch testValueType {
					case 1:
						testCondInfo.testValue = int64(testCondInfo_C.testValue.choice[0])
					case 2:
						testCondInfo.testValue = int64(testCondInfo_C.testValue.choice[1])
					case 3:
						testCondInfo.testValue = int32(testCondInfo_C.testValue.choice[2])
					case 4:
						valueBitS := BitString{}

						valueBitS_C := (*C.BIT_STRING_t)(unsafe.Pointer(&testCondInfo_C.testValue.choice[0]))

						valueBitS.Buf = C.GoBytes(unsafe.Pointer(valueBitS_C.buf), C.int(valueBitS_C.size))
						valueBitS.Size = int(valueBitS_C.size)
						valueBitS.BitsUnused = int(valueBitS_C.bits_unused)

						testCondInfo.testValue = valueBitS
					case 5:
						valueOctS := OctetString{}

						valueOctS_C := (*C.OCTET_STRING_t)(unsafe.Pointer(&testCondInfo_C.testValue.choice[0]))

						valueOctS.Buf = C.GoBytes(unsafe.Pointer(valueOctS_C.buf), C.int(valueOctS_C.size))
						valueOctS.Size = int(valueOctS_C.size)

						testCondInfo.testValue = valueOctS
					case 6:
						valuePrtS := PrintableString{}

						valuePrtS_C := (*C.PrintableString_t)(unsafe.Pointer(&testCondInfo_C.testValue.choice[0]))

						valuePrtS.Buf = C.GoBytes(unsafe.Pointer(valuePrtS_C.buf), C.int(valuePrtS_C.size))
						valuePrtS.Size = int(valuePrtS_C.size)

						testCondInfo.testValue = valuePrtS
					default:
					}

					matchingCondItem = testCondInfo
				default:
				}
				measCondUEidItem.matchingCond = append(measCondUEidItem.matchingCond, matchingCondItem)
			}

			if MeasurementCondUEidItem_C.matchingUEidList != nil {
				matchingUEidList := []MatchingUEidItem{}
				matchingUEidItem := MatchingUEidItem{}

				// Iteratively parse each item to list with j
				for j := 0; j < int(MeasurementCondUEidItem_C.matchingUEidList.list.count); j++ {
					var sizeof_MatchingUEidItem_t *C.MatchingUEidItem_t
					//Try
					MatchingUEidItem_C := *(**C.MatchingUEidItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(MeasurementCondUEidItem_C.matchingUEidList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_MatchingUEidItem_t)))
					//MatchingUEidItem_C := unsafe.Pointer(uintptr(unsafe.Pointer(MeasurementCondUEidItem_C.matchingUEidList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_MatchingUEidItem_t))
					matchingUEidItem.ueID = ParseUeId(unsafe.Pointer(&MatchingUEidItem_C.ueID))
					matchingUEidList = append(matchingUEidList, matchingUEidItem)
				}
				measCondUEidItem.matchingUEidList = &matchingUEidList
			}

			IndiMsgFmt2.measCondUEidList = append(IndiMsgFmt2.measCondUEidList, measCondUEidItem)
		}

		// Handle GranularityPeriod *Optional*
		if E2SM_KPM_IndicationMessage_Format2_C.granulPeriod != nil {
			granulPeriod := uint64(*E2SM_KPM_IndicationMessage_Format2_C.granulPeriod)
			IndiMsgFmt2.granulPeriod = &granulPeriod
		}

		IndiMsg.indicationMessage_Formats = IndiMsgFmt2

	case 3: //Indication Message Format 3
		IndiMsgFmt3 := &E2SM_KPM_IndicationMessage_Format3{}
		E2SM_KPM_IndicationMessage_Format3_C := *(**C.E2SM_KPM_IndicationMessage_Format3_t)(unsafe.Pointer(&DecodedIndiMsg.indicationMessage_formats.choice[2]))

		IndiMsgFmt3.ueMeasReportList = []UEMeasurementReportItem{}
		ueMeasReportItem := UEMeasurementReportItem{}

		// Iteratively parse each item to list with i
		for i := 0; i < int(E2SM_KPM_IndicationMessage_Format3_C.ueMeasReportList.list.count); i++ {
			var sizeof_UEMeasurementReportItem_t *C.UEMeasurementReportItem_t
			UEMeasurementReportItem_C := *(**C.UEMeasurementReportItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(E2SM_KPM_IndicationMessage_Format3_C.ueMeasReportList.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_UEMeasurementReportItem_t)))

			//Handle Indication Message Format 1
			IndiMsgFmt1 := E2SM_KPM_IndicationMessage_Format1{}
			E2SM_KPM_IndicationMessage_Format1_C := *(**C.E2SM_KPM_IndicationMessage_Format1_t)(unsafe.Pointer(&UEMeasurementReportItem_C.measReport))

			IndiMsgFmt1.measData = []MeasurementDataItem{}
			measDataItem := MeasurementDataItem{}

			// Iteratively parse each item to list with j
			for j := 0; j < int(E2SM_KPM_IndicationMessage_Format1_C.measData.list.count); j++ {
				var sizeof_MeasurementDataItem_t *C.MeasurementDataItem_t
				MeasurementDataItem_C := *(**C.MeasurementDataItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(E2SM_KPM_IndicationMessage_Format1_C.measData.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_MeasurementDataItem_t)))

				if MeasurementDataItem_C.incompleteFlag != nil {
					incompleteFlag := int64(*MeasurementDataItem_C.incompleteFlag)

					measDataItem.incompleteFlag = &incompleteFlag
				}

				measDataItem.measRecord = []MeasurementRecordItem{}
				var measRecordItem interface{}

				// Iteratively parse each item to list with k
				for k := 0; k < int(MeasurementDataItem_C.measRecord.list.count); k++ {
					var sizeof_MeasurementRecordItem_t *C.MeasurementRecordItem_t
					MeasurementRecordItem_C := *(**C.MeasurementRecordItem_t)(unsafe.Pointer(uintptr(unsafe.Pointer(MeasurementDataItem_C.measRecord.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_MeasurementRecordItem_t)))

					MeasurementRecordType := int32(MeasurementRecordItem_C.present)
					switch MeasurementRecordType {
					case 1:
						integer := uint64(MeasurementRecordItem_C.choice[0])
						measRecordItem = integer
					case 2:
						real := float64(MeasurementRecordItem_C.choice[1])
						measRecordItem = real
					case 3:
						noValue := int32(MeasurementRecordItem_C.choice[2])
						measRecordItem = noValue
					default:
					}

					measDataItem.measRecord = append(measDataItem.measRecord, measRecordItem)
				}

				IndiMsgFmt1.measData = append(IndiMsgFmt1.measData, measDataItem)
			}

			ueMeasReportItem.measReport = IndiMsgFmt1

			//Handle UEID Parsing
			ueMeasReportItem.ueID = ParseUeId(unsafe.Pointer(&UEMeasurementReportItem_C.ueID))

			IndiMsgFmt3.ueMeasReportList = append(IndiMsgFmt3.ueMeasReportList, ueMeasReportItem)
		}

		IndiMsg.indicationMessage_Formats = IndiMsgFmt3

	default: //Mot supported //Indication Message Format
		return IndiMsg, errors.New("Unknown RIC Indication Message Format")
	}
	return

}

func ParseUeId(ptr unsafe.Pointer) (ueID UEID) {
	UEID_C := (*C.UEID_t)(ptr)

	UEIdType := int32(UEID_C.present)
	switch UEIdType {
	case 1: //gNB_UEID
		gNB_UEID := &UEID_GNB{}
		UEID_GNB_C := *(**C.UEID_GNB_t)(unsafe.Pointer(&UEID_C.choice[0]))

		gNB_UEID.amf_UE_NGAP_ID.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_C.amf_UE_NGAP_ID.buf), C.int(UEID_GNB_C.amf_UE_NGAP_ID.size))
		gNB_UEID.amf_UE_NGAP_ID.Size = int(UEID_GNB_C.amf_UE_NGAP_ID.size)

		gNB_UEID.guami.pLMNIdentity.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_C.guami.pLMNIdentity.buf), C.int(UEID_GNB_C.guami.pLMNIdentity.size))
		gNB_UEID.guami.pLMNIdentity.Size = int(UEID_GNB_C.guami.pLMNIdentity.size)

		gNB_UEID.guami.aMFRegionID.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_C.guami.aMFRegionID.buf), C.int(UEID_GNB_C.guami.aMFRegionID.size))
		gNB_UEID.guami.aMFRegionID.Size = int(UEID_GNB_C.guami.aMFRegionID.size)
		gNB_UEID.guami.aMFRegionID.BitsUnused = int(UEID_GNB_C.guami.aMFRegionID.bits_unused)

		gNB_UEID.guami.aMFSetID.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_C.guami.aMFSetID.buf), C.int(UEID_GNB_C.guami.aMFSetID.size))
		gNB_UEID.guami.aMFSetID.Size = int(UEID_GNB_C.guami.aMFSetID.size)
		gNB_UEID.guami.aMFSetID.BitsUnused = int(UEID_GNB_C.guami.aMFSetID.bits_unused)

		gNB_UEID.guami.aMFPointer.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_C.guami.aMFPointer.buf), C.int(UEID_GNB_C.guami.aMFPointer.size))
		gNB_UEID.guami.aMFPointer.Size = int(UEID_GNB_C.guami.aMFPointer.size)
		gNB_UEID.guami.aMFPointer.BitsUnused = int(UEID_GNB_C.guami.aMFPointer.bits_unused)

		if UEID_GNB_C.gNB_CU_UE_F1AP_ID_List != nil {
			gNB_CU_UE_F1AP_ID_List := []UEID_GNB_CU_CP_F1AP_ID_Item{}
			gNB_CU_UE_F1AP_ID_Item := UEID_GNB_CU_CP_F1AP_ID_Item{}

			// Iteratively parse each item to list with j
			for j := 0; j < int(UEID_GNB_C.gNB_CU_UE_F1AP_ID_List.list.count); j++ {
				var sizeof_UEID_GNB_CU_CP_F1AP_ID_Item_t *C.UEID_GNB_CU_CP_F1AP_ID_Item_t
				UEID_GNB_CU_CP_F1AP_ID_Item_C := *(**C.UEID_GNB_CU_CP_F1AP_ID_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(UEID_GNB_C.gNB_CU_UE_F1AP_ID_List.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_UEID_GNB_CU_CP_F1AP_ID_Item_t)))

				gNB_CU_UE_F1AP_ID_Item.gNB_CU_UE_F1AP_ID = uint64(UEID_GNB_CU_CP_F1AP_ID_Item_C.gNB_CU_UE_F1AP_ID)
				gNB_CU_UE_F1AP_ID_List = append(gNB_CU_UE_F1AP_ID_List, gNB_CU_UE_F1AP_ID_Item)
			}

			gNB_UEID.gNB_CU_UE_F1AP_ID_List = &gNB_CU_UE_F1AP_ID_List
		}

		if UEID_GNB_C.gNB_CU_CP_UE_E1AP_ID_List != nil {
			gNB_CU_CP_UE_E1AP_ID_List := []UEID_GNB_CU_CP_E1AP_ID_Item{}
			gNB_CU_CP_UE_E1AP_ID_Item := UEID_GNB_CU_CP_E1AP_ID_Item{}

			// Iteratively parse each item to list with j
			for j := 0; j < int(UEID_GNB_C.gNB_CU_CP_UE_E1AP_ID_List.list.count); j++ {
				var sizeof_UEID_GNB_CU_CP_E1AP_ID_Item_t *C.UEID_GNB_CU_CP_E1AP_ID_Item_t
				UEID_GNB_CU_CP_E1AP_ID_Item_C := *(**C.UEID_GNB_CU_CP_E1AP_ID_Item_t)(unsafe.Pointer(uintptr(unsafe.Pointer(UEID_GNB_C.gNB_CU_CP_UE_E1AP_ID_List.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_UEID_GNB_CU_CP_E1AP_ID_Item_t)))

				gNB_CU_CP_UE_E1AP_ID_Item.gNB_CU_CP_UE_E1AP_ID = uint64(UEID_GNB_CU_CP_E1AP_ID_Item_C.gNB_CU_CP_UE_E1AP_ID)
				gNB_CU_CP_UE_E1AP_ID_List = append(gNB_CU_CP_UE_E1AP_ID_List, gNB_CU_CP_UE_E1AP_ID_Item)
			}

			gNB_UEID.gNB_CU_CP_UE_E1AP_ID_List = &gNB_CU_CP_UE_E1AP_ID_List
		}

		if UEID_GNB_C.ran_UEID != nil {
			RANUEID := OctetString{}

			RANUEID.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_C.ran_UEID.buf), C.int(UEID_GNB_C.ran_UEID.size))
			RANUEID.Size = int(UEID_GNB_C.ran_UEID.size)

			gNB_UEID.ran_UEID = &RANUEID
		}

		if UEID_GNB_C.m_NG_RAN_UE_XnAP_ID != nil {
			m_NG_RAN_UE_XnAP_ID := uint64(*UEID_GNB_C.m_NG_RAN_UE_XnAP_ID)

			gNB_UEID.m_NG_RAN_UE_XnAP_ID = &m_NG_RAN_UE_XnAP_ID
		}

		if UEID_GNB_C.globalGNB_ID != nil {
			globalGNB_ID := GlobalGNB_ID{}

			globalGNB_ID.pLMNIdentity.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_C.globalGNB_ID.pLMNIdentity.buf), C.int(UEID_GNB_C.globalGNB_ID.pLMNIdentity.size))
			globalGNB_ID.pLMNIdentity.Size = int(UEID_GNB_C.globalGNB_ID.pLMNIdentity.size)

			gNB_ID_Type := int32(UEID_GNB_C.globalGNB_ID.gNB_ID.present)
			if gNB_ID_Type == 1 {
				gNB_ID := GNB_ID{}

				gNB_ID_C := (*C.BIT_STRING_t)(unsafe.Pointer(&UEID_GNB_C.globalGNB_ID.gNB_ID.choice[0]))

				gNB_ID.gNB_ID.Buf = C.GoBytes(unsafe.Pointer(gNB_ID_C.buf), C.int(gNB_ID_C.size))
				gNB_ID.gNB_ID.Size = int(gNB_ID_C.size)
				gNB_ID.gNB_ID.BitsUnused = int(gNB_ID_C.bits_unused)

				globalGNB_ID.gNB_ID = &gNB_ID
			}

			gNB_UEID.globalGNB_ID = &globalGNB_ID
		}

		if UEID_GNB_C.globalNG_RANNode_ID != nil {
			var globalNG_RANNode_ID interface{}

			globalNG_RANNode_Type := int32(UEID_GNB_C.globalNG_RANNode_ID.present)
			switch globalNG_RANNode_Type {
			case 1:
				gNB := GlobalGNB_ID{}
				GlobalGNB_ID_C := *(**C.GlobalGNB_ID_t)(unsafe.Pointer(&UEID_GNB_C.globalNG_RANNode_ID.choice[0]))

				gNB.pLMNIdentity.Buf = C.GoBytes(unsafe.Pointer(GlobalGNB_ID_C.pLMNIdentity.buf), C.int(GlobalGNB_ID_C.pLMNIdentity.size))
				gNB.pLMNIdentity.Size = int(GlobalGNB_ID_C.pLMNIdentity.size)

				gNB_ID_Type := int32(GlobalGNB_ID_C.gNB_ID.present)
				if gNB_ID_Type == 1 {
					gNB_ID := &GNB_ID{}

					gNB_ID_C := (*C.BIT_STRING_t)(unsafe.Pointer(&GlobalGNB_ID_C.gNB_ID.choice[0]))

					gNB_ID.gNB_ID.Buf = C.GoBytes(unsafe.Pointer(gNB_ID_C.buf), C.int(gNB_ID_C.size))
					gNB_ID.gNB_ID.Size = int(gNB_ID_C.size)
					gNB_ID.gNB_ID.BitsUnused = int(gNB_ID_C.bits_unused)

					gNB.gNB_ID = gNB_ID
				}

				globalNG_RANNode_ID = gNB

			case 2:
				ng_eNB := &GlobalNgENB_ID{}
				GlobalNgENB_ID_C := *(**C.GlobalNgENB_ID_t)(unsafe.Pointer(&UEID_GNB_C.globalNG_RANNode_ID.choice[1]))

				ng_eNB.pLMNIdentity.Buf = C.GoBytes(unsafe.Pointer(GlobalNgENB_ID_C.pLMNIdentity.buf), C.int(GlobalNgENB_ID_C.pLMNIdentity.size))
				ng_eNB.pLMNIdentity.Size = int(GlobalNgENB_ID_C.pLMNIdentity.size)

				ngENB_ID_Type := int32(GlobalNgENB_ID_C.ngENB_ID.present)
				switch ngENB_ID_Type {
				case 1:
					macroNgENB_ID := &BitString{}

					macroNgENB_ID_C := (*C.BIT_STRING_t)(unsafe.Pointer(&GlobalNgENB_ID_C.ngENB_ID.choice[0]))

					macroNgENB_ID.Buf = C.GoBytes(unsafe.Pointer(macroNgENB_ID_C.buf), C.int(macroNgENB_ID_C.size))
					macroNgENB_ID.Size = int(macroNgENB_ID_C.size)
					macroNgENB_ID.BitsUnused = int(macroNgENB_ID_C.bits_unused)

					ng_eNB.ngENB_ID = macroNgENB_ID

				case 2:
					shortMacroNgENB_ID := &BitString{}

					shortMacroNgENB_ID_C := (*C.BIT_STRING_t)(unsafe.Pointer(&GlobalNgENB_ID_C.ngENB_ID.choice[0]))

					shortMacroNgENB_ID.Buf = C.GoBytes(unsafe.Pointer(shortMacroNgENB_ID_C.buf), C.int(shortMacroNgENB_ID_C.size))
					shortMacroNgENB_ID.Size = int(shortMacroNgENB_ID_C.size)
					shortMacroNgENB_ID.BitsUnused = int(shortMacroNgENB_ID_C.bits_unused)

					ng_eNB.ngENB_ID = shortMacroNgENB_ID

				case 3:
					longMacroNgENB_ID := &BitString{}

					longMacroNgENB_ID_C := (*C.BIT_STRING_t)(unsafe.Pointer(&GlobalNgENB_ID_C.ngENB_ID.choice[0]))

					longMacroNgENB_ID.Buf = C.GoBytes(unsafe.Pointer(longMacroNgENB_ID_C.buf), C.int(longMacroNgENB_ID_C.size))
					longMacroNgENB_ID.Size = int(longMacroNgENB_ID_C.size)
					longMacroNgENB_ID.BitsUnused = int(longMacroNgENB_ID_C.bits_unused)

					ng_eNB.ngENB_ID = longMacroNgENB_ID

				default:
				}

				globalNG_RANNode_ID = ng_eNB

			default:
			}
			gNB_UEID.globalNG_RANNode_ID = &globalNG_RANNode_ID
		}
		ueID = gNB_UEID

	case 2: //gNB_DU_UEID
		gNB_DU_UEID := &UEID_GNB_DU{}
		UEID_GNB_DU_C := *(**C.UEID_GNB_DU_t)(unsafe.Pointer(&UEID_C.choice[1]))

		gNB_DU_UEID.gNB_CU_UE_F1AP_ID = uint64(UEID_GNB_DU_C.gNB_CU_UE_F1AP_ID)

		if UEID_GNB_DU_C.ran_UEID != nil {
			RANUEID := &OctetString{}

			RANUEID.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_DU_C.ran_UEID.buf), C.int(UEID_GNB_DU_C.ran_UEID.size))
			RANUEID.Size = int(UEID_GNB_DU_C.ran_UEID.size)

			gNB_DU_UEID.ran_UEID = RANUEID
		}
		ueID = gNB_DU_UEID

	case 3: //gNB_CU_UP_UEID
		gNB_CU_UP_UEID := &UEID_GNB_CU_UP{}
		UEID_GNB_CU_UP_C := *(**C.UEID_GNB_CU_UP_t)(unsafe.Pointer(&UEID_C.choice[2]))

		gNB_CU_UP_UEID.gNB_CU_CP_UE_E1AP_ID = uint64(UEID_GNB_CU_UP_C.gNB_CU_CP_UE_E1AP_ID)

		if UEID_GNB_CU_UP_C.ran_UEID != nil {
			RANUEID := &OctetString{}

			RANUEID.Buf = C.GoBytes(unsafe.Pointer(UEID_GNB_CU_UP_C.ran_UEID.buf), C.int(UEID_GNB_CU_UP_C.ran_UEID.size))
			RANUEID.Size = int(UEID_GNB_CU_UP_C.ran_UEID.size)

			gNB_CU_UP_UEID.ran_UEID = RANUEID
		}
		ueID = gNB_CU_UP_UEID

	case 4: //Todo: ng_eNB_UEID
	case 5: //Todo: ng_eNB_DU_UEID
	case 6: //Todo: en_gNB_UEID
	case 7: //Todo: eNB_UEID
	default:
	}
	return
}
