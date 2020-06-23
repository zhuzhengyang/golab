#-gcflags=all=-l
go build -gcflags=-l -o ./tmp/main main.go
./tmp/main > ./tmp/main.log &
pid=$!
echo $pid
sleep 3
go run -gcflags=-l generate_plugin.go
sleep 3
rm tmp/example_patch.so
sleep 3
kill -9 $pid
cat ./tmp/main.log
