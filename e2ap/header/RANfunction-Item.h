/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-CommonDataTypes"
 * 	found in "O-RAN.WG3.E2APv2.01.asn1"
 * 	`asn1c -pdu=auto -fcompound-names -fno-include-deps -findirect-choice -gen-PER -gen-OER -no-gen-example`
 */

#ifndef	_RANfunction_Item_H_
#define	_RANfunction_Item_H_


#include <asn_application.h>

/* Including external dependencies */
#include "RANfunctionID.h"
#include "RANfunctionDefinition.h"
#include "RANfunctionRevision.h"
#include "RANfunctionOID.h"
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* RANfunction-Item */
typedef struct RANfunction_Item {
	RANfunctionID_t	 ranFunctionID;
	RANfunctionDefinition_t	 ranFunctionDefinition;
	RANfunctionRevision_t	 ranFunctionRevision;
	RANfunctionOID_t	 ranFunctionOID;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RANfunction_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RANfunction_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_RANfunction_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_RANfunction_Item_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _RANfunction_Item_H_ */
#include <asn_internal.h>