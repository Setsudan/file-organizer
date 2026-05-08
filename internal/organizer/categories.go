package organizer

import "strings"

type Category string

const (
	CategoryImages     Category = "images"
	CategoryVideos     Category = "videos"
	CategoryAudio      Category = "audio"
	CategoryFiles      Category = "files"
	CategoryArchives   Category = "archives"
	CategoryExecutable Category = "executable"
	CategoryCode       Category = "code"
	CategoryOthers     Category = "others"
)

var extensionCategoryMap = map[string]Category{
	".jpg": CategoryImages, ".jpeg": CategoryImages, ".png": CategoryImages, ".gif": CategoryImages,
	".bmp": CategoryImages, ".tiff": CategoryImages, ".svg": CategoryImages, ".webp": CategoryImages,
	".heic": CategoryImages, ".heif": CategoryImages, ".avif": CategoryImages,

	".mp4": CategoryVideos, ".mov": CategoryVideos, ".avi": CategoryVideos, ".mkv": CategoryVideos,
	".flv": CategoryVideos, ".wmv": CategoryVideos, ".webm": CategoryVideos, ".mpeg": CategoryVideos,
	".mpg": CategoryVideos, ".m4v": CategoryVideos,

	".mp3": CategoryAudio, ".flac": CategoryAudio, ".wav": CategoryAudio, ".aac": CategoryAudio,
	".ogg": CategoryAudio, ".wma": CategoryAudio, ".m4a": CategoryAudio, ".opus": CategoryAudio,

	".pdf": CategoryFiles, ".doc": CategoryFiles, ".docx": CategoryFiles, ".txt": CategoryFiles,
	".xls": CategoryFiles, ".xlsx": CategoryFiles, ".ppt": CategoryFiles, ".pptx": CategoryFiles,
	".odt": CategoryFiles, ".rtf": CategoryFiles, ".md": CategoryFiles, ".csv": CategoryFiles,
	".epub": CategoryFiles, ".mobi": CategoryFiles,

	".zip": CategoryArchives, ".7z": CategoryArchives, ".tar": CategoryArchives, ".gz": CategoryArchives,
	".rar": CategoryArchives, ".xz": CategoryArchives, ".bz2": CategoryArchives, ".zst": CategoryArchives,

	".exe": CategoryExecutable, ".msi": CategoryExecutable, ".bat": CategoryExecutable, ".cmd": CategoryExecutable,
	".dmg": CategoryExecutable, ".pkg": CategoryExecutable, ".deb": CategoryExecutable, ".rpm": CategoryExecutable,
	".appimage": CategoryExecutable,

	".go": CategoryCode, ".js": CategoryCode, ".ts": CategoryCode, ".tsx": CategoryCode, ".jsx": CategoryCode,
	".py": CategoryCode, ".java": CategoryCode, ".kt": CategoryCode, ".cs": CategoryCode, ".cpp": CategoryCode,
	".c": CategoryCode, ".h": CategoryCode, ".hpp": CategoryCode, ".rs": CategoryCode, ".php": CategoryCode,
	".rb": CategoryCode, ".swift": CategoryCode, ".sh": CategoryCode, ".sql": CategoryCode, ".json": CategoryCode,
	".yaml": CategoryCode, ".yml": CategoryCode, ".toml": CategoryCode, ".xml": CategoryCode,
}

func categoryForExtension(ext string) Category {
	cat, ok := extensionCategoryMap[strings.ToLower(ext)]
	if !ok {
		return CategoryOthers
	}
	return cat
}
