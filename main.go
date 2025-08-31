package main

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

const (
	MaxUploadSize = 100 * 1024 * 1024 // 100 MB
)

func getUploadBucketName() string {
    bucketName := os.Getenv("UPLOAD_BUCKET_NAME")
    if bucketName == "" {
        bucketName = "video-upload-zoaewe3s" // fallback default
    }
    return bucketName
}

func uploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		methodNotAllowedHandler(w,r)
		return
	}

	r.ParseMultipartForm(MaxUploadSize)

	file, handler, err := r.FormFile("video")
	if err!=nil{
		logJSON("error", "Failed to get file from form", map[string]interface{}{
            "error": err.Error(),
        })
		http.Error(w, "Failed to get file", http.StatusBadRequest)
        return
	}
	defer file.Close()

	if handler.Size > MaxUploadSize{
		logJSON("error", "File size exceeds the limits", map[string]interface{}{
			"path": r.URL.Path,
			"method": r.Method,
		})
		return 
	}

	if !isValidMP4(handler.Filename){
		logJSON("error", "Invalid file format", map[string]interface{}{
			"path": r.URL.Path,
			"method": r.Method,
		})
		http.Error(w, "Only MP4 files are allowed", http.StatusBadRequest)
		return
	}

	videoID := uuid.New().String()
	s3Key := videoID + ".mp4"

	// Upload to S3
    if err := uploadToS3(file, s3Key, handler.Filename, handler.Size); err != nil {
        logJSON("error", "S3 upload failed", map[string]interface{}{
            "error": err.Error(),
            "video_id": videoID,
        })
        http.Error(w, "Upload failed", http.StatusInternalServerError)
        return
    }

	    response := map[string]interface{}{
        "video_id": videoID,
        "original_filename": handler.Filename,
        "file_size": handler.Size,
        "status": "uploaded",
        "upload_time": time.Now().Format(time.RFC3339),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)

    logJSON("info", "Video uploaded successfully", map[string]interface{}{
        "video_id": videoID,
        "filename": handler.Filename,
        "size": handler.Size,
    })
}

func isValidMP4(filename string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    return ext == ".mp4"
}

func uploadToS3(file multipart.File, s3Key, originalFilename string, fileSize int64) error {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
    if err != nil {
        return err
    }

    svc := s3.New(sess)

    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(getUploadBucketName()),
        Key:    aws.String(s3Key),
        Body:   file,
        Metadata: map[string]*string{
            "original-filename": aws.String(originalFilename),
            "upload-time":      aws.String(time.Now().Format(time.RFC3339)),
            "file-size":        aws.String(strconv.FormatInt(fileSize, 10)),
        },
        ContentType: aws.String("video/mp4"),
    })

    return err
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logJSON("info","hello endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
		"name_param": r.URL.Query().Get("name"),
	})
	

	w.Header().Set("Content-Type", "text/plain")
	name := r.URL.Query().Get("name")
	if name == ""{
		name = "stranger"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	logJSON("info","about endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
	})

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "This is the about page.")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request){
	logJSON("info","not found endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
	})

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w,"Page not found!")
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request){
	logJSON("info","method not allowed endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
	})
	
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w,"Method not allowed!")
}

func logJSON(level string,message string,data  map[string]interface{}){
	logEntry := map[string]interface{}{
		"time":time.Now().Format(time.RFC3339),
		"message":message,
		"level":level,
	}
	for k,v := range data{
		logEntry[k] = v
	}
	jsonByte, _ := json.Marshal(logEntry)
	fmt.Println(string(jsonByte))
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":"healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
    logJSON("info", "version endpoint accessed", map[string]interface{}{
        "path": r.URL.Path,
        "method": r.Method,
    })
    
    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
		"version": "1.0.0",
        "status": "running",
        "timestamp": time.Now().Format(time.RFC3339),
    }
    json.NewEncoder(w).Encode(response)
}

func testAWSHandler(w http.ResponseWriter, r *http.Request){
	logJSON("info", "version endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
        "method": r.Method,
    })
	w.Header().Set("Content-Type", "application/json")

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err!=nil{
		http.Error(w, "Error connecting to AWS console", http.StatusInternalServerError)
		return
	}

	svc := s3.New(sess)
	result, err := svc.ListBuckets(nil)
	if err != nil {
        fmt.Println("Unable to list buckets:", err)
        return
    }

	var buckets []map[string]interface{}
    fmt.Println("Buckets:")
	for _, b := range result.Buckets {
    	bucket := map[string]interface{}{
			"name": *b.Name,
			"creation_date": *b.CreationDate,
		}
		buckets = append(buckets, bucket)
    }
	err = json.NewEncoder(w).Encode(buckets)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func main() {

	if err := godotenv.Load(); err != nil {
        logJSON("info", "No .env file found, using environment variables", map[string]interface{}{})
    }

	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}
	c := chi.NewRouter()
	c.Get("/hello", helloHandler)
	c.Get("/about", aboutHandler)
	c.Get("/health", healthHandler)
	c.Get("/version", versionHandler)
	c.Get("/test-aws", testAWSHandler)
	c.Post("/upload", uploadHandler)

	c.NotFound(notFoundHandler)
	c.MethodNotAllowed(methodNotAllowedHandler)


	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, c)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		return
	}
}
