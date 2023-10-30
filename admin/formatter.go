package admin

import (
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"time"
)

type UserFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatUser(user user.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Occupation = user.Occupation
	userFormatter.Email = user.Email
	userFormatter.Avatar = user.Avatar
	userFormatter.Role = user.Role
	userFormatter.IsVerified = user.IsVerified
	userFormatter.CreatedAt = user.CreatedAt

	return userFormatter
}

func FormatterUsers(users []user.User) []UserFormatter {
	var usersFormatter []UserFormatter

	for _, user := range users {
		formatUser := FormatUser(user)
		usersFormatter = append(usersFormatter, formatUser)
	}

	return usersFormatter
}

type TransactionFormatter struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	CampaignID int       `json:"campaign_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentURL string    `json:"payment_url"`
	User       string    `json:"user"`
	Campaign   string    `json:"campaign_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatTransaction(transaction transaction.Transaction) TransactionFormatter {
	transactionFormatter := TransactionFormatter{}
	transactionFormatter.ID = transaction.ID
	transactionFormatter.UserID = transaction.UserID
	transactionFormatter.CampaignID = transaction.CampaignID
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.Status = transaction.Status
	transactionFormatter.Code = transaction.Code
	transactionFormatter.PaymentURL = transaction.PaymentURL
	transactionFormatter.User = transaction.User.Name
	transactionFormatter.Campaign = transaction.Campaign.Name
	transactionFormatter.CreatedAt = transaction.CreatedAt
	transactionFormatter.UpdatedAt = transaction.UpdatedAt

	return transactionFormatter
}

func FormatterTransactions(transaction []transaction.Transaction) []TransactionFormatter {
	var transactionsFormatter []TransactionFormatter

	for _, tr := range transaction {
		formatTransaction := FormatTransaction(tr)
		transactionsFormatter = append(transactionsFormatter, formatTransaction)
	}

	return transactionsFormatter
}
