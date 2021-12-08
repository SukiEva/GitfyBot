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

func AddRepo(from, owner, name string) {
	repo := Repo{
		From:         from,
		Owner:        owner,
		Name:         name,
		Tags:         []Tag{},
		WatchedUsers: map[string]bool{},
	}
	_, err := rdb.HSet(ctx, "repos", from+":"+owner+":"+name, repo.MarshalBinary()).Result()
	if dropErr(err) {
		return
	}
	logger.Info(from + " Repo " + owner + "/" + name + " added")
}

func RemoveRepo(from, owner, name string) {
	_, err := rdb.HDel(ctx, "repos", from+":"+owner+":"+name).Result()
	if dropErr(err) {
		return
	}
	logger.Info(from + " Repo " + owner + "/" + name + " removed")
}

func BindUserToRepo(userid, from, owner, name string) {
	repo := getRepo(from, owner, name)
	if repo == nil {
		return
	}
	repo.WatchedUsers[userid] = true
	if updateRepo(repo) {
		logger.Info("User " + userid + " bind to " + from + ":" + owner + "/" + name)
	}
}

func UnBindUserFromRepo(userid, from, owner, name string) {
	repo := getRepo(from, owner, name)
	if repo == nil {
		return
	}
	delete(repo.WatchedUsers, userid)
	if updateRepo(repo) {
		logger.Info("User " + userid + " unbind from " + from + ":" + owner + "/" + name)
	}
}

func getUserSubscriptions(userId string) []Repo {
	user := GetUser(userId)
	if user != nil {
		return user.Subscriptions
	}
	return nil
}
