// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_twap

import ag_binary "github.com/gagliardetto/binary"

type TWAPOracle struct {
	ExpectedValue         uint64
	InitialSlot           uint64
	LastUpdatedSlot       uint64
	LastObservedSlot      uint64
	LastObservation       uint64
	ObservationAggregator ag_binary.Uint128
}

func (obj TWAPOracle) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ExpectedValue` param:
	err = encoder.Encode(obj.ExpectedValue)
	if err != nil {
		return err
	}
	// Serialize `InitialSlot` param:
	err = encoder.Encode(obj.InitialSlot)
	if err != nil {
		return err
	}
	// Serialize `LastUpdatedSlot` param:
	err = encoder.Encode(obj.LastUpdatedSlot)
	if err != nil {
		return err
	}
	// Serialize `LastObservedSlot` param:
	err = encoder.Encode(obj.LastObservedSlot)
	if err != nil {
		return err
	}
	// Serialize `LastObservation` param:
	err = encoder.Encode(obj.LastObservation)
	if err != nil {
		return err
	}
	// Serialize `ObservationAggregator` param:
	err = encoder.Encode(obj.ObservationAggregator)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TWAPOracle) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ExpectedValue`:
	err = decoder.Decode(&obj.ExpectedValue)
	if err != nil {
		return err
	}
	// Deserialize `InitialSlot`:
	err = decoder.Decode(&obj.InitialSlot)
	if err != nil {
		return err
	}
	// Deserialize `LastUpdatedSlot`:
	err = decoder.Decode(&obj.LastUpdatedSlot)
	if err != nil {
		return err
	}
	// Deserialize `LastObservedSlot`:
	err = decoder.Decode(&obj.LastObservedSlot)
	if err != nil {
		return err
	}
	// Deserialize `LastObservation`:
	err = decoder.Decode(&obj.LastObservation)
	if err != nil {
		return err
	}
	// Deserialize `ObservationAggregator`:
	err = decoder.Decode(&obj.ObservationAggregator)
	if err != nil {
		return err
	}
	return nil
}

type PlaceOrderArgs struct {
	Side                      Side
	PriceLots                 int64
	MaxBaseLots               int64
	MaxQuoteLotsIncludingFees int64
	ClientOrderId             uint64
	OrderType                 PlaceOrderType
	ExpiryTimestamp           uint64
	SelfTradeBehavior         SelfTradeBehavior
	Limit                     uint8
}

