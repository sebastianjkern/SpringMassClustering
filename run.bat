protoc --go_out=./ --python_out=./python --pyi_out=./python ./spring.proto
protoc --go_out=./ --python_out=./python --pyi_out=./python ./mass.proto
protoc --go_out=./ --python_out=./python --pyi_out=./python ./trajectory.proto

go build .

cd ./python/
python preprocess.py

cd ..
SpringMassClustering.exe

cd ./python/
python app.py

cd ..
del SpringMassClustering.exe
