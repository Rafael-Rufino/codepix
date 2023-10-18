package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixkeyRepositoryInterface interface {
	RegisterKey(pixkey *Pixkey) (*Pixkey, error)
	FindKeyByKind(key string, kind string) (*Pixkey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type Pixkey struct {
	Base `valid:"required"`
	Kind string `json:"kind" valid:"notnull"`
	Key string `json:"key" valid:"notnull"`
	AccountID string `json:"account_id" valid:"notnull"`
	Account *Account `valid:"-"`
	Status string `json:"status" valid:"notnull"`

}

func (pixkey *Pixkey) isValid() error {
	_, err := govalidator.ValidateStruct(pixkey)

	if pixkey.Kind != "email" && pixkey.Kind != "cpf" {
		return errors.New("Invalid type of key")
	}

	if pixkey.Status != "active" && pixkey.Status != "inactive" {
		return errors.New("Invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixkey(kind string, key string, account *Account) (*Pixkey, error) {
	pixkey := Pixkey{
		Kind: kind,
		Key: key,
		Account: account,
		Status: "active",
	}

	pixkey.ID =  uuid.NewV4().String()
	pixkey.CreatedAt = time.Now()

	err := pixkey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixkey, nil
}