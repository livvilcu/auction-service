package auctionservice


import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetOject(ctx, keeper, msg)
			// aici ar trebui sa vina si alte tipuri de mesaje, e.g. cel pentru cumparator 
			// cel pentru donatii
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsg{.Action}
func handleMsgSetOject(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.owner.Equals(keeper.GetOwner(ctx, msg.object)) {
		// cel care a trimis mesajul nu este acelasi ca si current owner => aruncam o eroare
	}
	keeper.SetObject(ctx, msg.owner, msg.object)
	return sdk.Result{}

}
