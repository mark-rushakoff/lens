package query

import (
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/strangelove-ventures/lens/client"
)

type Query struct {
	Client  *client.ChainClient
	Options *QueryOptions
}

func (q *Query) Balances(address string) (*bankTypes.QueryAllBalancesResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return BalanceWithAddressRPC(q, address)
}

func (q *Query) TotalSupply() (*bankTypes.QueryTotalSupplyResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return TotalSupplyRPC(q)
}

func (q *Query) DenomsMetadata() (*bankTypes.QueryDenomsMetadataResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return DenomsMetadataRPC(q)
}

func (q *Query) Delegation(delegator, validator string) (*types.DelegationResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return Delegation(q, delegator, validator)
}

func (q *Query) Delegations(delegator string) (*types.QueryDelegatorDelegationsResponse, error) {
	/// TODO: In the future have some logic to route the query to the appropriate client (gRPC or RPC)
	return Delegations(q, delegator)
}
