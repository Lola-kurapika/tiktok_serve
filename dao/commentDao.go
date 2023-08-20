package dao

import(
	"time"
	"errors"
	"log"
)

//定义 Comment 模型
type Comment struct{
	Id			int64 		//评论记录 id
	UserId 		int64		//发布评论的用户 id
	VideoId 	int64		//评论视频的 id
	Content		string		//评论的内容
	ActionType	int64		//评论行为，1 表示已发布评论，2 表示删除评论
	CreatedTime	time.Time	//评论发布时间
	UpdatedTime time.Time	//评论更新时间
}

//为 Comment 定义表名：
func (c Comment) TableName() string {
	return "comment"
}

//创建评论
func CreateComment(comment Comment)(Comment,error){
	if err:=Db.Model(Comment{}).Create(&comment).Error;err!=nil{
		log.Println(err.Error())
		return Comment{},err
	}
	return comment,nil
}

//更新评论
func UpdateComment(comment_Id int64,update_content Content)(Comment,error){
	var comment Comment 
	//更新前先查询该评论是否存在
	result:=Db.Where("id=?",comment_Id).First(&comment) //获得 id 为 CommentId 的第一条记录（主键升序）
	if result.Error!=nil{
		return errors.New("upd comment does not exist")
	}
	//该评论存在，则更新评论，将评论内容 content 置为 update_content
	result=Db.Model(Comment{}).Where("id=?",comment_Id).Update("content",update_content)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}

//删除评论
func DeleteComment(comment_Id int64) error {
	var comment Comment 
	//删除前先查询该评论是否存在
	result:=Db.Where("id=?",comment_Id).First(&comment) //获得 id 为 CommentId 的第一条记录（主键升序）
	if result.Error!=nil{
		return errors.New("del comment does not exist")
	}
	//该评论存在，则删除评论，将评论行为 action_type 置为2
	//此删除为[软删除]，并不会真正删除数据库中的数据
	result=Db.Model(Comment{}).Where("id=?",comment_Id).Update("action_type",2)
	if result.Error!=nil{
		return result.Error
	}
	return nil
}

//查询评论
func GetCommentList(video_Id int64)([]Comment,error){
	var comment_List []Comment
	result:=Db.Model(Comment{}).Where(map[string]interface{}{"video_id":video_Id,"action_type":1}).Order("created_time desc").Find(&comment_List)
	if result.Error!=nil{
		log.Println(result.Error)
		return comment_List,errors.New("fail to get comment list")
	}
	return comment_List,nil
}

//获取评论数量
func GetCommentCount(video_Id int64)(int64,error){
	var count int64
	result:=Db.Model(Comment{}).Where(map[string]interface{}{"video_id":video_Id,"action_type":1}).Count(&count)
	if result.Error!=nil{
		return 0,errors.New("fail to find comments count")
	}
	return count,nil
}