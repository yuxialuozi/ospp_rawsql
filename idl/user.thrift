include "video.thrift"

struct User {
    1: i64 Id (go.tag="sql:\"id,omitempty\"")
    2: string Username (go.tag="sql:\"username\"")
    3: i32 Age (go.tag="sql:\"age\"")
    4: string City (go.tag="sql:\"city\"")
    5: bool Banned (go.tag="sql:\"banned\"")
}

(
mysql.InsertOne = "InsertOne(ctx context.Context, user *user.User) (interface{}, error)"
mysql.InsertMany = "InsertMany(ctx context.Context, user []*user.User) ([]interface{}, error)"
mysql.FindUsernameOrderbyIdSkipLimitAll = "FindUsernames(ctx context.Context, skip, limit int64) ([]*user.User, error)"
mysql.FindByLbLbUsernameEqualOrUsernameEqualRbAndAgeGreaterThanRb = "FindByUsernameAge(ctx context.Context, name1, name2 string, age int32) (*user.User, error)"
mysql.UpdateContactByIdEqual = "UpdateContact(ctx context.Context, contact *user.UserContact, id int64) (bool, error)"
mysql.DeleteByYdEqual = "DeleteById(ctx context.Context, yd []user.YDType) (int, error)"
mysql.CountByAgeBetween = "CountByAge(ctx context.Context, age1, age2 int32) (int, error)"
)