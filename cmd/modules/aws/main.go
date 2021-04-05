package main



func main() {
	// bucket := "dt-gbr-prod-ftp-7km3u"
	// region := "eu-central-1"

	//
	// size, err := awsSize(bucket, region, "gonyi/private/test.in/.dat")
	// if err != nil {
	// 	println(err.Error())
	// }
	//
	// println("size:", size)

	println(splitAwsPath("s3://dt-gbr-prod-ftp-7km3u/naudin/outbound/Test_QIN001_Output_QA_Step2_20210318_21.pip.gz"))

	println(splitAwsPath("naudin/outbound/Test_QIN001_Output_QA_Step2_20210318_21.pip.gz"))


	// size, err := awsRm(bucket, region, "gonyi/private/dob-cl-output.csv")
	// if err != nil {
	// 	println(err.Error())
	// }
	// println("size", size)


	// res, err := awsLs(bucket, region, "gonyi/outbound/Test_QIN001_Output_QA_Step2_20210222_20.pip")
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }
	// for _, v := range res {
	// 	println(v.String())
	// }
}