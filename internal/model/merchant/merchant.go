package merchant

// Globally unique identifier for merchants
type MerchantId struct {
	id string
}

func NewMerchantId(id string) MerchantId {
	return MerchantId{id: id}
}

func (merchantId MerchantId) String() string {
	return merchantId.id
}

// Identifier for deals/benefits offered to customers by a merchant
// Need only be unique among all offerings associated with a given merchant
type OfferingId struct {
	id string
}

func NewOfferingId(id string) OfferingId {
	return OfferingId{id: id}
}

func (offeringId OfferingId) String() string {
	return offeringId.id
}

// The (MerchantId, OfferingId) pair which uniquely identifies any offering (globally unique)
type MerchantOfferingId struct {
	merchantId MerchantId
	offeringId OfferingId
}

func NewMerchantOfferingId(merchantId MerchantId, offeringId OfferingId) MerchantOfferingId {
	return MerchantOfferingId{merchantId: merchantId, offeringId: offeringId}
}

func (merchantOfferingId MerchantOfferingId) MerchantId() MerchantId {
	return merchantOfferingId.merchantId
}

func (merchantOfferingId MerchantOfferingId) OfferingId() OfferingId {
	return merchantOfferingId.offeringId
}

// Information regarding the deal/benefit
// TODO: What else belongs here?
type OfferingDetails struct {
	description string
}

func NewOfferingDetails(description string) *OfferingDetails {
	return &OfferingDetails{description: description}
}

func (offeringDetails *OfferingDetails) Description() string {
	return offeringDetails.description
}

// Base-level object for representing a deals/benefits offered by merchants
type MerchantOffering struct {
	id MerchantOfferingId
	details *OfferingDetails
}

func NewMerchantOffering(id MerchantOfferingId, details *OfferingDetails) MerchantOffering {
	return MerchantOffering{id: id, details: details}
}

func (merchantOffering *MerchantOffering) Id() MerchantOfferingId {
	return merchantOffering.id
}

func (merchantOffering *MerchantOffering) Details() *OfferingDetails {
	return merchantOffering.details
}