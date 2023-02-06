// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AddressType string

// Enum values for AddressType
const (
	AddressTypeShippingAddress  AddressType = "SHIPPING_ADDRESS"
	AddressTypeOperatingAddress AddressType = "OPERATING_ADDRESS"
)

// Values returns all known values for AddressType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AddressType) Values() []AddressType {
	return []AddressType{
		"SHIPPING_ADDRESS",
		"OPERATING_ADDRESS",
	}
}

type AssetState string

// Enum values for AssetState
const (
	AssetStateActive   AssetState = "ACTIVE"
	AssetStateRetiring AssetState = "RETIRING"
)

// Values returns all known values for AssetState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AssetState) Values() []AssetState {
	return []AssetState{
		"ACTIVE",
		"RETIRING",
	}
}

type AssetType string

// Enum values for AssetType
const (
	AssetTypeCompute AssetType = "COMPUTE"
)

// Values returns all known values for AssetType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AssetType) Values() []AssetType {
	return []AssetType{
		"COMPUTE",
	}
}

type CatalogItemClass string

// Enum values for CatalogItemClass
const (
	CatalogItemClassRack   CatalogItemClass = "RACK"
	CatalogItemClassServer CatalogItemClass = "SERVER"
)

// Values returns all known values for CatalogItemClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CatalogItemClass) Values() []CatalogItemClass {
	return []CatalogItemClass{
		"RACK",
		"SERVER",
	}
}

type CatalogItemStatus string

// Enum values for CatalogItemStatus
const (
	CatalogItemStatusAvailable    CatalogItemStatus = "AVAILABLE"
	CatalogItemStatusDiscontinued CatalogItemStatus = "DISCONTINUED"
)

// Values returns all known values for CatalogItemStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CatalogItemStatus) Values() []CatalogItemStatus {
	return []CatalogItemStatus{
		"AVAILABLE",
		"DISCONTINUED",
	}
}

type ComputeAssetState string

// Enum values for ComputeAssetState
const (
	ComputeAssetStateActive   ComputeAssetState = "ACTIVE"
	ComputeAssetStateIsolated ComputeAssetState = "ISOLATED"
	ComputeAssetStateRetiring ComputeAssetState = "RETIRING"
)

// Values returns all known values for ComputeAssetState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComputeAssetState) Values() []ComputeAssetState {
	return []ComputeAssetState{
		"ACTIVE",
		"ISOLATED",
		"RETIRING",
	}
}

type FiberOpticCableType string

// Enum values for FiberOpticCableType
const (
	FiberOpticCableTypeSingleMode FiberOpticCableType = "SINGLE_MODE"
	FiberOpticCableTypeMultiMode  FiberOpticCableType = "MULTI_MODE"
)

// Values returns all known values for FiberOpticCableType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FiberOpticCableType) Values() []FiberOpticCableType {
	return []FiberOpticCableType{
		"SINGLE_MODE",
		"MULTI_MODE",
	}
}

type LineItemStatus string

// Enum values for LineItemStatus
const (
	LineItemStatusPreparing  LineItemStatus = "PREPARING"
	LineItemStatusBuilding   LineItemStatus = "BUILDING"
	LineItemStatusShipped    LineItemStatus = "SHIPPED"
	LineItemStatusDelivered  LineItemStatus = "DELIVERED"
	LineItemStatusInstalling LineItemStatus = "INSTALLING"
	LineItemStatusInstalled  LineItemStatus = "INSTALLED"
	LineItemStatusError      LineItemStatus = "ERROR"
	LineItemStatusCancelled  LineItemStatus = "CANCELLED"
	LineItemStatusReplaced   LineItemStatus = "REPLACED"
)

// Values returns all known values for LineItemStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LineItemStatus) Values() []LineItemStatus {
	return []LineItemStatus{
		"PREPARING",
		"BUILDING",
		"SHIPPED",
		"DELIVERED",
		"INSTALLING",
		"INSTALLED",
		"ERROR",
		"CANCELLED",
		"REPLACED",
	}
}

