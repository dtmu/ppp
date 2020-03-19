go install
gox
rm $PWD/bin/*
mv ppp_* $PWD/bin/
