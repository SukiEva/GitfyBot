package db

import (
	"GitfyBot/internal"
	"GitfyBot/internal/logger"
)

func AddUser(userId string) {
	user := User{
		Id:            userId,
		Subscriptions: []Repo{},
		IsAdmin:       userId == internal.Config.Admin,
	}
	_, err := rdb.HSet(ctx, "users", userId, user.MarshalBinary()).Result()
	if dropErr(err) {
		return
	}
	logger.Info("User " + userId + " added")
}

func AddRepo(owner, name string) {
	repo := Repo{
		Owner:        owner,
		Name:         name,
		Tags:         []Tag{},
		Releases:     []Release{},
		WatchedUsers: map[string]bool{},
	}
	_, err := rdb.HSet(ctx, "repos", owner+":"+name, repo.MarshalBinary()).Result()
	if dropErr(err) {
		return
	}
	logger.Info("Repo " + owner + "/" + name + " added")
}

func BindUserToRepo(userid, owner, name string) {
	repo := getRepo(owner, name)
	if repo == nil {
		return
	}
	repo.WatchedUsers[userid] = true
	if updateRepo(repo) {
		logger.Info("User " + userid + " bind to " + owner + "/" + name)
	}
}

func UnBindUserFromRepo(userid, owner, name string) {
	repo := getRepo(owner, name)
	if repo == nil {
		return
	}
	delete(repo.WatchedUsers, userid)
	if updateRepo(repo) {
		logger.Info("User " + userid + " unbind from " + owner + "/" + name)
	}
}
