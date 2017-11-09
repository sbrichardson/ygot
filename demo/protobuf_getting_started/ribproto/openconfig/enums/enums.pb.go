// Code generated by protoc-gen-go.
// source: ribproto/openconfig/enums/enums.proto
// DO NOT EDIT!

/*
Package openconfig_enums is a generated protocol buffer package.

It is generated from these files:
	ribproto/openconfig/enums/enums.proto

It has these top-level messages:
*/
package openconfig_enums

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/openconfig/ygot/proto/ywrapper"
import _ "github.com/openconfig/ygot/proto/yext"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OpenconfigBgpTypesAFISAFITYPE represents an enumerated type generated for the YANG identity AFI_SAFI_TYPE.
type OpenconfigBgpTypesAFISAFITYPE int32

const (
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_UNSET                OpenconfigBgpTypesAFISAFITYPE = 0
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV4_MULTICAST OpenconfigBgpTypesAFISAFITYPE = 53601692
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV6_MULTICAST OpenconfigBgpTypesAFISAFITYPE = 171618750
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_IPV4_UNICAST         OpenconfigBgpTypesAFISAFITYPE = 196909421
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_L2VPN_EVPN           OpenconfigBgpTypesAFISAFITYPE = 213517208
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_IPV6_UNICAST         OpenconfigBgpTypesAFISAFITYPE = 270342995
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_IPV4_LABELED_UNICAST OpenconfigBgpTypesAFISAFITYPE = 279108253
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV4_UNICAST   OpenconfigBgpTypesAFISAFITYPE = 358537365
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_IPV6_LABELED_UNICAST OpenconfigBgpTypesAFISAFITYPE = 420698035
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV6_UNICAST   OpenconfigBgpTypesAFISAFITYPE = 421890363
	OpenconfigBgpTypesAFISAFITYPE_OPENCONFIGBGPTYPESAFISAFITYPE_L2VPN_VPLS           OpenconfigBgpTypesAFISAFITYPE = 533405094
)

var OpenconfigBgpTypesAFISAFITYPE_name = map[int32]string{
	0:         "OPENCONFIGBGPTYPESAFISAFITYPE_UNSET",
	53601692:  "OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV4_MULTICAST",
	171618750: "OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV6_MULTICAST",
	196909421: "OPENCONFIGBGPTYPESAFISAFITYPE_IPV4_UNICAST",
	213517208: "OPENCONFIGBGPTYPESAFISAFITYPE_L2VPN_EVPN",
	270342995: "OPENCONFIGBGPTYPESAFISAFITYPE_IPV6_UNICAST",
	279108253: "OPENCONFIGBGPTYPESAFISAFITYPE_IPV4_LABELED_UNICAST",
	358537365: "OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV4_UNICAST",
	420698035: "OPENCONFIGBGPTYPESAFISAFITYPE_IPV6_LABELED_UNICAST",
	421890363: "OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV6_UNICAST",
	533405094: "OPENCONFIGBGPTYPESAFISAFITYPE_L2VPN_VPLS",
}
var OpenconfigBgpTypesAFISAFITYPE_value = map[string]int32{
	"OPENCONFIGBGPTYPESAFISAFITYPE_UNSET":                0,
	"OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV4_MULTICAST": 53601692,
	"OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV6_MULTICAST": 171618750,
	"OPENCONFIGBGPTYPESAFISAFITYPE_IPV4_UNICAST":         196909421,
	"OPENCONFIGBGPTYPESAFISAFITYPE_L2VPN_EVPN":           213517208,
	"OPENCONFIGBGPTYPESAFISAFITYPE_IPV6_UNICAST":         270342995,
	"OPENCONFIGBGPTYPESAFISAFITYPE_IPV4_LABELED_UNICAST": 279108253,
	"OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV4_UNICAST":   358537365,
	"OPENCONFIGBGPTYPESAFISAFITYPE_IPV6_LABELED_UNICAST": 420698035,
	"OPENCONFIGBGPTYPESAFISAFITYPE_L3VPN_IPV6_UNICAST":   421890363,
	"OPENCONFIGBGPTYPESAFISAFITYPE_L2VPN_VPLS":           533405094,
}

