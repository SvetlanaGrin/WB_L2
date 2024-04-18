package main

// FileScheme represents the implementation of a URL scheme and gives access to
// fetching files of that scheme.
//
// For example, an http FileScheme implementation would fetch files using
// the HTTP protocol.

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"pack.ag/tftp"
	"strings"
)

type Reader interface {
	ReadData(ctx context.Context, u *url.URL) error
}

func Read(r Reader, ctx context.Context, u *url.URL) error {
	return r.ReadData(ctx, u)
}

var (
	DefaultTFTPClient = NewTFTPClient(tftp.ClientMode(tftp.ModeOctet), tftp.ClientBlocksize(1450), tftp.ClientWindowsize(65535))
)

type HTTPRequest struct {
	u *url.URL
}

func newHTTPRequest() *HTTPRequest {
	return &HTTPRequest{}
}
func (h *HTTPRequest) ReadData(ctx context.Context, u *url.URL) error {
	req, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer req.Body.Close()
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", b)
	return nil
}

type FileRequest struct {
	buf bytes.Buffer
}

func newFileRequest() *FileRequest {
	return &FileRequest{}
}

func (f *FileRequest) ReadData(ctx context.Context, u *url.URL) error {
	file, err := os.Open(u.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (h *TFTPClient) ReadData(ctx context.Context, u *url.URL) error {
	client, err := tftp.NewClient()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	resp, err := client.Get(u.String())
	if err != nil {
		log.Fatalln(err)
		return err
	}

	_, err = io.Copy(os.Stdout, resp)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

type TFTPClient struct {
	opts []tftp.ClientOpt
}

func NewTFTPClient(opts ...tftp.ClientOpt) *TFTPClient {
	return &TFTPClient{
		opts: opts,
	}
}

var (
	outPath = flag.String("O", "", "output file")
)

func usage() {
	log.Printf("Usage: %s [ARGS] URL\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	if os.Args[1] != "wget:" {
		usage()
	}
	argURL := os.Args[2]
	if argURL == "" {
		log.Fatalln("Empty URL")
	}

	url, err := url.Parse(argURL)
	if err != nil {
		log.Fatalln(err)
	}
	if strings.HasPrefix(url.String(), "http") || strings.HasPrefix(url.String(), "https") {
		ht := newHTTPRequest()
		ht.ReadData(context.Background(), url)
	} else if strings.HasPrefix(url.String(), "tftp") {
		tf := DefaultTFTPClient
		tf.ReadData(context.Background(), url)
	} else {
		file := newFileRequest()
		file.ReadData(context.Background(), url)
	}

}
