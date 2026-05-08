package organizer

import "testing"

func TestCategoryForExtension(t *testing.T) {
	tests := []struct {
		name string
		ext  string
		want Category
	}{
		{name: "image", ext: ".png", want: CategoryImages},
		{name: "video", ext: ".mp4", want: CategoryVideos},
		{name: "audio", ext: ".mp3", want: CategoryAudio},
		{name: "files", ext: ".pdf", want: CategoryFiles},
		{name: "archive", ext: ".zip", want: CategoryArchives},
		{name: "executable", ext: ".exe", want: CategoryExecutable},
		{name: "code", ext: ".go", want: CategoryCode},
		{name: "unknown", ext: ".nope", want: CategoryOthers},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := categoryForExtension(tt.ext)
			if got != tt.want {
				t.Fatalf("categoryForExtension(%q) = %q, want %q", tt.ext, got, tt.want)
			}
		})
	}
}
