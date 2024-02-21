// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_twap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// EditOrder is the `editOrder` instruction.
type EditOrder struct {
	ClientOrderId      *uint64
	ExpectedCancelSize *int64
	PlaceOrder         *PlaceOrderArgs

	// [0] = [SIGNER] signer
	//
	// [1] = [WRITE] openOrdersAccount
	//
	// [2] = [WRITE] twapMarket
	//
	// [3] = [WRITE] userTokenAccount
	//
	// [4] = [WRITE] market
	//
	// [5] = [WRITE] bids
	//
	// [6] = [WRITE] asks
	//
	// [7] = [WRITE] eventHeap
	//
	// [8] = [WRITE] marketVault
	//
	// [9] = [] tokenProgram
	//
	// [10] = [] openbookProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEditOrderInstructionBuilder creates a new `EditOrder` instruction builder.
func NewEditOrderInstructionBuilder() *EditOrder {
	nd := &EditOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetClientOrderId sets the "clientOrderId" parameter.
func (inst *EditOrder) SetClientOrderId(clientOrderId uint64) *EditOrder {
	inst.ClientOrderId = &clientOrderId
	return inst
}

// SetExpectedCancelSize sets the "expectedCancelSize" parameter.
func (inst *EditOrder) SetExpectedCancelSize(expectedCancelSize int64) *EditOrder {
	inst.ExpectedCancelSize = &expectedCancelSize
	return inst
}

// SetPlaceOrder sets the "placeOrder" parameter.
func (inst *EditOrder) SetPlaceOrder(placeOrder PlaceOrderArgs) *EditOrder {
	inst.PlaceOrder = &placeOrder
	return inst
}

// SetSignerAccount sets the "signer" account.
func (inst *EditOrder) SetSignerAccount(signer ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *EditOrder) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersAccountAccount sets the "openOrdersAccount" account.
func (inst *EditOrder) SetOpenOrdersAccountAccount(openOrdersAccount ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrdersAccount).WRITE()
	return inst
}

// GetOpenOrdersAccountAccount gets the "openOrdersAccount" account.
func (inst *EditOrder) GetOpenOrdersAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTwapMarketAccount sets the "twapMarket" account.
func (inst *EditOrder) SetTwapMarketAccount(twapMarket ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(twapMarket).WRITE()
	return inst
}

// GetTwapMarketAccount gets the "twapMarket" account.
func (inst *EditOrder) GetTwapMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserTokenAccountAccount sets the "userTokenAccount" account.
func (inst *EditOrder) SetUserTokenAccountAccount(userTokenAccount ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userTokenAccount).WRITE()
	return inst
}

// GetUserTokenAccountAccount gets the "userTokenAccount" account.
func (inst *EditOrder) GetUserTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMarketAccount sets the "market" account.
func (inst *EditOrder) SetMarketAccount(market ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *EditOrder) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetBidsAccount sets the "bids" account.
func (inst *EditOrder) SetBidsAccount(bids ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *EditOrder) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAsksAccount sets the "asks" account.
func (inst *EditOrder) SetAsksAccount(asks ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *EditOrder) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetEventHeapAccount sets the "eventHeap" account.
func (inst *EditOrder) SetEventHeapAccount(eventHeap ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(eventHeap).WRITE()
	return inst
}

// GetEventHeapAccount gets the "eventHeap" account.
func (inst *EditOrder) GetEventHeapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMarketVaultAccount sets the "marketVault" account.
func (inst *EditOrder) SetMarketVaultAccount(marketVault ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(marketVault).WRITE()
	return inst
}

// GetMarketVaultAccount gets the "marketVault" account.
func (inst *EditOrder) GetMarketVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *EditOrder) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *EditOrder) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetOpenbookProgramAccount sets the "openbookProgram" account.
func (inst *EditOrder) SetOpenbookProgramAccount(openbookProgram ag_solanago.PublicKey) *EditOrder {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(openbookProgram)
	return inst
}

// GetOpenbookProgramAccount gets the "openbookProgram" account.
func (inst *EditOrder) GetOpenbookProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst EditOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EditOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EditOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EditOrder) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ClientOrderId == nil {
			return errors.New("ClientOrderId parameter is not set")
		}
		if inst.ExpectedCancelSize == nil {
			return errors.New("ExpectedCancelSize parameter is not set")
		}
		if inst.PlaceOrder == nil {
			return errors.New("PlaceOrder parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Signer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OpenOrdersAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TwapMarket is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.EventHeap is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MarketVault is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.OpenbookProgram is not set")
		}
	}
	return nil
}

func (inst *EditOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EditOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     ClientOrderId", *inst.ClientOrderId))
						paramsBranch.Child(ag_format.Param("ExpectedCancelSize", *inst.ExpectedCancelSize))
						paramsBranch.Child(ag_format.Param("        PlaceOrder", *inst.PlaceOrder))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         signer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     openOrders", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     twapMarket", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      userToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         market", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           bids", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           asks", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("      eventHeap", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    marketVault", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("openbookProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj EditOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ClientOrderId` param:
	err = encoder.Encode(obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Serialize `ExpectedCancelSize` param:
	err = encoder.Encode(obj.ExpectedCancelSize)
	if err != nil {
		return err
	}
	// Serialize `PlaceOrder` param:
	err = encoder.Encode(obj.PlaceOrder)
	if err != nil {
		return err
	}
	return nil
}
func (obj *EditOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ClientOrderId`:
	err = decoder.Decode(&obj.ClientOrderId)
	if err != nil {
		return err
	}
	// Deserialize `ExpectedCancelSize`:
	err = decoder.Decode(&obj.ExpectedCancelSize)
	if err != nil {
		return err
	}
	// Deserialize `PlaceOrder`:
	err = decoder.Decode(&obj.PlaceOrder)
	if err != nil {
		return err
	}
	return nil
}

// NewEditOrderInstruction declares a new EditOrder instruction with the provided parameters and accounts.
func NewEditOrderInstruction(
	// Parameters:
	clientOrderId uint64,
	expectedCancelSize int64,
	placeOrder PlaceOrderArgs,
	// Accounts:
	signer ag_solanago.PublicKey,
	openOrdersAccount ag_solanago.PublicKey,
	twapMarket ag_solanago.PublicKey,
	userTokenAccount ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	eventHeap ag_solanago.PublicKey,
	marketVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	openbookProgram ag_solanago.PublicKey) *EditOrder {
	return NewEditOrderInstructionBuilder().
		SetClientOrderId(clientOrderId).
		SetExpectedCancelSize(expectedCancelSize).
		SetPlaceOrder(placeOrder).
		SetSignerAccount(signer).
		SetOpenOrdersAccountAccount(openOrdersAccount).
		SetTwapMarketAccount(twapMarket).
		SetUserTokenAccountAccount(userTokenAccount).
		SetMarketAccount(market).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetEventHeapAccount(eventHeap).
		SetMarketVaultAccount(marketVault).
		SetTokenProgramAccount(tokenProgram).
		SetOpenbookProgramAccount(openbookProgram)
}