type MaximumSupportedWeightLbs string

// Enum values for MaximumSupportedWeightLbs
const (
	MaximumSupportedWeightLbsNoLimit    MaximumSupportedWeightLbs = "NO_LIMIT"
	MaximumSupportedWeightLbsMax1400Lbs MaximumSupportedWeightLbs = "MAX_1400_LBS"
	MaximumSupportedWeightLbsMax1600Lbs MaximumSupportedWeightLbs = "MAX_1600_LBS"
	MaximumSupportedWeightLbsMax1800Lbs MaximumSupportedWeightLbs = "MAX_1800_LBS"
	MaximumSupportedWeightLbsMax2000Lbs MaximumSupportedWeightLbs = "MAX_2000_LBS"
)

// Values returns all known values for MaximumSupportedWeightLbs. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (MaximumSupportedWeightLbs) Values() []MaximumSupportedWeightLbs {
	return []MaximumSupportedWeightLbs{
		"NO_LIMIT",
		"MAX_1400_LBS",
		"MAX_1600_LBS",
		"MAX_1800_LBS",
		"MAX_2000_LBS",
	}
}

type OpticalStandard string

// Enum values for OpticalStandard
const (
	OpticalStandardOptic10gbaseSr      OpticalStandard = "OPTIC_10GBASE_SR"
	OpticalStandardOptic10gbaseIr      OpticalStandard = "OPTIC_10GBASE_IR"
	OpticalStandardOptic10gbaseLr      OpticalStandard = "OPTIC_10GBASE_LR"
	OpticalStandardOptic40gbaseSr      OpticalStandard = "OPTIC_40GBASE_SR"
	OpticalStandardOptic40gbaseEsr     OpticalStandard = "OPTIC_40GBASE_ESR"
	OpticalStandardOptic40gbaseIr4Lr4l OpticalStandard = "OPTIC_40GBASE_IR4_LR4L"
	OpticalStandardOptic40gbaseLr4     OpticalStandard = "OPTIC_40GBASE_LR4"
	OpticalStandardOptic100gbaseSr4    OpticalStandard = "OPTIC_100GBASE_SR4"
	OpticalStandardOptic100gbaseCwdm4  OpticalStandard = "OPTIC_100GBASE_CWDM4"
	OpticalStandardOptic100gbaseLr4    OpticalStandard = "OPTIC_100GBASE_LR4"
	OpticalStandardOptic100gPsm4Msa    OpticalStandard = "OPTIC_100G_PSM4_MSA"
	OpticalStandardOptic1000baseLx     OpticalStandard = "OPTIC_1000BASE_LX"
	OpticalStandardOptic1000baseSx     OpticalStandard = "OPTIC_1000BASE_SX"
)

// Values returns all known values for OpticalStandard. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OpticalStandard) Values() []OpticalStandard {
	return []OpticalStandard{
		"OPTIC_10GBASE_SR",
		"OPTIC_10GBASE_IR",
		"OPTIC_10GBASE_LR",
		"OPTIC_40GBASE_SR",
		"OPTIC_40GBASE_ESR",
		"OPTIC_40GBASE_IR4_LR4L",
		"OPTIC_40GBASE_LR4",
		"OPTIC_100GBASE_SR4",
		"OPTIC_100GBASE_CWDM4",
		"OPTIC_100GBASE_LR4",
		"OPTIC_100G_PSM4_MSA",
		"OPTIC_1000BASE_LX",
		"OPTIC_1000BASE_SX",
	}
}

type OrderStatus string

// Enum values for OrderStatus
const (
	OrderStatusReceived   OrderStatus = "RECEIVED"
	OrderStatusPending    OrderStatus = "PENDING"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusInstalling OrderStatus = "INSTALLING"
	OrderStatusFulfilled  OrderStatus = "FULFILLED"
	OrderStatusCancelled  OrderStatus = "CANCELLED"
	OrderStatusPreparing  OrderStatus = "PREPARING"
	OrderStatusInProgress OrderStatus = "IN_PROGRESS"
	OrderStatusCompleted  OrderStatus = "COMPLETED"
	OrderStatusError      OrderStatus = "ERROR"
)

