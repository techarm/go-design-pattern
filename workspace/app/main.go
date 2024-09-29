package main

import (
	"fmt"
	"streamer"
)

func main() {
	// Define number of workers and jobs
	const numJobs = 4
	const numWorkers = 4

	// Create channels for work and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool.
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println("wp:", wp)

	// Start the worker pool.
	wp.Run()

	// create a video that coverts mp4 to web ready format
	video1 := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// create second video that should fail.
	video2 := wp.NewVideo(2, "./input/bad.txt", "./output", "mp4", notifyChan, nil)

	// create a third video that coverts mp4 to hls
	ops := &streamer.VideoOptions{
		RenameOutput:    true,
		SegmentDuration: 10,
		MaxRate1080p:    "1200k",
		MaxRate720p:     "600k",
		MaxRate480p:     "400k",
	}
	video3 := wp.NewVideo(3, "./input/puppy2.mp4", "./output", "hls", notifyChan, ops)

	// create a fourth video that converts mp4 to web ready format
	video4 := wp.NewVideo(4, "./input/puppy2.mp4", "./output", "mp4", notifyChan, ops)

	// Send the videos to the work pool.
	videoQueue <- streamer.VideoProcessingJob{Video: video1}
	videoQueue <- streamer.VideoProcessingJob{Video: video2}
	videoQueue <- streamer.VideoProcessingJob{Video: video3}
	videoQueue <- streamer.VideoProcessingJob{Video: video4}

	// Print out result.
	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i: ", i, "msg:", msg)
	}

	fmt.Println("Done")
}
