package main

import (
	"cloud.google.com/go/vision/apiv1"
	"io"
	"context"
	"os"
	"log"
	"fmt"
	"google.golang.org/api/option"
)

func detectTexts(ctx context.Context, r io.Reader) ([]string, error) {
	opts := option.WithCredentialsFile("secret.json")
	client, err := vision.NewImageAnnotatorClient(ctx, opts)
	if err != nil {
		return nil, err
	}

	image, err := vision.NewImageFromReader(r)
	if err != nil {
		return nil, err
	}

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return nil, err
	}
	texts := make([]string, len(annotations))
	for i, ann := range annotations {
		texts[i] = ann.Description
	}
	return texts, nil
}


func main() {
	f, err := os.Open("test.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	ctx := context.Background()
	texts, err := detectTexts(ctx, f)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v", texts)
}
