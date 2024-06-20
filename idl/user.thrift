include "video.thrift"

struct User {
    1: i64 Id (thrift.sql="id", thrift.auto_increment="true", thrift.primary_key="true")
    2: string Username (thrift.sql="username", thrift.unique="true", thrift.not_null="true")
    3: i32 Age (thrift.sql="age", thrift.not_null="true")
    4: string City (thrift.sql="city")
    5: bool Banned (thrift.sql="banned", thrift.default="false")
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