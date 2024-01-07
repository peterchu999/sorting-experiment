echo "run script with $1 as parameter" 

cd sorting
go build main.go
chmod +x ./main
./main $1

cd ../visualization
python3 main.py


