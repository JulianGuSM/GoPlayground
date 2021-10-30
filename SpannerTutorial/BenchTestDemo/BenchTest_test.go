package main

import (
	"bytes"
	"cloud.google.com/go/spanner"
	"compress/gzip"
	"context"
	"encoding/base64"
	jsoniter "github.com/json-iterator/go"
	"log"
	"testing"
)

const DbUrl = "projects/aftership-dev/instances/d-automizely-asea1/databases/cn-d-core"

func BenchmarkWriteArrayToSpanner(b *testing.B) {
	testArray, _ := createData()
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}

	defer client.Close()
	j, _ := jsoniter.Marshal(testArray)
	b.Logf("length: %.5f M", float64(len(j))/1024.00/1024.00)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		err := WriteArrayToSpanner(client, ctx, testArray)
		if err != nil {
			b.Error("error: ", err)
			return
		}
	}
}

func BenchmarkParallelWriteArrayToSpanner(b *testing.B) {
	testArray, _ := createData()
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}

	defer client.Close()
	j, _ := jsoniter.Marshal(testArray)
	b.Logf("length: %.5f M", float64(len(j))/1024.00/1024.00)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := WriteArrayToSpanner(client, ctx, testArray)
			if err != nil {
				b.Error("error: ", err)
				return
			}
		}
	})

}

func BenchmarkWriteStringToSpanner(b *testing.B) {
	_, testStr := createData()
	//fmt.Printf("string size: %d", float64() len(testStr)
	b.Logf("length: %.5f M", float64(len(testStr))/1024.00/1024.00)

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		err := WriteStringToSpanner(client, ctx, testStr)
		if err != nil {
			b.Log("error", err)
			return
		}
	}
}

func BenchmarkParallelWriteStringToSpanner(b *testing.B) {
	_, testStr := createData()
	//fmt.Printf("string size: %d", float64() len(testStr)
	b.Logf("length: %.5f M", float64(len(testStr))/1024.00/1024.00)

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := WriteStringToSpanner(client, ctx, testStr)
			if err != nil {
				b.Log("error", err)
				return
			}
		}
	})

}

func BenchmarkWriteZipStringToSpanner(b *testing.B) {
	_, testStr := createData()
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	_, err := zw.Write([]byte(testStr))
	if err := zw.Flush(); err != nil {
		panic(err)
	}
	if err := zw.Close(); err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	b.Logf("gzip length: %.5f M", float64(len(buf.Bytes()))/1024.00/1024.00)

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	//fmt.Printf("string size: %d", len(encoded))
	b.Log("encode length:", float64(len(encoded))/1024.00/1024.00, " M")

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := WriteStringToSpanner(client, ctx, encoded)
		if err != nil {
			b.Log("error: ", err)
			return
		}
	}
}

func BenchmarkParallelWriteZipStringToSpanner(b *testing.B) {
	_, testStr := createData()
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	_, err := zw.Write([]byte(testStr))
	if err := zw.Flush(); err != nil {
		panic(err)
	}
	if err := zw.Close(); err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	b.Logf("gzip length: %.5f M", float64(len(buf.Bytes()))/1024.00/1024.00)

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	//fmt.Printf("string size: %d", len(encoded))
	b.Log("encode length:", float64(len(encoded))/1024.00/1024.00, " M")

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {

	})

}

func BenchmarkCreateZipStringToSpanner(b *testing.B) {
	_, testStr := createData()
	var buf bytes.Buffer

	zw := gzip.NewWriter(&buf)
	_, err := zw.Write([]byte(testStr))
	if err := zw.Flush(); err != nil {
		panic(err)
	}
	if err := zw.Close(); err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("string size: %d", float64() len(testStr)
	b.Logf("zip string length: %.5f M", float64(len(buf.Bytes()))/1024.00/1024.00)
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := CreateStringToSpanner(client, ctx, encoded)
		if err != nil {
			b.Log("error: ", err)
			return
		}
	}
}

func BenchmarkCreateStringToSpanner(b *testing.B) {
	_, testStr := createData()
	//fmt.Printf("string size: %d", float64() len(testStr)
	b.Logf("stirng length: %.5f M", float64(len(testStr))/1024.00/1024.00)

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}
	defer client.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := CreateStringToSpanner(client, ctx, testStr)
		if err != nil {
			b.Log("error: ", err)
			return
		}
	}
}

func BenchmarkCreateArrayToSpanner(b *testing.B) {
	testArray, _ := createData()
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}

	defer client.Close()
	j, _ := jsoniter.Marshal(testArray)
	b.Logf("array length: %.5f M", float64(len(j))/1024.00/1024.00)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := CreateArrayToSpanner(client, ctx, testArray)
			if err != nil {
				b.Log("error: ", err)
				return
			}
		}
	})
}

func BenchmarkSelectArrayFromSpanner(b *testing.B) {

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}

	defer client.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := SelectArrayFromSpanner(client, ctx)
		if err != nil {
			return
		}
	}
}

func BenchmarkSelectStringFromSpanner(b *testing.B) {

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, DbUrl)
	if err != nil {
		return
	}

	defer client.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := SelectStringFromSpanner(client, ctx)
		if err != nil {
			return
		}
	}
}
