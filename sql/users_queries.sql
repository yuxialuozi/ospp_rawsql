
-- name: insert:exec
INSERT INTO Users (Id, Age, Banned, DefaultAge)
VALUES (?, ?, ?, ?);



-- name: DeleteByID:exec
DELETE FROM Users WHERE Id = ?;



-- name: UpdateUsersById:exec
UPDATE Users SET  Id = ?, Age = ?, Banned = ?, DefaultAge = ?
WHERE Id = ?;



-- name: GetUsersById:many
SELECT * FROM Users
WHERE Id = ?;



-- name: GetUsersByAge:many
SELECT * FROM Users
WHERE Age = ?;



-- name: GetUsersByBanned:many
SELECT * FROM Users
WHERE Banned = ?;



-- name: GetUsersByDefaultAge:many
SELECT * FROM Users
WHERE DefaultAge = ?;



-- name: DeleteUsersById:exec
DELETE FROM Users
WHERE Id = ?;



-- name: DeleteUsersByAge:exec
DELETE FROM Users
WHERE Age = ?;



-- name: DeleteUsersByBanned:exec
DELETE FROM Users
WHERE Banned = ?;


-- name: DeleteUsersByDefaultAge:exec
DELETE FROM Users
WHERE DefaultAge = ?;

-- name: InsertUser:exec
INSERT INTO Users (Id, Age, Banned, DefaultAge)
VALUES (?, ?, ?, ? ;  -- 缺少右括号

-- name: InsertUser:exec
INSERT INTO Users (Id, Age, Banned, DefaultAge)
VALUES ?, ?, ?, ? ;  -- 缺少右括号
