package dao

import "newsCenter/article/infrastructure/persistence/database/gorms"

type Transaction struct {
	conn gorms.DbConn
}

func (t *Transaction) Action(f func(conn gorms.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction() *Transaction {
	return &Transaction{
		conn: gorms.NewTran(),
	}
}
