// <snippet_imports>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
)

// </snippet_imports>

/*  The examples in this quickstart are for the Computer Vision API for Microsoft
 *  Cognitive Services with the following tasks:
 *  - Describing images
 *  - Categorizing images
 *  - Tagging images
 *  - Detecting faces
 *  - Detecting adult or racy content
 *  - Detecting the color scheme
 *  - Detecting domain-specific content (celebrities/landmarks)
 *  - Detecting image types (clip art/line drawing)
 *  - Detecting objects
 *  - Detecting brands
 *  - Generate Thumbnail
 *  - Recognizing printed and handwritten text with the Batch Read API
 *	- Recognizing printed text with OCR
 *
 *  Prerequisites:
 *    Import the required libraries. From the command line, you will need to 'go get'
 *    the azure-sdk-for-go and go-autorest packages from Github.
 *    For example:
 *	  go get github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision
 *
 *    Download images faces.jpg, handwritten_text.jpg, objects.jpg, cheese_clipart.png,
 *    printed_text.jpg, and gray-shirt-logo.jpg, then add to your root folder from here:
 *    https://github.com/Azure-Samples/cognitive-services-sample-data-files/tree/master/ComputerVision/Images
 *
 *    Add your Azure Computer Vision subscription key and endpoint to your environment variables with names:
 *    COMPUTER_VISION_SUBSCRIPTION_KEY and COMPUTER_VISION_ENDPOINT
 *
 *  How to run:
 *	  From command line: go run ComptuerVisionQuickstart.go
 *
 *  References:
 *    - SDK reference:
 *      https://godoc.org/github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision
 *    - Computer Vision documentation:
 * 		https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/index
 *    - Computer Vision API:
 *      https://westus.dev.cognitive.microsoft.com/docs/services/5cd27ec07268f6c679a3e641/operations/56f91f2e778daf14a499f21b
 */

// Declare global so don't have to pass it to all of the tasks.
var computerVisionContext context.Context

// </snippet_context>

func main() {
	/*
	 * Configure the Computer Vision client
	 * Set environment variables for COMPUTER_VISION_SUBSCRIPTION_KEY and COMPUTER_VISION_ENDPOINT,
	 * then restart your command shell or your IDE for changes to take effect.
	 */
	computerVisionKey := os.Getenv("COMPUTER_VISION_SUBSCRIPTION_KEY")

	if computerVisionKey == "" {
		log.Fatal("\n\nPlease set a COMPUTER_VISION_SUBSCRIPTION_KEY environment variable.\n" +
			"**You may need to restart your shell or IDE after it's set.**\n")
	}

	endpointURL := "https://425streamjan.cognitiveservices.azure.com/"

	computerVisionClient := computervision.New(endpointURL)
	// We should use the BearerAuthorizer here:
	// https://pkg.go.dev/github.com/Azure/go-autorest/autorest#NewBearerAuthorizer
	//
	// The adal.OauthTokenProvider will need to be adapted to work with MSAL:
	// https://pkg.go.dev/github.com/Azure/go-autorest/autorest/adal#OAuthTokenProvider
	// computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)
	oauthTokenProvider := new(tokenProvider)
	computerVisionClient.Authorizer = autorest.NewBearerAuthorizer(oauthTokenProvider)

	computerVisionContext = context.Background()

	// printedImageURL := "https://res.cloudinary.com/demo/image/upload/sample_document.docx.png"
	printedImageURL := "https://www.velvetjobs.com/resume/document-processing-resume-sample.jpg"
	RecognizePrintedOCRRemoteImage(computerVisionClient, printedImageURL)

	fmt.Println("-----------------------------------------")
	fmt.Println("End of quickstart.")
}

/*
 *  Recognize Printed Text with OCR - remote
 */
func RecognizePrintedOCRRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	fmt.Println("-----------------------------------------")
	fmt.Println("RECOGNIZE PRINTED TEXT - remote")
	fmt.Println()
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	fmt.Println("Recognizing text in a remote image with OCR ...")
	ocrResult, err := client.RecognizePrintedText(computerVisionContext, true, remoteImage, computervision.En)
	if err != nil {
		log.Fatal(err)
	}

	// Get orientation of text.
	fmt.Printf("Text angle: %.4f\n", *ocrResult.TextAngle)

	// Get bounding boxes for each line of text and print text.
	for _, region := range *ocrResult.Regions {
		for _, line := range *region.Lines {
			fmt.Printf("\nBounding box: %v\n", *line.BoundingBox)
			s := ""
			for _, word := range *line.Words {
				s += *word.Text + " "
			}
			fmt.Printf("Text: %v", s)
		}
	}
	fmt.Println()
	fmt.Println()
}
