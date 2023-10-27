package transaction

import "github.com/RianIhsan/raise-unity/user"

type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
