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

	last_poll_log, err := FetchLastPollLog(track.ID)
	if err != nil {
		fmt.Println("Couldn't fetch last poll log from database.")
	}
	var since time.Time
	if last_poll_log != nil {
		since = last_poll_log.CreatedAt
	} else {
		// This is the fall back option if this is the first poll
		since = time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	}

	until := time.Now()
	// Fetching Commits
	commits := utilities.GetCommits(track.Owner, track.Repository, track.Branch, since, until)

	number_of_commits := 0
	number_of_files := 0
	for i, commit := range commits {
		number_of_commits++
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
				number_of_files++
				f_content, err := file_content.GetContent()
				if err != nil {
					panic(err)
				}
				f.WriteString(f_content)
			}
		}

		commit_history_log := models.CommitHistory{
			SHA:         commit.GetSHA(),
			URL:         commit.GetURL(),
			Date:        commit.GetAuthor().CreatedAt.Time,
			Message:     *commit.Commit.Message,
			AuthorName:  *commit.GetAuthor().Login,
			AuthorEmail: *commit.GetAuthor().Email,
			TrackID:     track.ID,
		}
		err := CreateNewCommitHistoryEntry(&commit_history_log)
		if err != nil {
			panic(err)
		}
	}
	poll_log := models.PollLogs{
		NumberOfFiles:   number_of_files,
		NumberOfCommits: number_of_commits,
		TrackID:         track.ID,
	}

	if err := CreatePollLog(&poll_log); err != nil {
		panic(err)
	}
}

func SchedulePoll(track *models.Track) cron.EntryID {
	entry_id, err := utilities.Cron.AddFunc("@every "+track.PollInterval, func() {
		Poll(track)
	})
	if err != nil {
		panic(err)
	}
	return entry_id
}

func GetAllScheduledEntries() []cron.Entry {
	return utilities.Cron.Entries()
}
