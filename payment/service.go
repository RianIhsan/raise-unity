package payment

import (
	"github.com/RianIhsan/raise-unity/user"
	"github.com/veritrans/go-midtrans"
	"os"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}
func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	server := os.Getenv("MSERVER")
	client := os.Getenv("MCLIENT")
	midclient.ServerKey = server
	midclient.ClientKey = client
	midclient.APIEnvType = midtrans.Sandbox
	orderID := strconv.Itoa(transaction.ID)
	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(transaction.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
