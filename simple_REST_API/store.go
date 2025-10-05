package main

import "sync"

var (
	users = map[int]User{}
	posts = map[int]Post{}

	userMu = sync.RWMutex{}
	postMu = sync.RWMutex{}

	nextUserID = 1
	nextPostID = 1
)
