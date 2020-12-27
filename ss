curl -H "Content-Type: application/json" -X POST -d '{"email":"aaaa@gmai.com","password" : "abcd","city":"kochi","age":80}' http://localhost:8080/user


mongo --port 27012 -u "test-user" -p "test-pass" --authenticationDatabase "admin"

db.createUser({
 user: "admin",
 pwd: "password", 
 roles: [{ role: "readWrite", db: "users"}]
});

mongo --host "replicaset-example/localhost:27011" -u "admin" -p "password" --authenticationDatabase "admin"