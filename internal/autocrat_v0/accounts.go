// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package autocrat_v0

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type DAO struct {
	TreasuryPdaBump          uint8
	Treasury                 ag_solanago.PublicKey
	MetaMint                 ag_solanago.PublicKey
	UsdcMint                 ag_solanago.PublicKey
	ProposalCount            uint32
	LastProposalSlot         uint64
	PassThresholdBps         uint16
	BaseBurnLamports         uint64
	BurnDecayPerSlotLamports uint64
	SlotsPerProposal         uint64
	MarketTakerFee           int64
	TwapExpectedValue        uint64
}

var DAODiscriminator = [8]byte{242, 60, 23, 196, 237, 48, 173, 129}

func (obj DAO) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DAODiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `TreasuryPdaBump` param:
	err = encoder.Encode(obj.TreasuryPdaBump)
	if err != nil {
		return err
	}
	// Serialize `Treasury` param:
	err = encoder.Encode(obj.Treasury)
	if err != nil {
		return err
	}
	// Serialize `MetaMint` param:
	err = encoder.Encode(obj.MetaMint)
	if err != nil {
		return err
	}
	// Serialize `UsdcMint` param:
	err = encoder.Encode(obj.UsdcMint)
	if err != nil {
		return err
	}
	// Serialize `ProposalCount` param:
	err = encoder.Encode(obj.ProposalCount)
	if err != nil {
		return err
	}
	// Serialize `LastProposalSlot` param:
	err = encoder.Encode(obj.LastProposalSlot)
	if err != nil {
		return err
	}
	// Serialize `PassThresholdBps` param:
	err = encoder.Encode(obj.PassThresholdBps)
	if err != nil {
		return err
	}
	// Serialize `BaseBurnLamports` param:
	err = encoder.Encode(obj.BaseBurnLamports)
	if err != nil {
		return err
	}
	// Serialize `BurnDecayPerSlotLamports` param:
	err = encoder.Encode(obj.BurnDecayPerSlotLamports)
	if err != nil {
		return err
	}
	// Serialize `SlotsPerProposal` param:
	err = encoder.Encode(obj.SlotsPerProposal)
	if err != nil {
		return err
	}
	// Serialize `MarketTakerFee` param:
	err = encoder.Encode(obj.MarketTakerFee)
	if err != nil {
		return err
	}
	// Serialize `TwapExpectedValue` param:
	err = encoder.Encode(obj.TwapExpectedValue)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DAO) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DAODiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[242 60 23 196 237 48 173 129]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TreasuryPdaBump`:
	err = decoder.Decode(&obj.TreasuryPdaBump)
	if err != nil {
		return err
	}
	// Deserialize `Treasury`:
	err = decoder.Decode(&obj.Treasury)
	if err != nil {
		return err
	}
	// Deserialize `MetaMint`:
	err = decoder.Decode(&obj.MetaMint)
	if err != nil {
		return err
	}
	// Deserialize `UsdcMint`:
	err = decoder.Decode(&obj.UsdcMint)
	if err != nil {
		return err
	}
	// Deserialize `ProposalCount`:
	err = decoder.Decode(&obj.ProposalCount)
	if err != nil {
		return err
	}
	// Deserialize `LastProposalSlot`:
	err = decoder.Decode(&obj.LastProposalSlot)
	if err != nil {
		return err
	}
	// Deserialize `PassThresholdBps`:
	err = decoder.Decode(&obj.PassThresholdBps)
	if err != nil {
		return err
	}
	// Deserialize `BaseBurnLamports`:
	err = decoder.Decode(&obj.BaseBurnLamports)
	if err != nil {
		return err
	}
	// Deserialize `BurnDecayPerSlotLamports`:
	err = decoder.Decode(&obj.BurnDecayPerSlotLamports)
	if err != nil {
		return err
	}
	// Deserialize `SlotsPerProposal`:
	err = decoder.Decode(&obj.SlotsPerProposal)
	if err != nil {
		return err
	}
	// Deserialize `MarketTakerFee`:
	err = decoder.Decode(&obj.MarketTakerFee)
	if err != nil {
		return err
	}
	// Deserialize `TwapExpectedValue`:
	err = decoder.Decode(&obj.TwapExpectedValue)
	if err != nil {
		return err
	}
	return nil
}

type Proposal struct {
	Number                 uint32
	Proposer               ag_solanago.PublicKey
	DescriptionUrl         string
	SlotEnqueued           uint64
	State                  ProposalState
	Instruction            ProposalInstruction
	OpenbookTwapPassMarket ag_solanago.PublicKey
	OpenbookTwapFailMarket ag_solanago.PublicKey
	OpenbookPassMarket     ag_solanago.PublicKey
	OpenbookFailMarket     ag_solanago.PublicKey
	BaseVault              ag_solanago.PublicKey
	QuoteVault             ag_solanago.PublicKey
}

var ProposalDiscriminator = [8]byte{26, 94, 189, 187, 116, 136, 53, 33}

func (obj Proposal) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ProposalDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Number` param:
	err = encoder.Encode(obj.Number)
	if err != nil {
		return err
	}
	// Serialize `Proposer` param:
	err = encoder.Encode(obj.Proposer)
	if err != nil {
		return err
	}
	// Serialize `DescriptionUrl` param:
	err = encoder.Encode(obj.DescriptionUrl)
	if err != nil {
		return err
	}
	// Serialize `SlotEnqueued` param:
	err = encoder.Encode(obj.SlotEnqueued)
	if err != nil {
		return err
	}
	// Serialize `State` param:
	err = encoder.Encode(obj.State)
	if err != nil {
		return err
	}
	// Serialize `Instruction` param:
	err = encoder.Encode(obj.Instruction)
	if err != nil {
		return err
	}
	// Serialize `OpenbookTwapPassMarket` param:
	err = encoder.Encode(obj.OpenbookTwapPassMarket)
	if err != nil {
		return err
	}
	// Serialize `OpenbookTwapFailMarket` param:
	err = encoder.Encode(obj.OpenbookTwapFailMarket)
	if err != nil {
		return err
	}
	// Serialize `OpenbookPassMarket` param:
	err = encoder.Encode(obj.OpenbookPassMarket)
	if err != nil {
		return err
	}
	// Serialize `OpenbookFailMarket` param:
	err = encoder.Encode(obj.OpenbookFailMarket)
	if err != nil {
		return err
	}
	// Serialize `BaseVault` param:
	err = encoder.Encode(obj.BaseVault)
	if err != nil {
		return err
	}
	// Serialize `QuoteVault` param:
	err = encoder.Encode(obj.QuoteVault)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Proposal) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ProposalDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[26 94 189 187 116 136 53 33]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Number`:
	err = decoder.Decode(&obj.Number)
	if err != nil {
		return err
	}
	// Deserialize `Proposer`:
	err = decoder.Decode(&obj.Proposer)
	if err != nil {
		return err
	}
	// Deserialize `DescriptionUrl`:
	err = decoder.Decode(&obj.DescriptionUrl)
	if err != nil {
		return err
	}
	// Deserialize `SlotEnqueued`:
	err = decoder.Decode(&obj.SlotEnqueued)
	if err != nil {
		return err
	}
	// Deserialize `State`:
	err = decoder.Decode(&obj.State)
	if err != nil {
		return err
	}
	// Deserialize `Instruction`:
	err = decoder.Decode(&obj.Instruction)
	if err != nil {
		return err
	}
	// Deserialize `OpenbookTwapPassMarket`:
	err = decoder.Decode(&obj.OpenbookTwapPassMarket)
	if err != nil {
		return err
	}
	// Deserialize `OpenbookTwapFailMarket`:
	err = decoder.Decode(&obj.OpenbookTwapFailMarket)
	if err != nil {
		return err
	}
	// Deserialize `OpenbookPassMarket`:
	err = decoder.Decode(&obj.OpenbookPassMarket)
	if err != nil {
		return err
	}
	// Deserialize `OpenbookFailMarket`:
	err = decoder.Decode(&obj.OpenbookFailMarket)
	if err != nil {
		return err
	}
	// Deserialize `BaseVault`:
	err = decoder.Decode(&obj.BaseVault)
	if err != nil {
		return err
	}
	// Deserialize `QuoteVault`:
	err = decoder.Decode(&obj.QuoteVault)
	if err != nil {
		return err
	}
	return nil
}