// Values returns all known values for OrderStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OrderStatus) Values() []OrderStatus {
	return []OrderStatus{
		"RECEIVED",
		"PENDING",
		"PROCESSING",
		"INSTALLING",
		"FULFILLED",
		"CANCELLED",
		"PREPARING",
		"IN_PROGRESS",
		"COMPLETED",
		"ERROR",
	}
}

type OrderType string

// Enum values for OrderType
const (
	OrderTypeOutpost     OrderType = "OUTPOST"
	OrderTypeReplacement OrderType = "REPLACEMENT"
)

// Values returns all known values for OrderType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OrderType) Values() []OrderType {
	return []OrderType{
		"OUTPOST",
		"REPLACEMENT",
	}
}

type PaymentOption string

// Enum values for PaymentOption
const (
	PaymentOptionAllUpfront     PaymentOption = "ALL_UPFRONT"
	PaymentOptionNoUpfront      PaymentOption = "NO_UPFRONT"
	PaymentOptionPartialUpfront PaymentOption = "PARTIAL_UPFRONT"
)

// Values returns all known values for PaymentOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PaymentOption) Values() []PaymentOption {
	return []PaymentOption{
		"ALL_UPFRONT",
		"NO_UPFRONT",
		"PARTIAL_UPFRONT",
	}
}

type PaymentTerm string

// Enum values for PaymentTerm
const (
	PaymentTermThreeYears PaymentTerm = "THREE_YEARS"
	PaymentTermOneYear    PaymentTerm = "ONE_YEAR"
)

// Values returns all known values for PaymentTerm. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PaymentTerm) Values() []PaymentTerm {
	return []PaymentTerm{
		"THREE_YEARS",
		"ONE_YEAR",
	}
}

type PowerConnector string

// Enum values for PowerConnector
const (
	PowerConnectorL630p    PowerConnector = "L6_30P"
	PowerConnectorIec309   PowerConnector = "IEC309"
	PowerConnectorAh530p7w PowerConnector = "AH530P7W"
	PowerConnectorAh532p6w PowerConnector = "AH532P6W"
)

// Values returns all known values for PowerConnector. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PowerConnector) Values() []PowerConnector {
	return []PowerConnector{
		"L6_30P",
		"IEC309",
		"AH530P7W",
		"AH532P6W",
	}
}

type PowerDrawKva string

// Enum values for PowerDrawKva
const (
	PowerDrawKvaPower5Kva  PowerDrawKva = "POWER_5_KVA"
	PowerDrawKvaPower10Kva PowerDrawKva = "POWER_10_KVA"
	PowerDrawKvaPower15Kva PowerDrawKva = "POWER_15_KVA"
	PowerDrawKvaPower30Kva PowerDrawKva = "POWER_30_KVA"
)

// Values returns all known values for PowerDrawKva. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PowerDrawKva) Values() []PowerDrawKva {
	return []PowerDrawKva{
		"POWER_5_KVA",
		"POWER_10_KVA",
		"POWER_15_KVA",
		"POWER_30_KVA",
	}
}

type PowerFeedDrop string

// Enum values for PowerFeedDrop
const (
	PowerFeedDropAboveRack PowerFeedDrop = "ABOVE_RACK"
	PowerFeedDropBelowRack PowerFeedDrop = "BELOW_RACK"
)

// Values returns all known values for PowerFeedDrop. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PowerFeedDrop) Values() []PowerFeedDrop {
	return []PowerFeedDrop{
		"ABOVE_RACK",
		"BELOW_RACK",
	}
}

type PowerPhase string

// Enum values for PowerPhase
const (
	PowerPhaseSinglePhase PowerPhase = "SINGLE_PHASE"
	PowerPhaseThreePhase  PowerPhase = "THREE_PHASE"
)

