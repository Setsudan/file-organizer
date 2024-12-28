package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Category mapping
var fileCategories = map[string]string{
	// Documents
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".txt":  "Documents",
	".xls":  "Documents",
	".xlsx": "Documents",
	".ppt":  "Documents",
	".pptx": "Documents",
	".odt":  "Documents",
	".rtf":  "Documents",
	// Videos
	".mp4":  "Videos",
	".mov":  "Videos",
	".avi":  "Videos",
	".mkv":  "Videos",
	".flv":  "Videos",
	".wmv":  "Videos",
	".webm": "Videos",
	// Audio
	".mp3":  "Audio",
	".flac": "Audio",
	".wav":  "Audio",
	".aac":  "Audio",
	".ogg":  "Audio",
	".wma":  "Audio",
	// Pictures
	".jpg":  "Pictures",
	".jpeg": "Pictures",
	".png":  "Pictures",
	".gif":  "Pictures",
	".bmp":  "Pictures",
	".tiff": "Pictures",
	".svg":  "Pictures",
}

type FileOrganizer struct {
	downloadDir string
	docDir      string
	videoDir    string
	audioDir    string
	picDir      string
	otherDir    string
}

// NewFileOrganizer sets up the main directories
func NewFileOrganizer(downDir, doc, vid, aud, pic, oth string) *FileOrganizer {
	return &FileOrganizer{
		downloadDir: downDir,
		docDir:      doc,
		videoDir:    vid,
		audioDir:    aud,
		picDir:      pic,
		otherDir:    oth,
	}
}

// organizeFiles demonstrates a mix of functional & OOP style
func (fo *FileOrganizer) organizeFiles() error {
	files, err := os.ReadDir(fo.downloadDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))
		targetParent := determineCategory(ext, fo)
		targetFolder := filepath.Join(targetParent, "Downloaded "+filepath.Base(targetParent))
		if err := os.MkdirAll(targetFolder, os.ModePerm); err != nil {
			return err
		}
		oldPath := filepath.Join(fo.downloadDir, file.Name())
		newPath := filepath.Join(targetFolder, file.Name())
		if err := os.Rename(oldPath, newPath); err != nil {
			return fmt.Errorf("failed to move file %s: %w", file.Name(), err)
		}
	}

	return nil
}

// determineCategory is a functional helper
func determineCategory(ext string, fo *FileOrganizer) string {
	switch fileCategories[ext] {
	case "Documents":
		return fo.docDir
	case "Videos":
		return fo.videoDir
	case "Audio":
		return fo.audioDir
	case "Pictures":
		return fo.picDir
	default:
		return fo.otherDir
	}
}

func main() {
	// Adjust paths as needed
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}

	downloadDir := filepath.Join(homeDir, "Downloads")
	docDir := filepath.Join(homeDir, "Documents")
	videoDir := filepath.Join(homeDir, "Videos")
	audioDir := filepath.Join(homeDir, "Music")
	picDir := filepath.Join(homeDir, "Pictures")
	otherDir := filepath.Join(homeDir, "Others")

	organizer := NewFileOrganizer(downloadDir, docDir, videoDir, audioDir, picDir, otherDir)
	if err := organizer.organizeFiles(); err != nil {
		fmt.Println("Error organizing files:", err)
	} else {
		fmt.Println("Files moved successfully!")
	}
}