func (obj PlaceOrderArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.Encode(obj.Side)
	if err != nil {
		return err
	}
	// Serialize `PriceLots` param:
	err = encoder.Encode(obj.PriceLots)
	if err != nil {
		return err
	}
	// Serialize `MaxBaseLots` param:
	err = encoder.Encode(obj.MaxBaseLots)
	if err != nil {
		return err
	}
	// Serialize `MaxQuoteLotsIncludingFees` param:
	err = encoder.Encode(obj.MaxQuoteLotsIncludingFees)
	if err != nil {
		return err
	}
	// Serialize `ClientOrderId` param:
	err = encoder.Encode(obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Serialize `OrderType` param:
	err = encoder.Encode(obj.OrderType)
	if err != nil {
		return err
	}
	// Serialize `ExpiryTimestamp` param:
	err = encoder.Encode(obj.ExpiryTimestamp)
	if err != nil {
		return err
	}
	// Serialize `SelfTradeBehavior` param:
	err = encoder.Encode(obj.SelfTradeBehavior)
	if err != nil {
		return err
	}
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PlaceOrderArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	err = decoder.Decode(&obj.Side)
	if err != nil {
		return err
	}
	// Deserialize `PriceLots`:
	err = decoder.Decode(&obj.PriceLots)
	if err != nil {
		return err
	}
	// Deserialize `MaxBaseLots`:
	err = decoder.Decode(&obj.MaxBaseLots)
	if err != nil {
		return err
	}
	// Deserialize `MaxQuoteLotsIncludingFees`:
	err = decoder.Decode(&obj.MaxQuoteLotsIncludingFees)
	if err != nil {
		return err
	}
	// Deserialize `ClientOrderId`:
	err = decoder.Decode(&obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Deserialize `OrderType`:
	err = decoder.Decode(&obj.OrderType)
	if err != nil {
		return err
	}
	// Deserialize `ExpiryTimestamp`:
	err = decoder.Decode(&obj.ExpiryTimestamp)
	if err != nil {
		return err
	}
	// Deserialize `SelfTradeBehavior`:
	err = decoder.Decode(&obj.SelfTradeBehavior)
	if err != nil {
		return err
	}
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

type PlaceTakeOrderArgs struct {
	Side                      Side
	PriceLots                 int64
	MaxBaseLots               int64
	MaxQuoteLotsIncludingFees int64
	OrderType                 PlaceOrderType
	Limit                     uint8
}

func (obj PlaceTakeOrderArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Side` param:
	err = encoder.Encode(obj.Side)
	if err != nil {
		return err
	}
	// Serialize `PriceLots` param:
	err = encoder.Encode(obj.PriceLots)
	if err != nil {
		return err
	}
	// Serialize `MaxBaseLots` param:
	err = encoder.Encode(obj.MaxBaseLots)
	if err != nil {
		return err
	}
	// Serialize `MaxQuoteLotsIncludingFees` param:
	err = encoder.Encode(obj.MaxQuoteLotsIncludingFees)
	if err != nil {
		return err
	}
	// Serialize `OrderType` param:
	err = encoder.Encode(obj.OrderType)
	if err != nil {
		return err
	}
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PlaceTakeOrderArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Side`:
	err = decoder.Decode(&obj.Side)
	if err != nil {
		return err
	}
	// Deserialize `PriceLots`:
	err = decoder.Decode(&obj.PriceLots)
	if err != nil {
		return err
	}
	// Deserialize `MaxBaseLots`:
	err = decoder.Decode(&obj.MaxBaseLots)
	if err != nil {
		return err
	}
	// Deserialize `MaxQuoteLotsIncludingFees`:
	err = decoder.Decode(&obj.MaxQuoteLotsIncludingFees)
	if err != nil {
		return err
	}
	// Deserialize `OrderType`:
	err = decoder.Decode(&obj.OrderType)
	if err != nil {
		return err
	}
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

type SelfTradeBehavior ag_binary.BorshEnum

const (
	SelfTradeBehaviorDecrementTake SelfTradeBehavior = iota
	SelfTradeBehaviorCancelProvide
	SelfTradeBehaviorAbortTransaction
)

func (value SelfTradeBehavior) String() string {
	switch value {
	case SelfTradeBehaviorDecrementTake:
		return "DecrementTake"
	case SelfTradeBehaviorCancelProvide:
		return "CancelProvide"
	case SelfTradeBehaviorAbortTransaction:
		return "AbortTransaction"
	default:
		return ""
	}
}

type PlaceOrderType ag_binary.BorshEnum

const (
	PlaceOrderTypeLimit PlaceOrderType = iota
	PlaceOrderTypeImmediateOrCancel
	PlaceOrderTypePostOnly
	PlaceOrderTypeMarket
	PlaceOrderTypePostOnlySlide
)

func (value PlaceOrderType) String() string {
	switch value {
	case PlaceOrderTypeLimit:
		return "Limit"
	case PlaceOrderTypeImmediateOrCancel:
		return "ImmediateOrCancel"
	case PlaceOrderTypePostOnly:
		return "PostOnly"
	case PlaceOrderTypeMarket:
		return "Market"
	case PlaceOrderTypePostOnlySlide:
		return "PostOnlySlide"
	default:
		return ""
	}
}

type Side ag_binary.BorshEnum

const (
	SideBid Side = iota
	SideAsk
)

func (value Side) String() string {
	switch value {
	case SideBid:
		return "Bid"
	case SideAsk:
		return "Ask"
	default:
		return ""
	}
}