// Values returns all known values for PowerPhase. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PowerPhase) Values() []PowerPhase {
	return []PowerPhase{
		"SINGLE_PHASE",
		"THREE_PHASE",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeOutpost ResourceType = "OUTPOST"
	ResourceTypeOrder   ResourceType = "ORDER"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"OUTPOST",
		"ORDER",
	}
}

type ShipmentCarrier string

// Enum values for ShipmentCarrier
const (
	ShipmentCarrierDhl   ShipmentCarrier = "DHL"
	ShipmentCarrierDbs   ShipmentCarrier = "DBS"
	ShipmentCarrierFedex ShipmentCarrier = "FEDEX"
	ShipmentCarrierUps   ShipmentCarrier = "UPS"
)

// Values returns all known values for ShipmentCarrier. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShipmentCarrier) Values() []ShipmentCarrier {
	return []ShipmentCarrier{
		"DHL",
		"DBS",
		"FEDEX",
		"UPS",
	}
}

type SupportedHardwareType string

// Enum values for SupportedHardwareType
const (
	SupportedHardwareTypeRack   SupportedHardwareType = "RACK"
	SupportedHardwareTypeServer SupportedHardwareType = "SERVER"
)

// Values returns all known values for SupportedHardwareType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SupportedHardwareType) Values() []SupportedHardwareType {
	return []SupportedHardwareType{
		"RACK",
		"SERVER",
	}
}

type SupportedStorageEnum string

// Enum values for SupportedStorageEnum
const (
	SupportedStorageEnumEbs SupportedStorageEnum = "EBS"
	SupportedStorageEnumS3  SupportedStorageEnum = "S3"
)

// Values returns all known values for SupportedStorageEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SupportedStorageEnum) Values() []SupportedStorageEnum {
	return []SupportedStorageEnum{
		"EBS",
		"S3",
	}
}

type UplinkCount string

// Enum values for UplinkCount
const (
	UplinkCountUplinkCount1  UplinkCount = "UPLINK_COUNT_1"
	UplinkCountUplinkCount2  UplinkCount = "UPLINK_COUNT_2"
	UplinkCountUplinkCount3  UplinkCount = "UPLINK_COUNT_3"
	UplinkCountUplinkCount4  UplinkCount = "UPLINK_COUNT_4"
	UplinkCountUplinkCount5  UplinkCount = "UPLINK_COUNT_5"
	UplinkCountUplinkCount6  UplinkCount = "UPLINK_COUNT_6"
	UplinkCountUplinkCount7  UplinkCount = "UPLINK_COUNT_7"
	UplinkCountUplinkCount8  UplinkCount = "UPLINK_COUNT_8"
	UplinkCountUplinkCount12 UplinkCount = "UPLINK_COUNT_12"
	UplinkCountUplinkCount16 UplinkCount = "UPLINK_COUNT_16"
)

// Values returns all known values for UplinkCount. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UplinkCount) Values() []UplinkCount {
	return []UplinkCount{
		"UPLINK_COUNT_1",
		"UPLINK_COUNT_2",
		"UPLINK_COUNT_3",
		"UPLINK_COUNT_4",
		"UPLINK_COUNT_5",
		"UPLINK_COUNT_6",
		"UPLINK_COUNT_7",
		"UPLINK_COUNT_8",
		"UPLINK_COUNT_12",
		"UPLINK_COUNT_16",
	}
}

type UplinkGbps string

// Enum values for UplinkGbps
const (
	UplinkGbpsUplink1g   UplinkGbps = "UPLINK_1G"
	UplinkGbpsUplink10g  UplinkGbps = "UPLINK_10G"
	UplinkGbpsUplink40g  UplinkGbps = "UPLINK_40G"
	UplinkGbpsUplink100g UplinkGbps = "UPLINK_100G"
)

// Values returns all known values for UplinkGbps. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UplinkGbps) Values() []UplinkGbps {
	return []UplinkGbps{
		"UPLINK_1G",
		"UPLINK_10G",
		"UPLINK_40G",
		"UPLINK_100G",
	}
}
