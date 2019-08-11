package nameservice

import(
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/types"
)

// toate mesajele trebuie sa implementeze interfata
type Msg interface {
	// Tipul mesajului.In functie de tip apelam handler ul corespunzator
	Type() string

	// ruta - nu am inteles exact
	Route() string

	// validare 
	ValidateBasic() types.Any

	// ?
	GetSignBytes() []byte

	// ? 
	GetSigners() []sdk.AccAddress
}
