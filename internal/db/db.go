package db

import (
	"GitfyBot/internal"
	"GitfyBot/internal/logger"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     internal.Config.Database.URL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Fatal("Fail to connect redis: " + err.Error())
	} else {
		logger.Info("Connected successfully to redis")
	}
}

func dropErr(e error) bool {
	if e != nil {
		logger.Error(e.Error())
		return true
	}
	return false
}

func GetUser(userId string) *User {
	result, err := rdb.HGet(ctx, "users", userId).Result()
	if dropErr(err) {
		logger.Info("Fail to get user")
		return nil
	}
	var user *User
	err = json.Unmarshal([]byte(result), &user)
	if dropErr(err) {
		logger.Info("Fail to parse user json")
		return nil
	}
	return user
}

func updateUser(user *User) bool {
	_, err := rdb.HSet(ctx, "users", user.Id, user.MarshalBinary()).Result()
	if dropErr(err) {
		logger.Info("Fail to update user")
		return false
	}
	return true
}

func getRepo(owner, name string) *Repo {
	result, err := rdb.HGet(ctx, "repos", owner+":"+name).Result()
	if dropErr(err) {
		logger.Info("Fail to get repo")
		return nil
	}
	var repo *Repo
	err = json.Unmarshal([]byte(result), &repo)
	if dropErr(err) {
		logger.Info("Fail to parse repo json")
		return nil
	}
	return repo
}

func updateRepo(repo *Repo) bool {
	_, err := rdb.HSet(ctx, "repos", repo.Owner+":"+repo.Name, repo.MarshalBinary()).Result()
	if dropErr(err) {
		logger.Info("Fail to update repo")
		return false
	}
	return true
}
