package taskrunner

import (
	"errors"
	"log"
	"sync"
	"os"
	"go_dev/video_server/scheduler/dbops"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)

	if err != nil && !os.IsNotExist(err) {
		log.Printf("deleting video error :%v", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)

	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {

	errMap := &sync.Map{}  //线程安全的

	forloop:
		for {
			select {
				case vid :=<- dc:
					//这里使用匿名函数，得主动传入参数，如果在函数中直接使用vid，会导致获取的值不一致，goroute获取的是瞬时的值
					go func(id interface{}) {
						if err := deleteVideo(id.(string)); err != nil {
							errMap.Store(id, err)
							return
						}
						if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
							errMap.Store(id, err)
							return
						}
					}(vid)  //这需要传进去，而不是直接在函数取vid值
				default:
					break forloop
			}
		}

	errMap.Range(func(k, v interface{}) bool {
		err := v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return nil
}