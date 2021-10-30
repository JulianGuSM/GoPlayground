package main

import (
	"bytes"
	"cloud.google.com/go/spanner"
	"context"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/iterator"
	"math/rand"
	"strconv"
)

func main() {

	logger := initLogger()
	ctx := context.Background()
	const DbUrl = "projects/aftership-dev/instances/d-automizely-asea1/databases/cn-d-core"
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()

	arr, str := createData()
	err = WriteArrayToSpanner(client, ctx, arr)
	err = WriteStringToSpanner(client, ctx, str)

	sugar := logger.Sugar()

	if err != nil {
		sugar.Errorw("Failed write to Spanner",
			"error", err)
		return
	}
}

func initLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}

func WriteArrayToSpanner(client *spanner.Client, ctx context.Context, testArray []string) error {

	//stmt := spanner.Statement{
	//	SQL:    "SELECT ID, ArrayCol FROM testTabel limit 1",
	//	Params: nil,
	//}

	//iter := client.Single().Query(ctx, stmt)
	//defer iter.Stop()
	//row, err := iter.Next()
	//if err != nil {
	//	return err
	//}

	//if row != nil {
	_, err := client.ReadWriteTransaction(ctx,
		func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			//cols := []string{"ID", "ArrayCol"}

			stmt2 := spanner.Statement{
				SQL: `UPDATE testTabel SET ArrayCol = @ArrayCol WHERE ID = @ID`,
				Params: map[string]interface{}{
					"ArrayCol": testArray,
					"ID":       "1",
				},
			}

			_, err := txn.Update(ctx, stmt2)
			if err != nil {
				return err
			}
			//fmt.Println("write to spanner successfully!")
			return nil
		})
	//}

	return err
}

func WriteStringToSpanner(client *spanner.Client, ctx context.Context, testStr string) error {

	//stmt := spanner.Statement{
	//	SQL:    "SELECT ID, ArrayCol FROM testTabel limit 1",
	//	Params: nil,
	//}

	//iter := client.Single().Query(ctx, stmt)
	//defer iter.Stop()
	//row, err := iter.Next()
	//if err != nil {
	//	return err
	//}

	//if row != nil {
	_, err := client.ReadWriteTransaction(ctx,
		func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			//cols := []string{"ID", "ArrayCol"}

			stmt2 := spanner.Statement{
				SQL: `UPDATE testTabel SET StringCol = @StringCol WHERE ID = @ID`,
				Params: map[string]interface{}{
					"StringCol": testStr,
					"ID":        "1",
				},
			}

			_, err := txn.Update(ctx, stmt2)
			if err != nil {
				return err
			}
			//fmt.Println("write to spanner successfully!")
			return nil
		})
	//}

	return err
}

func CreateStringToSpanner(client *spanner.Client, ctx context.Context, testStr string) error {

	//stmt := spanner.Statement{
	//	SQL:    "SELECT ID, ArrayCol FROM testTabel limit 1",
	//	Params: nil,
	//}

	//iter := client.Single().Query(ctx, stmt)
	//defer iter.Stop()
	//row, err := iter.Next()
	//if err != nil {
	//	return err
	//}

	//if row != nil {
	_, err := client.ReadWriteTransaction(ctx,
		func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			//cols := []string{"ID", "ArrayCol"}

			stmt2 := spanner.Statement{
				SQL: `INSERT INTO testTabel (ID, StringCol) VALUES(@ID, @StringCol)`,
				Params: map[string]interface{}{
					"StringCol": testStr,
					"ID":        uuid.New().String(),
				},
			}

			_, err := txn.Update(ctx, stmt2)
			if err != nil {
				return err
			}
			//fmt.Println("write to spanner successfully!")
			return nil
		})
	//}

	return err
}

func CreateArrayToSpanner(client *spanner.Client, ctx context.Context, testArray []string) error {

	//stmt := spanner.Statement{
	//	SQL:    "SELECT ID, ArrayCol FROM testTabel limit 1",
	//	Params: nil,
	//}

	//iter := client.Single().Query(ctx, stmt)
	//defer iter.Stop()
	//row, err := iter.Next()
	//if err != nil {
	//	return err
	//}

	//if row != nil {
	_, err := client.ReadWriteTransaction(ctx,
		func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			//cols := []string{"ID", "ArrayCol"}

			stmt2 := spanner.Statement{
				SQL: `INSERT INTO testTabel (ID, ArrayCol) VALUES(@ID, @ArrayCol)`,
				Params: map[string]interface{}{
					"ArrayCol": testArray,
					"ID":       uuid.New().String(),
				},
			}

			_, err := txn.Update(ctx, stmt2)
			if err != nil {
				return err
			}
			//fmt.Println("write to spanner successfully!")
			return nil
		})
	//}

	return err
}

func SelectArrayFromSpanner(client *spanner.Client, ctx context.Context) error {

	//stmt := spanner.Statement{
	//	SQL:    "SELECT ID, ArrayCol FROM testTabel limit 1",
	//	Params: nil,
	//}

	//iter := client.Single().Query(ctx, stmt)
	//defer iter.Stop()
	//row, err := iter.Next()
	//if err != nil {
	//	return err
	//}

	//if row != nil {
	_, err := client.ReadWriteTransaction(ctx,
		func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			//cols := []string{"ID", "ArrayCol"}

			stmt2 := spanner.Statement{
				SQL: `SELECT ID, ArrayCol FROM testTabel WHERE ID = @ID`,
				Params: map[string]interface{}{
					"ID": "1",
				},
			}

			iter := txn.Query(ctx, stmt2)
			defer iter.Stop()
			for {
				_, err := iter.Next()
				if err == iterator.Done {
					return nil
				}
				if err != nil {
					return err
				}
				//var ArrayCol []string
				//var ID string
				//if err := row.Columns(&ID, &ArrayCol); err != nil {
				//	return err
				//}
				//fmt.Printf("ID: %s ArrayCol %v \n", ID, ArrayCol)
			}
		})
	//}

	return err
}

func SelectStringFromSpanner(client *spanner.Client, ctx context.Context) error {

	_, err := client.ReadWriteTransaction(ctx,
		func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
			//cols := []string{"ID", "ArrayCol"}

			stmt2 := spanner.Statement{
				SQL: `SELECT ID, StringCol FROM testTabel WHERE ID = @ID`,
				Params: map[string]interface{}{
					"ID": "1",
				},
			}

			iter := txn.Query(ctx, stmt2)
			defer iter.Stop()
			for {
				_, err := iter.Next()
				if err == iterator.Done {
					return nil
				}
				if err != nil {
					return err
				}

				//var ID, StringCol string
				//if err := row.Columns(&ID, &StringCol); err != nil {
				//	return err
				//}
				//fmt.Printf("ID: %s StringCol %v \n", ID, StringCol)
			}
		})

	return err
}

func createData() ([]string, string) {
	length := 100000

	var arr []string
	for i := 0; i < length; i++ {
		buffer := bytes.Buffer{}
		for j := 0; j < 13; j++ {

			buffer.Write([]byte(strconv.Itoa(rand.Intn(10))))
		}
		arr = append(arr, buffer.String())
	}
	result, _ := jsoniter.Marshal(arr)
	return arr, string(result)
}
