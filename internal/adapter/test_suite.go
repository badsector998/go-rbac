package adapter

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/badsector998/go-rbac/internal/lib/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestSuite struct {
	client       *gorm.DB
	mockClient   sqlmock.Sqlmock
	txClient     *gorm.DB
	tx           transaction.Manager
	txMockClient sqlmock.Sqlmock
}

func (s *TestSuite) Setup(t *testing.T) {
	client, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to create sql mock client")
	}

	c, err := gorm.Open(postgres.New(postgres.Config{
		Conn: client,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("failed to open gorm client")
	}

	s.client = c
	s.mockClient = mock

	txClient, txMock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to create sql mock tx client")
	}

	cTx, err := gorm.Open(postgres.New(postgres.Config{
		Conn: txClient,
	}), &gorm.Config{})
	if err != nil {
		t.Errorf("failed to open gorm client tx")
	}

	s.txClient = cTx
	s.txMockClient = txMock

	s.tx = NewTransactionAdapter(cTx)
}

func (s *TestSuite) Teardown(t *testing.T) {
	s.mockClient.ExpectClose()
	s.client = nil
	err := s.mockClient.ExpectationsWereMet()
	if err != nil {
		t.Errorf("mock client expectation weren't met: %s", err.Error())
	}

	s.txMockClient.ExpectClose()
	s.txClient = nil
	err = s.txMockClient.ExpectationsWereMet()
	if err != nil {
		t.Errorf("mock client expectation weren't met: %s", err.Error())
	}
}
