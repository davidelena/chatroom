package process

import "fmt"

var (
	UserManager *UserMgr
)

type UserMgr struct {
	OnlineUsers map[int]*UserProcessor
}

func init() {
	UserManager = &UserMgr{OnlineUsers: make(map[int]*UserProcessor, 1024)}
}

func (this *UserMgr) AddOrUpdateOnlineUser(processor *UserProcessor) {
	this.OnlineUsers[processor.UserId] = processor
}

func (this *UserMgr) DeleteOnlineUser(userId int) {
	delete(this.OnlineUsers, userId)
}

func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcessor {
	return this.OnlineUsers
}

func (this *UserMgr) GetUserById(userId int) (processor *UserProcessor, err error) {
	if processor, ok := this.OnlineUsers[userId]; !ok {
		err := fmt.Errorf("当前用户[%d]不在线", userId)
		return nil, err
	} else {
		return processor, nil
	}
}
