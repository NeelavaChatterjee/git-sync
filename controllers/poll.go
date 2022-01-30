package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/NeelavaChatterjee/git-sync/models"
	"github.com/NeelavaChatterjee/git-sync/utilities"
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

func createNewDirectory(track *models.Track) {
	path := filepath.Join(track.Owner, track.Repository, track.Branch)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Poll(track *models.Track) {
	// Start by fetching the commits
	last_poll_log, err := FetchLastPollLog(track.ID)
	if err != nil {
		fmt.Println("Couldn't fetch last poll log from database.")
		return
	}
	// since := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	since := last_poll_log.CreatedAt
	until := time.Now()
	commits := utilities.GetCommits(track.Owner, track.Repository, track.Branch, since, until)
	// Fetch the filenames and urls of the commits
	filemap := make(map[string]string)
	for i, commit := range commits {
		fmt.Println(
			i,
			commit.GetSHA(),
			/* commit.Commit.GetMessage(), */
			commit.Commit.GetAuthor().GetName(),
			commit.GetAuthor().GetLogin(),
		)
		files := utilities.GetCommitFiles(track.Owner, track.Repository, commit.GetSHA())
		for i, file := range files {
			filemap[file.GetFilename()] = file.GetRawURL()
			fmt.Println(
				i,
				file.GetFilename(),
				file.GetBlobURL(),
				file.GetContentsURL(),
				file.GetRawURL(),
			)
		}

	}
	// Create the directory
	createNewDirectory(track)
	// Fetch all the files in the directory

}
