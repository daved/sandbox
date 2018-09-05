package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

const (
	cmdCreate = "create"
	cmdUpdate = "update"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}

func run() error {
	cmd := os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)[0 : len(os.Args)-1 : len(os.Args)-1]

	c := newConfig()
	c.attachFlags(flag.Parse)
	if err := c.validate(cmd); err != nil {
		return err
	}

	sessOpts := session.Options{SharedConfigState: session.SharedConfigEnable}
	sess, err := session.NewSessionWithOptions(sessOpts)
	if err != nil {
		return err
	}

	switch cmd {
	case cmdCreate:
		return createFunction(c, sess)

	case cmdUpdate:
		return updateFunction(c, sess)

	default:
		return fmt.Errorf("unknown/missing command %q", cmd)
	}
}

type config struct {
	fn      string
	zip     string
	region  string
	bucket  string
	handler string
	rsrcArn string
	runtime string
}

func newConfig() *config {
	return &config{
		zip:     "main",
		region:  "us-west-2",
		runtime: "go",
	}
}

func (c *config) attachFlags(parseFn func()) {
	flag.StringVar(&c.fn, "fn", c.fn, "The name of the Lambda function")
	flag.StringVar(&c.zip, "zip", c.zip, "The name of the ZIP file (without the .zip extension)")
	flag.StringVar(&c.region, "reg", c.region, "The name of the ZIP file (without the .zip extension)")
	flag.StringVar(&c.bucket, "bkt", c.bucket, "the name of bucket to which the ZIP file is uploaded")
	flag.StringVar(&c.handler, "hdl", c.handler, "The name of the package.class handling the call")
	flag.StringVar(&c.rsrcArn, "arn", c.rsrcArn, "The ARN of the role that calls the function")
	flag.StringVar(&c.runtime, "rt", c.runtime, "The runtime for the function.")
	if parseFn != nil {
		parseFn()
	}
}

func (c *config) validate(cmd string) error {
	switch cmd {
	case cmdCreate, cmdUpdate:
	default:
		return fmt.Errorf("unknown/missing command %q", cmd)
	}

	werr := func(msg string) error {
		return fmt.Errorf("validate config: %s", msg)
	}

	if c.zip == "" {
		return werr("zip file value is required")
	}
	if c.fn == "" {
		return werr("function name is required")
	}

	if cmd == cmdCreate {
		if c.region == "" {
			return werr("region name is required")
		}
		if c.bucket == "" {
			return werr("bucket name is required")
		}
		if c.handler == "" {
			return werr("handler name is required")
		}
		if c.rsrcArn == "" {
			return werr("resource arn is required")
		}
		if c.runtime == "" {
			return werr("runtime type is required")
		}
	}

	return nil
}

func createFunction(cnf *config, sess *session.Session) error {
	svc := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

	contents, err := ioutil.ReadFile(cnf.zip + ".zip")

	if err != nil {
		fmt.Println("Could not read " + cnf.zip + ".zip")
		os.Exit(0)
	}

	createCode := &lambda.FunctionCode{
		S3Bucket:        aws.String(cnf.bucket),
		S3Key:           aws.String(cnf.zip),
		S3ObjectVersion: aws.String(""),
		ZipFile:         contents,
	}

	createArgs := &lambda.CreateFunctionInput{
		Code:         createCode,
		FunctionName: aws.String(cnf.fn),
		Handler:      aws.String(cnf.handler),
		Role:         aws.String(cnf.rsrcArn),
		Runtime:      aws.String(cnf.runtime),
	}

	result, err := svc.CreateFunction(createArgs)

	if err != nil {
		fmt.Println("Cannot create function: " + err.Error())
	} else {
		fmt.Println(result)
	}

	return nil
}

func updateFunction(cnf *config, sess *session.Session) error {
	return nil
}
