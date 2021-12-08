package db

import "encoding/json"

type User struct {
	Id            string `json:"id"`
	Subscriptions []Repo `json:"subscriptions"`
	IsAdmin       bool   `json:"is_admin"`
}

type Release struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	IsPrerelease bool   `json:"is_prerelease"`
}

type Tag struct {
	Version string `json:"version"`
	Release `json:"release"`
}

type Repo struct {
	Owner        string          `json:"owner"`
	Name         string          `json:"name"`
	Tags         []Tag           `json:"tags"`
	Releases     []Release       `json:"releases"`
	WatchedUsers map[string]bool `json:"watched_users"`
}

func (u *User) MarshalBinary() string {
	data, _ := json.Marshal(u)
	return string(data)
}

func (r *Repo) MarshalBinary() string {
	data, _ := json.Marshal(r)
	return string(data)
}
