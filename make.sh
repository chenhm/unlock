windres rc/go.rc -o temp-rc.o

export GOPKG=$GOPATH/pkg/windows_386
export SRC="main.go"
export OUT=unlock

go tool 8g -I$GOPKG $SRC

go tool pack grc _go_.8 main.8 temp-rc.o

go tool 8l -L$GOPKG -s -o tmp.exe _go_.8
#go tool 8l -L$GOPKG -s -Hwindowsgui -o tmp.exe _go_.8

rm *.8 *.o

cat tmp.exe cat.exe > $OUT.exe

rm tmp.exe
