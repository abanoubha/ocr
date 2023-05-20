# Code Snippets & Automation Scripts

## git push one file

```fish
function gitpf --description "gitpf <message> <file>"
    git add $argv[2]
    git commit -m $argv[1]
    git push
end
```

## automate pushing one file

```fish
for f in img/binarize/0(seq 100 207).png;gitpf "[img] add
 test image for binarize/threshold process" $f;end;
```

## automate pushing one file w/ sleep

```fish
for f in img/binarize/0{067,071,072,073,074,075,076,077,078,079,080,081,082,083,084,085,086,087,088,089,090,091,092,093,094,095,09
6,097,098,099,100}.png;gitpf "[img] add test image for binarize/threshold process" $f;sleep 30;end;
```

** no need to write all of them , `seq` works well.

### with seq

```fish
for f in img/binarize/00(seq 67 99).png;gitpf "[img] add test image for binarize/threshold process" $f;sleep 30;end;
```
