# API Reference

Complete documentation for the Video Processing Platform API endpoints.

## Base URL
```
http://<your-ec2-ip>:3000
```

## Authentication
Currently no authentication required (development phase).

## Endpoints

### Health Check
Check if the service is running.

**GET** `/health`

**Response**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z"
}
```

### Upload Video
Upload a video file for processing.

**POST** `/upload`

**Request**
- Content-Type: `multipart/form-data`
- Body: Video file in `video` field

**Response**
```json
{
  "video_id": "uuid-here",
  "status": "uploaded",
  "original_filename": "my-video.mp4",
  "upload_time": "2024-01-15T10:30:00Z"
}
```

### Get Video Status
Check processing status of uploaded video.

**GET** `/video/{video_id}/status`

**Response**
```json
{
  "video_id": "uuid-here",
  "status": "processing", // uploaded, processing, completed, failed
  "resolutions_available": ["320p", "480p"],
  "processing_progress": 75
}
```

### Download Video
Download processed video in specified resolution.

**GET** `/video/{video_id}/download?resolution=320p`

**Query Parameters**
- `resolution`: `original`, `320p`, `480p`

**Response**
- Content-Type: `video/mp4`
- Video file stream

### List Videos
Get list of all uploaded videos.

**GET** `/videos`

**Response**
```json
{
  "videos": [
    {
      "video_id": "uuid-1",
      "filename": "video1.mp4",
      "status": "completed",
      "upload_time": "2024-01-15T10:30:00Z",
      "resolutions": ["320p", "480p"]
    }
  ]
}
```

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid file format",
  "code": "INVALID_FORMAT"
}
```

### 404 Not Found
```json
{
  "error": "Video not found",
  "code": "VIDEO_NOT_FOUND"
}
```

### 500 Internal Server Error
```json
{
  "error": "Processing failed",
  "code": "PROCESSING_ERROR"
}
```

## Usage Examples

### Upload a video
```bash
curl -X POST \
  -F "video=@my-video.mp4" \
  http://your-server:3000/upload
```

### Check processing status
```bash
curl http://your-server:3000/video/uuid-here/status
```

### Download 320p version
```bash
curl -o video-320p.mp4 \
  "http://your-server:3000/video/uuid-here/download?resolution=320p"
```

## Rate Limits
- Upload: 10 videos per hour per IP
- Download: 100 requests per hour per IP

## File Constraints
- Maximum file size: 100MB
- Supported formats: MP4, AVI, MOV
- Maximum duration: 10 minutes

---

**Next**: [Architecture Guide](architecture.md)