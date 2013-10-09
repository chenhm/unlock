windres rc/go.rc -o temp-rc.o -F pe-i386 

export GOPKG=$GOPATH/pkg/windows_386
export SRC="main.go"
export OUT=unlock

go tool 8g -I $GOPKG $SRC

go tool pack grc _go_.8 main.8 temp-rc.o

go tool 8l -L $GOPKG -s -H windowsgui -o tmp.exe _go_.8

rm -f *.8 *.o

cat tmp.exe cat.exe > $OUT.exe

rm tmp.exe
