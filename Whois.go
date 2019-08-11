package nameservice
import (
sdk "github.com/cosmos/cosmos-sdk/types"
)

type Whois struct {
	Owner sdk.AccAddress // un tip din sdk. (!!!!)
	Price sdk.Coins //
}

func NewWhois() Whois {
	return Whois{}
}