package berlioz

import dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"

// THIS FILE IS GENERATED USING ROBOT
// DO NOT MODIFY THIS FILE DIRECTLY

// Wrapper for *dynamodb.DynamoDB::CreateBackup.
//   In: , Type: *dynamodb.CreateBackupInput
//   Out: , Type: *dynamodb.CreateBackupOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) CreateBackup(in1 *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.CreateBackup(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.CreateBackupOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::CreateTable.
//   In: , Type: *dynamodb.CreateTableInput
//   Out: , Type: *dynamodb.CreateTableOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) CreateTable(in1 *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.CreateTable(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.CreateTableOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::DeleteItem.
//   In: , Type: *dynamodb.DeleteItemInput
//   Out: , Type: *dynamodb.DeleteItemOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) DeleteItem(in1 *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.DeleteItem(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.DeleteItemOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::DeleteTable.
//   In: , Type: *dynamodb.DeleteTableInput
//   Out: , Type: *dynamodb.DeleteTableOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) DeleteTable(in1 *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.DeleteTable(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.DeleteTableOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::DescribeContinuousBackups.
//   In: , Type: *dynamodb.DescribeContinuousBackupsInput
//   Out: , Type: *dynamodb.DescribeContinuousBackupsOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) DescribeContinuousBackups(in1 *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.DescribeContinuousBackups(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.DescribeContinuousBackupsOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::DescribeTable.
//   In: , Type: *dynamodb.DescribeTableInput
//   Out: , Type: *dynamodb.DescribeTableOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) DescribeTable(in1 *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.DescribeTable(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.DescribeTableOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::DescribeTimeToLive.
//   In: , Type: *dynamodb.DescribeTimeToLiveInput
//   Out: , Type: *dynamodb.DescribeTimeToLiveOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) DescribeTimeToLive(in1 *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.DescribeTimeToLive(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.DescribeTimeToLiveOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::GetItem.
//   In: , Type: *dynamodb.GetItemInput
//   Out: , Type: *dynamodb.GetItemOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) GetItem(in1 *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.GetItem(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.GetItemOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::ListBackups.
//   In: , Type: *dynamodb.ListBackupsInput
//   Out: , Type: *dynamodb.ListBackupsOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) ListBackups(in1 *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.ListBackups(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.ListBackupsOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::PutItem.
//   In: , Type: *dynamodb.PutItemInput
//   Out: , Type: *dynamodb.PutItemOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) PutItem(in1 *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.PutItem(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.PutItemOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::Query.
//   In: , Type: *dynamodb.QueryInput
//   Out: , Type: *dynamodb.QueryOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) Query(in1 *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.Query(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.QueryOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::QueryPages.
//   In: , Type: *dynamodb.QueryInput
//   In: , Type: func(*dynamodb.QueryOutput, bool) bool
//   Out: error, Type: error
func (x DynamoDBAccessor) QueryPages(in1 *dynamodb.QueryInput, in2 func(*dynamodb.QueryOutput, bool) bool) error {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 0)
		err = client.QueryPages(in1, in2)
		return res, err
	}
	_, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return execErr
	}
	return nil
}

// Wrapper for *dynamodb.DynamoDB::Scan.
//   In: , Type: *dynamodb.ScanInput
//   Out: , Type: *dynamodb.ScanOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) Scan(in1 *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.Scan(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.ScanOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::ScanPages.
//   In: , Type: *dynamodb.ScanInput
//   In: , Type: func(*dynamodb.ScanOutput, bool) bool
//   Out: error, Type: error
func (x DynamoDBAccessor) ScanPages(in1 *dynamodb.ScanInput, in2 func(*dynamodb.ScanOutput, bool) bool) error {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 0)
		err = client.ScanPages(in1, in2)
		return res, err
	}
	_, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return execErr
	}
	return nil
}

// Wrapper for *dynamodb.DynamoDB::UpdateContinuousBackups.
//   In: , Type: *dynamodb.UpdateContinuousBackupsInput
//   Out: , Type: *dynamodb.UpdateContinuousBackupsOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) UpdateContinuousBackups(in1 *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.UpdateContinuousBackups(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.UpdateContinuousBackupsOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::UpdateItem.
//   In: , Type: *dynamodb.UpdateItemInput
//   Out: , Type: *dynamodb.UpdateItemOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) UpdateItem(in1 *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.UpdateItem(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.UpdateItemOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::UpdateTable.
//   In: , Type: *dynamodb.UpdateTableInput
//   Out: , Type: *dynamodb.UpdateTableOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) UpdateTable(in1 *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.UpdateTable(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.UpdateTableOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::UpdateTimeToLive.
//   In: , Type: *dynamodb.UpdateTimeToLiveInput
//   Out: , Type: *dynamodb.UpdateTimeToLiveOutput
//   Out: error, Type: error
func (x DynamoDBAccessor) UpdateTimeToLive(in1 *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 1)
		res[0], err = client.UpdateTimeToLive(in1)
		return res, err
	}
	execRes, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return nil, execErr
	}
	return execRes[0].(*dynamodb.UpdateTimeToLiveOutput), nil
}

// Wrapper for *dynamodb.DynamoDB::WaitUntilTableExists.
//   In: , Type: *dynamodb.DescribeTableInput
//   Out: error, Type: error
func (x DynamoDBAccessor) WaitUntilTableExists(in1 *dynamodb.DescribeTableInput) error {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 0)
		err = client.WaitUntilTableExists(in1)
		return res, err
	}
	_, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return execErr
	}
	return nil
}

// Wrapper for *dynamodb.DynamoDB::WaitUntilTableNotExists.
//   In: , Type: *dynamodb.DescribeTableInput
//   Out: error, Type: error
func (x DynamoDBAccessor) WaitUntilTableNotExists(in1 *dynamodb.DescribeTableInput) error {
	action := func(rawPeer interface{}) ([]interface{}, error) {
		client, peer, err := x.Client(rawPeer)
		if err != nil {
			return nil, err
		}
		in1.TableName = &peer.Name
		res := make([]interface{}, 0)
		err = client.WaitUntilTableNotExists(in1)
		return res, err
	}
	_, execErr := execute(x.info.kind, x.info.path, action)
	if execErr != nil {
		return execErr
	}
	return nil
}