func (x OpenconfigBgpTypesAFISAFITYPE) String() string {
	return proto.EnumName(OpenconfigBgpTypesAFISAFITYPE_name, int32(x))
}
func (OpenconfigBgpTypesAFISAFITYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

// OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY represents an enumerated type generated for the YANG identity BGP_WELL_KNOWN_STD_COMMUNITY.
type OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY int32

const (
	OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_UNSET               OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY = 0
	OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_EXPORT           OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY = 7425995
	OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_EXPORT_SUBCONFED OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY = 75227749
	OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NOPEER              OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY = 77476850
	OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_ADVERTISE        OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY = 126797620
)

var OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_name = map[int32]string{
	0:         "OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_UNSET",
	7425995:   "OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_EXPORT",
	75227749:  "OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_EXPORT_SUBCONFED",
	77476850:  "OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NOPEER",
	126797620: "OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_ADVERTISE",
}
var OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_value = map[string]int32{
	"OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_UNSET":               0,
	"OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_EXPORT":           7425995,
	"OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_EXPORT_SUBCONFED": 75227749,
	"OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NOPEER":              77476850,
	"OPENCONFIGBGPTYPESBGPWELLKNOWNSTDCOMMUNITY_NO_ADVERTISE":        126797620,
}

func (x OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY) String() string {
	return proto.EnumName(OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_name, int32(x))
}
func (OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

// OpenconfigPolicyTypesINSTALLPROTOCOLTYPE represents an enumerated type generated for the YANG identity INSTALL_PROTOCOL_TYPE.
type OpenconfigPolicyTypesINSTALLPROTOCOLTYPE int32

const (
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_UNSET              OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 0
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_STATIC             OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 164198026
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_LOCAL_AGGREGATE    OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 167243989
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_OSPF               OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 187006448
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_ISIS               OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 308545656
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_BGP                OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 436650411
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_DIRECTLY_CONNECTED OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 463126862
	OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_OSPF3              OpenconfigPolicyTypesINSTALLPROTOCOLTYPE = 470106339
)

var OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_name = map[int32]string{
	0:         "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_UNSET",
	164198026: "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_STATIC",
	167243989: "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_LOCAL_AGGREGATE",
	187006448: "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_OSPF",
	308545656: "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_ISIS",
	436650411: "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_BGP",
	463126862: "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_DIRECTLY_CONNECTED",
	470106339: "OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_OSPF3",
}
var OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_value = map[string]int32{
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_UNSET":              0,
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_STATIC":             164198026,
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_LOCAL_AGGREGATE":    167243989,
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_OSPF":               187006448,
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_ISIS":               308545656,
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_BGP":                436650411,
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_DIRECTLY_CONNECTED": 463126862,
	"OPENCONFIGPOLICYTYPESINSTALLPROTOCOLTYPE_OSPF3":              470106339,
}

func (x OpenconfigPolicyTypesINSTALLPROTOCOLTYPE) String() string {
	return proto.EnumName(OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_name, int32(x))
}
func (OpenconfigPolicyTypesINSTALLPROTOCOLTYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

// OpenconfigRibBgpAsPathSegmentType represents an enumerated type generated for the YANG enumerated type as-path-segment-type.
type OpenconfigRibBgpAsPathSegmentType int32

const (
	OpenconfigRibBgpAsPathSegmentType_OPENCONFIGRIBBGPASPATHSEGMENTTYPE_UNSET              OpenconfigRibBgpAsPathSegmentType = 0
	OpenconfigRibBgpAsPathSegmentType_OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_SEQ             OpenconfigRibBgpAsPathSegmentType = 1
	OpenconfigRibBgpAsPathSegmentType_OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_SET             OpenconfigRibBgpAsPathSegmentType = 2
	OpenconfigRibBgpAsPathSegmentType_OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_CONFED_SEQUENCE OpenconfigRibBgpAsPathSegmentType = 3
	OpenconfigRibBgpAsPathSegmentType_OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_CONFED_SET      OpenconfigRibBgpAsPathSegmentType = 4
)

var OpenconfigRibBgpAsPathSegmentType_name = map[int32]string{
	0: "OPENCONFIGRIBBGPASPATHSEGMENTTYPE_UNSET",
	1: "OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_SEQ",
	2: "OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_SET",
	3: "OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_CONFED_SEQUENCE",
	4: "OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_CONFED_SET",
}
var OpenconfigRibBgpAsPathSegmentType_value = map[string]int32{
	"OPENCONFIGRIBBGPASPATHSEGMENTTYPE_UNSET":              0,
	"OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_SEQ":             1,
	"OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_SET":             2,
	"OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_CONFED_SEQUENCE": 3,
	"OPENCONFIGRIBBGPASPATHSEGMENTTYPE_AS_CONFED_SET":      4,
}

func (x OpenconfigRibBgpAsPathSegmentType) String() string {
	return proto.EnumName(OpenconfigRibBgpAsPathSegmentType_name, int32(x))
}
func (OpenconfigRibBgpAsPathSegmentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

// OpenconfigRibBgpBgpOriginAttrType represents an enumerated type generated for the YANG enumerated type bgp-origin-attr-type.
type OpenconfigRibBgpBgpOriginAttrType int32

const (
	OpenconfigRibBgpBgpOriginAttrType_OPENCONFIGRIBBGPBGPORIGINATTRTYPE_UNSET      OpenconfigRibBgpBgpOriginAttrType = 0
	OpenconfigRibBgpBgpOriginAttrType_OPENCONFIGRIBBGPBGPORIGINATTRTYPE_IGP        OpenconfigRibBgpBgpOriginAttrType = 1
	OpenconfigRibBgpBgpOriginAttrType_OPENCONFIGRIBBGPBGPORIGINATTRTYPE_EGP        OpenconfigRibBgpBgpOriginAttrType = 2
	OpenconfigRibBgpBgpOriginAttrType_OPENCONFIGRIBBGPBGPORIGINATTRTYPE_INCOMPLETE OpenconfigRibBgpBgpOriginAttrType = 3
)

var OpenconfigRibBgpBgpOriginAttrType_name = map[int32]string{
	0: "OPENCONFIGRIBBGPBGPORIGINATTRTYPE_UNSET",
	1: "OPENCONFIGRIBBGPBGPORIGINATTRTYPE_IGP",
	2: "OPENCONFIGRIBBGPBGPORIGINATTRTYPE_EGP",
	3: "OPENCONFIGRIBBGPBGPORIGINATTRTYPE_INCOMPLETE",
}
var OpenconfigRibBgpBgpOriginAttrType_value = map[string]int32{
	"OPENCONFIGRIBBGPBGPORIGINATTRTYPE_UNSET":      0,
	"OPENCONFIGRIBBGPBGPORIGINATTRTYPE_IGP":        1,
	"OPENCONFIGRIBBGPBGPORIGINATTRTYPE_EGP":        2,
	"OPENCONFIGRIBBGPBGPORIGINATTRTYPE_INCOMPLETE": 3,
}

func (x OpenconfigRibBgpBgpOriginAttrType) String() string {
	return proto.EnumName(OpenconfigRibBgpBgpOriginAttrType_name, int32(x))
}
func (OpenconfigRibBgpBgpOriginAttrType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4}
}

// OpenconfigRibBgpTypesINVALIDROUTEREASON represents an enumerated type generated for the YANG identity INVALID_ROUTE_REASON.
type OpenconfigRibBgpTypesINVALIDROUTEREASON int32

const (
	OpenconfigRibBgpTypesINVALIDROUTEREASON_OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_UNSET                OpenconfigRibBgpTypesINVALIDROUTEREASON = 0
	OpenconfigRibBgpTypesINVALIDROUTEREASON_OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_AS_LOOP      OpenconfigRibBgpTypesINVALIDROUTEREASON = 281940072
	OpenconfigRibBgpTypesINVALIDROUTEREASON_OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_ORIGINATOR   OpenconfigRibBgpTypesINVALIDROUTEREASON = 329543457
	OpenconfigRibBgpTypesINVALIDROUTEREASON_OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_CLUSTER_LOOP OpenconfigRibBgpTypesINVALIDROUTEREASON = 363552896
	OpenconfigRibBgpTypesINVALIDROUTEREASON_OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_CONFED       OpenconfigRibBgpTypesINVALIDROUTEREASON = 404003798
)

var OpenconfigRibBgpTypesINVALIDROUTEREASON_name = map[int32]string{
	0:         "OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_UNSET",
	281940072: "OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_AS_LOOP",
	329543457: "OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_ORIGINATOR",
	363552896: "OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_CLUSTER_LOOP",
	404003798: "OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_CONFED",
}
var OpenconfigRibBgpTypesINVALIDROUTEREASON_value = map[string]int32{
	"OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_UNSET":                0,
	"OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_AS_LOOP":      281940072,
	"OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_ORIGINATOR":   329543457,
	"OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_CLUSTER_LOOP": 363552896,
	"OPENCONFIGRIBBGPTYPESINVALIDROUTEREASON_INVALID_CONFED":       404003798,
}

func (x OpenconfigRibBgpTypesINVALIDROUTEREASON) String() string {
	return proto.EnumName(OpenconfigRibBgpTypesINVALIDROUTEREASON_name, int32(x))
}
func (OpenconfigRibBgpTypesINVALIDROUTEREASON) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func init() {
	proto.RegisterEnum("openconfig.enums.OpenconfigBgpTypesAFISAFITYPE", OpenconfigBgpTypesAFISAFITYPE_name, OpenconfigBgpTypesAFISAFITYPE_value)
	proto.RegisterEnum("openconfig.enums.OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY", OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_name, OpenconfigBgpTypesBGPWELLKNOWNSTDCOMMUNITY_value)
	proto.RegisterEnum("openconfig.enums.OpenconfigPolicyTypesINSTALLPROTOCOLTYPE", OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_name, OpenconfigPolicyTypesINSTALLPROTOCOLTYPE_value)
	proto.RegisterEnum("openconfig.enums.OpenconfigRibBgpAsPathSegmentType", OpenconfigRibBgpAsPathSegmentType_name, OpenconfigRibBgpAsPathSegmentType_value)
	proto.RegisterEnum("openconfig.enums.OpenconfigRibBgpBgpOriginAttrType", OpenconfigRibBgpBgpOriginAttrType_name, OpenconfigRibBgpBgpOriginAttrType_value)
	proto.RegisterEnum("openconfig.enums.OpenconfigRibBgpTypesINVALIDROUTEREASON", OpenconfigRibBgpTypesINVALIDROUTEREASON_name, OpenconfigRibBgpTypesINVALIDROUTEREASON_value)
}

func init() { proto.RegisterFile("ribproto/openconfig/enums/enums.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x6d, 0x68, 0x1c, 0x45,
	0x18, 0x76, 0x93, 0x18, 0xed, 0xd0, 0xda, 0x71, 0xfc, 0xc2, 0x05, 0x7f, 0x48, 0x2d, 0x29, 0xa7,
	0x26, 0xcd, 0x07, 0x69, 0xd1, 0x98, 0x3a, 0xbb, 0x37, 0x59, 0x07, 0x37, 0x33, 0x93, 0x9d, 0xb9,
	0xab, 0x27, 0x84, 0xa3, 0x09, 0xe7, 0xf5, 0xc0, 0xdc, 0x1d, 0x97, 0x2b, 0x9a, 0x7f, 0x75, 0x45,
	0x05, 0x05, 0xa9, 0x4a, 0x44, 0xd0, 0xe2, 0x07, 0xe8, 0x0f, 0xb5, 0x88, 0xdf, 0x42, 0xa0, 0x28,
	0xf8, 0x49, 0x51, 0xfc, 0x40, 0x7f, 0x88, 0xf8, 0x41, 0x11, 0x8b, 0x52, 0x94, 0x22, 0x58, 0x28,
	0x88, 0xec, 0x6e, 0xb2, 0xbb, 0xf7, 0x91, 0x64, 0xd7, 0xfe, 0x39, 0x6e, 0xde, 0xe7, 0x7d, 0x9e,
	0x79, 0xf6, 0x7d, 0xe7, 0x9d, 0x5d, 0xb0, 0xb3, 0x51, 0x99, 0xab, 0x37, 0x6a, 0xcd, 0xda, 0x50,
	0xad, 0x5e, 0xaa, 0xce, 0xd7, 0xaa, 0x77, 0x55, 0xca, 0x43, 0xa5, 0xea, 0xa1, 0x85, 0xc5, 0xe0,
	0x77, 0xd0, 0x07, 0x11, 0x8c, 0xd0, 0x41, 0x3f, 0xae, 0xef, 0x2d, 0x57, 0x9a, 0x07, 0x0f, 0xcd,
	0x0d, 0xce, 0xd7, 0x16, 0xe2, 0xd4, 0xa5, 0x72, 0xad, 0x39, 0x14, 0x08, 0x2e, 0xdd, 0xd3, 0x38,
	0x50, 0xaf, 0x97, 0x1a, 0xe1, 0x9f, 0x40, 0x4b, 0xdf, 0xbd, 0x39, 0xb3, 0x74, 0x6f, 0xd3, 0xff,
	0x09, 0x18, 0x99, 0x8f, 0xfb, 0xc1, 0x35, 0x3c, 0xcc, 0x34, 0xca, 0x75, 0xb5, 0x54, 0x2f, 0x2d,
	0xe2, 0x29, 0x2a, 0xf1, 0x14, 0x55, 0x05, 0x41, 0xd0, 0x00, 0xd8, 0xc1, 0x05, 0x61, 0x26, 0x67,
	0x53, 0xd4, 0x32, 0x2c, 0xe1, 0x05, 0x65, 0x2c, 0xa1, 0x98, 0x63, 0x92, 0x28, 0x78, 0x01, 0x72,
	0xc0, 0xc8, 0xc6, 0x89, 0xf6, 0x68, 0x5e, 0xb0, 0x22, 0x15, 0xf9, 0xb1, 0xe2, 0x74, 0xce, 0x56,
	0xd4, 0xc4, 0x52, 0xc1, 0x67, 0x4e, 0x7c, 0x72, 0xb5, 0x7e, 0x95, 0x8b, 0x2f, 0xef, 0x86, 0xa6,
	0xd0, 0x1c, 0x8f, 0x69, 0x1e, 0x3f, 0xf9, 0xfb, 0x4c, 0xab, 0x66, 0x0c, 0x45, 0x06, 0xc8, 0x6c,
	0xac, 0xe9, 0x7b, 0xc8, 0xb1, 0x40, 0xeb, 0xf4, 0x1b, 0x67, 0x66, 0xf5, 0xed, 0x2e, 0xde, 0x1a,
	0x8f, 0xa2, 0x49, 0xb0, 0x6b, 0x13, 0x5f, 0x23, 0xde, 0xce, 0x24, 0x2f, 0x18, 0x7c, 0xea, 0xa1,
	0x53, 0x25, 0x7d, 0x9b, 0x8b, 0x41, 0x14, 0x43, 0x66, 0x02, 0x0f, 0xe3, 0xa1, 0x87, 0x6f, 0xde,
	0xf9, 0xfb, 0xb0, 0xb6, 0x66, 0x22, 0x0c, 0x23, 0xb9, 0x59, 0x71, 0x7c, 0xcb, 0x36, 0x36, 0x88,
	0x4d, 0xb2, 0xa1, 0xd8, 0xd1, 0xb7, 0x1f, 0x79, 0x40, 0xf3, 0xab, 0xd3, 0x0d, 0x46, 0x1c, 0xec,
	0x4e, 0xdc, 0xc5, 0x35, 0xc9, 0xe5, 0xd7, 0xcf, 0xbd, 0xac, 0xe9, 0x57, 0xb8, 0x18, 0x75, 0x82,
	0x89, 0x5c, 0x8e, 0x77, 0xb8, 0x7c, 0xf3, 0xb5, 0xcf, 0x3f, 0x0d, 0x5d, 0x8e, 0x9f, 0x87, 0xcb,
	0xa8, 0x8a, 0x2b, 0x8f, 0x3f, 0xf9, 0x59, 0x9b, 0xcb, 0xa8, 0x96, 0xfb, 0x92, 0x35, 0x34, 0x2f,
	0x6c, 0x09, 0x5f, 0x5a, 0x79, 0xf5, 0x5f, 0x2d, 0xd6, 0x51, 0x2f, 0x98, 0x79, 0xa1, 0x17, 0x64,
	0x3a, 0x07, 0xc9, 0xb0, 0xc4, 0x7e, 0x62, 0xdb, 0xb7, 0x33, 0xbe, 0x9f, 0x49, 0x95, 0x35, 0xf9,
	0xf4, 0x74, 0x8e, 0x51, 0x55, 0x40, 0x63, 0xdd, 0x1e, 0x60, 0xbd, 0xec, 0x70, 0xc4, 0x28, 0x18,
	0x4b, 0xc1, 0x62, 0xbc, 0x48, 0xee, 0x10, 0xdc, 0x51, 0xf0, 0xc4, 0xb3, 0x1f, 0xf5, 0xea, 0x5b,
	0x5d, 0xbc, 0x25, 0x0c, 0xa1, 0x59, 0x30, 0xf9, 0x7f, 0xa4, 0x8a, 0x32, 0x67, 0x78, 0x1c, 0x92,
	0x85, 0xbf, 0x7e, 0xf8, 0xe7, 0x0e, 0xfd, 0x4a, 0x17, 0x5f, 0xd6, 0x05, 0x44, 0x26, 0x18, 0x4e,
	0x25, 0x2f, 0x08, 0x71, 0xe0, 0x99, 0xdf, 0xce, 0x5e, 0xa7, 0x6f, 0x71, 0x71, 0x7f, 0xb0, 0x46,
	0x1c, 0xec, 0x49, 0xe7, 0x11, 0x67, 0xf3, 0xc4, 0x51, 0x54, 0x12, 0xf8, 0xd6, 0xa3, 0x2b, 0x13,
	0xfe, 0xc4, 0xc4, 0xa3, 0x99, 0xaf, 0xfa, 0xc0, 0xae, 0xa8, 0x49, 0xa2, 0x76, 0x77, 0x65, 0x7e,
	0xc9, 0xef, 0x13, 0x65, 0x52, 0x61, 0xdb, 0x16, 0x0e, 0x57, 0xdc, 0xe4, 0xb6, 0x7f, 0xf1, 0x8d,
	0x80, 0xc1, 0x68, 0x77, 0xc1, 0x6d, 0x6a, 0x16, 0x7c, 0x03, 0x5d, 0x72, 0xc3, 0x06, 0x61, 0x30,
	0x94, 0x98, 0x23, 0x15, 0x56, 0xd4, 0x84, 0x0f, 0x9f, 0x7e, 0x91, 0xf9, 0x0f, 0x1d, 0xac, 0x91,
	0x04, 0x7b, 0x13, 0x4b, 0xd8, 0xdc, 0xc4, 0x76, 0x11, 0x5b, 0x96, 0x43, 0x2c, 0xac, 0x08, 0xfc,
	0xf6, 0x97, 0x9f, 0xb8, 0x8e, 0x5c, 0xbc, 0xbd, 0x0d, 0x40, 0xb7, 0x80, 0x1b, 0x13, 0x8b, 0x72,
	0x29, 0xa6, 0xe0, 0x5f, 0xe7, 0x96, 0x0b, 0xfa, 0x45, 0x2e, 0xee, 0xf3, 0x56, 0x68, 0x32, 0x05,
	0x9d, 0x4a, 0x2a, 0xe1, 0xd9, 0x23, 0x47, 0x9e, 0xd0, 0x7c, 0xbe, 0xb7, 0x44, 0x13, 0xe0, 0x86,
	0xc4, 0x7c, 0xc3, 0x12, 0xf0, 0x95, 0xfb, 0x9f, 0xfe, 0x52, 0xd3, 0xfb, 0x5d, 0xdc, 0x6b, 0x58,
	0x02, 0xdd, 0x09, 0x6e, 0x4e, 0xcc, 0xce, 0x52, 0x87, 0x98, 0xca, 0x2e, 0x14, 0x4d, 0xce, 0x18,
	0x31, 0x15, 0xc9, 0xc2, 0x2f, 0xdc, 0x3f, 0x7e, 0x08, 0xe6, 0xbe, 0x13, 0x44, 0xb7, 0xa6, 0x68,
	0xb2, 0x57, 0x8a, 0x51, 0x78, 0xf2, 0xbe, 0xe5, 0x9f, 0x35, 0xfd, 0x62, 0x17, 0x5f, 0xe8, 0xaf,
	0x33, 0xdf, 0xf7, 0x80, 0x6b, 0xa3, 0x33, 0xe5, 0x54, 0xe6, 0x8c, 0x72, 0x1d, 0x2f, 0x8a, 0x03,
	0xcd, 0x83, 0xb2, 0x54, 0x5e, 0x28, 0x55, 0x9b, 0xde, 0x01, 0x43, 0xd7, 0x83, 0x81, 0x68, 0x1f,
	0x87, 0x1a, 0x86, 0x25, 0xb0, 0x14, 0x58, 0xdd, 0x26, 0x89, 0x35, 0x4d, 0x98, 0x6a, 0x39, 0x45,
	0x7b, 0xe2, 0x97, 0xd1, 0x7a, 0xc9, 0x58, 0x16, 0x25, 0x99, 0x81, 0x9a, 0x7f, 0x76, 0x82, 0x45,
	0x0a, 0xa2, 0x82, 0x3d, 0x11, 0xd1, 0xbb, 0x4f, 0xc7, 0x12, 0x11, 0x83, 0xe9, 0xf6, 0xf6, 0xca,
	0x11, 0x66, 0x12, 0xd8, 0xeb, 0xd7, 0xb5, 0x13, 0x40, 0x24, 0x3e, 0x08, 0x49, 0x04, 0x15, 0xec,
	0xd3, 0xa1, 0x8b, 0xb7, 0xb5, 0xc4, 0x32, 0xff, 0x68, 0x9d, 0xc5, 0x35, 0xca, 0x75, 0xde, 0xa8,
	0x94, 0x2b, 0x55, 0xdc, 0x6c, 0x36, 0xd6, 0x2b, 0xae, 0x61, 0x09, 0xee, 0x50, 0x8b, 0x32, 0xac,
	0x94, 0xd3, 0x52, 0xdc, 0x61, 0xb0, 0x73, 0xf3, 0x64, 0x6a, 0x09, 0x18, 0x1c, 0x40, 0x6a, 0x89,
	0x64, 0x14, 0x62, 0x09, 0xd8, 0xe3, 0x53, 0x88, 0x25, 0xd0, 0xbe, 0xf8, 0x89, 0x5f, 0x77, 0x17,
	0x66, 0xf2, 0x69, 0x61, 0x13, 0xe5, 0x15, 0xd2, 0x7b, 0x9f, 0x44, 0x81, 0xcc, 0xf1, 0x5e, 0x30,
	0xd0, 0xfe, 0xe4, 0xab, 0x57, 0x55, 0x1e, 0xdb, 0x34, 0xeb, 0xf0, 0x9c, 0x22, 0x0e, 0xc1, 0x92,
	0x33, 0x34, 0x1c, 0x1f, 0xcf, 0x60, 0xb3, 0xd5, 0x43, 0xdc, 0x9e, 0x1a, 0x56, 0x41, 0xc6, 0xaf,
	0xd6, 0x8d, 0x29, 0xab, 0x21, 0xaf, 0x5b, 0x36, 0xe7, 0x02, 0x9e, 0x7a, 0xee, 0xbd, 0x07, 0x35,
	0xff, 0x96, 0x69, 0x43, 0x50, 0x01, 0xdc, 0x94, 0x56, 0x74, 0xad, 0x22, 0xdc, 0x81, 0xcf, 0xff,
	0xf8, 0xd8, 0xd1, 0x60, 0x4e, 0x3b, 0x41, 0x34, 0x0b, 0x26, 0xd2, 0x4a, 0x9b, 0x76, 0x4e, 0x2a,
	0xe2, 0x04, 0xa6, 0x0f, 0x7f, 0x70, 0xec, 0xd8, 0xea, 0xf7, 0x44, 0x17, 0x18, 0xcd, 0x80, 0xf1,
	0xd4, 0xf2, 0xc1, 0x5b, 0xf0, 0xbb, 0x77, 0xbf, 0x7e, 0x5f, 0xd3, 0x2f, 0x75, 0xf1, 0x25, 0xad,
	0xc0, 0x5c, 0xbf, 0xff, 0x81, 0x3d, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x69, 0xd5,
	0x36, 0x07, 0x0c, 0x00, 0x00,
}
