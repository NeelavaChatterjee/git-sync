package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/NeelavaChatterjee/git-sync/models"
	"github.com/NeelavaChatterjee/git-sync/utilities"
	"github.com/robfig/cron/v3"
)

//TODO
// use cron
// do the actual polling (github)
// schedule polling (cron)
// check all the scheduled polls and their times (cron)
// check all the scheduled polls and their times just for a repo and branch (can be combined with above) (cron)
// stop a scheduled poll (cron)
// actual polling
//   - create directory
//   - fetch commit details
//   - fetch all files

func createNewDirectory(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Poll(track *models.Track) {
	// Create the directory
	poll_time := time.Now().String()
	path := filepath.Join("repos", track.Owner, track.Repository, track.Branch, poll_time)
	createNewDirectory(path)

	// TODO
	// Start by fetching the commits
	// last_poll_log, err := FetchLastPollLog(track.ID)
	// if err != nil {
	// 	fmt.Println("Couldn't fetch last poll log from database.")
	// 	return
	// }
	since := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	// since := last_poll_log.CreatedAt
	until := time.Now()
	commits := utilities.GetCommits(track.Owner, track.Repository, track.Branch, since, until)

	for i, commit := range commits {
		fmt.Println(i)
		files := utilities.GetCommitFiles(track.Owner, track.Repository, commit.GetSHA())
		for i, file := range files {
			fmt.Println(i)
			file_remote_dir := file.GetFilename()
			file_content := utilities.GetFileContents(track.Owner, track.Repository, track.Branch, file_remote_dir)
			file_local_path := filepath.Join(path, file_remote_dir)
			file_local_dir := filepath.Dir(file_local_path)
			createNewDirectory(file_local_dir)

			f, err := os.Create(file_local_path)
			if err != nil {
				panic(err)
			}

			// incase its a 404 that is file was deleted we just skip the file
			if file_content != nil {
				f_content, err := file_content.GetContent()
				if err != nil {
					panic(err)
				}
				f.WriteString(f_content)
			}
		}
	}
}

// TODO store the id
func SchedulePoll(track *models.Track, poll_interval string) {
	utilities.Cron.AddFunc("@every "+poll_interval, func() {
		Poll(track)
	})
}

func RemoveScheduledPoll(id cron.EntryID) {
	utilities.Cron.Remove(id)
}

func GetAllScheduledEntries() []cron.Entry {
	return utilities.Cron.Entries()
}
