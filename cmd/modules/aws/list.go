package main
//
// import (
// 	"fmt"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	"github.com/gonyyi/afmt"
// 	"github.com/gonyyi/examples.go/x_already_exists_change_dt_func_name/plugin"
// 	"strings"
// )
//
// // default aws credential location: ~/.aws/credentials
//
// func list() {
// 	msg, err := plugin.GetMessage()
// 	if err != nil {
// 		plugin.NewResponse().AsError(err.Error())
// 		return
// 	}
//
// 	s := strings.TrimSpace(strings.TrimPrefix(msg.Text(), "aws-dir "))
// 	ss := strings.SplitN(s, " ", 2)
//
// 	if len(ss) != 2 {
// 		plugin.NewResponse().AsError("Incorrect command. Usage: aws-dir <env> <path>\nEg: `aws-dir uk gonyi/`")
// 		return
// 	}
//
// 	var bucket, prefix, region string
// 	prefix = ss[1]
// 	switch ss[0] {
// 	case "gbr", "uk":
// 		bucket = "dt-gbr-prod-ftp-7km3u"
// 		region = "eu-central-1"
// 	}
// 	if bucket == "" || region == "" ||  prefix == "" {
// 		plugin.NewResponse().AsError("Incorrect command. Usage: aws-dir <env> <path>\nEg: `aws-dir uk gonyi/`")
// 		return
// 	}
//
// 	out, err := PrintAwsDir(bucket, region, prefix)
// 	if err != nil {
// 		plugin.NewResponse().AsError(err.Error())
// 		return
// 	}
//
// 	plugin.NewResponse().AsText("```\n" + out + "\n```").Send()
// }
//
// func PrintAwsDir(bucket, region, prefix string) (string, error) {
// 	out, err := awsDir(bucket, region, prefix)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	maxNameLength := 0
// 	maxSizeLength := 0
//
// 	for _, v := range out {
// 		if tmpLen := len(v.Name); tmpLen > maxNameLength {
// 			maxNameLength = tmpLen
// 		}
// 		if tmpLen := len(afmt.HumanBytes(v.Size, 1)); tmpLen > maxSizeLength {
// 			maxSizeLength = tmpLen
// 		}
// 	}
// 	outFmt := fmt.Sprintf(`%%-%ds | %%%ds | %%s`, maxNameLength, maxSizeLength)
//
// 	var finalOut []string
//
// 	for _, v := range out {
// 		finalOut = append(finalOut, fmt.Sprintf(outFmt, v.Name, afmt.HumanBytes(v.Size, 1), v.LastModified.Local().Format("2006-01-02 15:04:05")))
// 	}
// 	finalOut = append(finalOut, fmt.Sprintf("---\nTotal: %d files", len(out)))
//
// 	return strings.Join(finalOut, "\n"), nil
// }
// //
// // type awsFiles struct {
// // 	Name         string
// // 	LastModified time.Time
// // 	Size         int64
// // 	StorageClass string
// // }
//
// func awsDir(bucket, region, prefix string) ([]awsFiles, error) {
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: &region},
// 	)
//
// 	// Create S3 service client
// 	svc := s3.New(sess)
//
// 	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket), Prefix: &prefix})
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var out []awsFiles
//
// 	for _, item := range resp.Contents {
// 		out = append(out, awsFiles{
// 			Name:         strings.TrimPrefix(*item.Key, prefix),
// 			LastModified: *item.LastModified,
// 			Size:         *item.Size,
// 			StorageClass: *item.StorageClass,
// 		})
// 	}
// 	return out, nil
// }
