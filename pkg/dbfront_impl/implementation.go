package dbfront

import (
	"context"
	"database/sql"
	"strconv"

	dbfront "github.com/anz-bank/test-rig-example/gen/dbfront"
)

func GetEndpoint(ctx context.Context, req *dbfront.GetEndpointRequest, client dbfront.GetEndpointClient) (*dbfront.StatusMsg, error) {
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return nil, err
	}
	statement := `SELECT txt FROM data WHERE id=$1;`
	var txt string
	row := client.conn.QueryRowContext(ctx, statement, id)
	switch err := row.Scan(&txt); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return &dbfront.StatusMsg{Status: txt}, nil
	default:
		return nil, err
	}

	return nil, nil
}
func PostEndpointWithArg(ctx context.Context, req *dbfront.PostEndpointWithArgRequest, client dbfront.PostEndpointWithArgClient) (*dbfront.StatusMsg, error) {
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		return nil, err
	}
	statement := `INSERT INTO data (id, txt) VALUES($1, $2)`
	_, err = client.conn.ExecContext(ctx, statement, id, "kekeke")
	if err != nil {
		return nil, err
	}
	return &dbfront.StatusMsg{Status: "done"}, nil
}

func GetServiceInterfaceImplementation() dbfront.ServiceInterface {
	impl := dbfront.ServiceInterface{}
	impl.GetEndpoint = GetEndpoint
	impl.PostEndpointWithArg = PostEndpointWithArg
	return impl
}
