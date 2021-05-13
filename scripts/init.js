db = new Mongo().getDB("dataimpact")

db.createCollection("user_auth")
db.user_auth.createIndex({"id": 1})
db.user_auth.insert({
    "id": "root",
    "hash": "$2a$10$AvZZf4PrUp/59oYlSfGHz.ipb0fOhl304BxgjPpt6ofPYi8IW92x."
})

db.createCollection("users")
db.users.createIndex({"id": 1})
