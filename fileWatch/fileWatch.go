package fileWatch

import (
	"time"
)

type FileWatcher struct {
	FilePath string
	OriginContent string
	LastUpdateTime time.Time
}


func NewWatcher(name string) *FileWatcher{
	return &FileWatcher{
		name,nil,nil,
	}
}
//
//func (f *FileWatcher) Run() {
//	for {
//		select {
//		case event :=<- watcher.Events:
//			// fsnotify sometimes sends a bunch of events without name or operation.
//			// It's unclear what they are and why they are sent - filter them out.
//			if len(event.Name) == 0 {
//				break
//			}
//			// Everything but a chmod requires rereading.
//			if event.Op^fsnotify.Chmod == 0 {
//				break
//			}
//			watcher.Add(file/directory)
//			//logic
//		case err :=<-watcher.Errors:
//		}
//	}
//
//}


