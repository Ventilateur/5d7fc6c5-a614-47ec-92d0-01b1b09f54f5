db = new Mongo().getDB("dataimpact")

db.createCollection("user_auth")
db.user_auth.createIndex({"id": 1})
db.user_auth.insert({
    "id": "root",
    "hash": "$2a$10$O/tAYhALkST9lHDYf9Yq4eXbasb9K.aYvIGUJu/.HcjLYBJlqyWiK"
})

db.createCollection("users")
db.users.createIndex({"id": 1})
