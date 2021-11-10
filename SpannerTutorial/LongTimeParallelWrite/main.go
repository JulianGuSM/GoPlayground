package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/iterator"
)

func main() {
	logger := initLogger()
	//sugar := logger.Sugar()
	ctx := context.Background()
	const DbUrl = "projects/aftership-dev/instances/d-automizely-asea1/databases/cn-d-core"
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()

	var avgConsumeTime int64 = 0

	var wg sync.WaitGroup

	// 第一个命令行参数表示并发量
	parallelNum, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}
	// 第二个命令行参数表示数据量,xxx 条 products_id
	dataNum, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		panic(err)
		return
	}
	limit := make(chan struct{}, parallelNum)

	arr, _ := createData(int(dataNum))
	j, _ := jsoniter.Marshal(arr)
	logger.Info("array size", zap.String("size(M)", fmt.Sprintf("%.4f", float64(len(j))/1024.00/1024.00)))

	_, testStr := createData(int(dataNum))
	logger.Info("string size", zap.String("size(M)", fmt.Sprintf("%.4f", float64(len(testStr))/1024.00/1024.00)))

	zipStr := zip(testStr)
	logger.Info("zip string size", zap.String("size(M)", fmt.Sprintf("%.4f", float64(len(zipStr))/1024.00/1024.00)))

	s := time.Now()
	var id int
	for {
		id++
		wg.Add(1)
		limit <- struct{}{}
		go func(i int) {
			defer wg.Done()
			start := time.Now()
			//err := CreateArrayToSpanner(client, ctx, arr)
			//err := CreateStringToSpanner(client, ctx, testStr)
			err := CreateStringToSpanner(client, ctx, zipStr)
			if err != nil {
				logger.Error("error: %v", zap.Error(err))
				return
			}
			costTime := time.Now().Sub(start).Milliseconds()
			avgConsumeTime = avgConsumeTime + costTime
			logger.Info("create", zap.Int("num", i), zap.Int64("execution time(ms)", costTime))
			<-limit
		}(id)
	}

	wg.Wait()
	logger.Info("create", zap.Float64("avg execution time(ms)", float64(avgConsumeTime/(parallelNum*10))))
	logger.Info("create", zap.Float64("execution time(s)", time.Now().Sub(s).Seconds()))

}

func zip(str string) string {

	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	_, err := zw.Write([]byte(str))
	if err := zw.Flush(); err != nil {
		panic(err)
	}
	if err := zw.Close(); err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encoded
}

func waitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
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
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

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

func createData(length int) ([]string, string) {

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
