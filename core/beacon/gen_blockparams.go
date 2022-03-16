// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package beacon

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*payloadAttributesMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (p PayloadAttributesV1) MarshalJSON() ([]byte, error) {
	type PayloadAttributesV1 struct {
		Timestamp             hexutil.Uint64  `json:"timestamp"     gencodec:"required"`
		Random                common.Hash     `json:"prevRandao"        gencodec:"required"`
		SuggestedFeeRecipient common.Address  `json:"suggestedFeeRecipient"  gencodec:"required"`
		Transactions          []hexutil.Bytes `json:"transactions,omitempty"  gencodec:"optional"`
		NoTxPool              bool            `json:"noTxPool,omitempty" gencodec:"optional"`
	}
	var enc PayloadAttributesV1
	enc.Timestamp = hexutil.Uint64(p.Timestamp)
	enc.Random = p.Random
	enc.SuggestedFeeRecipient = p.SuggestedFeeRecipient
	if p.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(p.Transactions))
		for k, v := range p.Transactions {
			enc.Transactions[k] = v
		}
	}
	enc.NoTxPool = p.NoTxPool
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (p *PayloadAttributesV1) UnmarshalJSON(input []byte) error {
	type PayloadAttributesV1 struct {
		Timestamp             *hexutil.Uint64 `json:"timestamp"     gencodec:"required"`
		Random                *common.Hash    `json:"prevRandao"        gencodec:"required"`
		SuggestedFeeRecipient *common.Address `json:"suggestedFeeRecipient"  gencodec:"required"`
		Transactions          []hexutil.Bytes `json:"transactions,omitempty"  gencodec:"optional"`
		NoTxPool              *bool           `json:"noTxPool,omitempty" gencodec:"optional"`
	}
	var dec PayloadAttributesV1
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for PayloadAttributesV1")
	}
	p.Timestamp = uint64(*dec.Timestamp)
	if dec.Random == nil {
		return errors.New("missing required field 'prevRandao' for PayloadAttributesV1")
	}
	p.Random = *dec.Random
	if dec.SuggestedFeeRecipient == nil {
		return errors.New("missing required field 'suggestedFeeRecipient' for PayloadAttributesV1")
	}
	p.SuggestedFeeRecipient = *dec.SuggestedFeeRecipient
	if dec.Transactions != nil {
		p.Transactions = make([][]byte, len(dec.Transactions))
		for k, v := range dec.Transactions {
			p.Transactions[k] = v
		}
	}
	if dec.NoTxPool != nil {
		p.NoTxPool = *dec.NoTxPool
	}
	return nil
}
