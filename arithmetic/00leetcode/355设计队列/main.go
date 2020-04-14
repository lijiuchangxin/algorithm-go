package main

type Twitter struct {
	Tweets 	[]int				// 记录所有的tweets， 数组的顺序则为推特创建的顺序
	UserTweets 	map[int][]int	// user与其发布的tweet映射关系
	Follows 	map[int][]int	// 关注列表映射关系
	IsFollowMy	map[int]bool	// 是否關注自己映射關係
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	// 初始化twitter
	var Tweets []int
	var UserTweets = make(map[int][]int)
	var Follows = make(map[int][]int)
	var IsFollows = make(map[int]bool)

	t :=  Twitter{
		Tweets:     Tweets,
		UserTweets: UserTweets,
		Follows:    Follows,
		IsFollowMy:  IsFollows,
	}

	return t
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	// 记录每一条的推特
	this.Tweets = append(this.Tweets, tweetId)
	// 将推特记录在用户名下
	this.UserTweets[userId] = append(this.UserTweets[userId], tweetId)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	// 获取该名用户所有的关注
	follows := this.Follows[userId]

	// 记录所有与该用户相关的推特（自己发表的推特，关注发表的推特）
	var allTweets []int
	// 获取关注用户的所有的推特并加入allTweets
	for _, id := range follows {
		allTweets = append(allTweets, this.UserTweets[id]...)
	}

	// 判断是否关注自己，如果没有关注则需要将自己所有的推特加入allTweets
	if !this.IsFollowMy[userId] {
		allTweets = append(allTweets, this.UserTweets[userId]...)
	}

	// 将推特排序存储
	var sortTweets []int
	aTLen := len(this.Tweets)
	s := 0
	// 按照发的推特进行倒叙
	for i := aTLen-1; i >= 0; i-- {
		if s >= 10 {
			break
		}
		for _, n := range allTweets {
			if this.Tweets[i] == n && s < 10 {
				s++
				sortTweets = append(sortTweets, n)
			}
		}
	}
	return sortTweets
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followeeId == followerId {
		this.IsFollowMy[followerId] = true
	}

	var isFol bool
	for _, follId := range this.Follows[followerId] {
		if follId == followeeId {
			isFol = true
		}
	}

	if !isFol {
		this.Follows[followerId] = append(this.Follows[followerId], followeeId)
	}

}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followeeId == followerId {
		this.IsFollowMy[followerId] = false
	}

	var temp []int
	for _, v := range this.Follows[followerId] {
		if v != followeeId {
			temp = append(temp, v)
		}
	}

	this.Follows[followerId] = temp
}


/**
* Your Twitter object will be instantiated and called as such:
* obj := Constructor();
* obj.PostTweet(userId,tweetId);
* param_2 := obj.GetNewsFeed(userId);
* obj.Follow(followerId,followeeId);
* obj.Unfollow(followerId,followeeId);
*/