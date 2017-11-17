package transliterate

import (
	"fmt"
	"log"
	"net/http"
	"mime/multipart"
	// Imports the Google Cloud Speech API client package.
	"golang.org/x/net/context"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func HandleTransliterateQuery(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	fmt.Println(r)
	r.ParseMultipartForm(32 << 20)
	var data []byte
	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			var infile multipart.File
			infile, _ = hdr.Open()
			data = make([]byte, hdr.Size)
			infile.Read(data)
		}
	}

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println("Created Client")

	// Detects speech in the audio file.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_FLAC,
			LanguageCode:    "ta-IN",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})
	if err != nil {
		log.Println("failed to recognize: %v", err)
	}
	fmt.Println("Client recognized speech")

	// Prints the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}

}
