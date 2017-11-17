package transliterate

import (
	"bufio"
	"encoding/json"
	"fmt"
	// Imports the Google Cloud Speech API client package.
	"golang.org/x/net/context"
	// "cloud.google.com/go/translate"
	// "golang.org/x/text/language"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"os"
	"mime/multipart"
	"net/http"
	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"strings"
	"time"
)

func HandleSilakkiDumma(w http.ResponseWriter, r *http.Request) {
	c := cache.New(5*time.Hour, 10*time.Hour)

	file, err := os.Open("data/only_tamil_uniq_sorted_words.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		c.Set(strings.TrimSpace(scanner.Text()), true, cache.NoExpiration)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println(len(c.Items()))


	jsonOutput := make(map[string]string)
	ctx := context.Background()
	fmt.Println(r)


	audiofile, _ := os.Open("data/silakki.flac")
	data, err := ioutil.ReadAll(audiofile)

	if err != nil {
		log.Fatal(err)
	}

	// r.ParseMultipartForm(32 << 20)
	// var data []byte
	// for _, fheaders := range r.MultipartForm.File {
	// 	for _, hdr := range fheaders {
	// 		var infile multipart.File
	// 		infile, _ = hdr.Open()
	// 		data = make([]byte, hdr.Size)
	// 		infile.Read(data)
	// 	}
	// }

	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Println("Failed to create client: %v", err)

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

	// Prints the results.strings.Join(s, ", ")
	for _, result := range resp.Results {
		fmt.Println(result);
		for _, alt := range result.Alternatives {
			transcript := alt.Transcript
			jsonOutput["transcript"] = transcript
			// TODO: ADD if tamil clause
			if (alt.Confidence < 0.70) {
				stringList := strings.Fields(transcript)
				var unMeaningful []string
				for _, word := range stringList {
					_, found := c.Get(word)
					if !found {
						unMeaningful = append(unMeaningful, word)
					}
				}
				jsonOutput["learn"] = strings.Join(unMeaningful, " ")
			}
 		}
	}
	// fmt.Fprintf(w, jsonOutput["Transcript"])
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(jsonOutput)
}

func HandleTransliterateQuery(w http.ResponseWriter, r *http.Request) {

	c := cache.New(5*time.Hour, 10*time.Hour)

	file, err := os.Open("data/only_tamil_uniq_sorted_words.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		c.Set(strings.TrimSpace(scanner.Text()), true, cache.NoExpiration)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println(len(c.Items()))


	jsonOutput := make(map[string]string)
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
		log.Println("Failed to create client: %v", err)

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

	// Prints the results.strings.Join(s, ", ")
	for _, result := range resp.Results {
		fmt.Println(result);
		for _, alt := range result.Alternatives {
			transcript := alt.Transcript
			jsonOutput["Transcript"] = transcript
			// TODO: ADD if tamil clause
			if (alt.Confidence < 0.70) {
				stringList := strings.Fields(transcript)
				var unMeaningful []string
				for _, word := range stringList {
					_, found := c.Get(word)
					if !found {
						unMeaningful = append(unMeaningful, word)
					}
				}
				jsonOutput["Learn"] = strings.Join(unMeaningful, " ")
			}
 		}
	}
	// fmt.Fprintf(w, jsonOutput["Transcript"])
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(jsonOutput)
}

func HandleQuestions(w http.ResponseWriter, r *http.Request) {

	jsonOutput := make(map[string]string)
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
		log.Println("Failed to create client: %v", err)

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

	// Prints the results.strings.Join(s, ", ")
	var answer string
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			transcript := alt.Transcript
			// lang, err := language.Parse("en-US")
			// if err != nil {
			// 	fmt.Println("No translation parser for target lang")
			// }
			//
			// client, err := translate.NewClient(ctx)
			// if err != nil {
			// 	fmt.Println("No translation client ")
			// }
			// defer client.Close()
			//
			// resp, err := client.Translate(ctx, []string{transcript}, lang, nil)
			// if err != nil {
			// 	fmt.Println("No translation yet")
			// }

			fmt.Println(transcript)
			if strings.Contains(transcript, "மணி") {
				answer = time.Now().Format("20060102150405")
			} else if strings.Contains(transcript, "வானிலை") || strings.Contains(transcript, "climate") {
				answer = "12'C"
			} else {
				answer = "I am not yet, good a personal assistant :( Go ask someone else"
			}
 		}
	}
	// fmt.Fprintf(w, jsonOutput["Transcript"])
	jsonOutput["answer"] = answer
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(jsonOutput)
}
