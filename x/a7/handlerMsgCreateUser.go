package a7

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/a7/a7/x/a7/types"
	"github.com/a7/a7/x/a7/keeper"
)

func handleMsgCreateUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateUser) (*sdk.Result, error) {
	var user = types.User{
		Creator: msg.Creator,
		ID:      msg.ID,
    Name: msg.Name,
    Email: msg.Email,
	}
	k.CreateUser(ctx, user)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
