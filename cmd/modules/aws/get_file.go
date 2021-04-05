package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gonyyi/examples.go/x_already_exists_change_dt_func_name/plugin"
	"github.com/orangenumber/arand"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFile() {
	// test()
	// testGzip()
	// return

	msg, err := plugin.GetMessage()
	if err != nil {
		plugin.NewResponse().AsError(err.Error())
		return
	}

	s := strings.TrimSpace(strings.TrimPrefix(msg.Text(), "aws-get-sample "))

	// ss := strings.SplitN(s, " ", 2)
	ss := strings.SplitN(s, " ", 4)

	// if len(ss) < 2 {
	if len(ss) < 4 {
		plugin.NewResponse().AsError("Incorrect command. Usage: aws-get-sample <env> <path> <numOfBytes> <secret>\nEg: `aws-get-sample uk gonyi/abc.txt.gz 10240 secretKey`")
		return
	}

	if ss[3] != "gde12321" {
		plugin.NewResponse().AsError("Incorrect secret key used.")
		return
	}

	var bucket, filename, region string

	numOfBytes := 1024*4
	{
		tmp, err := strconv.Atoi(ss[2])
		if err == nil {
			if numOfBytes > 1024*512 {
				numOfBytes = 1024*512
			} else {
				numOfBytes = tmp
			}
		}
	}

	filename = ss[1]
	switch ss[0] {
	case "gbr", "uk":
		bucket = "dt-gbr-prod-ftp-7km3u"
		region = "eu-central-1"
	}
	if bucket == "" || region == "" || filename == "" {
		plugin.NewResponse().AsError("Incorrect command. Usage: aws-get-sample <env> <path> <numOfBytes> <secret>\nEg: `aws-get-sample uk gonyi/abc.txt.gz 10240 secretKeye`")
		return
	}

	out, err := awsDnHead(bucket, region, filename, numOfBytes)
	if err != nil {
		plugin.NewResponse().AsError(err.Error())
		return
	}

	tmpSaveAsName := arand.RandStr(arand.R_ALPHANUMERIC, 20)

	if strings.HasSuffix(filename, ".gz") {
		tmpout, _ := gzipToByte(out)
		out = tmpout
	}

	if err := ioutil.WriteFile(tmpSaveAsName, out, 0755); err != nil {
		plugin.NewResponse().AsError(err.Error())
		return
	}

	plugin.NewResponse().AddFiles(tmpSaveAsName).Send()
}

func test() {
	bucket := "dt-gbr-prod-ftp-7km3u"
	region := "eu-central-1"
	filename := "gonyi/private/test3.txt.gz"
	filename = "gonyi/private/test2.txt"

	out, err := awsDnHead(bucket, region, filename, 1024*4)
	if err != nil {
		println(err.Error())
	}
	ioutil.WriteFile("./aws_test/getfile/test2.txt.gz", out, 0755)
}

func testGzip() {
	fin, err := ioutil.ReadFile("./aws_test/getfile/test2.txt.gz")
	if err != nil {
		println(1, err.Error())
		return
	}

	fout, err := gzipToByte(fin)
	if err != nil {
		println("failed", err.Error())
		return
	}
	println("out:", string(fout))
}

func gzipToByte(g []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(g))
	if err != nil {
		return nil, err
	}
	out := bytes.Buffer{}

	buf := make([]byte, 512)
	for {
		n, err := r.Read(buf)
		out.WriteString(string(buf[:n]))
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return out.Bytes(), nil
		} else if err != nil {
			return out.Bytes(), err
		}
	}
	return out.Bytes(), nil
}

func awsDnHead(bucket, region, filename string, numOfBytes int) ([]byte, error) {
	return awsDownload(bucket, region, filename, 0, int64(numOfBytes))
}

// This needs to be updated at some point.
func awsDownload(bucket, region, filename string, rangeStart, rangeEnd int64) ([]byte, error) {
	sess, err := session.NewSession(&aws.Config{Region: &region})
	if err != nil {
		return nil, err
	}

	dn := s3manager.NewDownloader(sess)
	buf := aws.NewWriteAtBuffer([]byte{})

	var rge string
	if rangeEnd != 0 {
		rge = fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd)
	}

	_, err = dn.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Range:  &rge,
	})

	return buf.Bytes(), err
}
